package repository

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data/statements"
	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/stretchr/testify/assert"
)

func TestMySQLSensorTypeRepository_CRUD_InPkg(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	cfg := &config.Configuration{DbType: "mysql"}
	stmt := statements.NewSensorTypeStatements(cfg)
	rows := sqlmock.NewRows([]string{"Id", "TypeName"}).AddRow(1, "Temp")
	mock.ExpectQuery(regexp.QuoteMeta(stmt.SelectAll())).WillReturnRows(rows)
	mock.ExpectQuery(regexp.QuoteMeta(stmt.Insert())).WithArgs("New").WillReturnRows(sqlmock.NewRows([]string{"last_insert_id"}).AddRow(11))
	r := NewMySQLSensorTypeRepository(db, cfg)
	list, err := r.SelectAll()
	assert.NoError(t, err)
	assert.Len(t, list, 1)
	v := &models.SensorType{TypeName: "New"}
	inserted, err := r.Insert(v)
	assert.NoError(t, err)
	assert.Equal(t, 11, inserted.(*models.SensorType).Id)
	mock.ExpectQuery(regexp.QuoteMeta(stmt.SelectById())).WithArgs(11).WillReturnRows(sqlmock.NewRows([]string{"Id", "TypeName"}).AddRow(11, "New"))
	m, err := r.SelectById(11)
	assert.NoError(t, err)
	_ = m.(models.SensorType)
	mock.ExpectExec(regexp.QuoteMeta(stmt.Update())).WithArgs("Updated", 11).WillReturnResult(sqlmock.NewResult(0, 1))
	err = r.Update(&models.SensorType{Id: 11, TypeName: "Updated"})
	assert.NoError(t, err)
	mock.ExpectExec(regexp.QuoteMeta(stmt.Delete())).WithArgs(11).WillReturnResult(sqlmock.NewResult(0, 1))
	err = r.Delete(11)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
func TestMySQLSensorTypeRepository_SelectAll_EmptyResultSet(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	cfg := &config.Configuration{DbType: "mysql"}
	stmt := statements.NewSensorTypeStatements(cfg)
	mock.ExpectQuery(regexp.QuoteMeta(stmt.SelectAll())).WillReturnRows(sqlmock.NewRows([]string{"Id", "TypeName"}))
	r := NewMySQLSensorTypeRepository(db, cfg)
	list, err := r.SelectAll()
	assert.NoError(t, err)
	assert.Len(t, list, 0)
	assert.NoError(t, mock.ExpectationsWereMet())
}
func TestMySQLSensorTypeRepository_Insert_InvalidInput(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	cfg := &config.Configuration{DbType: "mysql"}
	stmt := statements.NewSensorTypeStatements(cfg)
	mock.ExpectQuery(regexp.QuoteMeta(stmt.Insert())).WithArgs("").WillReturnError(sql.ErrNoRows)
	r := NewMySQLSensorTypeRepository(db, cfg)
	v := &models.SensorType{TypeName: ""}
	inserted, err := r.Insert(v)
	assert.Error(t, err)
	assert.Nil(t, inserted)
	assert.NoError(t, mock.ExpectationsWereMet())
}
func TestMySQLSensorTypeRepository_SelectById_InvalidID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	cfg := &config.Configuration{DbType: "mysql"}
	stmt := statements.NewSensorTypeStatements(cfg)
	mock.ExpectQuery(regexp.QuoteMeta(stmt.SelectById())).WithArgs(0).WillReturnError(sql.ErrNoRows)
	r := NewMySQLSensorTypeRepository(db, cfg)
	m, err := r.SelectById(0)
	assert.Error(t, err)
	assert.Nil(t, m)
	assert.NoError(t, mock.ExpectationsWereMet())
}
