package mysql

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"test.com/project_project/internal/data"
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

func (p *ProjectDao) FindProjectByMemId(ctx context.Context, memId int64, condition string, page int64, size int64) ([]*data.ProjectAndMember, int64, error) {
	var pms []*data.ProjectAndMember
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

func (p *ProjectDao) FindCollectProjectByMemId(ctx context.Context, memId int64, page int64, size int64) ([]*data.ProjectAndMember, int64, error) {
	session := p.conn.Session(ctx)
	index := (page - 1) * size
	sql := fmt.Sprintf("select * from ms_project where id in (select project_code from ms_project_collection where member_code=? ) order by sort limit ?,?")
	db := session.Raw(sql, memId, index, size)
	var mp []*data.ProjectAndMember
	err := db.Scan(&mp).Error
	var total int64
	query := fmt.Sprintf("member_code=?")
	session.Model(&data.ProjectCollection{}).Where(query, memId).Count(&total)
	return mp, total, err
}

func (p *ProjectDao) FindProjectByPIdAndMemId(ctx context.Context, projectCode int64, memberId int64) (*data.ProjectAndMember, error) {
	var pms *data.ProjectAndMember
	session := p.conn.Session(ctx)
	sql := fmt.Sprintf("select a.*,b.project_code,b.member_code,b.join_time,b.is_owner,b.authorize from ms_project a, ms_project_member b where a.id = b.project_code and b.member_code=? and b.project_code =? limit 1")
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
	var pms *data.ProjectAndMember
	session := p.conn.Session(ctx)
	sql := fmt.Sprintf("select * from ms_project_member where id=? limit 1")
	raw := session.Raw(sql, cipherIdCode)
	err := raw.Scan(&pms).Error
	return pms.ProjectCode, err
}

func (p *ProjectDao) SaveProject(conn database.DbConn, ctx context.Context, pr *data.Project) error {
	//使用事务	需要使用同一连接 不然没法做事务操作
	p.conn = conn.(*gorms.GormConn)
	//return p.conn.Tx(ctx).Save(&pr).Error
	return p.conn.Tx(ctx).Create(&pr).Error
}

func (p *ProjectDao) SaveProjectMember(conn database.DbConn, ctx context.Context, pm *data.ProjectMember) error {
	//使用事务
	p.conn = conn.(*gorms.GormConn)
	//return p.conn.Tx(ctx).Save(&pm).Error
	return p.conn.Tx(ctx).Create(&pm).Error
}

func (p *ProjectDao) DeleteProject(ctx context.Context, id int64) error {
	//err := p.conn.Session(ctx).Model(&project.Project{}).Where("id=?", id).Update("deleted", 1).Error
	err := p.conn.Session(ctx).Model(&data.Project{}).Where("id=?", id).Updates(map[string]interface{}{"deleted": 1, "deleted_time": time.Now().UnixMilli()}).Error
	return err
}

func (p *ProjectDao) RecoveryProject(ctx context.Context, id int64) error {
	err := p.conn.Session(ctx).Model(&data.Project{}).Where("id=?", id).Updates(map[string]interface{}{"deleted": 0, "deleted_time": 0}).Error
	return err
}

func (p *ProjectDao) CollectProject(ctx context.Context, pc *data.ProjectCollection) error {
	return p.conn.Session(ctx).Save(&pc).Error
}

func (p *ProjectDao) CancelCollectProject(ctx context.Context, projectCode int64, memberId int64) error {
	return p.conn.Session(ctx).Where("project_code = ? and member_code = ?", projectCode, memberId).Delete(data.ProjectCollection{}).Error
}

func (p *ProjectDao) UpdateProject(ctx context.Context, proj *data.Project) error {
	return p.conn.Session(ctx).Updates(&proj).Error
}

func (p *ProjectDao) FindProjectByPid(ctx context.Context, projectCode int64) (list []*data.ProjectMember, total int64, err error) {
	session := p.conn.Session(ctx)
	err = session.Model(&data.ProjectMember{}).Where("project_code = ?", projectCode).Find(&list).Error
	err = session.Model(&data.ProjectMember{}).Where("project_code = ?", projectCode).Count(&total).Error
	return
}

func (p *ProjectDao) FindProjectById(ctx context.Context, projectCode int64) (pj *data.Project, err error) {
	session := p.conn.Session(ctx)
	err = session.Where("id = ?", projectCode).Find(&pj).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	return
}

func (p *ProjectDao) FindProjectByIds(ctx context.Context, pids []int64) (list []*data.Project, err error) {
	session := p.conn.Session(ctx)
	err = session.Model(&data.Project{}).Where("id in (?)", pids).Find(&list).Error
	return
}
