package data

import (
	"database/sql"
	"log"

	"github.com/ltruelove/gohome/internal/app/models"
)

func VerifyControlPointIdIsNew(nodeId int, db *sql.DB) (bool, error) {
	node, err := FetchControlPoint(nodeId, db)
	if err != nil {
		log.Println("Error fetching control point")
		return false, err
	}

	return node.Id < 1, nil
}

func FetchAllControlPoints(db *sql.DB) ([]models.ControlPoint, error) {
	stmt, err := db.Prepare(`SELECT Id, Name, IpAddress, Mac FROM ControlPoint`)
	if err != nil {
		log.Println("Error preparing all control points sql")
		return nil, err
	}

	var controlPoints []models.ControlPoint

	rows, err := stmt.Query()
	if err != nil {
		log.Println("Error querying for all control points")
		return nil, err
	}

	defer stmt.Close()

	for rows.Next() {
		var controlPoint models.ControlPoint

		rows.Scan(&controlPoint.Id,
			&controlPoint.Name,
			&controlPoint.IpAddress,
			&controlPoint.Mac)

		controlPoints = append(controlPoints, controlPoint)
	}

	return controlPoints, nil
}

func FetchAllAvailableControlPoints(db *sql.DB) ([]models.ControlPoint, error) {
	// Only select control points that aren't maxed out for nodes, 20 is the max
	stmt, err := db.Prepare(`SELECT Id, Name, IpAddress, Mac
	FROM ControlPoint AS c
	WHERE  (SELECT COUNT(Id) FROM ControlPointNodes WHERE ControlPointId = c.Id) < 20`)
	if err != nil {
		log.Println("Error preparing all control points sql")
		return nil, err
	}

	var controlPoints []models.ControlPoint

	rows, err := stmt.Query()
	if err != nil {
		log.Println("Error querying for all control points")
		return nil, err
	}

	defer stmt.Close()

	for rows.Next() {
		var controlPoint models.ControlPoint

		rows.Scan(&controlPoint.Id,
			&controlPoint.Name,
			&controlPoint.IpAddress,
			&controlPoint.Mac)

		controlPoints = append(controlPoints, controlPoint)
	}

	return controlPoints, nil
}

func FetchAllControlPointNodes(db *sql.DB, controlPointId int) ([]models.Node, error) {
	stmt, err := db.Prepare(`SELECT
			Node.Id,
			Node.Name,
			Node.Mac
		FROM ControlPointNodes
		INNER JOIN Node ON Node.Id = ControlPointNodes.NodeId
		WHERE ControlPointNodes.ControlPointId = ?`)
	if err != nil {
		log.Println("Error preparing all control point nodes sql")
		return nil, err
	}

	var nodes []models.Node

	rows, err := stmt.Query(controlPointId)
	if err != nil {
		log.Println("Error querying for all control point nodes")
		return nil, err
	}

	defer stmt.Close()

	for rows.Next() {
		var node models.Node

		rows.Scan(&node.Id,
			&node.Name,
			&node.Mac)

		nodes = append(nodes, node)
	}

	return nodes, nil
}

func FetchControlPoint(controlPointId int, db *sql.DB) (models.ControlPoint, error) {
	var controlPoint models.ControlPoint

	stmt, err := db.Prepare("SELECT Id, Name, IpAddress, Mac FROM ControlPoint WHERE id = ?")
	if err != nil {
		log.Printf("Error preparing fetch control point sql: %v", err)
		return controlPoint, err
	}

	err = stmt.QueryRow(controlPointId).Scan(&controlPoint.Id,
		&controlPoint.Name,
		&controlPoint.IpAddress,
		&controlPoint.Mac)

	defer stmt.Close()

	if err != nil {
		log.Println("Error querying for control point")
		return controlPoint, err
	}

	return controlPoint, nil
}

func FetchControlPointByNode(nodeId int, db *sql.DB) (models.ControlPoint, error) {
	var controlPoint models.ControlPoint

	stmt, err := db.Prepare(`SELECT
	cp.Id,
	cp.Name,
	cp.IpAddress,
	cp.Mac FROM ControlPointNodes AS cpn
	INNER JOIN ControlPoint AS cp ON cp.Id = cpn.ControlPointId
	WHERE cpn.NodeId = ?`)

	if err != nil {
		log.Printf("Error preparing fetch control point by node sql: %v", err)
		return controlPoint, err
	}

	err = stmt.QueryRow(nodeId).Scan(&controlPoint.Id,
		&controlPoint.Name,
		&controlPoint.IpAddress,
		&controlPoint.Mac)

	defer stmt.Close()

	if err != nil {
		log.Println("Error querying for control point")
		return controlPoint, err
	}

	return controlPoint, nil
}

