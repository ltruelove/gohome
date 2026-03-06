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

func TestMySQLNodeSensorRepository_SelectAllAndInsert(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	cfg := &config.Configuration{DbType: "mysql"}
	stmt := statements.NewNodeSensorDataStatements(cfg)

	rows := sqlmock.NewRows([]string{"Id", "NodeId", "SensorTypeId", "SensorName", "Pin", "DHTType"}).AddRow(1, 5, 2, "s1", 7, 11)
	mock.ExpectQuery(regexp.QuoteMeta(stmt.SelectAll())).WillReturnRows(rows)

	// Insert returns last id
	mock.ExpectQuery(regexp.QuoteMeta(stmt.Insert())).WithArgs(5, 2, "s2", 8, 22).WillReturnRows(sqlmock.NewRows([]string{"last_insert_id"}).AddRow(55))

	r := repo.NewMySQLNodeSensorRepository(db, cfg)

	list, err := r.SelectAll()
	assert.NoError(t, err)
	assert.Len(t, list, 1)

	ns := &models.NodeSensor{NodeId: 5, SensorTypeId: 2, Name: "s2", Pin: 8, DHTType: 22}
	inserted, err := r.Insert(ns)
	assert.NoError(t, err)
	assert.Equal(t, 55, inserted.(*models.NodeSensor).Id)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestMySQLNodeSensorRepository_SelectByParentAndCRUD(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	cfg := &config.Configuration{DbType: "mysql"}
	stmt := statements.NewNodeSensorDataStatements(cfg)

	rows := sqlmock.NewRows([]string{"Id", "NodeId", "SensorTypeId", "SensorName", "Pin", "DHTType"}).AddRow(2, 6, 3, "s3", 9, 11)
	mock.ExpectQuery(regexp.QuoteMeta(stmt.SelectByParentId())).WithArgs(6).WillReturnRows(rows)

	r := repo.NewMySQLNodeSensorRepository(db, cfg)

	list, err := r.SelectByParentId(6)
	assert.NoError(t, err)
	assert.Len(t, list, 1)

	// SelectById
	single := sqlmock.NewRows([]string{"Id", "NodeId", "SensorTypeId", "SensorName", "Pin", "DHTType"}).AddRow(2, 6, 3, "s3", 9, 11)
	mock.ExpectQuery(regexp.QuoteMeta(stmt.SelectById())).WithArgs(2).WillReturnRows(single)
	m, err := r.SelectById(2)
	assert.NoError(t, err)
	_ = m.(models.NodeSensor)

	// Update
	mock.ExpectExec(regexp.QuoteMeta(stmt.Update())).WithArgs(6, 3, "s3-up", 9, 11, 2).WillReturnResult(sqlmock.NewResult(0, 1))
	err = r.Update(&models.NodeSensor{Id: 2, NodeId: 6, SensorTypeId: 3, Name: "s3-up", Pin: 9, DHTType: 11})
	assert.NoError(t, err)

	// Delete
	mock.ExpectExec(regexp.QuoteMeta(stmt.Delete())).WithArgs(2).WillReturnResult(sqlmock.NewResult(0, 1))
	err = r.Delete(2)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}
