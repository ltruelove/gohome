package data

type Statements struct {
	DbType string
}

func (statements *Statements) SelectAllSensorTypes() string {
	if statements.DbType == "mysql" {
		return "SELECT Id, TypeName FROM SensorType"
	} else {
		return "SELECT id, name FROM sensortype"
	}

}

func (statements *Statements) SelectSensorTypeById() string {
	if statements.DbType == "mysql" {
		return "SELECT Id, TypeName FROM SensorType WHERE Id = ?"
	} else {
		return "SELECT id, name FROM sensortype WHERE id = $1"
	}
}

func (statements *Statements) SelectSensorTypeData() string {
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
