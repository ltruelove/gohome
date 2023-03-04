package data

import (
	"database/sql"
	"log"

	"github.com/ltruelove/gohome/internal/app/dto"
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
	stmt, err := db.Prepare(`SELECT id, name, ipaddress, mac FROM controlpoint`)
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
	stmt, err := db.Prepare(`SELECT id, name, ipaddress, mac
	FROM controlpoint AS c
	WHERE  (SELECT COUNT(id) FROM controlpointnodes WHERE controlpointid = c.id) < 20`)
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

func FetchAllControlPointNodes(controlPointId int, db *sql.DB) ([]dto.ControlPointNode, error) {
	stmt, err := db.Prepare(`SELECT
			node.id,
			node.name,
			node.mac,
			cpn.id AS relationid
		FROM controlpointnodes AS cpn
		INNER JOIN node ON node.id = cpn.nodeid
		WHERE cpn.controlpointid = $1`)

	if err != nil {
		log.Println("Error preparing all control point nodes sql")
		return nil, err
	}

	var nodes []dto.ControlPointNode

	rows, err := stmt.Query(controlPointId)
	if err != nil {
		log.Println("Error querying for all control point nodes")
		return nil, err
	}

	defer stmt.Close()

	for rows.Next() {
		var record dto.ControlPointNode

		rows.Scan(&record.Id,
			&record.Name,
			&record.Mac,
			&record.RelationId)

		nodes = append(nodes, record)
	}

	return nodes, nil
}

func FetchControlPoint(controlPointId int, db *sql.DB) (models.ControlPoint, error) {
	var controlPoint models.ControlPoint

	stmt, err := db.Prepare("SELECT id, name, ipaddress, mac FROM controlpoint WHERE id = $1")
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
	cp.id,
	cp.name,
	cp.ipaddress,
	cp.mac FROM controlpointnodes AS cpn
	INNER JOIN controlpoint AS cp ON cp.id = cpn.controlpointid
	WHERE cpn.nodeid = $1`)

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

	stmt, err := db.Prepare("SELECT id, name, ipaddress, mac FROM controlpoint WHERE mac = $1")
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
	stmt, err := db.Prepare(`INSERT INTO controlpoint
	(name, ipaddress, mac)
	VALUES ($1, $2, $3) RETURNING id`)

	if err != nil {
		log.Println("Error preparing create control point sql")
		return err
	}

	lastInsertId := 0

	err = stmt.QueryRow(controlPoint.Name,
		controlPoint.IpAddress,
		controlPoint.Mac).Scan(&lastInsertId)

	if err != nil {
		log.Println("Error creating control point")
		return err
	}

	defer stmt.Close()

	if err != nil {
		log.Println("Error getting the id of the inserted control point")
		return err
	}

	controlPoint.Id = int(lastInsertId)

	return nil
}

func UpdateControlPointIp(controlPoint *models.ControlPoint, db *sql.DB) error {
	stmt, err := db.Prepare(`UPDATE controlpoint
	SET ipaddress = $1
	WHERE id = $2`)

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
	stmt, err := db.Prepare(`UPDATE controlpoint
	SET name = $1, ipaddress = $2, mac = $3
	WHERE id = $4`)

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
	stmt, err := db.Prepare(`DELETE FROM controlpoint
	WHERE id = $1`)

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
	stmt, err := db.Prepare(`INSERT INTO controlpointnodes
	(controlpointid, nodeid)
	VALUES ($1, $2) RETURNING id`)

	if err != nil {
		log.Println("Error preparing create control point node sql")
		return err
	}

	lastInsertId := 0

	err = stmt.QueryRow(controlPointNode.ControlPointId,
		controlPointNode.NodeId).Scan(&lastInsertId)

	if err != nil {
		log.Println("Error adding node to control point ")
		return err
	}

	defer stmt.Close()

	if err != nil {
		log.Println("Error getting the id of the inserted control point node")
		return err
	}

	controlPointNode.Id = int(lastInsertId)

	return nil
}

func RemoveNodeFromControlPoint(nodeId int, db *sql.DB) error {
	stmt, err := db.Prepare(`DELETE FROM controlpointnodes
	WHERE nodeid = $1`)

	if err != nil {
		log.Println("Error preparing delete control point node sql")
		return err
	}

	_, err = stmt.Exec(nodeId)

	if err != nil {
		log.Println("Error removing the node from the control point ")
		return err
	}

	defer stmt.Close()

	return nil
}

func DeleteControlPointNode(id int, db *sql.DB) error {
	stmt, err := db.Prepare(`DELETE FROM controlpointnodes
	WHERE id = $1`)

	if err != nil {
		log.Println("Error preparing delete control point node sql")
		return err
	}

	_, err = stmt.Exec(id)

	if err != nil {
		log.Println("Error removing the node from the control point ")
		return err
	}

	defer stmt.Close()

	return nil
}

func DeleteControlPointNodeByNodeId(nodeId int, db *sql.DB) error {
	stmt, err := db.Prepare(`DELETE FROM controlpointnodes
	WHERE nodeid = $1`)

	if err != nil {
		log.Println("Error preparing delete control point node sql")
		return err
	}

	_, err = stmt.Exec(nodeId)

	if err != nil {
		log.Println("Error removing the node from the control point ")
		return err
	}

	defer stmt.Close()

	return nil
}
