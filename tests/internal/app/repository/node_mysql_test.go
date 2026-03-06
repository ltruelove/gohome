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

func TestMySQLNodeRepository_FetchById(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "mac", "name", "ipaddress"}).AddRow(1, "aa:bb:cc", "node1", "192.168.1.2")

	query := "SELECT id, mac, name, ipaddress FROM node WHERE id = ?"
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(1).WillReturnRows(rows)

	r := repo.NewMySQLNodeRepository(db)
	n, err := r.FetchById(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, n.Id)
	assert.Equal(t, "aa:bb:cc", n.Mac)
	assert.Equal(t, "node1", n.Name)
}

func TestMySQLNodeRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO node (mac, name, ipaddress) VALUES (?, ?, ?)")).WithArgs("aa:bb:cc", "node1", "192.168.1.2").WillReturnResult(sqlmock.NewResult(42, 1))

	r := repo.NewMySQLNodeRepository(db)
	node := &models.Node{Mac: "aa:bb:cc", Name: "node1", IpAddress: "192.168.1.2"}
	err = r.Create(node)

	assert.NoError(t, err)
	assert.Equal(t, 42, node.Id)
}

func TestMySQLNodeRepository_FetchAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "mac", "name", "ipaddress"}).AddRow(1, "aa:bb", "n1", "1.2.3.4").AddRow(2, "cc:dd", "n2", "1.2.3.5")
	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, mac, name, ipaddress FROM node")).WillReturnRows(rows)

	r := repo.NewMySQLNodeRepository(db)
	list, err := r.FetchAll()

	assert.NoError(t, err)
	assert.Len(t, list, 2)
}

func TestMySQLNodeRepository_UpdateAndDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta("UPDATE node SET mac = ?, name = ?, ipaddress = ? WHERE id = ?")).WithArgs("aa:bb:cc", "node1-up", "192.168.1.3", 42).WillReturnResult(sqlmock.NewResult(0, 1))

	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM node WHERE id = ?")).WithArgs(42).WillReturnResult(sqlmock.NewResult(0, 1))

	r := repo.NewMySQLNodeRepository(db)
	node := &models.Node{Id: 42, Mac: "aa:bb:cc", Name: "node1-up", IpAddress: "192.168.1.3"}

	err = r.Update(node)
	assert.NoError(t, err)

	err = r.Delete(42)
	assert.NoError(t, err)
}

func TestMySQLNodeRepository_FetchNodeSwitchesAndFetchNodeSwitch(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "nodeid", "switchtypeid", "name", "pin", "momentarypressduration", "isclosedon"}).
		AddRow(10, 5, 2, "switch1", 7, 0, 1).
		AddRow(11, 5, 3, "switch2", 8, 100, 0)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, nodeid, switchtypeid, name, pin, momentarypressduration, isclosedon FROM nodeswitch WHERE nodeid = ?")).WithArgs(5).WillReturnRows(rows)

	single := sqlmock.NewRows([]string{"id", "nodeid", "switchtypeid", "name", "pin", "momentarypressduration", "isclosedon"}).AddRow(10, 5, 2, "switch1", 7, 0, 1)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, nodeid, switchtypeid, name, pin, momentarypressduration, isclosedon FROM nodeswitch WHERE id = ?")).WithArgs(10).WillReturnRows(single)

	r := repo.NewMySQLNodeRepository(db)

	list, err := r.FetchNodeSwitches(5)
	assert.NoError(t, err)
	assert.Len(t, list, 2)
	assert.Equal(t, 10, list[0].Id)

	ns, err := r.FetchNodeSwitch(10)
	assert.NoError(t, err)
	assert.Equal(t, 10, ns.Id)
	assert.Equal(t, 7, ns.Pin)
}

func TestMySQLNodeRepository_FetchControlPointByNode_VerifyIdIsNew(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// FetchControlPointByNode
	row := sqlmock.NewRows([]string{"id", "name", "ipaddress", "mac"}).AddRow(21, "cp-node", "10.0.0.2", "aa:bb:11")
	mock.ExpectQuery(regexp.QuoteMeta("SELECT cp.id, cp.name, cp.ipaddress, cp.mac FROM controlpointnodes AS cpn INNER JOIN controlpoint AS cp ON cp.id = cpn.controlpointid WHERE cpn.nodeid = ?")).WithArgs(5).WillReturnRows(row)

	r := repo.NewMySQLNodeRepository(db)
	cp, err := r.FetchControlPointByNode(5)
	assert.NoError(t, err)
	assert.Equal(t, 21, cp.Id)

	// VerifyIdIsNew when not found
	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, mac, name, ipaddress FROM node WHERE id = ?")).WithArgs(999).WillReturnError(sql.ErrNoRows)
	ok, err := r.VerifyIdIsNew(999)
	assert.NoError(t, err)
	assert.True(t, ok)

	// VerifyIdIsNew when exists
	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, mac, name, ipaddress FROM node WHERE id = ?")).WithArgs(21).WillReturnRows(sqlmock.NewRows([]string{"id", "mac", "name", "ipaddress"}).AddRow(21, "aa:bb", "n", "1.2.3.4"))
	ok2, err := r.VerifyIdIsNew(21)
	assert.NoError(t, err)
	assert.False(t, ok2)
}
