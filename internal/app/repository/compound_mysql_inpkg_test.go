package repository

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data/statements"
	"github.com/ltruelove/gohome/internal/app/viewModels"
	"github.com/stretchr/testify/assert"
)

func TestMySQLCompoundRepository_FetchViewNodeSensorAndSwitchData_InPkg(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	cfg := &config.Configuration{DbType: "mysql"}
	stmt := statements.NewCompoundStatements(cfg)

	rows := sqlmock.NewRows([]string{"Id", "NodeId", "ViewId", "NodeSensorId", "Name", "NodeName", "SensorName", "SensorTypeName"}).AddRow(1, 2, 3, 4, "S", "NodeA", "SensorA", "Temp")
	mock.ExpectQuery(regexp.QuoteMeta(stmt.SelectViewNodeSensorDataByViewId())).WithArgs(3).WillReturnRows(rows)

	rows2 := sqlmock.NewRows([]string{"Id", "NodeId", "ViewId", "NodeSwitchId", "Name", "NodeName", "SwitchName", "SwitchTypeName"}).AddRow(2, 5, 3, 6, "SW", "NodeB", "SwitchA", "Toggle")
	mock.ExpectQuery(regexp.QuoteMeta(stmt.SelectViewNodeSwitchDataByViewId())).WithArgs(3).WillReturnRows(rows2)

	r := NewMySQLCompoundRepository(db, cfg)

	sensors, err := r.FetchViewNodeSensorDataByViewId(3)
	assert.NoError(t, err)
	assert.Len(t, sensors, 1)

	switches, err := r.FetchViewNodeSwitchDataByViewId(3)
	assert.NoError(t, err)
	assert.Len(t, switches, 1)

	_ = viewModels.ViewNodeSensorVM{}
	_ = viewModels.ViewNodeSwitchVM{}

	assert.NoError(t, mock.ExpectationsWereMet())
}
