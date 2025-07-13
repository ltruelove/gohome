package statements

import "github.com/ltruelove/gohome/config"

type SensorTypeStatements struct {
	DbType string
}

func NewSensorTypeStatements(config *config.Configuration) CrudStatement {
	return &SensorTypeStatements{DbType: config.DbType}
}

func (statements *SensorTypeStatements) SelectAll() string {
	if statements.DbType == "mysql" {
		return "SELECT Id, TypeName FROM SensorType"
	} else {
		return "SELECT id, name FROM sensortype"
	}

}

func (statements *SensorTypeStatements) SelectById() string {
	if statements.DbType == "mysql" {
		return "SELECT Id, TypeName FROM SensorType WHERE Id = ?"
	} else {
		return "SELECT id, name FROM sensortype WHERE id = $1"
	}
}

func (statements *SensorTypeStatements) SelectByParentId() string {
	return "no parent ID selection for SensorType"
}

func (statements *SensorTypeStatements) SelectBySecondParentId() string {
	return "no second parent ID selection for SensorType"
}

func (statements *SensorTypeStatements) Insert() string {
	if statements.DbType == "mysql" {
		return "INSERT INTO SensorType (TypeName) VALUES (?)"
	} else {
		return "INSERT INTO sensortype (name) VALUES ($1)"
	}
}

func (statements *SensorTypeStatements) Update() string {
	if statements.DbType == "mysql" {
		return "UPDATE SensorType SET TypeName = ? WHERE Id = ?"
	} else {
		return "UPDATE sensortype SET name = $1 WHERE id = $2"
	}
}

func (statements *SensorTypeStatements) Delete() string {
	if statements.DbType == "mysql" {
		return "DELETE FROM SensorType WHERE Id = ?"
	} else {
		return "DELETE FROM sensortype WHERE id = $1"
	}
}

func (statements *SensorTypeStatements) DeleteAll() string {
	if statements.DbType == "mysql" {
		return "DELETE FROM SensorType"
	} else {
		return "DELETE FROM sensortype"
	}
}

func (statements *SensorTypeStatements) DeleteByParentId() string {
	return "no parent ID deletion for SensorType"
}
