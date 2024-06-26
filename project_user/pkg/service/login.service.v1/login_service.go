package login_service_v1

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"strconv"
	"strings"
	common "test.com/project_common"
	"test.com/project_common/encrypts"
	"test.com/project_common/errs"
	"test.com/project_common/jwts"
	"test.com/project_common/tms"
	"test.com/project_grpc/user/login"
	"test.com/project_user/config"
	"test.com/project_user/internal/dao"
	"test.com/project_user/internal/dao/mysql"
	"test.com/project_user/internal/data/member"
	data "test.com/project_user/internal/data/organization"
	"test.com/project_user/internal/database"
	"test.com/project_user/internal/database/tran"
	"test.com/project_user/internal/repo"
	"test.com/project_user/pkg/model"
	"test.com/project_user/util"
	"time"
)

// LoginService grpc 登陆服务 实现
type LoginService struct {
	login.UnimplementedLoginServiceServer
	cache            repo.Cache
	memberRepo       repo.MemberRepo
	organizationRepo repo.OrganizationRepo
	transaction      tran.Transaction
}

func New() *LoginService {
	return &LoginService{
		cache:            dao.Rc,
		memberRepo:       mysql.NewMemberDao(),
		organizationRepo: mysql.NewOrganizationDao(),
		transaction:      dao.NewTransaction(),
	}
}

func (ls *LoginService) GetCaptcha(ctx context.Context, req *login.CaptchaRequest) (*login.CaptchaResponse, error) {
	//1.获取参数
	mobile := req.Mobile
	//2.校验参数
	if !common.VerifyMobile(mobile) {
		return nil, errs.GrpcError(model.NoLegalMobile) //使用自定义错误码进行处理
	}
	//3.生成验证码(随机四位1000-9999或者六位100000-999999)
	code := util.CreateCaptcha(6) //生成随机六位数字验证码
	//4.调用短信平台(第三方 放入go func 协程 接口可以快速响应
	go func() {
		time.Sleep(2 * time.Second)
		zap.L().Info("短信平台调用成功，发送短信")
		//logs.LG.Debug("短信平台调用成功，发送短信 debug")
		//zap.L().Debug("短信平台调用成功，发送短信 debug")
		//zap.L().Error("短信平台调用成功，发送短信 error")
		//redis存储	假设后续缓存可能存在mysql当中,也可以存在mongo当中,也可能存在memcache当中
		//使用接口 达到低耦合高内聚
		//5.存储验证码 redis 当中,过期时间15分钟
		//redis.Set"REGISTER_"+mobile, code)
		c, cancel := context.WithTimeout(context.Background(), 2*time.Second) //编写上下文 最多允许两秒超时
		defer cancel()
		err := ls.cache.Put(c, model.RegisterRedisKey+mobile, code, 5*time.Minute)
		if err != nil {
			zap.L().Error("验证码存入redis出错", zap.Error(err))

		}
		zap.L().Debug("将手机号和验证码存入redis成功：REGISTER_" + mobile + " : " + code + "\n")
	}()
	//注意code一般不发送
	//这里是做了简化处理 由于短信平台目前对于个人不好使用
	return &login.CaptchaResponse{Code: code}, nil
}

