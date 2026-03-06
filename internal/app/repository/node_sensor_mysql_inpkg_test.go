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

func TestMySQLNodeSensorRepository_CRUD_InPkg(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	cfg := &config.Configuration{DbType: "mysql"}
	stmt := statements.NewNodeSensorDataStatements(cfg)

	rows := sqlmock.NewRows([]string{"Id", "NodeId", "SensorTypeId", "Name", "Pin", "DHTType"}).AddRow(1, 2, 3, "S", 4, 11)
	mock.ExpectQuery(regexp.QuoteMeta(stmt.SelectAll())).WillReturnRows(rows)

	mock.ExpectQuery(regexp.QuoteMeta(stmt.Insert())).WithArgs(2, 3, "n", 5, 22).WillReturnRows(sqlmock.NewRows([]string{"last_insert_id"}).AddRow(55))

	r := NewMySQLNodeSensorRepository(db, cfg)

	list, err := r.SelectAll()
	assert.NoError(t, err)
	assert.Len(t, list, 1)

	v := &models.NodeSensor{NodeId: 2, SensorTypeId: 3, Name: "n", Pin: 5, DHTType: 22}
	inserted, err := r.Insert(v)
	assert.NoError(t, err)
	assert.Equal(t, 55, inserted.(*models.NodeSensor).Id)

	mock.ExpectQuery(regexp.QuoteMeta(stmt.SelectById())).WithArgs(55).WillReturnRows(sqlmock.NewRows([]string{"Id", "NodeId", "SensorTypeId", "Name", "Pin", "DHTType"}).AddRow(55, 2, 3, "n", 5, 22))
	m, err := r.SelectById(55)
	assert.NoError(t, err)
	_ = m.(models.NodeSensor)

	mock.ExpectExec(regexp.QuoteMeta(stmt.Update())).WithArgs(2, 3, "n-up", 5, 22, 55).WillReturnResult(sqlmock.NewResult(0, 1))
	err = r.Update(&models.NodeSensor{Id: 55, NodeId: 2, SensorTypeId: 3, Name: "n-up", Pin: 5, DHTType: 22})
	assert.NoError(t, err)

	mock.ExpectExec(regexp.QuoteMeta(stmt.Delete())).WithArgs(55).WillReturnResult(sqlmock.NewResult(0, 1))
	err = r.Delete(55)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}
