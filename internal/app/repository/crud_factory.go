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

// NewViewCrudRepository returns a CrudRepository for views, using a MySQL scaffold when available.
func NewViewCrudRepository(db *sql.DB, dbType string, cfg *config.Configuration) CrudRepository {
	if dbType == "mysql" {
		return NewMySQLViewRepository(db, cfg)
	}
	return NewCrudRepositoryFromData(data.NewViewData(db, cfg))
}

// NewViewNodeSensorCrudRepository returns a CrudRepository for view node sensor data.
func NewViewNodeSensorCrudRepository(db *sql.DB, dbType string, cfg *config.Configuration) CrudRepository {
	if dbType == "mysql" {
		return NewMySQLViewNodeSensorRepository(db, cfg)
	}
	return NewCrudRepositoryFromData(data.NewViewNodeSensorData(db, cfg))
}

// NewViewNodeSwitchCrudRepository returns a CrudRepository for view node switch data.
func NewViewNodeSwitchCrudRepository(db *sql.DB, dbType string, cfg *config.Configuration) CrudRepository {
	if dbType == "mysql" {
		return NewMySQLViewNodeSwitchRepository(db, cfg)
	}
	return NewCrudRepositoryFromData(data.NewViewNodeSwitchData(db, cfg))
}

// NewSensorTypeCrudRepository returns a CrudRepository for sensor types.
func NewSensorTypeCrudRepository(db *sql.DB, dbType string, cfg *config.Configuration) CrudRepository {
	if dbType == "mysql" {
		return NewMySQLSensorTypeRepository(db, cfg)
	}
	return NewCrudRepositoryFromData(data.NewSensorType(db, cfg))
}

// NewSwitchTypeCrudRepository returns a CrudRepository for switch types.
func NewSwitchTypeCrudRepository(db *sql.DB, dbType string, cfg *config.Configuration) CrudRepository {
	if dbType == "mysql" {
		return NewMySQLSwitchTypeRepository(db, cfg)
	}
	return NewCrudRepositoryFromData(data.NewSwitchTypeData(db, cfg))
}
