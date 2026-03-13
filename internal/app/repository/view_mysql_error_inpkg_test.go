package repository

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ltruelove/gohome/config"
)

func TestViewMySQL_SelectById_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %v", err)
	}
	defer db.Close()

	cfg := &config.Configuration{}
	repo := NewMySQLViewRepository(db, cfg)

	mock.ExpectQuery(".*").WithArgs(5).WillReturnError(errors.New("query fail"))

	_, err = repo.SelectById(5)
	if err == nil {
		t.Fatalf("expected error when SelectById fails")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}