func FetchControlPointByMac(macAddress string, db *sql.DB) (models.ControlPoint, error) {
	var controlPoint models.ControlPoint

	stmt, err := db.Prepare("SELECT Id, Name, IpAddress, Mac FROM ControlPoint WHERE Mac = ?")
	if err != nil {
		log.Printf("Error preparing fetch control point by Mac sql: %v", err)
		return controlPoint, err
	}

	err = stmt.QueryRow(macAddress).Scan(&controlPoint.Id,
		&controlPoint.Name,
		&controlPoint.IpAddress,
		&controlPoint.Mac)

	defer stmt.Close()

	if err != nil {
		log.Println("Error querying for control point by Mac")
		return controlPoint, err
	}

	return controlPoint, nil
}

func CreateControlPoint(controlPoint *models.ControlPoint, db *sql.DB) error {
	stmt, err := db.Prepare(`INSERT INTO ControlPoint
	(Name, IpAddress, Mac)
	VALUES (?, ?, ?)`)

	if err != nil {
		log.Println("Error preparing create control point sql")
		return err
	}

	result, err := stmt.Exec(controlPoint.Name,
		controlPoint.IpAddress,
		controlPoint.Mac)

	if err != nil {
		log.Println("Error creating control point")
		return err
	}

	defer stmt.Close()

	lastInsertId, err := result.LastInsertId()

	if err != nil {
		log.Println("Error getting the id of the inserted control point")
		return err
	}

	controlPoint.Id = int(lastInsertId)

	return nil
}

func UpdateControlPointIp(controlPoint *models.ControlPoint, db *sql.DB) error {
	stmt, err := db.Prepare(`UPDATE ControlPoint
	SET IpAddress = ?
	WHERE id = ?`)

	if err != nil {
		log.Println("Error preparing update control point IP Address sql")
		return err
	}

	_, err = stmt.Exec(controlPoint.IpAddress,
		controlPoint.Id)

	if err != nil {
		log.Println("Error updating control point IP Address")
		return err
	}

	defer stmt.Close()

	return nil
}

func UpdateControlPoint(controlPoint *models.ControlPoint, db *sql.DB) error {
	stmt, err := db.Prepare(`UPDATE ControlPoint
	SET Name = ?, IpAddress = ?, Mac = ?
	WHERE id = ?`)

	if err != nil {
		log.Println("Error preparing update control point sql")
		return err
	}

	_, err = stmt.Exec(controlPoint.Name,
		controlPoint.IpAddress,
		controlPoint.Mac,
		controlPoint.Id)

	if err != nil {
		log.Println("Error updating control point")
		return err
	}

	defer stmt.Close()

	return nil
}

func DeleteControlPoint(controlPointId int, db *sql.DB) error {
	stmt, err := db.Prepare(`DELETE FROM ControlPoint
	WHERE id = ?`)

	if err != nil {
		log.Println("Error preparing delete control point sql")
		return err
	}

	_, err = stmt.Exec(controlPointId)

	if err != nil {
		log.Println("Error deleting control point")
		return err
	}

	defer stmt.Close()

	return nil
}

func AddNodeToControlPoint(controlPointNode *models.ControlPointNode, db *sql.DB) error {
	stmt, err := db.Prepare(`INSERT INTO ControlPointNodes
	(ControlPointId, NodeId)
	VALUES (?, ?)`)

	if err != nil {
		log.Println("Error preparing create control point node sql")
		return err
	}

	result, err := stmt.Exec(controlPointNode.ControlPointId,
		controlPointNode.NodeId)

	if err != nil {
		log.Println("Error adding node to control point ")
		return err
	}

	defer stmt.Close()

	lastInsertId, err := result.LastInsertId()

	if err != nil {
		log.Println("Error getting the id of the inserted control point node")
		return err
	}

	controlPointNode.Id = int(lastInsertId)

	return nil
}
