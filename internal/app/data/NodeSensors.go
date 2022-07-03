package data

import (
	"database/sql"

	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/ltruelove/gohome/internal/app/setup"
)

func VerifyNodeSensorIdIsNew(nodeId string, db *sql.DB) bool {
	node := FetchNodeSensor(nodeId, db)
	return node.Id > 0
}

func FetchAllNodeSensors(db *sql.DB) []models.NodeSensor {
	stmt, err := db.Prepare(`SELECT
		Id,
		NodeId,
		SensorTypeId,
		Name,
		Pin FROM NodeSensor`)

	setup.CheckErr(err)
	var sensors []models.NodeSensor

	rows, err := stmt.Query()
	setup.CheckErr(err)

	for rows.Next() {
		var sensor models.NodeSensor

		rows.Scan(&sensor.Id,
			&sensor.NodeId,
			&sensor.SensorTypeId,
			&sensor.Name,
			&sensor.Pin)

		sensors = append(sensors, sensor)
	}
	defer stmt.Close()

	return sensors
}

func FetchNodeSensor(nodeId string, db *sql.DB) models.NodeSensor {
	stmt, err := db.Prepare(`SELECT
		Id,
		NodeId,
		SensorTypeId,
		Name
		Pin FROM NodeSensor WHERE id = ?`)
	setup.CheckErr(err)
	defer stmt.Close()

	var sensor models.NodeSensor

	err = stmt.QueryRow(nodeId).Scan(&sensor.Id,
		&sensor.NodeId,
		&sensor.SensorTypeId,
		&sensor.Name,
		&sensor.Pin)

	if err != sql.ErrNoRows {
		setup.CheckErr(err)
	}

	return sensor
}

func CreateNodeSensor(sensor *models.NodeSensor, db *sql.DB) {
	stmt, err := db.Prepare(`INSERT INTO NodeSensor
	(Id, NodeId, SensorTypeId, Name, Pin)
	VALUES (?, ?, ?, ?, ?)`)

	setup.CheckErr(err)

	_, err = stmt.Exec(sensor.Id,
		sensor.NodeId,
		sensor.SensorTypeId,
		sensor.Name,
		sensor.Pin)

	defer stmt.Close()

	setup.CheckErr(err)
}

func UpdateNodeSensor(sensor *models.NodeSensor, db *sql.DB) {
	stmt, err := db.Prepare(`UPDATE NodeSensor
	SET NodeId = ?, SensorTypeId = ?, Name = ?, Pin = ?
	WHERE id = ?`)

	setup.CheckErr(err)

	_, err = stmt.Exec(sensor.NodeId,
		sensor.SensorTypeId,
		sensor.Name,
		sensor.Pin,
		sensor.Id)

	defer stmt.Close()

	setup.CheckErr(err)
}

func DeleteNodeSensor(sensorId string, db *sql.DB) {
	stmt, err := db.Prepare(`DELETE FROM NodeSensor
	WHERE id = ?`)

	setup.CheckErr(err)

	_, err = stmt.Exec(sensorId)

	defer stmt.Close()

	setup.CheckErr(err)
}
