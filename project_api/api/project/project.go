package project

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"test.com/project_api/api/rpc"
	"test.com/project_api/pkg/model"
	"test.com/project_api/pkg/model/menu"
	pro "test.com/project_api/pkg/model/project"
	common "test.com/project_common"
	"test.com/project_common/errs"
	"test.com/project_grpc/project"
	"time"
)

type HandlerProject struct {
}

func New() *HandlerProject {
	return &HandlerProject{}
}
func (p *HandlerProject) index(c *gin.Context) {
	result := &common.Result{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	msg := &project_service_v1.IndexRequest{}
	rsp, err := rpc.ProjectServiceClient.Index(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
	}
	menus := rsp.Menus
	var ms []*menu.Menu
	err = copier.Copy(&ms, menus)
	if err != nil {
		zap.L().Error("api模块menu复制失败", zap.Error(err))
		c.JSON(http.StatusOK, result.Fail(00000, "系统内部错误"))
	}
	c.JSON(http.StatusOK, result.Success(ms))
}

func (p *HandlerProject) myProjectList(c *gin.Context) {
	result := &common.Result{}
	//1. 获取参数
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	memberId := c.GetInt64("memberId")
	memberName := c.GetString("memberName")
	organizationCode := c.GetString("organizationCode")
	//memberId := memberIdStr.(int64) //转换
	//分页
	page := &model.Page{}
	page.Bind(c)
	selectBy := c.PostForm("selectBy")
	msg := &project_service_v1.ProjectRpcMessage{
		MemberId:         memberId,
		MemberName:       memberName,
		SelectBy:         selectBy,
		Page:             page.Page,
		PageSize:         page.PageSize,
		OrganizationCode: organizationCode,
	}
	myProjectResponse, err := rpc.ProjectServiceClient.FindProjectByMemId(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
	}

	var pms []*pro.ProAndMember
	err = copier.Copy(&pms, myProjectResponse.Pm)
	if err != nil {
		zap.L().Error("项目列表模块返回数据复制出错", zap.Error(err))
		c.JSON(http.StatusOK, result.Fail(502, "系统内部错误"))
	}
	//设定默认值
	if pms == nil {
		pms = []*pro.ProAndMember{}
	}
	c.JSON(http.StatusOK, result.Success(gin.H{
		"list":  pms, //不能返回 null nil, 空的话要返回[] 不然前端没法判断
		"total": myProjectResponse.Total,
	}))
}

func (p *HandlerProject) projectTemplate(c *gin.Context) {
	result := &common.Result{}
	//1. 获取参数
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	memberId := c.GetInt64("memberId")
	memberName := c.GetString("memberName")
	organizationCode := c.GetString("organizationCode")
	//memberId := memberIdStr.(int64) //转换
	//分页
	page := &model.Page{}
	page.Bind(c)
	viewTypeStr := c.PostForm("viewType")
	viewType, _ := strconv.ParseInt(viewTypeStr, 10, 64)
	msg := &project_service_v1.ProjectRpcMessage{
		MemberId:         memberId,
		MemberName:       memberName,
		ViewType:         int32(viewType),
		Page:             page.Page,
		PageSize:         page.PageSize,
		OrganizationCode: organizationCode,
	}
	templateResponse, err := rpc.ProjectServiceClient.FindProjectTemplate(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
	}

	var pms []*pro.ProjectTemplate
	err = copier.Copy(&pms, templateResponse.Ptm)
	if err != nil {
		zap.L().Error("项目模板模块返回数据复制出错", zap.Error(err))
		c.JSON(http.StatusOK, result.Fail(502, "系统内部错误"))
	}
	//设定默认值
	if pms == nil {
		pms = []*pro.ProjectTemplate{}
	}
	//设定默认值	避免出现返回值为nil
	for _, v := range pms {
		if v.TaskStages == nil {
			v.TaskStages = []*pro.TaskStagesOnlyName{}
		}
	}
	c.JSON(http.StatusOK, result.Success(gin.H{
		"list":  pms, //不能返回 null nil, 空的话要返回[] 不然前端没法判断
		"total": templateResponse.Total,
	}))
}

func (p *HandlerProject) projectSave(c *gin.Context) {
	result := &common.Result{}
	//1. 获取参数
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	memberId := c.GetInt64("memberId")
	organizationCode := c.GetString("organizationCode")
	var req *pro.SaveProjectRequest
	err := c.ShouldBind(&req)
	if err != nil {
		zap.L().Error("项目创建模块_模型绑定失败", zap.Error(err))
		c.JSON(http.StatusOK, result.Fail(502, "系统内部错误"))
	}
	msg := &project_service_v1.ProjectRpcMessage{
		MemberId:         memberId,
		OrganizationCode: organizationCode,
		TemplateCode:     req.TemplateCode,
		Name:             req.Name,
		Id:               int64(req.Id),
		Description:      req.Description,
	}
	saveProject, err := rpc.ProjectServiceClient.SaveProject(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
	}
	//var rsp *pro.SaveProject
	////copier.Copy(&rsp, saveProject)
	//err = copier.Copy(&rsp, saveProject)
	//if err != nil {
	//	zap.L().Error("项目创建模块返回数据复制出错", zap.Error(err))
	//	c.JSON(http.StatusOK, result.Fail(502, "系统内部错误"))
	//}
	rsp := &pro.SaveProject{
		Id:               saveProject.Id,
		Cover:            saveProject.Cover,
		Name:             saveProject.Name,
		Description:      saveProject.Description,
		Code:             saveProject.Code,
		CreateTime:       saveProject.CreateTime,
		TaskBoardTheme:   saveProject.TaskBoardTheme,
		OrganizationCode: saveProject.OrganizationCode,
	}

	c.JSON(http.StatusOK, result.Success(rsp))

}

func (p *HandlerProject) projectRead(c *gin.Context) {
	result := &common.Result{}
	//1. 获取参数
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	memberId := c.GetInt64("memberId")
	projectCode := c.PostForm("projectCode")
	detail, err := rpc.ProjectServiceClient.FindProjectDetail(ctx, &project_service_v1.ProjectRpcMessage{ProjectCode: projectCode, MemberId: memberId})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
	}
	pd := pro.ProjectDetail{}
	err = copier.Copy(&pd, detail)
	if err != nil {
		zap.L().Error("项目读取模块返回数据复制出错", zap.Error(err))
		c.JSON(http.StatusOK, result.Fail(502, "系统内部错误"))
	}
	c.JSON(http.StatusOK, result.Success(pd))
}

