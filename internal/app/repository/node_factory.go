package repository

import (
	"database/sql"
)

// NewNodeRepository returns an implementation of NodeRepository.
// Currently the MySQL scaffold is the canonical implementation.
func NewNodeRepository(db *sql.DB, dbType string) NodeRepository {
	return NewMySQLNodeRepository(db)
}
