package statements

import "github.com/ltruelove/gohome/config"

type MagneticLogDataStatements struct {
	DbType string
}

func NewMagneticLogDataStatements(config *config.Configuration) CrudStatement {
	return &MagneticLogDataStatements{DbType: config.DbType}
}

func (statements *MagneticLogDataStatements) SelectAll() string {
	if statements.DbType == "mysql" {
		return `SELECT
		Id,
		NodeSensorLogId,
		IsClosed
		FROM MagneticLog`
	} else {
		return `SELECT
		id,
		nodesensorlogid,
		isclosed
		FROM magneticlog`
	}
}

func (statements *MagneticLogDataStatements) SelectById() string {
	if statements.DbType == "mysql" {
		return `SELECT
		Id,
		NodeSensorLogId,
		IsClosed
		FROM MagneticLog
		WHERE Id = ?`
	} else {
		return `SELECT
		id,
		nodesensorlogid,
		isclosed
		FROM magneticlog
		WHERE id = $1`
	}
}

func (statements *MagneticLogDataStatements) SelectByParentId() string {
	if statements.DbType == "mysql" {
		return `SELECT
		Id,
		NodeSensorLogId,
		IsClosed
		FROM MagneticLog
		WHERE NodeSensorLogId = ?`
	} else {
		return `SELECT
		id,
		nodesensorlogid,
		isclosed
		FROM magneticlog
		WHERE nodesensorlogid = $1`
	}
}

func (statements *MagneticLogDataStatements) SelectBySecondParentId() string {
	return "MagneticLog table has no second parent ID"
}

func (statements *MagneticLogDataStatements) Insert() string {
	if statements.DbType == "mysql" {
		return `INSERT INTO MagneticLog
		(NodeSensorLogId, IsClosed)
		VALUES (?, ?); SELECT LAST_INSERT_ID();`
	} else {
		return `INSERT INTO magneticlog
		(nodesensorlogid, isclosed)
		VALUES ($1, $2) RETURNING Id`
	}
}

func (statements *MagneticLogDataStatements) Update() string {
	if statements.DbType == "mysql" {
		return `UPDATE MagneticLog SET NodeSensorLogId = ?, IsClosed = ? WHERE Id = ?`
	} else {
		return `UPDATE magneticlog SET nodesensorlogid = $1, isclosed = $2 WHERE id = $3`
	}
}

func (statements *MagneticLogDataStatements) Delete() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM MagneticLog WHERE Id = ?`
	} else {
		return `DELETE FROM magneticlog WHERE id = $1`
	}
}

func (statements *MagneticLogDataStatements) DeleteAll() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM MagneticLog`
	} else {
		return `DELETE FROM magneticlog`
	}
}

func (statements *MagneticLogDataStatements) DeleteByParentId() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM MagneticLog WHERE NodeSensorLogId = ?`
	} else {
		return `DELETE FROM magneticlog WHERE nodesensorlogid = $1`
	}
}

func (statements *MagneticLogDataStatements) DeleteBySecondParentId() string {
	return "MagneticLogData table has no second parent ID deletion"
}