func (p *HandlerProject) projectRecycle(c *gin.Context) {
	result := &common.Result{}
	//1. 获取参数
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	projectCode := c.PostForm("projectCode")
	_, err := rpc.ProjectServiceClient.RecycleProject(ctx, &project_service_v1.ProjectRpcMessage{ProjectCode: projectCode})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
	}
	c.JSON(http.StatusOK, result.Success([]int{}))
}

func (p *HandlerProject) projectRecovery(c *gin.Context) {
	result := &common.Result{}
	//1. 获取参数
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	projectCode := c.PostForm("projectCode")
	_, err := rpc.ProjectServiceClient.RecoveryProject(ctx, &project_service_v1.ProjectRpcMessage{ProjectCode: projectCode})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
	}
	c.JSON(http.StatusOK, result.Success([]int{}))
}

func (p *HandlerProject) projectCollect(c *gin.Context) {
	result := &common.Result{}
	//1. 获取参数
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	projectCode := c.PostForm("projectCode")
	collectType := c.PostForm("type")
	memberId := c.GetInt64("memberId")
	_, err := rpc.ProjectServiceClient.CollectProject(ctx, &project_service_v1.ProjectRpcMessage{ProjectCode: projectCode, CollectType: collectType, MemberId: memberId})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
	}
	c.JSON(http.StatusOK, result.Success([]int{}))
}

func (p *HandlerProject) projectEdit(c *gin.Context) {
	result := &common.Result{}
	//1. 获取参数
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	var req *pro.ProjectReq
	_ = c.ShouldBind(&req)
	memberId := c.GetInt64("memberId")
	msg := &project_service_v1.UpdateProjectMessage{}
	err := copier.Copy(msg, req)
	if err != nil {
		zap.L().Error("api projectEdit Copy error", zap.Error(err))
		c.JSON(http.StatusOK, result.Fail(errs.ParseGrpcError(err)))
	}
	msg.MemberId = memberId
	_, err = rpc.ProjectServiceClient.UpdateProject(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
	}
	c.JSON(http.StatusOK, result.Success([]int{}))
}
