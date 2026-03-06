package repository_test

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ltruelove/gohome/internal/app/models"
	repo "github.com/ltruelove/gohome/internal/app/repository"
	"github.com/stretchr/testify/assert"
)

func TestMySQLControlPointRepository_FetchByMac(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "ipaddress", "mac"}).AddRow(3, "cp1", "10.0.0.1", "aa:bb:cc")
	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, ipaddress, mac FROM controlpoint WHERE mac = ?")).WithArgs("aa:bb:cc").WillReturnRows(rows)

	r := repo.NewMySQLControlPointRepository(db)
	cp, err := r.FetchByMac("aa:bb:cc")

	assert.NoError(t, err)
	assert.Equal(t, 3, cp.Id)
	assert.Equal(t, "cp1", cp.Name)
}

func TestMySQLControlPointRepository_CreateAndFetchAllNodes(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO controlpoint (name, ipaddress, mac) VALUES (?, ?, ?)")).WithArgs("cp1", "10.0.0.1", "aa:bb:cc").WillReturnResult(sqlmock.NewResult(7, 1))

	r := repo.NewMySQLControlPointRepository(db)
	cp := &models.ControlPoint{Name: "cp1", IpAddress: "10.0.0.1", Mac: "aa:bb:cc"}
	err = r.Create(cp)
	assert.NoError(t, err)
	assert.Equal(t, 7, cp.Id)

	// fetch nodes
	rows := sqlmock.NewRows([]string{"id", "Name", "Mac", "RelationId"}).AddRow(1, "n1", "m1", 11)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT node.id, node.name, node.mac, cpn.id AS relationid FROM controlpointnodes AS cpn INNER JOIN node ON node.id = cpn.nodeid WHERE cpn.controlpointid = ?")).WithArgs(7).WillReturnRows(rows)

	list, err := r.FetchAllNodes(7)
	assert.NoError(t, err)
	assert.Len(t, list, 1)
	assert.Equal(t, 11, list[0].RelationId)
}

func TestMySQLControlPointRepository_FetchAllAvailable_VerifyAndDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// available
	query := `SELECT id, name, ipaddress, mac FROM controlpoint AS c WHERE (SELECT COUNT(id) FROM controlpointnodes WHERE controlpointid = c.id) < 20`
	rows := sqlmock.NewRows([]string{"id", "name", "ipaddress", "mac"}).AddRow(4, "cp-avail", "10.0.0.5", "aa:11:22")
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	// verify id not found
	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, ipaddress, mac FROM controlpoint WHERE id = ?")).WithArgs(999).WillReturnError(sql.ErrNoRows)

	// verify id found
	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, ipaddress, mac FROM controlpoint WHERE id = ?")).WithArgs(4).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "ipaddress", "mac"}).AddRow(4, "cp-avail", "10.0.0.5", "aa:11:22"))

	// delete
	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM controlpoint WHERE id = ?")).WithArgs(4).WillReturnResult(sqlmock.NewResult(0, 1))

	r := repo.NewMySQLControlPointRepository(db)

	list, err := r.FetchAllAvailable()
	assert.NoError(t, err)
	assert.Len(t, list, 1)

	ok, err := r.VerifyIdIsNew(999)
	assert.NoError(t, err)
	assert.True(t, ok)

	ok2, err := r.VerifyIdIsNew(4)
	assert.NoError(t, err)
	assert.False(t, ok2)

	err = r.Delete(4)
	assert.NoError(t, err)
}
