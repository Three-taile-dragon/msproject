package department

import (
	"github.com/jinzhu/copier"
	"test.com/project_common/encrypts"
	"test.com/project_common/tms"
)

type Department struct {
	Id               int64
	OrganizationCode int64
	Name             string
	Sort             int
	PCode            int64
	icon             string
	CreateTime       int64
	Path             string
}

func (*Department) TableName() string {
	return "ms_department"
}

type DepartmentDisplay struct {
	Id               int64
	OrganizationCode string
	Name             string
	Sort             int
	Pcode            string
	icon             string
	CreateTime       string
	Path             string
}

func (d *Department) ToDisplay() *DepartmentDisplay {
	dp := &DepartmentDisplay{}
	_ = copier.Copy(dp, d)
	dp.CreateTime = tms.FormatByMill(d.CreateTime)
	dp.OrganizationCode = encrypts.EncryptInt64NoErr(d.OrganizationCode)
	if d.PCode > 0 {
		dp.Pcode = encrypts.EncryptInt64NoErr(d.PCode)
	} else {
		dp.Pcode = ""
	}
	return dp
}
