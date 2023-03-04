package data

import (
	"database/sql"
	"log"

	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/ltruelove/gohome/internal/app/viewModels"
)

func VerifyNodeSensorIdIsNew(nodeId int, db *sql.DB) (bool, error) {
	node, err := FetchNodeSensor(nodeId, db)
	if err != nil {
		log.Println("Error fetching node sensor")
		return false, err
	}

	return node.Id > 0, nil
}

func FetchAllNodeSensors(db *sql.DB) ([]viewModels.NodeSensorVM, error) {
	stmt, err := db.Prepare(`SELECT
		ns.id,
		ns.nodeid,
		ns.sensortypeid,
		ns.name,
		ns.pin,
		ns.dhttype,
		st.name AS sensortypename
		FROM nodesensor AS ns
		INNER JOIN sensortype AS st ON st.id = ns.sensortypeid`)

	if err != nil {
		log.Println("Error preparing all node sensors sql")
		return nil, err
	}

	var sensors []viewModels.NodeSensorVM

	rows, err := stmt.Query()
	if err != nil {
		log.Println("Error querying for all node sensors")
		return nil, err
	}

	defer stmt.Close()

	for rows.Next() {
		var sensor viewModels.NodeSensorVM

		err := rows.Scan(&sensor.Id,
			&sensor.NodeId,
			&sensor.SensorTypeId,
			&sensor.Name,
			&sensor.Pin,
			&sensor.DHTType,
			&sensor.SensorTypeName)

		if err != nil {
			log.Println("Error scanning node sensor")
			return nil, err
		}

		sensors = append(sensors, sensor)
	}

	return sensors, nil
}

func FetchAllNodeSensorsByNode(nodeId int, db *sql.DB) ([]viewModels.NodeSensorVM, error) {
	stmt, err := db.Prepare(`SELECT
		ns.id,
		ns.nodeid,
		ns.sensortypeid,
		ns.name,
		ns.pin,
		ns.dhttype,
		st.name AS sensortypename
		FROM nodesensor AS ns
		INNER JOIN sensortype AS st ON st.id = ns.sensortypeid
		WHERE ns.nodeid = $1`)

	if err != nil {
		log.Println("Error preparing all node sensors by node sql")
		return nil, err
	}

	var sensors []viewModels.NodeSensorVM

	rows, err := stmt.Query(nodeId)
	if err != nil {
		log.Println("Error querying for all node sensors by node")
		return nil, err
	}

	defer stmt.Close()

	for rows.Next() {
		var sensor viewModels.NodeSensorVM

		err := rows.Scan(&sensor.Id,
			&sensor.NodeId,
			&sensor.SensorTypeId,
			&sensor.Name,
			&sensor.Pin,
			&sensor.DHTType,
			&sensor.SensorTypeName)

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
		id,
		nodeid,
		sensortypeid,
		name,
		pin,
		dhttype FROM nodesensor WHERE id = $1`)
	if err != nil {
		log.Println("Error preparing fetch node sensor sql")
		return sensor, err
	}

	err = stmt.QueryRow(nodeId).Scan(&sensor.Id,
		&sensor.NodeId,
		&sensor.SensorTypeId,
		&sensor.Name,
		&sensor.Pin,
		&sensor.DHTType)
	if err != nil {
		log.Println("Error querying for node sensor")
		return sensor, err
	}

	defer stmt.Close()

	return sensor, nil
}

func CreateNodeSensor(item *models.NodeSensor, db *sql.DB) error {
	stmt, err := db.Prepare(`INSERT INTO nodesensor
	(nodeid, sensortypeid, name, pin, dhttype)
	VALUES ($1, $2, $3, $4, $5) RETURNING id`)

	if err != nil {
		log.Println("Error preparing create node sensor sql")
		return err
	}

	lastInsertId := 0

	err = stmt.QueryRow(&item.NodeId,
		&item.SensorTypeId,
		&item.Name,
		&item.Pin,
		&item.DHTType).Scan(&lastInsertId)

	if err != nil {
		log.Println("Error creating node sensor")
		return err
	}

	defer stmt.Close()

	if err != nil {
		log.Println("Error getting the id of the inserted node sensor")
		return err
	}

	item.Id = int(lastInsertId)

	return nil
}

func UpdateNodeSensor(sensor *models.NodeSensor, db *sql.DB) error {
	stmt, err := db.Prepare(`UPDATE nodesensor
	SET nodeid = $1, sensortypeid = $2, name = $3, pin = $4, dhttype = $5
	WHERE id = $6`)

	if err != nil {
		log.Println("Error preparing update node sensor sql")
		return err
	}

	_, err = stmt.Exec(sensor.NodeId,
		sensor.SensorTypeId,
		sensor.Name,
		sensor.Pin,
		sensor.Id,
		sensor.DHTType)

	if err != nil {
		log.Println("Error updating node sensor")
		return err
	}

	defer stmt.Close()

	return nil
}

func DeleteNodeSensor(sensorId int, db *sql.DB) error {
	stmt, err := db.Prepare(`DELETE FROM nodesensor
	WHERE id = $1`)

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

func DeleteAllNodeSensors(nodeId int, db *sql.DB) error {
	stmt, err := db.Prepare(`DELETE FROM nodesensor
	WHERE nodeid = $1`)

	if err != nil {
		log.Println("Error preparing delete node sensor sql")
		return err
	}

	_, err = stmt.Exec(nodeId)

	if err != nil {
		log.Println("Error deleting node sensor")
		return err
	}

	defer stmt.Close()

	return nil
}
