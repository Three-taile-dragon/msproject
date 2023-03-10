package model

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Page struct {
	Page     int64 `json:"page" form:"page"`
	PageSize int64 `json:"pageSize" form:"pageSize"`
}

func (p *Page) Bind(c *gin.Context) {
	err := c.ShouldBind(&p)
	if err != nil {
		zap.L().Error("分页模型绑定失败", zap.Error(err))
	}
	//定义默认值
	if p.Page == 0 {
		p.Page = 1
	}
	if p.PageSize == 0 {
		p.PageSize = 10
	}
}
