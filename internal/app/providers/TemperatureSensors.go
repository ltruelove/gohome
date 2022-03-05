package providers

import (
	"github.com/ltruelove/gohome/internal/app/data"
	"github.com/ltruelove/gohome/internal/app/setup"
)

const tableName string = "TemperatureSensors"

func VerifyTemperatureSensorIdIsNew(sensorId string) bool {
	sensor := FetchTemperatureSensor(sensorId)
	return len(sensor.SensorId) == 0
}

func FetchAllTemperatureSensors() []data.TemperatureSensor {
	stmt, err := setup.DB.Prepare(`SELECT id, name, isGarage, ipAddress
	FROM ?`)

	setup.CheckErr(err)
	var sensors []data.TemperatureSensor

	rows, err := stmt.Query(tableName)
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

func FetchTemperatureSensor(sensorId string) data.TemperatureSensor {
	stmt, err := setup.DB.Prepare(`SELECT id, name, isGarage, ipAddress
	FROM ?
	WHERE id = ?`)

	setup.CheckErr(err)
	var sensor data.TemperatureSensor

	err = stmt.QueryRow(tableName, sensorId).Scan(&sensor.SensorId,
		&sensor.Name,
		&sensor.IsGarage,
		&sensor.IpAddress)
	defer stmt.Close()

	setup.CheckErr(err)

	return sensor
}

func AddNewTemperatureSensor(sensor *data.TemperatureSensor) {
	stmt, err := setup.DB.Prepare(`INSERT INTO ?
	(id, name, isGarage, ipAddress)
	VALUES (?, ?, ?, ?)`)

	setup.CheckErr(err)

	_, err = stmt.Query(tableName,
		sensor.SensorId,
		sensor.Name,
		sensor.IsGarage,
		sensor.IpAddress)

	defer stmt.Close()

	setup.CheckErr(err)
}

func UpdateTemperatureSensor(sensor *data.TemperatureSensor) {
	stmt, err := setup.DB.Prepare(`UPDATE ?
	SET name = ?, isGarage = ?, ipAddress = ?
	WHERE id = ?`)

	setup.CheckErr(err)

	_, err = stmt.Query(tableName,
		sensor.Name,
		sensor.IsGarage,
		sensor.IpAddress,
		sensor.SensorId)

	defer stmt.Close()

	setup.CheckErr(err)
}

func DeleteTemperatureSensor(sensorId string) {
	stmt, err := setup.DB.Prepare(`DELETE FROM ?
	WHERE id = ?`)

	setup.CheckErr(err)

	_, err = stmt.Query(tableName, sensorId)

	defer stmt.Close()

	setup.CheckErr(err)
}
