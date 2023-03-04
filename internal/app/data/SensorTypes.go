package data

import (
	"database/sql"
	"log"

	"github.com/ltruelove/gohome/internal/app/models"
)

func FetchAllSensorTypes(db *sql.DB) ([]models.SensorType, error) {
	stmt, err := db.Prepare(`SELECT Id, Name
	FROM SensorType`)
	if err != nil {
		log.Println("Error preparing all sensor types sql")
		return nil, err
	}

	var sensors []models.SensorType

	rows, err := stmt.Query()
	if err != nil {
		log.Println("Error querying for all sensor types")
		return nil, err
	}

	for rows.Next() {
		var sensor models.SensorType
		rows.Scan(&sensor.Id,
			&sensor.Name)
		sensors = append(sensors, sensor)
	}
	defer stmt.Close()

	return sensors, nil
}

func FetchSensorType(sensorTypeId int, db *sql.DB) (models.SensorType, error) {
	var sensor models.SensorType

	stmt, err := db.Prepare("SELECT id, name FROM sensortype WHERE id = $1")
	if err != nil {
		log.Println("Error preparing the fetch sensor type sql")
		return sensor, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(sensorTypeId).Scan(&sensor.Id,
		&sensor.Name)

	if err != nil {
		log.Println("Error querying for the sensor type")
		return sensor, err
	}

	return sensor, nil
}

func FetchSensorTypeData(sensorTypeId int, db *sql.DB) ([]models.SensorTypeData, error) {
	stmt, err := db.Prepare(`SELECT
		id, 
		name, 
		valuetype 
		FROM sensortypedata 
		WHERE sensortypeid = $1`)
	if err != nil {
		log.Println("Error preparing the fetch sensor type data sql")
		return nil, err
	}
	defer stmt.Close()

	var sensorData []models.SensorTypeData

	rows, err := stmt.Query(sensorTypeId)
	if err != nil {
		log.Println("Error querying for the sensor type data")
		return nil, err
	}

	for rows.Next() {
		var sensor models.SensorTypeData
		sensor.SensorTypeId = sensorTypeId

		rows.Scan(&sensor.Id,
			&sensor.Name,
			&sensor.ValueType)
		sensorData = append(sensorData, sensor)
	}
	defer stmt.Close()

	return sensorData, nil
}
