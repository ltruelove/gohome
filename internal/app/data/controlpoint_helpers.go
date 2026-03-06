package data

import (
	"database/sql"
	"log"

	"github.com/ltruelove/gohome/internal/app/models"
)

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
