package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ltruelove/gohome/config"
)

func TestSensorType_SelectByParentId_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %v", err)
	}
	defer db.Close()

	cfg := &config.Configuration{}
	repo := NewMySQLSensorTypeRepository(db, cfg)

	// SelectByParentId is intentionally unimplemented and should return nil,nil
	res, err := repo.SelectByParentId(42)
	if err != nil {
		t.Fatalf("expected no error from SelectByParentId, got: %v", err)
	}
	if res != nil {
		t.Fatalf("expected nil result from SelectByParentId, got: %v", res)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}
