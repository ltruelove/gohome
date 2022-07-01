package data

import (
	"database/sql"

	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/ltruelove/gohome/internal/app/setup"
)

func FetchAllSensorTypes(db *sql.DB) []models.SensorType {
	stmt, err := db.Prepare(`SELECT Id, Name
	FROM SensorTypes`)

	setup.CheckErr(err)
	var sensors []models.SensorType

	rows, err := stmt.Query()
	setup.CheckErr(err)

	for rows.Next() {
		var sensor models.SensorType
		rows.Scan(&sensor.Id,
			&sensor.Name)
		sensors = append(sensors, sensor)
	}
	defer stmt.Close()

	return sensors
}

func FetchSensorType(sensorTypeId int, db *sql.DB) models.SensorType {
	stmt, err := db.Prepare("SELECT Id, Name FROM SensorTypes WHERE id = ?")
	setup.CheckErr(err)
	defer stmt.Close()

	var sensor models.SensorType

	err = stmt.QueryRow(sensorTypeId).Scan(&sensor.Id,
		&sensor.Name)

	if err != sql.ErrNoRows {
		setup.CheckErr(err)
	}

	return sensor
}

func FetchSensorTypeData(sensorTypeId int, db *sql.DB) []models.SensorTypeData {
	stmt, err := db.Prepare(`SELECT
		Id, 
		Name, 
		ValueType 
		FROM SensorTypeData 
		WHERE SensorTypeId = ?`)
	setup.CheckErr(err)
	defer stmt.Close()
	var sensorData []models.SensorTypeData

	rows, err := stmt.Query(sensorTypeId)
	setup.CheckErr(err)

	for rows.Next() {
		var sensor models.SensorTypeData
		rows.Scan(&sensor.Id,
			&sensor.Name,
			&sensor.ValueType)
		sensorData = append(sensorData, sensor)
	}
	defer stmt.Close()

	return sensorData
}
