package model

import (
	"test.com/project_common/errs"
)

//定义状态码

var (
	NoLegalMobile = errs.NewError(2001, "手机号不合法") //不合法的手机号
)
