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

func TestMySQLControlPointRepository_CRUD_InPkg(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	cfg := &config.Configuration{DbType: "mysql"}
	_ = statements.NewControlPointDataStatements(cfg)

	// repository uses lower-case table/column names (controlpoint)
	selectAll := "SELECT id, name, ipaddress, mac FROM controlpoint"
	selectById := "SELECT id, name, ipaddress, mac FROM controlpoint WHERE id = ?"
	insert := "INSERT INTO controlpoint (name, ipaddress, mac) VALUES (?, ?, ?)"
	update := "UPDATE controlpoint SET name = ?, ipaddress = ?, mac = ? WHERE id = ?"
	delete := "DELETE FROM controlpoint WHERE id = ?"

	rows := sqlmock.NewRows([]string{"id", "name", "ipaddress", "mac"}).AddRow(21, "cp-node", "10.0.0.2", "aa:bb")
	mock.ExpectQuery(regexp.QuoteMeta(selectAll)).WillReturnRows(rows)

	mock.ExpectExec(regexp.QuoteMeta(insert)).WithArgs("cp-1", "10.0.0.3", "aa:11").WillReturnResult(sqlmock.NewResult(31, 1))

	r := NewMySQLControlPointRepository(db)

	list, err := r.FetchAll()
	assert.NoError(t, err)
	assert.Len(t, list, 1)

	v := &models.ControlPoint{Name: "cp-1", IpAddress: "10.0.0.3", Mac: "aa:11"}
	err = r.Create(v)
	assert.NoError(t, err)
	assert.Equal(t, 31, v.Id)

	mock.ExpectQuery(regexp.QuoteMeta(selectById)).WithArgs(31).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "ipaddress", "mac"}).AddRow(31, "cp-1", "10.0.0.3", "aa:11"))
	m, err := r.FetchById(31)
	assert.NoError(t, err)
	assert.Equal(t, 31, m.Id)

	mock.ExpectExec(regexp.QuoteMeta(update)).WithArgs("cp-up", "10.0.0.4", "", 31).WillReturnResult(sqlmock.NewResult(0, 1))
	err = r.Update(&models.ControlPoint{Id: 31, Name: "cp-up", IpAddress: "10.0.0.4"})
	assert.NoError(t, err)

	mock.ExpectExec(regexp.QuoteMeta(delete)).WithArgs(31).WillReturnResult(sqlmock.NewResult(0, 1))
	err = r.Delete(31)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}
