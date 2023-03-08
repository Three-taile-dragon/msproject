package mysql

import (
	"context"
	"test.com/project_project/internal/data/project"
	"test.com/project_project/internal/database/gorms"
)

type ProjectDao struct {
	conn *gorms.GormConn
}

func NewProjectDao() *ProjectDao {
	return &ProjectDao{
		conn: gorms.New(),
	}
}

func (p ProjectDao) FindProjectByMemId(ctx context.Context, memId int64, page int64, size int64) ([]*project.ProjectAndMember, int64, error) {
	var pms []*project.ProjectAndMember
	//数据库分表查询
	session := p.conn.Session(ctx)
	index := (page - 1) * size
	raw := session.Raw("select * from ms_project a, ms_project_member b where a.id=b.project_code and b.member_code=? limit ?,?", memId, index, size)
	err := raw.Scan(&pms).Error //扫描进 结构体
	var total int64
	session.Model(&project.MemberProject{}).Where("member_code=?", memId).Count(&total)
	return pms, total, err
}
