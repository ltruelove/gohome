package repository

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data/statements"
	"github.com/stretchr/testify/assert"
)

func TestMySQLViewRepository_DeleteAll_InPkg(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	cfg := &config.Configuration{DbType: "mysql"}
	stmt := statements.NewViewDataStatements(cfg)

	mock.ExpectExec(regexp.QuoteMeta(stmt.DeleteAll())).WillReturnResult(sqlmock.NewResult(0, 1))

	rv := NewMySQLViewRepository(db, cfg).(*mysqlViewRepository)
	err = rv.DeleteAll()
	assert.NoError(t, err)
}
