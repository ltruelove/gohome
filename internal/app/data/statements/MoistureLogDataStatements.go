package statements

import "github.com/ltruelove/gohome/config"

type MoistureLogDataStatements struct {
	DbType string
}

func NewMoistureLogDataStatements(config *config.Configuration) CrudStatement {
	return &MoistureLogDataStatements{DbType: config.DbType}
}

func (statements *MoistureLogDataStatements) SelectAll() string {
	if statements.DbType == "mysql" {
		return `SELECT
		Id,
		NodeSensorLogId,
		Moisture
		FROM MoistureLog`
	} else {
		return `SELECT
		id,
		nodesensorlogid,
		moisture
		FROM moisturelog`
	}
}

func (statements *MoistureLogDataStatements) SelectById() string {
	if statements.DbType == "mysql" {
		return `SELECT
		Id,
		NodeSensorLogId,
		Moisture
		FROM MoistureLog
		WHERE Id = ?`
	} else {
		return `SELECT
		id,
		nodesensorlogid,
		moisture
		FROM moisturelog
		WHERE id = $1`
	}
}

func (statements *MoistureLogDataStatements) SelectByParentId() string {
	if statements.DbType == "mysql" {
		return `SELECT
		Id,
		NodeSensorLogId,
		Moisture
		FROM MoistureLog
		WHERE NodeSensorLogId = ?`
	} else {
		return `SELECT
		id,
		nodesensorlogid,
		moisture
		FROM moisturelog
		WHERE nodesensorlogid = $1`
	}
}

func (statements *MoistureLogDataStatements) SelectBySecondParentId() string {
	return "MoistureLog table has no second parent ID"
}

func (statements *MoistureLogDataStatements) Insert() string {
	if statements.DbType == "mysql" {
		return `INSERT INTO MoistureLog
		(NodeSensorLogId, Moisture)
		VALUES (?, ?); SELECT LAST_INSERT_ID();`
	} else {
		return `INSERT INTO moisturelog
		(nodesensorlogid, moisture)
		VALUES ($1, $2) RETURNING Id`
	}
}

func (statements *MoistureLogDataStatements) Update() string {
	if statements.DbType == "mysql" {
		return `UPDATE MoistureLog SET NodeSensorLogId = ?, Moisture = ? WHERE Id = ?`
	} else {
		return `UPDATE moisturelog SET nodesensorlogid = $1, moisture = $2 WHERE id = $3`
	}
}

func (statements *MoistureLogDataStatements) Delete() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM MoistureLog WHERE Id = ?`
	} else {
		return `DELETE FROM moisturelog WHERE id = $1`
	}
}

func (statements *MoistureLogDataStatements) DeleteAll() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM MoistureLog`
	} else {
		return `DELETE FROM moisturelog`
	}
}

func (statements *MoistureLogDataStatements) DeleteByParentId() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM MoistureLog WHERE NodeSensorLogId = ?`
	} else {
		return `DELETE FROM moisturelog WHERE nodesensorlogid = $1`
	}
}

func (statements *MoistureLogDataStatements) DeleteBySecondParentId() string {
	return "MoistureLog table has no second parent ID"
}
