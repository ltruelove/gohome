package repository_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data/statements"
	repo "github.com/ltruelove/gohome/internal/app/repository"
	"github.com/stretchr/testify/assert"
)

func TestMySQLCompoundRepository_FetchViewNodeSensorAndSwitchData(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	cfg := &config.Configuration{DbType: "mysql"}
	stmt := statements.NewCompoundStatements(cfg)

	sensorRows := sqlmock.NewRows([]string{"Id", "NodeId", "ViewId", "NodeSensorId", "Name", "NodeName", "SensorName", "SensorTypeName"}).AddRow(1, 2, 3, 4, "sname", "nodeA", "temp", "dht22")
	mock.ExpectQuery(regexp.QuoteMeta(stmt.SelectViewNodeSensorDataByViewId())).WithArgs(3).WillReturnRows(sensorRows)

	switchRows := sqlmock.NewRows([]string{"Id", "NodeId", "ViewId", "NodeSwitchId", "Name", "NodeName", "SwitchName", "SwitchTypeName"}).AddRow(2, 5, 3, 6, "swname", "nodeB", "relay", "typeA")
	mock.ExpectQuery(regexp.QuoteMeta(stmt.SelectViewNodeSwitchDataByViewId())).WithArgs(3).WillReturnRows(switchRows)

	r := repo.NewMySQLCompoundRepository(db, cfg)

	sensors, err := r.FetchViewNodeSensorDataByViewId(3)
	assert.NoError(t, err)
	assert.Len(t, sensors, 1)
	assert.Equal(t, "sname", sensors[0].Name)

	switches, err := r.FetchViewNodeSwitchDataByViewId(3)
	assert.NoError(t, err)
	assert.Len(t, switches, 1)
	assert.Equal(t, "swname", switches[0].Name)

	assert.NoError(t, mock.ExpectationsWereMet())
}
