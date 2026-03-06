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

func TestMySQLSwitchTypeRepository_CRUD_InPkg(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	cfg := &config.Configuration{DbType: "mysql"}
	stmt := statements.NewSwitchTypeStatements(cfg)

	rows := sqlmock.NewRows([]string{"Id", "Name"}).AddRow(1, "Toggle")
	mock.ExpectQuery(regexp.QuoteMeta(stmt.SelectAll())).WillReturnRows(rows)

	mock.ExpectQuery(regexp.QuoteMeta(stmt.Insert())).WithArgs("New").WillReturnRows(sqlmock.NewRows([]string{"last_insert_id"}).AddRow(12))

	r := NewMySQLSwitchTypeRepository(db, cfg)

	list, err := r.SelectAll()
	assert.NoError(t, err)
	assert.Len(t, list, 1)

	v := &models.SwitchType{Name: "New"}
	inserted, err := r.Insert(v)
	assert.NoError(t, err)
	assert.Equal(t, 12, inserted.(*models.SwitchType).Id)

	mock.ExpectQuery(regexp.QuoteMeta(stmt.SelectById())).WithArgs(12).WillReturnRows(sqlmock.NewRows([]string{"Id", "Name"}).AddRow(12, "New"))
	m, err := r.SelectById(12)
	assert.NoError(t, err)
	_ = m.(models.SwitchType)

	mock.ExpectExec(regexp.QuoteMeta(stmt.Update())).WithArgs("Updated", 12).WillReturnResult(sqlmock.NewResult(0, 1))
	err = r.Update(&models.SwitchType{Id: 12, Name: "Updated"})
	assert.NoError(t, err)

	mock.ExpectExec(regexp.QuoteMeta(stmt.Delete())).WithArgs(12).WillReturnResult(sqlmock.NewResult(0, 1))
	err = r.Delete(12)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}
