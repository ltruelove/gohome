package statements

import "github.com/ltruelove/gohome/config"

type ViewDataStatements struct {
	DbType string
}

func NewViewDataStatements(config *config.Configuration) *ViewDataStatements {
	return &ViewDataStatements{DbType: config.DbType}
}

func (statements *ViewDataStatements) SelectAll() string {
	if statements.DbType == "mysql" {
		return `SELECT
			Id,
			ViewName
			FROM View`
	} else {
		return `SELECT
			id,
			name
			FROM view`
	}
}

func (statements *ViewDataStatements) SelectById() string {
	if statements.DbType == "mysql" {
		return `SELECT
			Id,
			ViewName
			FROM View
			WHERE Id = ?`
	} else {
		return `SELECT
			id,
			name
			FROM view
			WHERE id = $1`
	}
}

func (statements *ViewDataStatements) Insert() string {
	if statements.DbType == "mysql" {
		return `INSERT INTO View (ViewName) VALUES (?) RETURNING Id`
	} else {
		return `INSERT INTO view (name) VALUES ($1) RETURNING id`
	}
}

func (statements *ViewDataStatements) Update() string {
	if statements.DbType == "mysql" {
		return `UPDATE View SET ViewName = ? WHERE Id = ?`
	} else {
		return `UPDATE view SET name = $1 WHERE id = $2`
	}
}

func (statements *ViewDataStatements) Delete() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM View WHERE Id = ?`
	} else {
		return `DELETE FROM view WHERE id = $1`
	}
}
