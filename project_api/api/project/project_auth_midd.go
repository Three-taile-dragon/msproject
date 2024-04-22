package project

import (
	"github.com/gin-gonic/gin"
	"net/http"
	common "test.com/project_common"
	"test.com/project_common/errs"
)

func projectAuth() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		// 如果此用户不是项目的成员
		// 认为不能操作项目 直接报无权限
		result := &common.Result{}
		//在接口有权限的基础上，做项目权限，不是这个项目的成员，无权限查看项目和操作项目
		//检查是否有projectCode和taskCode这两个参数 如果都有 则进行项目权限判断
		isProjectAuth := false
		projectCode := c.PostForm("projectCode")
		if projectCode != "" {
			isProjectAuth = true
		}
		taskCode := c.PostForm("taskCode")
		if taskCode != "" {
			isProjectAuth = true
		}

		if isProjectAuth {
			memberId := c.GetInt64("memberId")
			p := New()
			pr, isMember, isOwner, err := p.FindProjectByMemberId(memberId, projectCode, taskCode)
			if err != nil {
				code, msg := errs.ParseGrpcError(err)
				c.JSON(http.StatusOK, result.Fail(code, msg))
				c.Abort()
				return
			}
			if !isMember && !isOwner {
				c.JSON(http.StatusOK, result.Fail(403, "不是项目成员，无操作权限"))
				c.Abort()
				return
			}
			if pr.Private == 1 {
				//私有项目
				if isOwner || isMember {
					c.Next()
					return
				} else {
					c.JSON(http.StatusOK, result.Fail(403, "私有项目，无操作权限"))
					c.Abort()
					return
				}
			}
		}

	}
}
