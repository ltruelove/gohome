package repository

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ltruelove/gohome/internal/app/models"
)

// Test FetchById returns sql.ErrNoRows -> repository should return error
func TestNodeMySQL_FetchById_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %v", err)
	}
	defer db.Close()

	repo := NewMySQLNodeRepository(db)

	// expectation: query for fetch by id returns no rows
	mock.ExpectQuery("SELECT .* FROM node WHERE id = ?").WithArgs(1).WillReturnError(sql.ErrNoRows)

	_, err = repo.FetchById(1)
	if err == nil {
		t.Fatalf("expected error when fetch by id not found")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}

// Test Create returns Exec error inside transaction and repo surfaces error
func TestNodeMySQL_Create_TxFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %v", err)
	}
	defer db.Close()

	repo := NewMySQLNodeRepository(db)

	// Exec returns error (no transaction in Create implementation)
	mock.ExpectExec("INSERT INTO node").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnError(sql.ErrConnDone)

	n := &models.Node{Name: "x", Mac: "m", IpAddress: "1.2.3.4"}
	err = repo.Create(n)
	if err == nil {
		t.Fatalf("expected error on create when exec fails")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}
