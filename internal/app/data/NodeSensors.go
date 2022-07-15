package data

import (
	"database/sql"
	"log"

	"github.com/ltruelove/gohome/internal/app/models"
)

func VerifyNodeSensorIdIsNew(nodeId int, db *sql.DB) (bool, error) {
	node, err := FetchNodeSensor(nodeId, db)
	if err != nil {
		log.Println("Error fetching node sensor")
		return false, err
	}

	return node.Id > 0, nil
}

func FetchAllNodeSensors(db *sql.DB) ([]models.NodeSensor, error) {
	stmt, err := db.Prepare(`SELECT
		Id,
		NodeId,
		SensorTypeId,
		Name,
		Pin FROM NodeSensor`)
	if err != nil {
		log.Println("Error preparing all node sensors sql")
		return nil, err
	}

	var sensors []models.NodeSensor

	rows, err := stmt.Query()
	if err != nil {
		log.Println("Error querying for all node sensors")
		return nil, err
	}

	defer stmt.Close()

	for rows.Next() {
		var sensor models.NodeSensor

		err := rows.Scan(&sensor.Id,
			&sensor.NodeId,
			&sensor.SensorTypeId,
			&sensor.Name,
			&sensor.Pin)

		if err != nil {
			log.Println("Error scanning node sensor")
			return nil, err
		}

		sensors = append(sensors, sensor)
	}

	return sensors, nil
}

func FetchNodeSensor(nodeId int, db *sql.DB) (models.NodeSensor, error) {
	var sensor models.NodeSensor

	stmt, err := db.Prepare(`SELECT
		Id,
		NodeId,
		SensorTypeId,
		Name
		Pin FROM NodeSensor WHERE id = ?`)
	if err != nil {
		log.Println("Error preparing fetch node sensor sql")
		return sensor, err
	}

	err = stmt.QueryRow(nodeId).Scan(&sensor.Id,
		&sensor.NodeId,
		&sensor.SensorTypeId,
		&sensor.Name,
		&sensor.Pin)
	if err != nil {
		log.Println("Error querying for node sensor")
		return sensor, err
	}

	defer stmt.Close()

	return sensor, nil
}

func CreateNodeSensor(item *models.NodeSensor, db *sql.DB) error {
	stmt, err := db.Prepare(`INSERT INTO NodeSensor
	(Id, NodeId, SensorTypeId, Name, Pin)
	VALUES (?, ?, ?, ?, ?)`)

	if err != nil {
		log.Println("Error preparing create node sensor sql")
		return err
	}

	result, err := stmt.Exec(item.Id,
		item.NodeId,
		item.SensorTypeId,
		item.Name,
		item.Pin)

	if err != nil {
		log.Println("Error creating node sensor")
		return err
	}

	defer stmt.Close()

	lastInsertId, err := result.LastInsertId()

	if err != nil {
		log.Println("Error getting the id of the inserted node sensor")
		return err
	}

	item.Id = int(lastInsertId)

	return nil
}

func UpdateNodeSensor(sensor *models.NodeSensor, db *sql.DB) error {
	stmt, err := db.Prepare(`UPDATE NodeSensor
	SET NodeId = ?, SensorTypeId = ?, Name = ?, Pin = ?
	WHERE id = ?`)

	if err != nil {
		log.Println("Error preparing update node sensor sql")
		return err
	}

	_, err = stmt.Exec(sensor.NodeId,
		sensor.SensorTypeId,
		sensor.Name,
		sensor.Pin,
		sensor.Id)

	if err != nil {
		log.Println("Error updating node sensor")
		return err
	}

	defer stmt.Close()

	return nil
}

func DeleteNodeSensor(sensorId int, db *sql.DB) error {
	stmt, err := db.Prepare(`DELETE FROM NodeSensor
	WHERE id = ?`)

	if err != nil {
		log.Println("Error preparing delete node sensor sql")
		return err
	}

	_, err = stmt.Exec(sensorId)

	if err != nil {
		log.Println("Error deleting node sensor")
		return err
	}

	defer stmt.Close()

	return nil
}
