package data

import "github.com/ltruelove/gohome/config"

type SensorTypeStatements struct {
	DbType string
}

func NewSensorTypeStatements(config *config.Configuration) *SensorTypeStatements {
	return &SensorTypeStatements{DbType: config.DbType}
}

func (statements *SensorTypeStatements) SelectAllSensorTypes() string {
	if statements.DbType == "mysql" {
		return "SELECT Id, TypeName FROM SensorType"
	} else {
		return "SELECT id, name FROM sensortype"
	}

}

func (statements *SensorTypeStatements) SelectSensorTypeById() string {
	if statements.DbType == "mysql" {
		return "SELECT Id, TypeName FROM SensorType WHERE Id = ?"
	} else {
		return "SELECT id, name FROM sensortype WHERE id = $1"
	}
}

func (statements *SensorTypeStatements) SelectSensorTypeData() string {
	if statements.DbType == "mysql" {
		return `SELECT
		Id, 
		TypeName, 
		ValueType 
		FROM SensorTypeData 
		WHERE SensorTypeId = ?`
	} else {
		return `SELECT
		id, 
		name, 
		valuetype 
		FROM sensortypedata 
		WHERE sensortypeid = $1`
	}
}
