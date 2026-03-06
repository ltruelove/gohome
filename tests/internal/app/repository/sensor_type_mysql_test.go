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

func TestMySQLSensorTypeRepository_CRUD(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	cfg := &config.Configuration{DbType: "mysql"}
	stmt := statements.NewSensorTypeStatements(cfg)

	rows := sqlmock.NewRows([]string{"Id", "TypeName"}).AddRow(1, "temperature")
	mock.ExpectQuery(regexp.QuoteMeta(stmt.SelectAll())).WillReturnRows(rows)

	mock.ExpectQuery(regexp.QuoteMeta(stmt.Insert())).WithArgs("humidity").WillReturnRows(sqlmock.NewRows([]string{"last_insert_id"}).AddRow(42))

	r := repo.NewMySQLSensorTypeRepository(db, cfg)

	list, err := r.SelectAll()
	assert.NoError(t, err)
	assert.Len(t, list, 1)

	s := &models.SensorType{TypeName: "humidity"}
	inserted, err := r.Insert(s)
	assert.NoError(t, err)
	assert.Equal(t, 42, inserted.(*models.SensorType).Id)

	// SelectById
	single := sqlmock.NewRows([]string{"Id", "TypeName"}).AddRow(42, "humidity")
	mock.ExpectQuery(regexp.QuoteMeta(stmt.SelectById())).WithArgs(42).WillReturnRows(single)
	m, err := r.SelectById(42)
	assert.NoError(t, err)
	_ = m.(models.SensorType)

	// Update
	mock.ExpectExec(regexp.QuoteMeta(stmt.Update())).WithArgs("humidity-up", 42).WillReturnResult(sqlmock.NewResult(0, 1))
	err = r.Update(&models.SensorType{Id: 42, TypeName: "humidity-up"})
	assert.NoError(t, err)

	// Delete
	mock.ExpectExec(regexp.QuoteMeta(stmt.Delete())).WithArgs(42).WillReturnResult(sqlmock.NewResult(0, 1))
	err = r.Delete(42)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}
