package repository_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data/statements"
	"github.com/ltruelove/gohome/internal/app/models"
	repo "github.com/ltruelove/gohome/internal/app/repository"
	"github.com/stretchr/testify/assert"
)

func TestMySQLViewNodeSwitchRepository_CRUD(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	cfg := &config.Configuration{DbType: "mysql"}
	stmt := statements.NewViewNodeSwitchDataStatements(cfg)

	rows := sqlmock.NewRows([]string{"Id", "NodeId", "ViewId", "NodeSwitchId", "Name"}).AddRow(11, 6, 3, 5, "sw-view")
	mock.ExpectQuery(regexp.QuoteMeta(stmt.SelectAll())).WillReturnRows(rows)

	mock.ExpectQuery(regexp.QuoteMeta(stmt.Insert())).WithArgs(7, 4, 6, "new-sw").WillReturnRows(sqlmock.NewRows([]string{"last_insert_id"}).AddRow(99))

	r := repo.NewMySQLViewNodeSwitchRepository(db, cfg)

	list, err := r.SelectAll()
	assert.NoError(t, err)
	assert.Len(t, list, 1)

	v := &models.ViewNodeSwitchData{NodeId: 7, ViewId: 4, NodeSwitchId: 6, Name: "new-sw"}
	inserted, err := r.Insert(v)
	assert.NoError(t, err)
	assert.Equal(t, 99, inserted.(*models.ViewNodeSwitchData).Id)

	// SelectByParentId
	mock.ExpectQuery(regexp.QuoteMeta(stmt.SelectByParentId())).WithArgs(7).WillReturnRows(sqlmock.NewRows([]string{"Id", "NodeId", "ViewId", "NodeSwitchId", "Name"}).AddRow(12, 7, 4, 6, "child-sw"))
	listByParent, err := r.SelectByParentId(7)
	assert.NoError(t, err)
	assert.Len(t, listByParent, 1)

	// SelectById
	mock.ExpectQuery(regexp.QuoteMeta(stmt.SelectById())).WithArgs(99).WillReturnRows(sqlmock.NewRows([]string{"Id", "NodeId", "ViewId", "NodeSwitchId", "Name"}).AddRow(99, 7, 4, 6, "new-sw"))
	m, err := r.SelectById(99)
	assert.NoError(t, err)
	_ = m.(models.ViewNodeSwitchData)

	// Update
	mock.ExpectExec(regexp.QuoteMeta(stmt.Update())).WithArgs(7, 4, 6, "new-sw-up", 99).WillReturnResult(sqlmock.NewResult(0, 1))
	err = r.Update(&models.ViewNodeSwitchData{Id: 99, NodeId: 7, ViewId: 4, NodeSwitchId: 6, Name: "new-sw-up"})
	assert.NoError(t, err)

	// Delete
	mock.ExpectExec(regexp.QuoteMeta(stmt.Delete())).WithArgs(99).WillReturnResult(sqlmock.NewResult(0, 1))
	err = r.Delete(99)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}
