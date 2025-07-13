package statements

import "github.com/ltruelove/gohome/config"

type ControlPointDataStatements struct {
	DbType string
}

func NewControlPointDataStatements(config *config.Configuration) CrudStatement {
	return &ControlPointDataStatements{DbType: config.DbType}
}

func (statements *ControlPointDataStatements) SelectAll() string {
	if statements.DbType == "mysql" {
		return `SELECT
		Id,
		PointName,
		IpAddress,
		Mac
		FROM ControlPoint`
	} else {
		return `SELECT
		id,
		pointname,
		ipaddress,
		mac
		FROM controlpoint`
	}
}

func (statements *ControlPointDataStatements) SelectById() string {
	if statements.DbType == "mysql" {
		return `SELECT
		Id,
		PointName,
		IpAddress,
		Mac
		FROM ControlPoint
		WHERE Id = ?`
	} else {
		return `SELECT
		id,
		pointname,
		ipaddress,
		mac
		FROM controlpoint
		WHERE id = $1`
	}
}

func (statements *ControlPointDataStatements) SelectByParentId() string {
	return "ControlPoint does not support parent ID queries"
}

func (statements *ControlPointDataStatements) SelectBySecondParentId() string {
	return "ControlPoint does not support second parent ID queries"
}

func (statements *ControlPointDataStatements) Insert() string {
	if statements.DbType == "mysql" {
		return `INSERT INTO ControlPoint (PointName, IpAddress, Mac)
		VALUES (?, ?, ?); SELECT LAST_INSERT_ID()`
	} else {
		return `INSERT INTO controlpoint (pointname, ipaddress, mac)
		VALUES ($1, $2, $3) RETURNING id`
	}
}

func (statements *ControlPointDataStatements) Update() string {
	if statements.DbType == "mysql" {
		return `UPDATE ControlPoint SET PointName = ?, IpAddress = ?, Mac = ? WHERE Id = ?`
	} else {
		return `UPDATE controlpoint SET pointname = $1, ipaddress = $2, mac = $3 WHERE id = $4`
	}
}

func (statements *ControlPointDataStatements) Delete() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM ControlPoint WHERE Id = ?`
	} else {
		return `DELETE FROM controlpoint WHERE id = $1`
	}
}

func (statements *ControlPointDataStatements) DeleteAll() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM ControlPoint`
	} else {
		return `DELETE FROM controlpoint`
	}
}

func (statements *ControlPointDataStatements) DeleteByParentId() string {
	return "ControlPoint does not support parent ID deletion"
}

func (statements *ControlPointDataStatements) DeleteBySecondParentId() string {
	return "ControlPoint does not support second parent ID deletion"
}
