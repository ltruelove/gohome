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

func TestDeleteOperations_Batch(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	cfg := &config.Configuration{DbType: "mysql"}

	// ViewNodeSensor: DeleteAll + DeleteByParentId + DeleteBySecondParentId
	vnsStmt := statements.NewViewNodeSensorDataStatements(cfg)
	mock.ExpectExec(regexp.QuoteMeta(vnsStmt.DeleteAll())).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(regexp.QuoteMeta(vnsStmt.DeleteByParentId())).WithArgs(7).WillReturnResult(sqlmock.NewResult(0, 2))
	mock.ExpectExec(regexp.QuoteMeta(vnsStmt.DeleteBySecondParentId())).WithArgs(8).WillReturnResult(sqlmock.NewResult(0, 1))

	rVns := repo.NewMySQLViewNodeSensorRepository(db, cfg)
	err = rVns.DeleteAll()
	assert.NoError(t, err)
	err = rVns.DeleteByParentId(7)
	assert.NoError(t, err)
	err = rVns.DeleteBySecondParentId(8)
	assert.NoError(t, err)

	// ViewNodeSwitch: DeleteAll + DeleteByParentId + DeleteBySecondParentId
	vnswStmt := statements.NewViewNodeSwitchDataStatements(cfg)
	mock.ExpectExec(regexp.QuoteMeta(vnswStmt.DeleteAll())).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(regexp.QuoteMeta(vnswStmt.DeleteByParentId())).WithArgs(9).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(regexp.QuoteMeta(vnswStmt.DeleteBySecondParentId())).WithArgs(10).WillReturnResult(sqlmock.NewResult(0, 1))

	rVnsw := repo.NewMySQLViewNodeSwitchRepository(db, cfg)
	err = rVnsw.DeleteAll()
	assert.NoError(t, err)
	err = rVnsw.DeleteByParentId(9)
	assert.NoError(t, err)
	err = rVnsw.DeleteBySecondParentId(10)
	assert.NoError(t, err)

	// NodeSensor: DeleteAll + DeleteByParentId
	nsStmt := statements.NewNodeSensorDataStatements(cfg)
	mock.ExpectExec(regexp.QuoteMeta(nsStmt.DeleteAll())).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(regexp.QuoteMeta(nsStmt.DeleteByParentId())).WithArgs(11).WillReturnResult(sqlmock.NewResult(0, 3))

	rNs := repo.NewMySQLNodeSensorRepository(db, cfg)
	err = rNs.DeleteAll()
	assert.NoError(t, err)
	err = rNs.DeleteByParentId(11)
	assert.NoError(t, err)

	// NodeSwitch: DeleteAll + DeleteByParentId
	nswStmt := statements.NewNodeSwitchDataStatements(cfg)
	mock.ExpectExec(regexp.QuoteMeta(nswStmt.DeleteAll())).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(regexp.QuoteMeta(nswStmt.DeleteByParentId())).WithArgs(12).WillReturnResult(sqlmock.NewResult(0, 2))

	rNsw := repo.NewMySQLNodeSwitchRepository(db, cfg)
	err = rNsw.DeleteAll()
	assert.NoError(t, err)
	err = rNsw.DeleteByParentId(12)
	assert.NoError(t, err)

	// View: DeleteAll
	vStmt := statements.NewViewDataStatements(cfg)
	mock.ExpectExec(regexp.QuoteMeta(vStmt.DeleteAll())).WillReturnResult(sqlmock.NewResult(0, 1))
	rV := repo.NewMySQLViewRepository(db, cfg)
	err = rV.DeleteAll()
	assert.NoError(t, err)

	// SensorType: DeleteAll
	sStmt := statements.NewSensorTypeStatements(cfg)
	mock.ExpectExec(regexp.QuoteMeta(sStmt.DeleteAll())).WillReturnResult(sqlmock.NewResult(0, 1))
	rS := repo.NewMySQLSensorTypeRepository(db, cfg)
	err = rS.DeleteAll()
	assert.NoError(t, err)

	// SwitchType: DeleteAll
	swStmt := statements.NewSwitchTypeStatements(cfg)
	mock.ExpectExec(regexp.QuoteMeta(swStmt.DeleteAll())).WillReturnResult(sqlmock.NewResult(0, 1))
	rSw := repo.NewMySQLSwitchTypeRepository(db, cfg)
	err = rSw.DeleteAll()
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}
