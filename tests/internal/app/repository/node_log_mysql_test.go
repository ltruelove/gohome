package repository_test

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ltruelove/gohome/internal/app/models"
	repo "github.com/ltruelove/gohome/internal/app/repository"
	"github.com/stretchr/testify/assert"
)

func TestMySQLNodeRepository_CreateNewLog(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Expect transaction begin
	mock.ExpectBegin()
	// Expect insert into NodeSensorLog returning id
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO NodeSensorLog (NodeId, DateLogged) VALUES (?, ?)")).WithArgs(5, sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1001, 1))
	// Expect temp insert
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO TempLog (NodeSensorLogId, TemperatureF, TemperatureC, Humidity) VALUES (?, ?, ?, ?)")).WithArgs(1001, 72.0, 22.2, 45.0).WillReturnResult(sqlmock.NewResult(1, 1))
	// Expect moisture insert
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO MoistureLog (NodeSensorLogId, Moisture) VALUES (?, ?)")).WithArgs(1001, 55.5).WillReturnResult(sqlmock.NewResult(1, 1))
	// Expect resistor insert
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO ResistorLog (NodeSensorLogId, ResistorValue) VALUES (?, ?)")).WithArgs(1001, 330).WillReturnResult(sqlmock.NewResult(1, 1))
	// Expect magnetic insert (IsClosed false => 0)
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO MagneticLog (NodeSensorLogId, IsClosed) VALUES (?, ?)")).WithArgs(1001, 0).WillReturnResult(sqlmock.NewResult(1, 1))
	// Expect commit
	mock.ExpectCommit()

	r := repo.NewMySQLNodeRepository(db)

	item := models.NodeData{
		NodeId:        5,
		TemperatureF:  72.0,
		TemperatureC:  22.2,
		Humidity:      45.0,
		Moisture:      55,
		ResistorValue: 330,
		IsClosed:      false,
	}

	err = r.CreateNewLog(item)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestMySQLNodeRepository_GetSensorLogData_WithChildren(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	start := time.Now().Add(-time.Hour)
	end := time.Now()

	// NodeSensorLog query returns one log id
	rows := sqlmock.NewRows([]string{"Id", "NodeId", "DateLogged"}).AddRow(2001, 5, time.Now())
	mock.ExpectQuery(regexp.QuoteMeta("SELECT Id, NodeId, DateLogged FROM NodeSensorLog WHERE NodeId = ? AND DateLogged BETWEEN ? AND ? ORDER BY DateLogged DESC")).WithArgs(5, start, end).WillReturnRows(rows)

	// TempLog rows
	tempRows := sqlmock.NewRows([]string{"Id", "NodeSensorLogId", "TemperatureF", "TemperatureC", "Humidity"}).AddRow(10, 2001, 70.0, 21.1, 40.0)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT Id, NodeSensorLogId, TemperatureF, TemperatureC, Humidity FROM TempLog WHERE NodeSensorLogId = ?")).WithArgs(2001).WillReturnRows(tempRows)

	// Moisture
	moistRows := sqlmock.NewRows([]string{"Id", "NodeSensorLogId", "Moisture"}).AddRow(11, 2001, 50.5)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT Id, NodeSensorLogId, Moisture FROM MoistureLog WHERE NodeSensorLogId = ?")).WithArgs(2001).WillReturnRows(moistRows)

	// Resistor
	resRows := sqlmock.NewRows([]string{"Id", "NodeSensorLogId", "ResistorValue"}).AddRow(12, 2001, 400)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT Id, NodeSensorLogId, ResistorValue FROM ResistorLog WHERE NodeSensorLogId = ?")).WithArgs(2001).WillReturnRows(resRows)

	// Magnetic
	magRows := sqlmock.NewRows([]string{"Id", "NodeSensorLogId", "IsClosed"}).AddRow(13, 2001, 1)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT Id, NodeSensorLogId, IsClosed FROM MagneticLog WHERE NodeSensorLogId = ?")).WithArgs(2001).WillReturnRows(magRows)

	r := repo.NewMySQLNodeRepository(db)

	logs, err := r.GetSensorLogData(5, start, end)
	assert.NoError(t, err)
	assert.Len(t, logs, 1)
	logEntry := logs[0]
	assert.Equal(t, 2001, logEntry.Id)
	assert.Len(t, logEntry.TemperatureEntries, 1)
	assert.Len(t, logEntry.MoistureEntries, 1)
	assert.Len(t, logEntry.ResistorEntries, 1)
	assert.Len(t, logEntry.MagneticEntries, 1)
	assert.NoError(t, mock.ExpectationsWereMet())
}
