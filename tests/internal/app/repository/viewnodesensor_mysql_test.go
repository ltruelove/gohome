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

func TestMySQLViewNodeSensorRepository_CRUD(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	cfg := &config.Configuration{DbType: "mysql"}
	stmt := statements.NewViewNodeSensorDataStatements(cfg)

	rows := sqlmock.NewRows([]string{"Id", "NodeId", "ViewId", "NodeSensorId", "Name"}).AddRow(21, 7, 4, 9, "sensor-view")
	mock.ExpectQuery(regexp.QuoteMeta(stmt.SelectAll())).WillReturnRows(rows)

	mock.ExpectQuery(regexp.QuoteMeta(stmt.Insert())).WithArgs(8, 5, 10, "new-sensor").WillReturnRows(sqlmock.NewRows([]string{"last_insert_id"}).AddRow(101))

	r := repo.NewMySQLViewNodeSensorRepository(db, cfg)

	list, err := r.SelectAll()
	assert.NoError(t, err)
	assert.Len(t, list, 1)

	v := &models.ViewNodeSensorData{NodeId: 8, ViewId: 5, NodeSensorId: 10, Name: "new-sensor"}
	inserted, err := r.Insert(v)
	assert.NoError(t, err)
	assert.Equal(t, 101, inserted.(*models.ViewNodeSensorData).Id)

	// SelectByParentId
	mock.ExpectQuery(regexp.QuoteMeta(stmt.SelectByParentId())).WithArgs(8).WillReturnRows(sqlmock.NewRows([]string{"Id", "NodeId", "ViewId", "NodeSensorId", "Name"}).AddRow(22, 8, 5, 10, "child-sensor"))
	listByParent, err := r.SelectByParentId(8)
	assert.NoError(t, err)
	assert.Len(t, listByParent, 1)

	// SelectById
	mock.ExpectQuery(regexp.QuoteMeta(stmt.SelectById())).WithArgs(101).WillReturnRows(sqlmock.NewRows([]string{"Id", "NodeId", "ViewId", "NodeSensorId", "Name"}).AddRow(101, 8, 5, 10, "new-sensor"))
	m, err := r.SelectById(101)
	assert.NoError(t, err)
	_ = m.(models.ViewNodeSensorData)

	// Update
	mock.ExpectExec(regexp.QuoteMeta(stmt.Update())).WithArgs(8, 5, 10, "new-sensor-up", 101).WillReturnResult(sqlmock.NewResult(0, 1))
	err = r.Update(&models.ViewNodeSensorData{Id: 101, NodeId: 8, ViewId: 5, NodeSensorId: 10, Name: "new-sensor-up"})
	assert.NoError(t, err)

	// Delete
	mock.ExpectExec(regexp.QuoteMeta(stmt.Delete())).WithArgs(101).WillReturnResult(sqlmock.NewResult(0, 1))
	err = r.Delete(101)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}
