package data

import (
	"database/sql"

	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/ltruelove/gohome/internal/app/setup"
)

func VerifyControlPointIdIsNew(nodeId int, db *sql.DB) bool {
	node := FetchControlPoint(nodeId, db)
	return node.Id > 0
}

func FetchAllControlPoints(db *sql.DB) []models.ControlPoint {
	stmt, err := db.Prepare(`SELECT Id, Name, IpAddress, Mac FROM ControlPoint`)

	setup.CheckErr(err)
	var controlPoints []models.ControlPoint

	rows, err := stmt.Query()
	setup.CheckErr(err)

	for rows.Next() {
		var controlPoint models.ControlPoint

		rows.Scan(&controlPoint.Id,
			&controlPoint.Name,
			&controlPoint.IpAddress,
			&controlPoint.Mac)

		controlPoints = append(controlPoints, controlPoint)
	}
	defer stmt.Close()

	return controlPoints
}

func FetchAllControlPointNodes(db *sql.DB, controlPointId int) []models.Node {
	stmt, err := db.Prepare(`SELECT
			Node.Id,
			Node.Name,
			Node.Mac
		FROM ControlPointNodes
		INNER JOIN Node ON Node.Id = ControlPointNodes.NodeId
		WHERE ControlPointNodes.ControlPointId = ?`)

	setup.CheckErr(err)
	var nodes []models.Node

	rows, err := stmt.Query(controlPointId)
	setup.CheckErr(err)

	for rows.Next() {
		var node models.Node

		rows.Scan(&node.Id,
			&node.Name,
			&node.Mac)

		nodes = append(nodes, node)
	}
	defer stmt.Close()

	return nodes
}

func FetchControlPoint(controlPointId int, db *sql.DB) models.ControlPoint {
	stmt, err := db.Prepare("SELECT Id, Name, IpAddress, Mac FROM ControlPoint WHERE id = ?")
	setup.CheckErr(err)
	defer stmt.Close()

	var controlPoint models.ControlPoint

	err = stmt.QueryRow(controlPointId).Scan(&controlPoint.Id,
		&controlPoint.Name,
		&controlPoint.IpAddress,
		&controlPoint.Mac)

	if err != sql.ErrNoRows {
		setup.CheckErr(err)
	}

	return controlPoint
}

func CreateControlPoint(controlPoint *models.ControlPoint, db *sql.DB) {
	stmt, err := db.Prepare(`INSERT INTO ControlPoint
	(Id, Name, IpAddress, Mac)
	VALUES (?, ?, ?, ?)`)

	setup.CheckErr(err)

	_, err = stmt.Exec(controlPoint.Id,
		controlPoint.Name,
		controlPoint.IpAddress,
		controlPoint.Mac)

	defer stmt.Close()

	setup.CheckErr(err)
}

func UpdateControlPoint(controlPoint *models.ControlPoint, db *sql.DB) {
	stmt, err := db.Prepare(`UPDATE ControlPoint
	SET Name = ?, IpAddress = ?, Mac = ?
	WHERE id = ?`)

	setup.CheckErr(err)

	_, err = stmt.Exec(controlPoint.Name,
		controlPoint.IpAddress,
		controlPoint.Mac,
		controlPoint.Id)

	defer stmt.Close()

	setup.CheckErr(err)
}

func DeleteControlPoint(controlPointId int, db *sql.DB) {
	stmt, err := db.Prepare(`DELETE FROM ControlPoint
	WHERE id = ?`)

	setup.CheckErr(err)

	_, err = stmt.Exec(controlPointId)

	defer stmt.Close()

	setup.CheckErr(err)
}
