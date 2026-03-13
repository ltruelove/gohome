package repository

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/stretchr/testify/assert"
)

func TestMySQLNodeRepository_CreateNewLog_TxFailure_InPkg(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Begin -> Insert NodeSensorLog succeeds
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO NodeSensorLog (NodeId, DateLogged) VALUES (?, ?)")).WithArgs(1, sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(10, 1))
	// Next insert fails (simulate DB error)
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO TempLog (NodeSensorLogId, TemperatureF, TemperatureC, Humidity) VALUES (?, ?, ?, ?)")).WillReturnError(sqlmock.ErrCancelled)
	// Expect rollback
	mock.ExpectRollback()

	r := NewMySQLNodeRepository(db)
	nd := models.NodeData{NodeId: 1, TemperatureF: 70.0, TemperatureC: 21.0, Humidity: 50.0, Moisture: 0, ResistorValue: 0, IsClosed: false}
	err = r.CreateNewLog(nd)
	assert.Error(t, err)
}
