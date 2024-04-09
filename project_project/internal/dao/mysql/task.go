package mysql

import (
	"context"
	"test.com/project_project/internal/data"
	"test.com/project_project/internal/database/gorms"
)

type TaskDao struct {
	conn *gorms.GormConn
}

func NewTaskDao() *TaskDao {
	return &TaskDao{
		conn: gorms.New(),
	}
}

func (t TaskDao) FindTaskByStageCode(ctx context.Context, stageCode int) (list []*data.Task, err error) {
	session := t.conn.Session(ctx)
	err = session.Model(&data.Task{}).Where("stage_code = ? and deleted = 0", stageCode).Order("sort asc").Find(&list).Error
	return
}
