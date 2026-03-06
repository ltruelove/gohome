package repository

import "database/sql"

// NewControlPointRepository returns the appropriate ControlPointRepository
// implementation. For now we always return the MySQL implementation which
// contains the concrete SQL logic. This lets us migrate package-level
// `data` helpers into repository implementations incrementally.
func NewControlPointRepository(db *sql.DB, dbType string) ControlPointRepository {
	return NewMySQLControlPointRepository(db)
}
