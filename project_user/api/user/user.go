package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlerUser struct {
}

func (*HandlerUser) getCaptcha(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "getCaptcha success")
}
