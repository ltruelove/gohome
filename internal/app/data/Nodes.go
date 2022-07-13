package data

import (
	"database/sql"

	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/ltruelove/gohome/internal/app/setup"
)

func VerifyNodeIdIsNew(nodeId int, db *sql.DB) bool {
	node := FetchNode(nodeId, db)
	return node.Id > 0
}

func FetchAllNodes(db *sql.DB) []models.Node {
	stmt, err := db.Prepare(`SELECT Id, Name FROM Node`)

	setup.CheckErr(err)
	var nodes []models.Node

	rows, err := stmt.Query()
	setup.CheckErr(err)

	for rows.Next() {
		var node models.Node
		rows.Scan(&node.Id,
			&node.Name)
		nodes = append(nodes, node)
	}
	defer stmt.Close()

	return nodes
}

func FetchNode(nodeId int, db *sql.DB) models.Node {
	stmt, err := db.Prepare("SELECT Id, Name FROM Node WHERE id = ?")
	setup.CheckErr(err)
	defer stmt.Close()

	var node models.Node

	err = stmt.QueryRow(nodeId).Scan(&node.Id,
		&node.Name)

	if err != sql.ErrNoRows {
		setup.CheckErr(err)
	}

	return node
}

func FetchNodeSensors(nodeId int, db *sql.DB) []models.NodeSensor {
	stmt, err := db.Prepare(`SELECT
		Id,
		SensorTypeId,
		Pin
		Name FROM NodeSenor WHERE id = ?`)
	setup.CheckErr(err)
	defer stmt.Close()

	var sensors []models.NodeSensor

	rows, err := stmt.Query()
	setup.CheckErr(err)

	for rows.Next() {
		var sensor models.NodeSensor
		sensor.NodeId = nodeId

		rows.Scan(&sensor.Id,
			&sensor.SensorTypeId,
			&sensor.Pin,
			&sensor.Name)

		sensors = append(sensors, sensor)
	}

	return sensors
}

func FetchNodeSwitches(nodeId int, db *sql.DB) []models.NodeSwitch {
	stmt, err := db.Prepare(`SELECT
		Id,
		SwitchTypeId,
		Pin
		Name FROM NodeSenor WHERE id = ?`)
	setup.CheckErr(err)
	defer stmt.Close()

	var nodeSwitches []models.NodeSwitch

	rows, err := stmt.Query()
	setup.CheckErr(err)

	for rows.Next() {
		var nodeSwitch models.NodeSwitch
		nodeSwitch.NodeId = nodeId

		rows.Scan(&nodeSwitch.Id,
			&nodeSwitch.SwitchTypeId,
			&nodeSwitch.Pin,
			&nodeSwitch.Name)

		nodeSwitches = append(nodeSwitches, nodeSwitch)
	}

	return nodeSwitches
}

func CreateNode(node *models.Node, db *sql.DB) {
	stmt, err := db.Prepare(`INSERT INTO Node
	(Id, Name, Mac)
	VALUES (?, ?, ?)`)

	setup.CheckErr(err)

	_, err = stmt.Exec(node.Id,
		node.Name,
		node.Mac)

	defer stmt.Close()

	setup.CheckErr(err)
}

func UpdateNode(node *models.Node, db *sql.DB) {
	stmt, err := db.Prepare(`UPDATE Node
	SET Name = ?, Mac = ?
	WHERE id = ?`)

	setup.CheckErr(err)

	_, err = stmt.Exec(node.Name,
		node.Mac,
		node.Id)

	defer stmt.Close()

	setup.CheckErr(err)
}

func DeleteNode(nodeId int, db *sql.DB) {
	stmt, err := db.Prepare(`DELETE FROM Node
	WHERE id = ?`)

	setup.CheckErr(err)

	_, err = stmt.Exec(nodeId)

	defer stmt.Close()

	setup.CheckErr(err)
}
