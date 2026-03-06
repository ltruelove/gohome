package repository

import (
	"database/sql"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data"
)

// NewNodeSensorCrudRepository returns a CrudRepository for node sensors.
// Currently it wraps the data implementation; can be replaced with a MySQL scaffold later.
func NewNodeSensorCrudRepository(db *sql.DB, dbType string, cfg *config.Configuration) CrudRepository {
	return NewCrudRepositoryFromData(data.NewNodeSensorData(db, cfg))
}

// NewNodeSwitchCrudRepository returns a CrudRepository for node switches.
func NewNodeSwitchCrudRepository(db *sql.DB, dbType string, cfg *config.Configuration) CrudRepository {
	return NewCrudRepositoryFromData(data.NewNodeSwitchData(db, cfg))
}
