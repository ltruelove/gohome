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

func TestMySQLNodeSwitchRepository_CRUD_InPkg(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	cfg := &config.Configuration{DbType: "mysql"}
	stmt := statements.NewNodeSwitchDataStatements(cfg)

	rows := sqlmock.NewRows([]string{"Id", "NodeId", "SwitchTypeId", "Name", "Pin", "MomentaryPressDuration", "IsClosedOn"}).AddRow(10, 5, 2, "switch1", 7, 0, 1)
	mock.ExpectQuery(regexp.QuoteMeta(stmt.SelectAll())).WillReturnRows(rows)

	mock.ExpectQuery(regexp.QuoteMeta(stmt.Insert())).WithArgs(5, 2, "ns", 9, 0, true).WillReturnRows(sqlmock.NewRows([]string{"last_insert_id"}).AddRow(66))

	r := NewMySQLNodeSwitchRepository(db, cfg)

	list, err := r.SelectAll()
	assert.NoError(t, err)
	assert.Len(t, list, 1)

	v := &models.NodeSwitch{NodeId: 5, SwitchTypeId: 2, Name: "ns", Pin: 9, MomentaryPressDuration: 0, IsClosedOn: true}
	inserted, err := r.Insert(v)
	assert.NoError(t, err)
	assert.Equal(t, 66, inserted.(*models.NodeSwitch).Id)

	mock.ExpectQuery(regexp.QuoteMeta(stmt.SelectById())).WithArgs(66).WillReturnRows(sqlmock.NewRows([]string{"Id", "NodeId", "SwitchTypeId", "Name", "Pin", "MomentaryPressDuration", "IsClosedOn"}).AddRow(66, 5, 2, "ns", 9, 0, 1))
	m, err := r.SelectById(66)
	assert.NoError(t, err)
	_ = m.(models.NodeSwitch)

	mock.ExpectExec(regexp.QuoteMeta(stmt.Update())).WithArgs(5, 2, "ns-up", 9, 0, false, 66).WillReturnResult(sqlmock.NewResult(0, 1))
	err = r.Update(&models.NodeSwitch{Id: 66, NodeId: 5, SwitchTypeId: 2, Name: "ns-up", Pin: 9, MomentaryPressDuration: 0, IsClosedOn: false})
	assert.NoError(t, err)

	mock.ExpectExec(regexp.QuoteMeta(stmt.Delete())).WithArgs(66).WillReturnResult(sqlmock.NewResult(0, 1))
	err = r.Delete(66)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}
