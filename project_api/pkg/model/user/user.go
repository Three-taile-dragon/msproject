package user

import (
	"errors"
	common "test.com/project_common"
)

// RegisterReq user 相关模型
// 注意加上 form 表单
type RegisterReq struct {
	Email     string `json:"email" form:"email"`
	Name      string `json:"name" form:"name"`
	Password  string `json:"password" form:"password"`
	Password2 string `json:"password2" form:"password2"`
	Mobile    string `json:"mobile" form:"mobile"`
	Captcha   string `json:"captcha" form:"captcha"`
}

func (r RegisterReq) VerifyPassword() bool {
	return r.Password == r.Password2
}

// Verify 验证参数
func (r RegisterReq) Verify() error {
	//验证 邮箱 手机号 密码 用户名等等是否合法
	if !common.VerifyEmailFormat(r.Email) {
		return errors.New("邮箱格式不正确")
	}
	if !common.VerifyMobile(r.Mobile) {
		return errors.New("手机号格式不正确")
	}
	if !r.VerifyPassword() {
		return errors.New("两次密输入不一致")
	}
	return nil
}

type LoginReq struct {
	Account  string `json:"account" form:"account"`
	Password string `json:"password" form:"password"`
}

type LoginRsp struct {
	Member           Member             `json:"member"`
	TokenList        TokenList          `json:"tokenList"`
	OrganizationList []OrganizationList `json:"organizationList"`
}
type Member struct {
	//Id     int64  `json:"id"`
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
	Status int    `json:"status"`
	Code   string `json:"code"` //使用加密后的ID代替
}

type TokenList struct {
	AccessToken    string `json:"accessToken"`
	RefreshToken   string `json:"refreshToken"`
	TokenType      string `json:"tokenType"`
	AccessTokenExp int64  `json:"accessTokenExp"`
}

type OrganizationList struct {
	//Id          int64  `json:"id"`
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
	MemberId    int64  `json:"memberId"`
	CreateTime  int64  `json:"createTime"`
	Personal    int32  `json:"personal"`
	Address     string `json:"address"`
	Province    int32  `json:"province"`
	City        int32  `json:"city"`
	Area        int32  `json:"area"`
	Code        string `json:"code"` //使用加密后的ID代替
}

func (l *LoginReq) VerifyAccount() error {
	if l.Account == "" {
		return errors.New("用户名不能为空")
	}
	return nil
}
