package repository

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/stretchr/testify/assert"
)

func TestMySQLNodeRepository_FetchAll_VerifyUpdateDelete_InPkg(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "mac", "name", "ipaddress"}).AddRow(1, "aa:bb:cc", "n1", "1.2.3.4").AddRow(2, "dd:ee:ff", "n2", "5.6.7.8")
	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, mac, name, ipaddress FROM node")).WillReturnRows(rows)

	mock.ExpectExec(regexp.QuoteMeta("UPDATE node SET mac = ?, name = ?, ipaddress = ? WHERE id = ?")).WithArgs("aa:bb:cc", "n1", "1.2.3.4", 1).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM node WHERE id = ?")).WithArgs(2).WillReturnResult(sqlmock.NewResult(0, 1))

	r := NewMySQLNodeRepository(db)
	list, err := r.FetchAll()
	assert.NoError(t, err)
	assert.Len(t, list, 2)

	// Update
	err = r.Update(&models.Node{Id: 1, Mac: "aa:bb:cc", Name: "n1", IpAddress: "1.2.3.4"})
	assert.NoError(t, err)

	// Delete
	err = r.Delete(2)
	assert.NoError(t, err)
}

func TestMySQLNodeRepository_VerifyIdIsNew_FetchControlPointAndSwitches_InPkg(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// VerifyIdIsNew -> FetchById returns sql.ErrNoRows
	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, mac, name, ipaddress FROM node WHERE id = ?")).WithArgs(99).WillReturnError(sqlmock.ErrCancelled)

	r := NewMySQLNodeRepository(db)
	// Because FetchById returns an error that is not sql.ErrNoRows, VerifyIdIsNew should return false with error
	ok, verr := r.VerifyIdIsNew(99)
	assert.False(t, ok)
	assert.Error(t, verr)

	// FetchControlPointByNode
	cpRow := sqlmock.NewRows([]string{"id", "name", "ipaddress", "mac"}).AddRow(10, "cp1", "127.0.0.1", "CPMAC")
	mock.ExpectQuery(regexp.QuoteMeta("SELECT cp.id, cp.name, cp.ipaddress, cp.mac FROM controlpointnodes AS cpn INNER JOIN controlpoint AS cp ON cp.id = cpn.controlpointid WHERE cpn.nodeid = ?")).WithArgs(1).WillReturnRows(cpRow)
	cp, err := r.FetchControlPointByNode(1)
	assert.NoError(t, err)
	assert.Equal(t, 10, cp.Id)

	// FetchNodeSwitches
	swRows := sqlmock.NewRows([]string{"id", "nodeid", "switchtypeid", "name", "pin", "momentarypressduration", "isclosedon"}).AddRow(5, 1, 2, "s1", 3, 50, 1)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, nodeid, switchtypeid, name, pin, momentarypressduration, isclosedon FROM nodeswitch WHERE nodeid = ?")).WithArgs(1).WillReturnRows(swRows)
	list, err := r.FetchNodeSwitches(1)
	assert.NoError(t, err)
	assert.Len(t, list, 1)
	assert.Equal(t, 5, list[0].Id)

	// FetchNodeSwitch
	swRow := sqlmock.NewRows([]string{"id", "nodeid", "switchtypeid", "name", "pin", "momentarypressduration", "isclosedon"}).AddRow(5, 1, 2, "s1", 3, 50, 1)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, nodeid, switchtypeid, name, pin, momentarypressduration, isclosedon FROM nodeswitch WHERE id = ?")).WithArgs(5).WillReturnRows(swRow)
	ns, err := r.FetchNodeSwitch(5)
	assert.NoError(t, err)
	assert.Equal(t, 5, ns.Id)
}

func TestMySQLNodeRepository_GetSensorLogData_CreateNewLog_InPkg(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Prepare NodeSensorLog row
	now := time.Now()
	nsRows := sqlmock.NewRows([]string{"Id", "NodeId", "DateLogged"}).AddRow(9, 1, now)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT Id, NodeId, DateLogged FROM NodeSensorLog WHERE NodeId = ? AND DateLogged BETWEEN ? AND ? ORDER BY DateLogged DESC")).WithArgs(1, sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnRows(nsRows)

	// temp entries
	tempRows := sqlmock.NewRows([]string{"Id", "NodeSensorLogId", "TemperatureF", "TemperatureC", "Humidity"}).AddRow(1, 9, 72.0, 22.0, 55.0)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT Id, NodeSensorLogId, TemperatureF, TemperatureC, Humidity FROM TempLog WHERE NodeSensorLogId = ?")).WithArgs(9).WillReturnRows(tempRows)

	// moisture
	moistRows := sqlmock.NewRows([]string{"Id", "NodeSensorLogId", "Moisture"}).AddRow(2, 9, 10)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT Id, NodeSensorLogId, Moisture FROM MoistureLog WHERE NodeSensorLogId = ?")).WithArgs(9).WillReturnRows(moistRows)

	// resistor
	resRows := sqlmock.NewRows([]string{"Id", "NodeSensorLogId", "ResistorValue"}).AddRow(3, 9, 123)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT Id, NodeSensorLogId, ResistorValue FROM ResistorLog WHERE NodeSensorLogId = ?")).WithArgs(9).WillReturnRows(resRows)

	// magnetic
	magRows := sqlmock.NewRows([]string{"Id", "NodeSensorLogId", "IsClosed"}).AddRow(4, 9, 1)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT Id, NodeSensorLogId, IsClosed FROM MagneticLog WHERE NodeSensorLogId = ?")).WithArgs(9).WillReturnRows(magRows)

	r := NewMySQLNodeRepository(db)
	logs, err := r.GetSensorLogData(1, now.Add(-time.Hour), now)
	assert.NoError(t, err)
	assert.Len(t, logs, 1)
	assert.Len(t, logs[0].TemperatureEntries, 1)

	// CreateNewLog (transaction with multiple inserts)
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO NodeSensorLog (NodeId, DateLogged) VALUES (?, ?)")).WithArgs(1, sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(55, 1))
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO TempLog (NodeSensorLogId, TemperatureF, TemperatureC, Humidity) VALUES (?, ?, ?, ?)")).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO MoistureLog (NodeSensorLogId, Moisture) VALUES (?, ?)")).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO ResistorLog (NodeSensorLogId, ResistorValue) VALUES (?, ?)")).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO MagneticLog (NodeSensorLogId, IsClosed) VALUES (?, ?)")).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	nd := models.NodeData{NodeId: 1, TemperatureF: 70.0, TemperatureC: 21.0, Humidity: 50.0, Moisture: 10, ResistorValue: 123, IsClosed: true}
	err = r.CreateNewLog(nd)
	assert.NoError(t, err)
}
