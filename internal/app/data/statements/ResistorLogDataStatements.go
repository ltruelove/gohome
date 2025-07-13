package statements

import "github.com/ltruelove/gohome/config"

type ResistorLogDataStatements struct {
	DbType string
}

func NewResistorLogDataStatements(config *config.Configuration) CrudStatement {
	return &ResistorLogDataStatements{DbType: config.DbType}
}

func (statements *ResistorLogDataStatements) SelectAll() string {
	if statements.DbType == "mysql" {
		return `SELECT
		Id,
		NodeSensorLogId,
		ResistorValue
		FROM ResistorLog`
	} else {
		return `SELECT
		id,
		nodesensorlogid,
		resistorvalue
		FROM resistorlog`
	}
}

func (statements *ResistorLogDataStatements) SelectById() string {
	if statements.DbType == "mysql" {
		return `SELECT
		Id,
		NodeSensorLogId,
		ResistorValue
		FROM ResistorLog
		WHERE Id = ?`
	} else {
		return `SELECT
		id,
		nodesensorlogid,
		resistorvalue
		FROM resistorlog
		WHERE id = $1`
	}
}

func (statements *ResistorLogDataStatements) SelectByParentId() string {
	if statements.DbType == "mysql" {
		return `SELECT
		Id,
		NodeSensorLogId,
		ResistorValue
		FROM ResistorLog
		WHERE NodeSensorLogId = ?`
	} else {
		return `SELECT
		id,
		nodesensorlogid,
		resistorvalue
		FROM resistorlog
		WHERE nodesensorlogid = $1`
	}
}

func (statements *ResistorLogDataStatements) SelectBySecondParentId() string {
	return "ResistorLog table has no second parent ID"
}

func (statements *ResistorLogDataStatements) Insert() string {
	if statements.DbType == "mysql" {
		return `INSERT INTO ResistorLog
		(NodeSensorLogId, Resistor)
		VALUES (?, ?); SELECT LAST_INSERT_ID();`
	} else {
		return `INSERT INTO resistorlog
		(nodesensorlogid, resistorvalue)
		VALUES ($1, $2) RETURNING Id`
	}
}

func (statements *ResistorLogDataStatements) Update() string {
	if statements.DbType == "mysql" {
		return `UPDATE ResistorLog SET NodeSensorLogId = ?, ResistorValue = ? WHERE Id = ?`
	} else {
		return `UPDATE resistorlog SET nodesensorlogid = $1, resistorvalue = $2 WHERE id = $3`
	}
}

func (statements *ResistorLogDataStatements) Delete() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM ResistorLog WHERE Id = ?`
	} else {
		return `DELETE FROM resistorlog WHERE id = $1`
	}
}

func (statements *ResistorLogDataStatements) DeleteAll() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM ResistorLog`
	} else {
		return `DELETE FROM resistorlog`
	}
}

func (statements *ResistorLogDataStatements) DeleteByParentId() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM ResistorLog WHERE NodeSensorLogId = ?`
	} else {
		return `DELETE FROM resistorlog WHERE nodesensorlogid = $1`
	}
}

func (statements *ResistorLogDataStatements) DeleteBySecondParentId() string {
	return "ResistorLog table has no second parent ID"
}
