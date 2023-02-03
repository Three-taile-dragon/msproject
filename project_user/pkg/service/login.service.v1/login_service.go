package login_service_v1

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	common "test.com/project_common"
	"test.com/project_common/encrypts"
	"test.com/project_common/errs"
	"test.com/project_grpc/user/login"
	"test.com/project_user/internal/dao"
	"test.com/project_user/internal/dao/mysql"
	"test.com/project_user/internal/data/member"
	data "test.com/project_user/internal/data/organization"
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
}

func New() *LoginService {
	return &LoginService{
		cache:            dao.Rc,
		memberRepo:       mysql.NewMemberDao(),
		organizationRepo: mysql.NewOrganizationDao(),
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
	if err == redis.Nil {
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
	err = ls.memberRepo.SaveMember(c, mem)
	if err != nil {
		zap.L().Error("注册模块member数据库存入出错", zap.Error(err))
		return nil, err
	}
	//存入组织
	org := &data.Organization{
		Name:       mem.Name + "个人项目",
		MemberId:   mem.Id,
		CreateTime: time.Now().UnixMilli(),
		Personal:   model.Personal,
		Avatar:     "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fc-ssl.dtstatic.com%2Fuploads%2Fblog%2F202103%2F31%2F20210331160001_9a852.thumb.1000_0.jpg&refer=http%3A%2F%2Fc-ssl.dtstatic.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1673017724&t=ced22fc74624e6940fd6a89a21d30cc5",
	}
	err = ls.organizationRepo.SaveOrganization(c, org)
	if err != nil {
		zap.L().Error("注册模块organization数据库存入失败", zap.Error(err))
		return nil, model.DBError
	}
	//5.返回

	return &login.RegisterResponse{}, nil
}
