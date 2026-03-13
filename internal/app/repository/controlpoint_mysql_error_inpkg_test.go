package repository

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestControlPointMySQL_FetchById_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %v", err)
	}
	defer db.Close()

	repo := NewMySQLControlPointRepository(db)

	// simulate query returning driver error
	mock.ExpectQuery(".*").WithArgs(5).WillReturnError(errors.New("driver fail"))

	_, err = repo.FetchById(5)
	if err == nil {
		t.Fatalf("expected error when query fails")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}
