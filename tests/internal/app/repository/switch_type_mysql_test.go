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

func TestMySQLSwitchTypeRepository_CRUD(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	cfg := &config.Configuration{DbType: "mysql"}
	stmt := statements.NewSwitchTypeStatements(cfg)

	rows := sqlmock.NewRows([]string{"Id", "Name"}).AddRow(1, "relay")
	mock.ExpectQuery(regexp.QuoteMeta(stmt.SelectAll())).WillReturnRows(rows)

	mock.ExpectQuery(regexp.QuoteMeta(stmt.Insert())).WithArgs("solid-state").WillReturnRows(sqlmock.NewRows([]string{"last_insert_id"}).AddRow(88))

	r := repo.NewMySQLSwitchTypeRepository(db, cfg)

	list, err := r.SelectAll()
	assert.NoError(t, err)
	assert.Len(t, list, 1)

	s := &models.SwitchType{Name: "solid-state"}
	inserted, err := r.Insert(s)
	assert.NoError(t, err)
	assert.Equal(t, 88, inserted.(*models.SwitchType).Id)

	// SelectById
	mock.ExpectQuery(regexp.QuoteMeta(stmt.SelectById())).WithArgs(88).WillReturnRows(sqlmock.NewRows([]string{"Id", "Name"}).AddRow(88, "solid-state"))
	m, err := r.SelectById(88)
	assert.NoError(t, err)
	_ = m.(models.SwitchType)

	// Update
	mock.ExpectExec(regexp.QuoteMeta(stmt.Update())).WithArgs("solid-up", 88).WillReturnResult(sqlmock.NewResult(0, 1))
	err = r.Update(&models.SwitchType{Id: 88, Name: "solid-up"})
	assert.NoError(t, err)

	// Delete
	mock.ExpectExec(regexp.QuoteMeta(stmt.Delete())).WithArgs(88).WillReturnResult(sqlmock.NewResult(0, 1))
	err = r.Delete(88)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}
