package data

import (
	"github.com/jinzhu/copier"
	"test.com/project_common/encrypts"
	"test.com/project_common/tms"
)

type TaskWorkTime struct {
	Id         int64
	TaskCode   int64
	MemberCode int64
	CreateTime int64
	Content    string
	BeginTime  int64
	Num        int
}

func (*TaskWorkTime) TableName() string {
	return "ms_task_work_time"
}

type TaskWorkTimeDisplay struct {
	Id         int64
	TaskCode   string
	MemberCode string
	CreateTime string
	Content    string
	BeginTime  string
	Num        int
	Member     Member
}

func (t *TaskWorkTime) ToDisplay() *TaskWorkTimeDisplay {
	td := &TaskWorkTimeDisplay{}
	_ = copier.Copy(td, t)
	td.MemberCode = encrypts.EncryptInt64NoErr(t.MemberCode)
	td.TaskCode = encrypts.EncryptInt64NoErr(t.TaskCode)
	td.CreateTime = tms.FormatByMill(t.CreateTime)
	td.BeginTime = tms.FormatByMill(t.BeginTime)
	return td
}
