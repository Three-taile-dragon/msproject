package domain

import (
	"context"
	"fmt"
	"test.com/project_common/encrypts"
	"test.com/project_common/errs"
	"test.com/project_project/internal/dao/mysql"
	"test.com/project_project/internal/data/account"
	"test.com/project_project/internal/repo"
	"test.com/project_project/pkg/model"
	"time"
)

type AccountDomain struct {
	accountRepo      repo.AccountRepo
	userRpcDomain    *UserRpcDomain
	departmentDomain *DepartmentDomain
}

func NewAccountDomain() *AccountDomain {
	return &AccountDomain{
		accountRepo:      mysql.NewMemberAccountDao(),
		userRpcDomain:    NewUserRpcDomain(),
		departmentDomain: NewDepartmentDomain(),
	}
}

func (d *AccountDomain) AccountList(organizationCode string, memberId int64, page int64, pageSize int64, departmentCode string, searchType int32) ([]*account.MemberAccountDisplay, int64, *errs.BError) {
	condition := ""
	organizationCodeId := encrypts.DecryptNoErr(organizationCode)
	departmentCodeId := encrypts.DecryptNoErr(departmentCode)
	switch searchType {
	case 1:
		// 查询 使用当中
		condition = "status = 1"
	case 2:
		// 查询 部门为空的
		condition = "department_code = NULL"
	case 3:
		// 查询 禁用的
		condition = "status = 0"
	case 4:
		// 查询 正在使用并且位于当前部门下
		condition = fmt.Sprintf("status = 1 and department_code = %d", departmentCodeId)
	default:
		condition = "status = 1"
	}
	// 调用
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	list, total, err := d.accountRepo.FindList(c, condition, organizationCodeId, departmentCodeId, page, pageSize)
	if err != nil {
		return nil, 0, model.DBError
	}
	// 数据转换 返回
	var dList []*account.MemberAccountDisplay
	for _, v := range list {
		display := v.ToDisplay()
		memberInfo, _ := d.userRpcDomain.MemberInfo(c, v.MemberCode)
		display.Avatar = memberInfo.Avatar
		if v.DepartmentCode > 0 {
			department, err := d.departmentDomain.FindDepartmentById(v.DepartmentCode)
			if err != nil {
				return nil, 0, err
			}
			display.Departments = department.Name
		}
		dList = append(dList, display)
	}
	return dList, total, nil
}
