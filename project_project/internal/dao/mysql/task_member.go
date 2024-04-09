package mysql

import (
	"context"
	"test.com/project_project/internal/data"
	"test.com/project_project/internal/database/gorms"
)

type TaskMemberDao struct {
	conn *gorms.GormConn
}

func NewTaskMemberDao() *TaskMemberDao {
	return &TaskMemberDao{
		conn: gorms.New(),
	}
}

func (t TaskMemberDao) FindTaskMemberByTaskId(ctx context.Context, taskCode int64, memberId int64) (task *data.TaskMember, err error) {
	session := t.conn.Session(ctx)
	err = session.Where("task_code = ? and member_code = ?", taskCode, memberId).Limit(1).Find(&task).Error
	return
}
