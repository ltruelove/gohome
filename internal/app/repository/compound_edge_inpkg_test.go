package repository

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data/statements"
	"github.com/stretchr/testify/assert"
)

func TestMySQLCompoundRepository_EmptyAndErrorPaths_InPkg(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	cfg := &config.Configuration{DbType: "mysql"}
	stmt := statements.NewCompoundStatements(cfg)

	// empty results
	mock.ExpectQuery(regexp.QuoteMeta(stmt.SelectViewNodeSensorDataByViewId())).WithArgs(5).WillReturnRows(sqlmock.NewRows([]string{"Id"}))
	mock.ExpectQuery(regexp.QuoteMeta(stmt.SelectViewNodeSwitchDataByViewId())).WithArgs(5).WillReturnRows(sqlmock.NewRows([]string{"Id"}))

	r := NewMySQLCompoundRepository(db, cfg)
	s, err := r.FetchViewNodeSensorDataByViewId(5)
	assert.NoError(t, err)
	assert.Len(t, s, 0)
	sw, err := r.FetchViewNodeSwitchDataByViewId(5)
	assert.NoError(t, err)
	assert.Len(t, sw, 0)

	// error path
	mock.ExpectQuery(regexp.QuoteMeta(stmt.SelectViewNodeSensorDataByViewId())).WithArgs(6).WillReturnError(sqlmock.ErrCancelled)
	_, err = r.FetchViewNodeSensorDataByViewId(6)
	assert.Error(t, err)
}