func (ls *LoginService) Register(ctx context.Context, req *login.RegisterRequest) (*login.RegisterResponse, error) {
	c := context.Background()
	//1.可以校验参数 这里 api服务中已经校验过 这里就不再校验
	//2.校验验证码
	redisCode, err := ls.cache.Get(c, model.RegisterRedisKey+req.Mobile)
	if errors.Is(err, redis.Nil) {
		return nil, errs.GrpcError(model.CaptchaNoExist)
	}
	if err != nil {
		zap.L().Error("Register 中 redis 读取错误", zap.Error(err))
		return nil, errs.GrpcError(model.RedisError)
	}
	if redisCode != req.Captcha {
		return nil, errs.GrpcError(model.CaptchaError)
	}
	//3.校验业务逻辑 (邮箱是否被注册 账号是否被注册 手机号是否被注册)
	//检验邮箱
	exist, err := ls.memberRepo.GetMemberByEmail(c, req.Email)
	if err != nil {
		zap.L().Error("数据库出错", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	if exist {
		return nil, errs.GrpcError(model.EmailExist)
	}
	//检验用户名
	exist, err = ls.memberRepo.GetMemberByAccount(c, req.Name)
	if err != nil {
		zap.L().Error("数据库出错", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	if exist {
		return nil, errs.GrpcError(model.AccountExist)
	}
	//检验手机号
	exist, err = ls.memberRepo.GetMemberByMobile(c, req.Mobile)
	if err != nil {
		zap.L().Error("数据库出错", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	if exist {
		return nil, errs.GrpcError(model.MobileExist)
	}
	//4.执行业务 将数据存入member表 生成一个数据 存入组织表 organization
	pwd := encrypts.Md5(req.Password) //加密部分
	mem := &member.Member{
		Account:       req.Name,
		Password:      pwd,
		Name:          req.Name,
		Mobile:        req.Mobile,
		Email:         req.Email,
		CreateTime:    time.Now().UnixMilli(),
		LastLoginTime: time.Now().UnixMilli(),
		Status:        model.Normal,
	}
	//将存入部分使用事务包裹 使得可以回滚数据库操作
	err = ls.transaction.Action(func(conn database.DbConn) error {
		err = ls.memberRepo.SaveMember(conn, c, mem)
		if err != nil {
			zap.L().Error("注册模块member数据库存入出错", zap.Error(err))
			return errs.GrpcError(model.DBError)
		}
		//存入组织
		org := &data.Organization{
			Name:       mem.Name + "个人项目",
			MemberId:   mem.Id,
			CreateTime: time.Now().UnixMilli(),
			Personal:   model.Personal,
			Avatar:     "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fc-ssl.dtstatic.com%2Fuploads%2Fblog%2F202103%2F31%2F20210331160001_9a852.thumb.1000_0.jpg&refer=http%3A%2F%2Fc-ssl.dtstatic.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1673017724&t=ced22fc74624e6940fd6a89a21d30cc5",
		}
		err = ls.organizationRepo.SaveOrganization(conn, c, org)
		if err != nil {
			zap.L().Error("注册模块organization数据库存入失败", zap.Error(err))
			return errs.GrpcError(model.DBError)
		}

		return nil
	})
	//将存入部分使用事务包裹 使得可以回滚数据库操作

	// 查询出新用户的 OrganizationId

	memOrg, err3 := ls.organizationRepo.FindOrganizationByMemId(c, mem.Id)
	if err3 != nil {
		zap.L().Error("user Register memberRepo.SaveMemberAccount DB Error", zap.Error(err3))
		return nil, errs.GrpcError(model.DBError)
	}
	err = ls.transaction.Action(func(conn database.DbConn) error {
		//生成 成员账户
		mAccount := &member.MemberAccount{
			OrganizationCode: memOrg[0].Id,
			MemberCode:       mem.Id,
			Authorize:        "3", // 默认权限为 管理员
			IsOwner:          1,
			Name:             mem.Name,
			Mobile:           mem.Mobile,
			Email:            mem.Email,
			CreateTime:       mem.CreateTime,
			LastLoginTime:    mem.LastLoginTime,
			Status:           1,
			Description:      mem.Description,
			Avatar:           mem.Avatar,
		}

		err = ls.memberRepo.SaveMemberAccount(conn, c, mAccount)
		if err != nil {
			zap.L().Error("user Register memberRepo.SaveMemberAccount DB Error", zap.Error(err))
			return errs.GrpcError(model.DBError)
		}
		return nil
	})

	//5.返回

	return &login.RegisterResponse{}, err
}

func (ls *LoginService) Login(ctx context.Context, req *login.LoginRequest) (*login.LoginResponse, error) {
	c := context.Background()
	//1.传入参数
	//2.校验参数
	//3.校验用户名
	//检验邮箱和用户名
	exist, err := ls.memberRepo.GetMemberByAccountAndEmail(c, req.Account)
	if err != nil {
		zap.L().Error("数据库出错", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	if !exist {
		return nil, errs.GrpcError(model.AccountNoExist)
	}
	//4.去数据库查询 账号密码是否正确
	pwd := encrypts.Md5(req.Password)
	mem, err := ls.memberRepo.FindMember(c, req.Account, pwd)
	if err != nil {
		zap.L().Error("登陆模块member数据库查询出错", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	if mem == nil {
		return nil, errs.GrpcError(model.AccountAndPwdError)
	}
	memMessage := &login.MemberMessage{}
	err = copier.Copy(&memMessage, mem)
	if err != nil {
		zap.L().Error("登陆模块memMessage赋值错误", zap.Error(err))
		return nil, errs.GrpcError(model.CopyError)
	}
	memMessage.Code, _ = encrypts.EncryptInt64(mem.Id, model.AESKey) //加密用户ID
	memMessage.LastLoginTime = tms.FormatByMill(mem.LastLoginTime)
	memMessage.CreateTime = tms.FormatByMill(mem.CreateTime)
	//5.根据用户id查组织
	orgs, err := ls.organizationRepo.FindOrganizationByMemId(c, mem.Id)
	if err != nil {
		zap.L().Error("登陆模块organization数据库查询出错", zap.Error(err))
		return nil, errs.GrpcError(model.OrganizationNoExist)
	}
	var orgsMessage []*login.OrganizationMessage
	err = copier.Copy(&orgsMessage, orgs)
	if err != nil {
		zap.L().Error("登陆模块orgs赋值错误", zap.Error(err))
		return nil, errs.GrpcError(model.CopyError)
	}
	for _, v := range orgsMessage {
		v.Code, _ = encrypts.EncryptInt64(v.Id, model.AESKey) //加密组织ID
		v.OwnerCode = memMessage.Code
		organization := data.ToMap(orgs)[v.Id]
		v.CreateTime = tms.FormatByMill(organization.CreateTime)
	}
	if len(orgs) > 0 {
		memMessage.OrganizationCode, _ = encrypts.EncryptInt64(orgs[0].Id, model.AESKey)
	}

	//6.用jwt生成token
	memIdStr := strconv.FormatInt(mem.Id, 10)
	exp := time.Duration(config.C.JC.AccessExp*3600*24) * time.Second
	rExp := time.Duration(config.C.JC.RefreshExp*3600*24) * time.Second
	token := jwts.CreateToken(memIdStr, exp, config.C.JC.AccessSecret, config.C.JC.RefreshSecret, rExp, req.Ip)
	// 可以给 token 做加密处理 增加安全性
	tokenList := &login.TokenMessage{
		AccessToken:    token.AccessToken,
		RefreshToken:   token.RefreshToken,
		TokenType:      "bearer",
		AccessTokenExp: token.AccessExp,
	}

	// 放入缓存 member organization
	go func() {
		memJson, _ := json.Marshal(mem)
		err2 := ls.cache.Put(c, model.Member+"::"+memIdStr, string(memJson), exp)
		if err2 != nil {
			zap.L().Error("login_service Login::MemID_Redis_Cache error", zap.Error(err))

		}
		orgsJson, _ := json.Marshal(orgs)
		err2 = ls.cache.Put(c, model.MemberOrganization+"::"+memIdStr, string(orgsJson), exp)
		if err2 != nil {
			zap.L().Error("login_service Login::Organization_Redis_Cache error", zap.Error(err))

		}
	}()
	//7.结果返回
	return &login.LoginResponse{
		Member:           memMessage,
		OrganizationList: orgsMessage,
		TokenList:        tokenList,
	}, nil
}

// TokenVerify token验证
func (ls *LoginService) TokenVerify(ctx context.Context, msg *login.TokenRequest) (*login.LoginResponse, error) {
	c := context.Background()
	token := msg.Token
	if strings.Contains(token, "bearer") {
		token = strings.ReplaceAll(token, "bearer ", "")
	}
	parseToken, err := jwts.ParseToken(token, config.C.JC.AccessSecret, msg.Ip)
	if err != nil {
		zap.L().Error("login_service Login TokenVerify::ParseToken error", zap.Error(err))
		return nil, errs.GrpcError(model.NoLogin)
	}

	// 从缓存中查询 memberId
	memJson, err := ls.cache.Get(c, model.Member+"::"+parseToken)
	if err != nil {
		zap.L().Error("login_service Login TokenVerify::cache get member error ", zap.Error(err))
		return nil, errs.GrpcError(model.NoLogin)
	}
	if memJson == "" { //过期了
		zap.L().Error("login_service Login TokenVerify::cache get member expire ")
		return nil, errs.GrpcError(model.NoLogin)
	}
	memberById := &member.Member{}
	err = json.Unmarshal([]byte(memJson), memberById)
	if err != nil {
		zap.L().Error("login_service Login TokenVerify::cache Unmarshal member error ", zap.Error(err))
		return nil, errs.GrpcError(model.SystemError)
	}

	//数据库查询 优化点 登陆之后应该把用户信息缓存起来
	//id, _ := strconv.ParseInt(parseToken, 10, 64)
	//memberById, err := ls.memberRepo.FindMemberById(c, id)
	//if err != nil {
	//	zap.L().Error("login_service Login TokenVerify::FindMemberById db error ", zap.Error(err))
	//	return nil, errs.GrpcError(model.DBError)
	//}

	memMessage := &login.MemberMessage{}
	err = copier.Copy(&memMessage, memberById)
	if err != nil {
		zap.L().Error("Token验证模块memMessage赋值错误", zap.Error(err))
		return nil, errs.GrpcError(model.CopyError)
	}
	memMessage.Code, _ = encrypts.EncryptInt64(memberById.Id, model.AESKey) //加密用户ID
	//从 redis缓存中读取organization
	orgsJson, err := ls.cache.Get(c, model.MemberOrganization+"::"+parseToken)
	if err != nil {
		zap.L().Error("login_service Login TokenVerify::cache get organization error ", zap.Error(err))
		return nil, errs.GrpcError(model.NoLogin)
	}
	if orgsJson == "" { //过期了
		zap.L().Error("login_service Login TokenVerify::cache get organization expire ")
		return nil, errs.GrpcError(model.NoLogin)
	}
	var orgs []*data.Organization
	err = json.Unmarshal([]byte(orgsJson), &orgs)
	if err != nil {
		zap.L().Error("login_service Login TokenVerify::cache Unmarshal organization error ", zap.Error(err))
		return nil, errs.GrpcError(model.SystemError)
	}

	//orgs, err := ls.organizationRepo.FindOrganizationByMemId(c, memMessage.Id)
	//if err != nil {
	//	zap.L().Error("Token验证模块organization数据库查询出错", zap.Error(err))
	//	return nil, errs.GrpcError(model.OrganizationNoExist)
	//}

	if len(orgs) > 0 {
		memMessage.OrganizationCode, _ = encrypts.EncryptInt64(orgs[0].Id, model.AESKey)

	}
	memMessage.CreateTime = tms.FormatByMill(memberById.CreateTime)
	return &login.LoginResponse{Member: memMessage}, nil
}

func (ls *LoginService) MyOrgList(ctx context.Context, msg *login.UserMessage) (*login.OrgListResponse, error) {
	memId := msg.MemId
	orgs, err := ls.organizationRepo.FindOrganizationByMemId(ctx, memId)
	if err != nil {
		zap.L().Error("用户模块组织列表获取失败", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	var orgsMessage []*login.OrganizationMessage
	err = copier.Copy(&orgsMessage, orgs)
	for _, org := range orgsMessage {
		org.Code, _ = encrypts.EncryptInt64(org.Id, model.AESKey)
	}
	return &login.OrgListResponse{OrganizationList: orgsMessage}, nil
}

// grpc 与 project 模块

func (ls *LoginService) FindMemInfoById(ctx context.Context, req *login.UserMessage) (*login.MemberMessage, error) {
	c := context.Background()
	memberById, err := ls.memberRepo.FindMemberById(c, req.MemId)
	if err != nil {
		zap.L().Error("login FindMemInfoById FindMemberById error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	memMessage := &login.MemberMessage{}
	err = copier.Copy(&memMessage, memberById)
	if err != nil {
		zap.L().Error("login FindMemInfoById Copy error", zap.Error(err))
		return nil, errs.GrpcError(model.CopyError)
	}
	memMessage.Code, _ = encrypts.EncryptInt64(memberById.Id, model.AESKey) //加密用户ID

	orgs, err := ls.organizationRepo.FindOrganizationByMemId(c, memMessage.Id)
	if err != nil {
		zap.L().Error("login FindMemInfoById FindOrganizationByMemId error", zap.Error(err))
		return nil, errs.GrpcError(model.OrganizationNoExist)
	}

	if len(orgs) > 0 {
		memMessage.OrganizationCode, _ = encrypts.EncryptInt64(orgs[0].Id, model.AESKey)
		memMessage.CreateTime = tms.FormatByMill(memberById.CreateTime)
	}
	return memMessage, nil
}

func (ls *LoginService) FindMemInfoByIds(ctx context.Context, req *login.UserMessage) (*login.MemberMessageList, error) {
	c := context.Background()
	memberList, err := ls.memberRepo.FindMemberByIds(c, req.MIds)
	if err != nil {
		zap.L().Error("login FindMemInfoByIds memberRepo.FindMemberByIds error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}

	if memberList == nil || len(memberList) <= 0 {
		return &login.MemberMessageList{List: nil}, nil
	}

	var memMsgs []*login.MemberMessage
	err = copier.Copy(&memMsgs, memberList)

	mMap := make(map[int64]*member.Member)
	for _, v := range memberList {
		mMap[v.Id] = v
	}

	for _, v := range memMsgs {
		m := mMap[v.Id]
		v.CreateTime = tms.FormatByMill(m.CreateTime)
		v.Code = encrypts.EncryptInt64NoErr(v.Id)
	}

	return &login.MemberMessageList{List: memMsgs}, nil
}
