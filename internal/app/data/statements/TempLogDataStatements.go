package statements

import "github.com/ltruelove/gohome/config"

type TempLogDataStatements struct {
	DbType string
}

func NewTempLogDataStatements(config *config.Configuration) CrudStatement {
	return &TempLogDataStatements{DbType: config.DbType}
}

func (statements *TempLogDataStatements) SelectAll() string {
	if statements.DbType == "mysql" {
		return `SELECT
		Id,
		NodeSensorLogId,
		TemperatureF,
		TemperatureC,
		Humidity
		FROM TempLog`
	} else {
		return `SELECT
		id,
		nodesensorlogid,
		temperaturef,
		temperaturec,
		humidity
		FROM templog`
	}
}

func (statements *TempLogDataStatements) SelectById() string {
	if statements.DbType == "mysql" {
		return `SELECT
		Id,
		NodeSensorLogId,
		TemperatureF,
		TemperatureC,
		Humidity
		FROM TempLog
		WHERE Id = ?`
	} else {
		return `SELECT
		id,
		nodesensorlogid,
		temperaturef,
		temperaturec,
		humidity
		FROM templog
		WHERE id = $1`
	}
}

func (statements *TempLogDataStatements) SelectByParentId() string {
	if statements.DbType == "mysql" {
		return `SELECT
		Id,
		NodeSensorLogId,
		TemperatureF,
		TemperatureC,
		Humidity
		FROM TempLog
		WHERE NodeSensorLogId = ?`
	} else {
		return `SELECT
		id,
		nodesensorlogid,
		temperaturef,
		temperaturec,
		humidity
		FROM templog
		WHERE nodesensorlogid = $1`
	}
}

func (statements *TempLogDataStatements) SelectBySecondParentId() string {
	return "TempLog table has no second parent ID"
}

func (statements *TempLogDataStatements) Insert() string {
	if statements.DbType == "mysql" {
		return `INSERT INTO TempLog
		(NodeSensorLogId, TemperatureF, TemperatureC, Humidity)
		VALUES (?, ?, ?, ?); SELECT LAST_INSERT_ID();`
	} else {
		return `INSERT INTO templog
		(nodesensorlogid, temperaturef, temperaturec, humidity)
		VALUES ($1, $2, $3, $4) RETURNING Id`
	}
}

func (statements *TempLogDataStatements) Update() string {
	if statements.DbType == "mysql" {
		return `UPDATE TempLog SET NodeSensorLogId = ?, TemperatureF = ?, TemperatureC = ?, Humidity = ? WHERE Id = ?`
	} else {
		return `UPDATE templog SET nodesensorlogid = $1, temperaturef = $2, temperaturec = $3, humidity = $4 WHERE id = $5`
	}
}

func (statements *TempLogDataStatements) Delete() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM TempLog WHERE Id = ?`
	} else {
		return `DELETE FROM templog WHERE id = $1`
	}
}

func (statements *TempLogDataStatements) DeleteAll() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM TempLog`
	} else {
		return `DELETE FROM templog`
	}
}

func (statements *TempLogDataStatements) DeleteByParentId() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM TempLog WHERE NodeSensorLogId = ?`
	} else {
		return `DELETE FROM templog WHERE nodesensorlogid = $1`
	}
}

func (statements *TempLogDataStatements) DeleteBySecondParentId() string {
	return "TempLog table has no second parent ID"
}
