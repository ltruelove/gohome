package repository

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/stretchr/testify/assert"
)

func TestMySQLNodeRepository_FetchById_InPkg(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "mac", "name", "ipaddress"}).AddRow(1, "aa:bb:cc", "node1", "192.168.1.2")

	query := "SELECT id, mac, name, ipaddress FROM node WHERE id = ?"
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(1).WillReturnRows(rows)

	r := NewMySQLNodeRepository(db)
	n, err := r.FetchById(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, n.Id)
	assert.Equal(t, "aa:bb:cc", n.Mac)
	assert.Equal(t, "node1", n.Name)
}

func TestMySQLNodeRepository_Create_InPkg(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO node (mac, name, ipaddress) VALUES (?, ?, ?)")).WithArgs("aa:bb:cc", "node1", "192.168.1.2").WillReturnResult(sqlmock.NewResult(42, 1))

	r := NewMySQLNodeRepository(db)
	node := &models.Node{Mac: "aa:bb:cc", Name: "node1", IpAddress: "192.168.1.2"}
	err = r.Create(node)

	assert.NoError(t, err)
	assert.Equal(t, 42, node.Id)
}
