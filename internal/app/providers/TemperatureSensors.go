package providers

import (
	"database/sql"
	"fmt"

	"github.com/ltruelove/gohome/internal/app/data"
	"github.com/ltruelove/gohome/internal/app/setup"
)

func VerifyTemperatureSensorIdIsNew(sensorId string, db *sql.DB) bool {
	sensor := FetchTemperatureSensor(sensorId, db)
	return len(sensor.SensorId) == 0
}

func FetchAllTemperatureSensors(db *sql.DB) []data.TemperatureSensor {
	stmt, err := db.Prepare(`SELECT id, name, isGarage, ipAddress
	FROM TemperatureSensors`)

	setup.CheckErr(err)
	var sensors []data.TemperatureSensor

	rows, err := stmt.Query()
	setup.CheckErr(err)

	for rows.Next() {
		var sensor data.TemperatureSensor
		rows.Scan(&sensor.SensorId,
			&sensor.Name,
			&sensor.IsGarage,
			&sensor.IpAddress)
		sensors = append(sensors, sensor)
	}
	defer stmt.Close()

	return sensors
}

func FetchTemperatureSensor(sensorId string, db *sql.DB) data.TemperatureSensor {
	stmt, err := db.Prepare("SELECT id, name, isGarage, ipAddress FROM TemperatureSensors WHERE id = ?")
	setup.CheckErr(err)
	defer stmt.Close()

	var sensor data.TemperatureSensor

	err = stmt.QueryRow(sensorId).Scan(&sensor.SensorId,
		&sensor.Name,
		&sensor.IsGarage,
		&sensor.IpAddress)

	if err != sql.ErrNoRows {
		setup.CheckErr(err)
	}

	return sensor
}

func AddNewTemperatureSensor(sensor *data.TemperatureSensor, db *sql.DB) {
	fmt.Println("Preparing insert")
	stmt, err := db.Prepare(`INSERT INTO TemperatureSensors
	(id, name, isGarage, ipAddress)
	VALUES (?, ?, ?, ?)`)

	setup.CheckErr(err)

	fmt.Println("Executing insert")
	_, err = stmt.Exec(sensor.SensorId,
		sensor.Name,
		sensor.IsGarage,
		sensor.IpAddress)

	defer stmt.Close()

	fmt.Println("Checking insert")
	setup.CheckErr(err)
	fmt.Println("Insert succeeded")
}

func UpdateTemperatureSensor(sensor *data.TemperatureSensor, db *sql.DB) {
	stmt, err := db.Prepare(`UPDATE TemperatureSensors
	SET name = ?, isGarage = ?, ipAddress = ?
	WHERE id = ?`)

	setup.CheckErr(err)

	_, err = stmt.Exec(sensor.Name,
		sensor.IsGarage,
		sensor.IpAddress,
		sensor.SensorId)

	defer stmt.Close()

	setup.CheckErr(err)
}

func DeleteTemperatureSensor(sensorId string, db *sql.DB) {
	stmt, err := db.Prepare(`DELETE FROM TemperatureSensors
	WHERE id = ?`)

	setup.CheckErr(err)

	_, err = stmt.Exec(sensorId)

	defer stmt.Close()

	setup.CheckErr(err)
}
