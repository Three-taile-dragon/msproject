package mysql

import (
	"context"
	"test.com/project_project/internal/data"
	"test.com/project_project/internal/database/gorms"
)

type ProjectTemplateDao struct {
	conn *gorms.GormConn
}

func NewProjectTemplateDao() *ProjectTemplateDao {
	return &ProjectTemplateDao{
		conn: gorms.New(),
	}
}

func (p *ProjectTemplateDao) FindProjectTemplateSystem(ctx context.Context, page int64, size int64) ([]data.ProjectTemplate, int64, error) {
	var pts []data.ProjectTemplate
	session := p.conn.Session(ctx)
	err := session.Model(&data.ProjectTemplate{}).Where("is_system=?", 1).Limit(int(size)).Offset(int((page - 1) * size)).Find(&pts).Error
	var total int64
	session.Model(&data.ProjectTemplate{}).Where("is_system=?", 1).Count(&total)
	return pts, total, err
}

func (p ProjectTemplateDao) FindProjectTemplateCustom(ctx context.Context, memId int64, organizationCode int64, page int64, size int64) ([]data.ProjectTemplate, int64, error) {
	var pts []data.ProjectTemplate
	session := p.conn.Session(ctx)
	err := session.Model(&data.ProjectTemplate{}).
		Where("is_system=? and member_code=? and organization_code=?", 0, memId, organizationCode).
		Limit(int(size)).
		Offset(int((page - 1) * size)).
		Find(&pts).Error
	var total int64
	session.Model(&data.ProjectTemplate{}).Where("is_system=? and member_code=? and organization_code=?", 0, memId, organizationCode).Count(&total)
	return pts, total, err
}

func (p ProjectTemplateDao) FindProjectTemplateAll(ctx context.Context, organizationCode int64, page int64, size int64) ([]data.ProjectTemplate, int64, error) {
	var pts []data.ProjectTemplate
	session := p.conn.Session(ctx)
	err := session.Model(&data.ProjectTemplate{}).
		Where("organization_code=?", organizationCode).
		Limit(int(size)).
		Offset(int((page - 1) * size)).
		Find(&pts).Error
	var total int64
	session.Model(&data.ProjectTemplate{}).Where("organization_code=?", organizationCode).Count(&total)
	return pts, total, err
}
