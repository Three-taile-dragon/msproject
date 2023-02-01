package main

import (
	"github.com/gin-gonic/gin"
	common "test.com/project_common"
)

func main() {
	r := gin.Default()
	//r.Run(":8080")
	common.Run(r, "project_user", ":80")
}
