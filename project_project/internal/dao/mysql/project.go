package mysql

import (
	"context"
	"fmt"
	"test.com/project_project/internal/data/project"
	"test.com/project_project/internal/database"
	"test.com/project_project/internal/database/gorms"
	"time"
)

type ProjectDao struct {
	conn *gorms.GormConn
}

func NewProjectDao() *ProjectDao {
	return &ProjectDao{
		conn: gorms.New(),
	}
}

func (p *ProjectDao) FindProjectByMemId(ctx context.Context, memId int64, condition string, page int64, size int64) ([]*project.ProjectAndMember, int64, error) {
	var pms []*project.ProjectAndMember
	session := p.conn.Session(ctx)
	index := (page - 1) * size
	sql := fmt.Sprintf("select * from ms_project a, ms_project_member b where a.id = b.project_code and b.member_code=? %s order by sort limit ?,?", condition)
	raw := session.Raw(sql, memId, index, size)
	raw.Scan(&pms)
	var total int64
	query := fmt.Sprintf("select count(*) from ms_project a, ms_project_member b where a.id = b.project_code and b.member_code=? %s", condition)
	tx := session.Raw(query, memId)
	err := tx.Scan(&total).Error
	return pms, total, err
}

func (p *ProjectDao) FindCollectProjectByMemId(ctx context.Context, memId int64, page int64, size int64) ([]*project.ProjectAndMember, int64, error) {
	session := p.conn.Session(ctx)
	index := (page - 1) * size
	sql := fmt.Sprintf("select * from ms_project where id in (select project_code from ms_project_collection where member_code=? ) order by sort limit ?,?")
	db := session.Raw(sql, memId, index, size)
	var mp []*project.ProjectAndMember
	err := db.Scan(&mp).Error
	var total int64
	query := fmt.Sprintf("member_code=?")
	session.Model(&project.ProjectCollection{}).Where(query, memId).Count(&total)
	return mp, total, err
}

func (p *ProjectDao) FindProjectByPIdAndMemId(ctx context.Context, projectCode int64, memberId int64) (*project.ProjectAndMember, error) {
	var pms *project.ProjectAndMember
	session := p.conn.Session(ctx)
	sql := fmt.Sprintf("select * from ms_project a, ms_project_member b where a.id = b.project_code and b.member_code=? and b.project_code =? limit 1")
	raw := session.Raw(sql, memberId, projectCode)
	err := raw.Scan(&pms).Error
	return pms, err
}

func (p *ProjectDao) FindCollectByPIdAndMemId(ctx context.Context, projectCode int64, memberId int64) (bool, error) {
	var count int64
	session := p.conn.Session(ctx)
	sql := fmt.Sprintf("select count(*) from ms_project_collection where member_code=? and project_code=?")
	raw := session.Raw(sql, memberId, projectCode)
	err := raw.Scan(&count).Error
	return count > 0, err
}

func (p *ProjectDao) FindProjectByCipId(ctx context.Context, cipherIdCode int64) (int64, error) {
	var pms *project.ProjectAndMember
	session := p.conn.Session(ctx)
	sql := fmt.Sprintf("select * from ms_project_member where id=? limit 1")
	raw := session.Raw(sql, cipherIdCode)
	err := raw.Scan(&pms).Error
	return pms.ProjectCode, err
}

func (p *ProjectDao) SaveProject(conn database.DbConn, ctx context.Context, pr *project.Project) error {
	//使用事务	需要使用同一连接 不然没法做事务操作
	p.conn = conn.(*gorms.GormConn)
	//return p.conn.Tx(ctx).Save(&pr).Error
	return p.conn.Tx(ctx).Create(&pr).Error
}

func (p *ProjectDao) SaveProjectMember(conn database.DbConn, ctx context.Context, pm *project.ProjectMember) error {
	//使用事务
	p.conn = conn.(*gorms.GormConn)
	//return p.conn.Tx(ctx).Save(&pm).Error
	return p.conn.Tx(ctx).Create(&pm).Error
}

func (p *ProjectDao) DeleteProject(ctx context.Context, id int64) error {
	//err := p.conn.Session(ctx).Model(&project.Project{}).Where("id=?", id).Update("deleted", 1).Error
	err := p.conn.Session(ctx).Model(&project.Project{}).Where("id=?", id).Updates(map[string]interface{}{"deleted": 1, "deleted_time": time.Now().UnixMilli()}).Error
	return err
}

func (p *ProjectDao) RecoveryProject(ctx context.Context, id int64) error {
	err := p.conn.Session(ctx).Model(&project.Project{}).Where("id=?", id).Updates(map[string]interface{}{"deleted": 0, "deleted_time": 0}).Error
	return err
}
