package data

import (
	"github.com/jinzhu/copier"
	"test.com/project_common/encrypts"
	"test.com/project_common/tms"
)

type ProjectLog struct {
	Id           int64
	MemberCode   int64
	Content      string
	Remark       string
	Type         string
	CreateTime   int64
	SourceCode   int64
	ActionType   string
	ToMemberCode int64
	IsComment    int
	ProjectCode  int64
	Icon         string
	IsRobot      int
}

func (*ProjectLog) TableName() string {
	return "ms_project_log"
}

type ProjectLogDisplay struct {
	Id           int64
	MemberCode   string
	Content      string
	Remark       string
	Type         string
	CreateTime   string
	SourceCode   string
	ActionType   string
	ToMemberCode string
	IsComment    int
	ProjectCode  string
	Icon         string
	IsRobot      int
	Member       Member
}

func (l *ProjectLog) ToDisplay() *ProjectLogDisplay {
	pd := &ProjectLogDisplay{}
	_ = copier.Copy(pd, l)
	pd.MemberCode = encrypts.EncryptInt64NoErr(l.MemberCode)
	pd.ToMemberCode = encrypts.EncryptInt64NoErr(l.ToMemberCode)
	pd.ProjectCode = encrypts.EncryptInt64NoErr(l.ProjectCode)
	pd.CreateTime = tms.FormatByMill(l.CreateTime)
	pd.SourceCode = encrypts.EncryptInt64NoErr(l.SourceCode)
	return pd
}
