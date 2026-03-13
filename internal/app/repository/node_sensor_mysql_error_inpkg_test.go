package repository

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ltruelove/gohome/config"
)

func TestNodeSensor_SelectAll_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %v", err)
	}
	defer db.Close()

	cfg := &config.Configuration{}
	repo := NewMySQLNodeSensorRepository(db, cfg)

	mock.ExpectQuery(".*").WillReturnError(errors.New("select fail"))

	_, err = repo.SelectAll()
	if err == nil {
		t.Fatalf("expected error when SelectAll fails")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}
