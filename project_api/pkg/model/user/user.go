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
