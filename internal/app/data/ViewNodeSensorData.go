package data

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ltruelove/gohome/internal/app/models"
)

const defaultViewNodeSensorDataSelect string = `SELECT
	Id,
	NodeId,
	ViewId,
	NodeSensorId,
	Name
	FROM View`

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

	query := fmt.Sprintf("%s %s", defaultViewNodeSensorDataSelect, "WHERE id = ?")
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
	stmt, err := db.Prepare(`INSERT INTO ViewNodeSensorData
	(NodeId, ViewId, NodeSensorId, Name)
	VALUES (?, ?, ?, ?)`)

	if err != nil {
		log.Println("Error preparing create node sensor data sql")
		return err
	}

	result, err := stmt.Exec(item.NodeId,
		item.ViewId,
		item.NodeSensorId,
		item.Name)

	if err != nil {
		log.Println("Error creating node sensor data")
		return err
	}

	lastInsertId, err := result.LastInsertId()

	if err != nil {
		log.Println("Error getting the id of the inserted view node sensor")
		return err
	}

	item.Id = int(lastInsertId)

	defer stmt.Close()

	return nil
}

func UpdateViewNodeSensorData(item *models.ViewNodeSensorData, db *sql.DB) error {
	stmt, err := db.Prepare(`UPDATE ViewNodeSensorData
	SET NodeId = ?, ViewId = ?, NodeSensorId = ?, Name = ?
	WHERE id = ?`)

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
	stmt, err := db.Prepare(`DELETE FROM ViewNodeSensorData
	WHERE id = ?`)

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
