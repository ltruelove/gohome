package repository

import (
	"database/sql"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/viewModels"
)

type CompoundRepository interface {
	FetchViewNodeSensorDataByViewId(viewId int) ([]viewModels.ViewNodeSensorVM, error)
	FetchViewNodeSwitchDataByViewId(viewId int) ([]viewModels.ViewNodeSwitchVM, error)
}

// NewCompoundRepository chooses an implementation based on dbType
func NewCompoundRepository(db *sql.DB, dbType string, config *config.Configuration) CompoundRepository {
	if dbType == "mysql" {
		return NewMySQLCompoundRepository(db, config)
	}
	// fallback to MySQL scaffold until legacy data is removed
	return NewMySQLCompoundRepository(db, config)
}
