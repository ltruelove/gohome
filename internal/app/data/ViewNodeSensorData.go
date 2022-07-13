package data

import (
	"database/sql"
	"fmt"

	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/ltruelove/gohome/internal/app/setup"
)

const defaultViewNodeSensorDataSelect string = `SELECT
	Id,
	NodeId,
	ViewId,
	NodeSensorId,
	SensorTypeDataId,
	Name
	FROM View`

func FetchAllViewNodeSensorData(db *sql.DB) []models.ViewNodeSensorData {
	stmt, err := db.Prepare(defaultViewNodeSensorDataSelect)

	setup.CheckErr(err)
	var listData []models.ViewNodeSensorData

	rows, err := stmt.Query()
	setup.CheckErr(err)

	for rows.Next() {
		var item models.ViewNodeSensorData
		rows.Scan(&item.Id,
			&item.NodeId,
			&item.ViewId,
			&item.NodeSensorId,
			&item.SensorTypeDataId,
			&item.Name)
		listData = append(listData, item)
	}
	defer stmt.Close()

	return listData
}

func FetchViewNodeSensorData(id int, db *sql.DB) models.ViewNodeSensorData {
	query := fmt.Sprintf("%s %s", defaultViewNodeSensorDataSelect, "WHERE id = ?")
	stmt, err := db.Prepare(query)

	setup.CheckErr(err)
	defer stmt.Close()

	var item models.ViewNodeSensorData

	err = stmt.QueryRow(id).Scan(&item.Id,
		&item.NodeId,
		&item.ViewId,
		&item.NodeSensorId,
		&item.SensorTypeDataId,
		&item.Name)

	if err != sql.ErrNoRows {
		setup.CheckErr(err)
	}

	return item
}

func CreateViewNodeSensorData(item *models.ViewNodeSensorData, db *sql.DB) {
	stmt, err := db.Prepare(`INSERT INTO ViewNodeSensorData
	(NodeId, ViewId, NodeSensorId, SensorTypeDataId, Name)
	VALUES (?, ?, ?, ?, ?, ?)`)

	setup.CheckErr(err)

	_, err = stmt.Exec(item.NodeId,
		item.ViewId,
		item.NodeSensorId,
		item.SensorTypeDataId,
		item.Name)

	defer stmt.Close()

	setup.CheckErr(err)
}

func UpdateViewNodeSensorData(item *models.ViewNodeSensorData, db *sql.DB) {
	stmt, err := db.Prepare(`UPDATE ViewNodeSensorData
	SET NodeId = ?, ViewId = ?, NodeSensorId = ?, SensorTypeDataId = ?, Name = ?
	WHERE id = ?`)

	setup.CheckErr(err)

	_, err = stmt.Exec(item.NodeId,
		item.ViewId,
		item.NodeSensorId,
		item.SensorTypeDataId,
		item.Name,
		item.Id)

	defer stmt.Close()

	setup.CheckErr(err)
}

func DeleteViewNodeSensorData(id int, db *sql.DB) {
	stmt, err := db.Prepare(`DELETE FROM ViewNodeSensorData
	WHERE id = ?`)

	setup.CheckErr(err)

	_, err = stmt.Exec(id)

	defer stmt.Close()

	setup.CheckErr(err)
}
