package repository

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data/statements"
	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/stretchr/testify/assert"
)

func TestMySQLViewNodeSwitchRepository_CRUD_InPkg(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	cfg := &config.Configuration{DbType: "mysql"}
	stmt := statements.NewViewNodeSwitchDataStatements(cfg)

	rows := sqlmock.NewRows([]string{"Id", "NodeId", "ViewId", "NodeSwitchId", "Name"}).AddRow(1, 2, 3, 4, "SW")
	mock.ExpectQuery(regexp.QuoteMeta(stmt.SelectAll())).WillReturnRows(rows)

	mock.ExpectQuery(regexp.QuoteMeta(stmt.Insert())).WithArgs(7, 4, 6, "nsv").WillReturnRows(sqlmock.NewRows([]string{"last_insert_id"}).AddRow(77))

	r := NewMySQLViewNodeSwitchRepository(db, cfg)

	list, err := r.SelectAll()
	assert.NoError(t, err)
	assert.Len(t, list, 1)

	v := &models.ViewNodeSwitchData{NodeId: 7, ViewId: 4, NodeSwitchId: 6, Name: "nsv"}
	inserted, err := r.Insert(v)
	assert.NoError(t, err)
	assert.Equal(t, 77, inserted.(*models.ViewNodeSwitchData).Id)

	mock.ExpectQuery(regexp.QuoteMeta(stmt.SelectById())).WithArgs(77).WillReturnRows(sqlmock.NewRows([]string{"Id", "NodeId", "ViewId", "NodeSwitchId", "Name"}).AddRow(77, 7, 4, 6, "nsv"))
	m, err := r.SelectById(77)
	assert.NoError(t, err)
	_ = m.(models.ViewNodeSwitchData)

	mock.ExpectExec(regexp.QuoteMeta(stmt.Update())).WithArgs(7, 4, 6, "nsv-up", 77).WillReturnResult(sqlmock.NewResult(0, 1))
	err = r.Update(&models.ViewNodeSwitchData{Id: 77, NodeId: 7, ViewId: 4, NodeSwitchId: 6, Name: "nsv-up"})
	assert.NoError(t, err)

	mock.ExpectExec(regexp.QuoteMeta(stmt.Delete())).WithArgs(77).WillReturnResult(sqlmock.NewResult(0, 1))
	err = r.Delete(77)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}
