package statements

import "github.com/ltruelove/gohome/config"

type SensorTypeDataStatements struct {
	DbType string
}

func NewSensorTypeDataStatements(config *config.Configuration) CrudStatement {
	return &SensorTypeDataStatements{DbType: config.DbType}
}

func (statements *SensorTypeDataStatements) SelectAll() string {
	if statements.DbType == "mysql" {
		return `SELECT Id, SensorTypeId, TypeName, ValueType FROM SensorTypeData`
	} else {
		return `SELECT id, sensortypeid, name, valuetype FROM sensortypedata`
	}
}

func (statements *SensorTypeDataStatements) SelectById() string {
	if statements.DbType == "mysql" {
		return `SELECT Id, SensorTypeId, TypeName, ValueType FROM SensorTypeData WHERE Id = ?`
	} else {
		return `SELECT id, sensortypeid, name, valuetype FROM sensortypedata WHERE id = $1`
	}
}

func (statements *SensorTypeDataStatements) SelectByParentId() string {
	if statements.DbType == "mysql" {
		return `SELECT Id, SensorTypeId, TypeName, ValueType FROM SensorTypeData WHERE SensorTypeId = ?`
	} else {
		return `SELECT id, sensortypeid, name, valuetype FROM sensortypedata WHERE sensortypeid = $1`
	}
}

func (statements *SensorTypeDataStatements) SelectBySecondParentId() string {
	return "no second parent ID selection for SensorType"
}

func (statements *SensorTypeDataStatements) Insert() string {
	if statements.DbType == "mysql" {
		return `INSERT INTO SensorTypeData (SensorTypeId, TypeName, ValueType) VALUES (?, ?, ?)`
	} else {
		return `INSERT INTO sensortypedata (sensortypeid, name, valuetype) VALUES ($1, $2, $3)`
	}
}

func (statements *SensorTypeDataStatements) Update() string {
	if statements.DbType == "mysql" {
		return `UPDATE SensorTypeData SET SensorTypeId=?, TypeName = ?, ValueType = ? WHERE Id = ?`
	} else {
		return `UPDATE sensortypedata SET sensortypeid = $1, name = $2, valuetype = $3 WHERE id = $4`
	}
}

func (statements *SensorTypeDataStatements) Delete() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM SensorTypeData WHERE Id = ?`
	} else {
		return `DELETE FROM sensortypedata WHERE id = $1`
	}
}

func (statements *SensorTypeDataStatements) DeleteByParentId() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM SensorTypeData WHERE SensorTypeId = ?`
	} else {
		return `DELETE FROM sensortypedata WHERE sensortypeid = $1`
	}
}

func (statements *SensorTypeDataStatements) DeleteAll() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM SensorTypeData`
	} else {
		return `DELETE FROM sensortypedata`
	}
}
