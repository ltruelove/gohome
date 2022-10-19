package data

import (
	"database/sql"
	"log"

	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/ltruelove/gohome/internal/app/setup"
	"github.com/ltruelove/gohome/internal/app/viewModels"
)

func VerifyNodeIdIsNew(nodeId int, db *sql.DB) (bool, error) {
	item, err := FetchNode(nodeId, db)
	if err != nil {
		log.Println("Error fetching node switch")
		return false, err
	}

	return item.Id < 1, nil
}

func FetchAllNodes(db *sql.DB) ([]models.Node, error) {
	stmt, err := db.Prepare(`SELECT Id, Name, Mac FROM Node`)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	var nodes []models.Node

	rows, err := stmt.Query()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		var node models.Node
		rows.Scan(&node.Id,
			&node.Name,
			&node.Mac)
		nodes = append(nodes, node)
	}
	defer stmt.Close()

	return nodes, nil
}

func FetchNode(nodeId int, db *sql.DB) (viewModels.NodeVM, error) {
	stmt, err := db.Prepare(`SELECT
	n.Id,
	n.Name,
	n.Mac,
	cp.Id,
	cp.IpAddress,
	cp.Name
	FROM Node AS n
	INNER JOIN ControlPointNodes AS cpn ON cpn.NodeId = n.Id
	INNER JOIN ControlPoint AS cp ON cp.Id = cpn.ControlPointId
	WHERE n.Id = ?`)
	setup.CheckErr(err)
	defer stmt.Close()

	var node viewModels.NodeVM

	err = stmt.QueryRow(nodeId).Scan(&node.Id,
		&node.Name,
		&node.Mac,
		&node.ControlPointId,
		&node.ControlPointIP,
		&node.ControlPointName)

	if err != nil {
		log.Println(err)
		return node, err
	}

	node.Sensors, err = FetchNodeSensors(node.Id, db)

	if err != nil {
		log.Println(err)
		return node, err
	}

	node.Switches, err = FetchNodeSwitches(node.Id, db)

	if err != nil {
		log.Println(err)
		return node, err
	}

	return node, nil
}

func FetchNodeSensors(nodeId int, db *sql.DB) ([]viewModels.NodeSensorVM, error) {
	stmt, err := db.Prepare(`SELECT
		ns.Id,
		ns.SensorTypeId,
		ns.Pin,
		ns.Name,
		st.Name AS SensorTypeName
		FROM NodeSensor AS ns
		INNER JOIN SensorType AS st ON st.Id = ns.SensorTypeId
		WHERE ns.NodeId = ?`)

	if err != nil {
		log.Printf("Error preparing select node sensors sql: %v", err)
		return nil, err
	}

	var sensors []viewModels.NodeSensorVM

	rows, err := stmt.Query(nodeId)

	if err != nil {
		log.Println("Error querying for node sensors")
		return nil, err
	}

	defer stmt.Close()

	for rows.Next() {
		var sensor viewModels.NodeSensorVM
		sensor.NodeId = nodeId

		err = rows.Scan(&sensor.Id,
			&sensor.SensorTypeId,
			&sensor.Pin,
			&sensor.Name,
			&sensor.SensorTypeName)

		if err != nil {
			log.Println("Error scanning node sensor")
			return nil, err
		}

		sensors = append(sensors, sensor)
	}

	return sensors, nil
}

func FetchNodeSwitches(nodeId int, db *sql.DB) ([]viewModels.NodeSwitchVM, error) {
	stmt, err := db.Prepare(`SELECT
		ns.Id,
		ns.SwitchTypeId,
		ns.Pin,
		ns.Name,
		ns.MomentaryPressDuration,
		ns.IsClosedOn,
		st.Name AS SwitchTypeName
		FROM NodeSwitch AS ns
		INNER JOIN SwitchType AS st ON st.Id = ns.SwitchTypeId
		WHERE ns.NodeId = ?`)

	if err != nil {
		log.Printf("Error preparing select node switches sql: %v", err)
		return nil, err
	}

	var nodeSwitches []viewModels.NodeSwitchVM

	rows, err := stmt.Query(nodeId)

	if err != nil {
		log.Println("Error querying for node switches")
		return nil, err
	}

	defer stmt.Close()

	for rows.Next() {
		var nodeSwitch viewModels.NodeSwitchVM
		nodeSwitch.NodeId = nodeId

		err = rows.Scan(&nodeSwitch.Id,
			&nodeSwitch.SwitchTypeId,
			&nodeSwitch.Pin,
			&nodeSwitch.Name,
			&nodeSwitch.MomentaryPressDuration,
			&nodeSwitch.IsClosedOn,
			&nodeSwitch.SwitchTypeName)

		if err != nil {
			log.Println("Error scanning node switch")
			return nil, err
		}

		nodeSwitches = append(nodeSwitches, nodeSwitch)
	}

	return nodeSwitches, nil
}

func CreateNode(item *models.Node, db *sql.DB) error {
	stmt, err := db.Prepare(`INSERT INTO Node
	(Name, Mac)
	VALUES (?, ?)`)

	if err != nil {
		log.Println("Error preparing create node sql")
		return err
	}

	result, err := stmt.Exec(&item.Name,
		&item.Mac)

	if err != nil {
		log.Println("Error creating node")
		return err
	}

	defer stmt.Close()

	lastInsertId, err := result.LastInsertId()

	if err != nil {
		log.Println("Error getting the id of the inserted node")
		return err
	}

	item.Id = int(lastInsertId)

	return nil
}

func UpdateNode(node *models.Node, db *sql.DB) error {
	stmt, err := db.Prepare(`UPDATE Node
	SET Name = ?, Mac = ?
	WHERE id = ?`)

	if err != nil {
		log.Println("Error preparing update node sql")
		return err
	}

	_, err = stmt.Exec(&node.Name,
		&node.Mac,
		&node.Id)

	if err != nil {
		log.Println("Error updating node")
		return err
	}

	defer stmt.Close()

	return nil
}

func DeleteNode(nodeId int, db *sql.DB) error {
	err := DeleteControlPointNodeByNodeId(nodeId, db)

	if err != nil {
		log.Printf("Error deleting the control point node: %v", err)
	}

	err = DeleteAllNodeSensors(nodeId, db)

	if err != nil {
		log.Printf("Error deleting the node sensors: %v", err)
	}

	err = DeleteAllNodeSwitches(nodeId, db)

	if err != nil {
		log.Printf("Error deleting the node switches: %v", err)
	}

	stmt, err := db.Prepare(`DELETE FROM Node WHERE id = ?;`)

	if err != nil {
		log.Println("Error preparing delete node sql")
		return err
	}

	_, err = stmt.Exec(nodeId)

	if err != nil {
		log.Println("Error deleting node")
		return err
	}

	defer stmt.Close()

	return nil
}
