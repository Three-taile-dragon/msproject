package account_service_v1

import (
	"context"
	"github.com/jinzhu/copier"
	"test.com/project_common/encrypts"
	"test.com/project_common/errs"
	"test.com/project_grpc/department"
	"test.com/project_project/internal/dao"
	"test.com/project_project/internal/database/tran"
	"test.com/project_project/internal/domain"
	"test.com/project_project/internal/repo"
)

// DepartmentService grpc 登陆服务 实现
type DepartmentService struct {
	department.UnimplementedDepartmentServiceServer
	cache            repo.Cache
	transaction      tran.Transaction
	departmentDomain *domain.DepartmentDomain
}

func New() *DepartmentService {
	return &DepartmentService{
		cache:            dao.Rc,
		transaction:      dao.NewTransaction(),
		departmentDomain: domain.NewDepartmentDomain(),
	}
}

func (d *DepartmentService) List(ctx context.Context, msg *department.DepartmentReqMessage) (*department.ListDepartmentMessage, error) {
	// 解析参数
	organizationCode := encrypts.DecryptNoErr(msg.OrganizationCode)
	var parentDepartmentCode int64
	if msg.ParentDepartmentCode != "" {
		parentDepartmentCode = encrypts.DecryptNoErr(msg.ParentDepartmentCode)
	}
	// 调用 domain 接口
	dps, total, err := d.departmentDomain.List(
		organizationCode,
		parentDepartmentCode,
		msg.Page,
		msg.PageSize)
	if err != nil {
		return nil, errs.GrpcError(err)
	}
	// 返回数据
	var list []*department.DepartmentMessage
	_ = copier.Copy(&list, dps)
	return &department.ListDepartmentMessage{List: list, Total: total}, nil
}
