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

func TestMySQLViewRepository_CRUD_InPkg(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	cfg := &config.Configuration{DbType: "mysql"}
	stmt := statements.NewViewDataStatements(cfg)

	rows := sqlmock.NewRows([]string{"Id", "Name"}).AddRow(1, "main")
	mock.ExpectQuery(regexp.QuoteMeta(stmt.SelectAll())).WillReturnRows(rows)

	mock.ExpectQuery(regexp.QuoteMeta(stmt.Insert())).WithArgs("side").WillReturnRows(sqlmock.NewRows([]string{"last_insert_id"}).AddRow(21))

	r := NewMySQLViewRepository(db, cfg)

	list, err := r.SelectAll()
	assert.NoError(t, err)
	assert.Len(t, list, 1)

	v := &models.View{Name: "side"}
	inserted, err := r.Insert(v)
	assert.NoError(t, err)
	assert.Equal(t, 21, inserted.(*models.View).Id)

	mock.ExpectQuery(regexp.QuoteMeta(stmt.SelectById())).WithArgs(21).WillReturnRows(sqlmock.NewRows([]string{"Id", "Name"}).AddRow(21, "side"))
	m, err := r.SelectById(21)
	assert.NoError(t, err)
	_ = m.(models.View)

	mock.ExpectExec(regexp.QuoteMeta(stmt.Update())).WithArgs("side-up", 21).WillReturnResult(sqlmock.NewResult(0, 1))
	err = r.Update(&models.View{Id: 21, Name: "side-up"})
	assert.NoError(t, err)

	mock.ExpectExec(regexp.QuoteMeta(stmt.Delete())).WithArgs(21).WillReturnResult(sqlmock.NewResult(0, 1))
	err = r.Delete(21)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}
