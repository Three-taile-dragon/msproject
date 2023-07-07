package project_service_v1

import (
	"context"
	"go.uber.org/zap"
	"strconv"
	"test.com/project_common/encrypts"
	"test.com/project_common/errs"
	"test.com/project_grpc/project"
	pro "test.com/project_project/internal/data/project"
	"test.com/project_project/pkg/model"
	"time"
)

func (ps *ProjectService) CollectProject(ctx context.Context, msg *project.ProjectRpcMessage) (*project.CollectProjectResponse, error) {
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second) //编写上下文 最多允许两秒超时
	defer cancel()
	cipherIdCodeStr, _ := encrypts.Decrypt(msg.ProjectCode, model.AESKey)
	cipherIdCode, _ := strconv.ParseInt(cipherIdCodeStr, 10, 64)
	projectCode, err := ps.projectRepo.FindProjectByCipId(c, cipherIdCode)
	if err != nil {
		zap.L().Error("project CollectProject FindProjectByCipId error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	collectType := msg.CollectType
	if collectType == "collect" {
		pc := &pro.ProjectCollection{
			ProjectCode: projectCode,
			MemberCode:  msg.MemberId,
			CreateTime:  time.Now().UnixMilli(),
		}
		err = ps.projectRepo.CollectProject(c, pc)
		if err != nil {
			zap.L().Error("project CollectProject CollectProject error", zap.Error(err))
			return nil, errs.GrpcError(model.DBError)
		}
	} else if collectType == "cancel" {
		err = ps.projectRepo.CancelCollectProject(c, projectCode, msg.MemberId)
		if err != nil {
			zap.L().Error("project CollectProject CancelCollectProject error", zap.Error(err))
			return nil, errs.GrpcError(model.DBError)
		}
	}

	return &project.CollectProjectResponse{}, nil
}
