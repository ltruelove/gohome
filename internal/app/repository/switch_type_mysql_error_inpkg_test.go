package repository

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ltruelove/gohome/config"
)

func TestSwitchType_SelectById_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %v", err)
	}
	defer db.Close()

	cfg := &config.Configuration{}
	repo := NewMySQLSwitchTypeRepository(db, cfg)

	mock.ExpectQuery(".*").WithArgs(99).WillReturnError(errors.New("select fail"))

	_, err = repo.SelectById(99)
	if err == nil {
		t.Fatalf("expected error when SelectById fails")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}
