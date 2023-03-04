package data

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ltruelove/gohome/internal/app/models"
)

const defaultViewNodeSensorDataSelect string = `SELECT
	id,
	nodeid,
	viewid,
	nodesensorid,
	name
	FROM view`

func FetchAllViewNodeSensorData(db *sql.DB) ([]models.ViewNodeSensorData, error) {
	stmt, err := db.Prepare(defaultViewNodeSensorDataSelect)

	if err != nil {
		log.Println("Error preparing all node sensor data sql")
		return nil, err
	}
	var listData []models.ViewNodeSensorData

	rows, err := stmt.Query()
	if err != nil {
		log.Println("Error querying for all node sensor data")
		return nil, err
	}

	for rows.Next() {
		var item models.ViewNodeSensorData
		err := rows.Scan(&item.Id,
			&item.NodeId,
			&item.ViewId,
			&item.NodeSensorId,
			&item.Name)

		if err != nil {
			log.Println("Error scanning node sensor data")
			return nil, err
		}
		listData = append(listData, item)
	}
	defer stmt.Close()

	return listData, nil
}

func FetchViewNodeSensorData(id int, db *sql.DB) (models.ViewNodeSensorData, error) {
	var item models.ViewNodeSensorData

	query := fmt.Sprintf("%s %s", defaultViewNodeSensorDataSelect, "WHERE id = $1")
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println("Error preparing fetch node sensor data sql")
		return item, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&item.Id,
		&item.NodeId,
		&item.ViewId,
		&item.NodeSensorId,
		&item.Name)

	if err != nil {
		log.Println("Error querying for node sensor data")
		return item, err
	}

	return item, nil
}

func CreateViewNodeSensorData(item *models.ViewNodeSensorData, db *sql.DB) error {
	stmt, err := db.Prepare(`INSERT INTO viewnodesensordata
	(nodeid, viewid, nodesensorid, name)
	VALUES ($1, $2, $3, $4) RETURNING id`)

	if err != nil {
		log.Println("Error preparing create node sensor data sql")
		return err
	}

	lastInsertId := 0

	err = stmt.QueryRow(item.NodeId,
		item.ViewId,
		item.NodeSensorId,
		item.Name).Scan((&lastInsertId))

	if err != nil {
		log.Println("Error creating node sensor data")
		return err
	}

	if err != nil {
		log.Println("Error getting the id of the inserted view node sensor")
		return err
	}

	item.Id = int(lastInsertId)

	defer stmt.Close()

	return nil
}

func UpdateViewNodeSensorData(item *models.ViewNodeSensorData, db *sql.DB) error {
	stmt, err := db.Prepare(`UPDATE viewnodesensordata
	SET nodeid = $1, viewid = $2, nodesensorid = $3, name = $4
	WHERE id = $5`)

	if err != nil {
		log.Println("Error preparing update node sensor data sql")
		return err
	}

	_, err = stmt.Exec(item.NodeId,
		item.ViewId,
		item.NodeSensorId,
		item.Name,
		item.Id)

	if err != nil {
		log.Println("Error updating node sensor data")
		return err
	}

	defer stmt.Close()

	return nil
}

func DeleteViewNodeSensorData(id int, db *sql.DB) error {
	stmt, err := db.Prepare(`DELETE FROM viewnodesensordata
	WHERE id = $1`)

	if err != nil {
		log.Println("Error preparing delete node sensor data sql")
		return err
	}

	_, err = stmt.Exec(id)

	if err != nil {
		log.Println("Error deleting node sensor data")
		return err
	}

	defer stmt.Close()

	return nil
}
