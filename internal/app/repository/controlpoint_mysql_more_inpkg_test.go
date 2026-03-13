package repository

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/stretchr/testify/assert"
)

func TestMySQLControlPointRepository_AdditionalQueries_InPkg(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// FetchAllAvailable
	availRows := sqlmock.NewRows([]string{"id", "name", "ipaddress", "mac"}).AddRow(11, "cpA", "1.2.3.4", "aa:11")
	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, ipaddress, mac FROM controlpoint AS c WHERE (SELECT COUNT(id) FROM controlpointnodes WHERE controlpointid = c.id) < 20")).WillReturnRows(availRows)

	r := NewMySQLControlPointRepository(db)
	list, err := r.FetchAllAvailable()
	assert.NoError(t, err)
	assert.Len(t, list, 1)

	// FetchAllNodes
	nodeRows := sqlmock.NewRows([]string{"id", "name", "mac", "relationid"}).AddRow(5, "nd1", "aa:bb", 77)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT node.id, node.name, node.mac, cpn.id AS relationid FROM controlpointnodes AS cpn INNER JOIN node ON node.id = cpn.nodeid WHERE cpn.controlpointid = ?")).WithArgs(10).WillReturnRows(nodeRows)
	cpNodes, err := r.FetchAllNodes(10)
	assert.NoError(t, err)
	assert.Len(t, cpNodes, 1)
	assert.Equal(t, 77, cpNodes[0].RelationId)

	// FetchByMac
	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, ipaddress, mac FROM controlpoint WHERE mac = ?")).WithArgs("aa:11").WillReturnRows(sqlmock.NewRows([]string{"id", "name", "ipaddress", "mac"}).AddRow(31, "cp-1", "10.0.0.3", "aa:11"))
	cp, err := r.FetchByMac("aa:11")
	assert.NoError(t, err)
	assert.Equal(t, 31, cp.Id)

	// VerifyIdIsNew: simulate sql.ErrNoRows
	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, ipaddress, mac FROM controlpoint WHERE id = ?")).WithArgs(999).WillReturnError(sql.ErrNoRows)
	ok, verr := r.VerifyIdIsNew(999)
	assert.True(t, ok)
	assert.NoError(t, verr)

	// AddNodeToControlPoint
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO controlpointnodes (controlpointid, nodeid) VALUES (?, ?)")).WithArgs(10, 5).WillReturnResult(sqlmock.NewResult(1, 1))
	err = r.AddNodeToControlPoint(&models.ControlPointNode{ControlPointId: 10, NodeId: 5})
	assert.NoError(t, err)
}
