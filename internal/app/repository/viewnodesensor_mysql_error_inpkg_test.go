package repository

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/models"
)

func TestViewNodeSensor_Insert_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %v", err)
	}
	defer db.Close()

	cfg := &config.Configuration{}
	repo := NewMySQLViewNodeSensorRepository(db, cfg)

	// Insert uses QueryRow().Scan — expect a query error and pass a pointer
	mock.ExpectQuery("INSERT INTO viewnodesensor").WithArgs(0, 1, 2, "").WillReturnError(errors.New("insert fail"))

	_, err = repo.Insert(&models.ViewNodeSensorData{ViewId: 1, NodeSensorId: 2})
	if err == nil {
		t.Fatalf("expected error when insert fails")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}
