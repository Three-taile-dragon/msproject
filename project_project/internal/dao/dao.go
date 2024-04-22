package dao

import (
	"errors"
	"test.com/project_common/errs"
	"test.com/project_project/internal/database"
	"test.com/project_project/internal/database/gorms"
)

//事务的具体实现

type TransactionTmpl struct {
	conn database.DbConn
}

func (t *TransactionTmpl) Action(f func(conn database.DbConn) error) error {
	t.conn.Begin()
	err := f(t.conn)

	var bErr *errs.BError
	if errors.Is(err, bErr) {
		errors.As(err, &bErr)
		if bErr != nil {
			t.conn.Rollback() //事务回滚
			return bErr
		} else {
			t.conn.Commit() //事务提交
			return nil
		}
	}

	if err != nil {
		t.conn.Rollback() //事务回滚
		return err
	}
	t.conn.Commit() //事务提交
	return nil
}

func NewTransaction() *TransactionTmpl {
	return &TransactionTmpl{
		conn: gorms.NewTran(),
	}
}
