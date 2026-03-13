package repository

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ltruelove/gohome/config"
)

func TestCompoundMySQL_FetchViewNodeSensorData_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %v", err)
	}
	defer db.Close()

	cfg := &config.Configuration{}
	repo := NewMySQLCompoundRepository(db, cfg)

	// simulate query error
	mock.ExpectQuery(".*").WithArgs(7).WillReturnError(errors.New("query fail"))

	_, err = repo.FetchViewNodeSensorDataByViewId(7)
	if err == nil {
		t.Fatalf("expected error when compound query fails")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}

func TestCompoundMySQL_FetchViewNodeSwitchData_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %v", err)
	}
	defer db.Close()

	cfg := &config.Configuration{}
	repo := NewMySQLCompoundRepository(db, cfg)

	mock.ExpectQuery(".*").WithArgs(9).WillReturnError(errors.New("query fail"))

	_, err = repo.FetchViewNodeSwitchDataByViewId(9)
	if err == nil {
		t.Fatalf("expected error when compound query fails")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}
