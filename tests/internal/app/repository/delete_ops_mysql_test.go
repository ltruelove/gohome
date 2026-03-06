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

	// ViewNodeSensor: Delete + DeleteByParentId + DeleteBySecondParentId
	vnsStmt := statements.NewViewNodeSensorDataStatements(cfg)
	mock.ExpectExec(regexp.QuoteMeta(vnsStmt.Delete())).WithArgs(21).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(regexp.QuoteMeta(vnsStmt.DeleteByParentId())).WithArgs(7).WillReturnResult(sqlmock.NewResult(0, 2))
	mock.ExpectExec(regexp.QuoteMeta(vnsStmt.DeleteBySecondParentId())).WithArgs(8).WillReturnResult(sqlmock.NewResult(0, 1))

	rVns := repo.NewMySQLViewNodeSensorRepository(db, cfg)
	err = rVns.Delete(21)
	assert.NoError(t, err)
	err = rVns.DeleteByParentId(7)
	assert.NoError(t, err)
	err = rVns.DeleteBySecondParentId(8)
	assert.NoError(t, err)

	// ViewNodeSwitch: Delete + DeleteByParentId + DeleteBySecondParentId
	vnswStmt := statements.NewViewNodeSwitchDataStatements(cfg)
	mock.ExpectExec(regexp.QuoteMeta(vnswStmt.Delete())).WithArgs(99).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(regexp.QuoteMeta(vnswStmt.DeleteByParentId())).WithArgs(9).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(regexp.QuoteMeta(vnswStmt.DeleteBySecondParentId())).WithArgs(10).WillReturnResult(sqlmock.NewResult(0, 1))

	rVnsw := repo.NewMySQLViewNodeSwitchRepository(db, cfg)
	err = rVnsw.Delete(99)
	assert.NoError(t, err)
	err = rVnsw.DeleteByParentId(9)
	assert.NoError(t, err)
	err = rVnsw.DeleteBySecondParentId(10)
	assert.NoError(t, err)

	// NodeSensor: Delete + DeleteByParentId
	nsStmt := statements.NewNodeSensorDataStatements(cfg)
	mock.ExpectExec(regexp.QuoteMeta(nsStmt.Delete())).WithArgs(33).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(regexp.QuoteMeta(nsStmt.DeleteByParentId())).WithArgs(11).WillReturnResult(sqlmock.NewResult(0, 3))

	rNs := repo.NewMySQLNodeSensorRepository(db, cfg)
	err = rNs.Delete(33)
	assert.NoError(t, err)
	err = rNs.DeleteByParentId(11)
	assert.NoError(t, err)

	// NodeSwitch: Delete + DeleteByParentId
	nswStmt := statements.NewNodeSwitchDataStatements(cfg)
	mock.ExpectExec(regexp.QuoteMeta(nswStmt.Delete())).WithArgs(44).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(regexp.QuoteMeta(nswStmt.DeleteByParentId())).WithArgs(12).WillReturnResult(sqlmock.NewResult(0, 2))

	rNsw := repo.NewMySQLNodeSwitchRepository(db, cfg)
	err = rNsw.Delete(44)
	assert.NoError(t, err)
	err = rNsw.DeleteByParentId(12)
	assert.NoError(t, err)

	// View: Delete (single)
	vStmt := statements.NewViewDataStatements(cfg)
	mock.ExpectExec(regexp.QuoteMeta(vStmt.Delete())).WithArgs(77).WillReturnResult(sqlmock.NewResult(0, 1))
	rV := repo.NewMySQLViewRepository(db, cfg)
	err = rV.Delete(77)
	assert.NoError(t, err)

	// SensorType: Delete (single)
	sStmt := statements.NewSensorTypeStatements(cfg)
	mock.ExpectExec(regexp.QuoteMeta(sStmt.Delete())).WithArgs(55).WillReturnResult(sqlmock.NewResult(0, 1))
	rS := repo.NewMySQLSensorTypeRepository(db, cfg)
	err = rS.Delete(55)
	assert.NoError(t, err)

	// SwitchType: Delete (single)
	swStmt := statements.NewSwitchTypeStatements(cfg)
	mock.ExpectExec(regexp.QuoteMeta(swStmt.Delete())).WithArgs(66).WillReturnResult(sqlmock.NewResult(0, 1))
	rSw := repo.NewMySQLSwitchTypeRepository(db, cfg)
	err = rSw.Delete(66)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}
