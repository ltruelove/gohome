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

func TestMySQLNodeSwitchRepository_SelectAllAndInsert(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	cfg := &config.Configuration{DbType: "mysql"}
	stmt := statements.NewNodeSwitchDataStatements(cfg)

	rows := sqlmock.NewRows([]string{"Id", "NodeId", "SwitchTypeId", "Name", "Pin", "MomentaryPressDuration", "IsClosedOn"}).AddRow(10, 5, 2, "sw1", 7, 0, true)
	mock.ExpectQuery(regexp.QuoteMeta(stmt.SelectAll())).WillReturnRows(rows)

	mock.ExpectQuery(regexp.QuoteMeta(stmt.Insert())).WithArgs(6, 3, "sw2", 8, 100, false).WillReturnRows(sqlmock.NewRows([]string{"last_insert_id"}).AddRow(77))

	r := repo.NewMySQLNodeSwitchRepository(db, cfg)

	list, err := r.SelectAll()
	assert.NoError(t, err)
	assert.Len(t, list, 1)

	ns := &models.NodeSwitch{NodeId: 6, SwitchTypeId: 3, Name: "sw2", Pin: 8, MomentaryPressDuration: 100, IsClosedOn: false}
	inserted, err := r.Insert(ns)
	assert.NoError(t, err)
	assert.Equal(t, 77, inserted.(*models.NodeSwitch).Id)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestMySQLNodeSwitchRepository_SelectByParentAndCRUD(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	cfg := &config.Configuration{DbType: "mysql"}
	stmt := statements.NewNodeSwitchDataStatements(cfg)

	rows := sqlmock.NewRows([]string{"Id", "NodeId", "SwitchTypeId", "Name", "Pin", "MomentaryPressDuration", "IsClosedOn"}).AddRow(11, 6, 3, "sw3", 9, 50, false)
	mock.ExpectQuery(regexp.QuoteMeta(stmt.SelectByParentId())).WithArgs(6).WillReturnRows(rows)

	r := repo.NewMySQLNodeSwitchRepository(db, cfg)

	list, err := r.SelectByParentId(6)
	assert.NoError(t, err)
	assert.Len(t, list, 1)

	// SelectById
	single := sqlmock.NewRows([]string{"Id", "NodeId", "SwitchTypeId", "Name", "Pin", "MomentaryPressDuration", "IsClosedOn"}).AddRow(11, 6, 3, "sw3", 9, 50, false)
	mock.ExpectQuery(regexp.QuoteMeta(stmt.SelectById())).WithArgs(11).WillReturnRows(single)
	m, err := r.SelectById(11)
	assert.NoError(t, err)
	_ = m.(models.NodeSwitch)

	// Update
	mock.ExpectExec(regexp.QuoteMeta(stmt.Update())).WithArgs(6, 3, "sw3-up", 9, 60, true, 11).WillReturnResult(sqlmock.NewResult(0, 1))
	err = r.Update(&models.NodeSwitch{Id: 11, NodeId: 6, SwitchTypeId: 3, Name: "sw3-up", Pin: 9, MomentaryPressDuration: 60, IsClosedOn: true})
	assert.NoError(t, err)

	// Delete
	mock.ExpectExec(regexp.QuoteMeta(stmt.Delete())).WithArgs(11).WillReturnResult(sqlmock.NewResult(0, 1))
	err = r.Delete(11)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}
