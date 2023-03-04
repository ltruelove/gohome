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

func FetchAllNodes(db *sql.DB) ([]viewModels.NodeVM, error) {
	stmt, err := db.Prepare(`SELECT
	n.id,
	n.name,
	n.mac,
	cp.id AS controlpointid,
	cp.ipaddress AS controlpointip,
	cp.name AS controlpointname
	FROM node AS n
	LEFT JOIN controlpointnodes AS cpn ON cpn.nodeid = n.id
	LEFT JOIN controlpoint AS cp ON cp.id = cpn.controlpointid`)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	var nodes []viewModels.NodeVM

	rows, err := stmt.Query()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		var node viewModels.NodeVM
		rows.Scan(&node.Id,
			&node.Name,
			&node.Mac,
			&node.ControlPointId,
			&node.ControlPointIP,
			&node.ControlPointName)

		node.Sensors, err = FetchNodeSensors(node.Id, db)

		if err != nil {
			log.Println(err)
		}

		node.Switches, err = FetchNodeSwitches(node.Id, db)

		if err != nil {
			log.Println(err)
		}

		nodes = append(nodes, node)
	}
	defer stmt.Close()

	return nodes, nil
}

func FetchNode(nodeId int, db *sql.DB) (viewModels.NodeVM, error) {
	stmt, err := db.Prepare(`SELECT
	n.id,
	n.name,
	n.mac,
	cp.id AS cpid,
	cp.ipaddress AS ipaddress,
	cp.name AS cpname
	FROM node AS n
	LEFT JOIN controlpointnodes AS cpn ON cpn.nodeid = n.id
	LEFT JOIN controlpoint AS cp ON cp.id = cpn.controlpointid
	WHERE n.id = $1`)
	log.Println("fetching node")
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
	log.Println("node found")

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
		ns.id,
		ns.sensortypeid,
		ns.pin,
		ns.name,
		st.name AS sensortypename
		FROM nodesensor AS ns
		INNER JOIN sensortype AS st ON st.id = ns.sensortypeid
		WHERE ns.nodeid = $1`)

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
		ns.id,
		ns.switchtypeid,
		ns.pin,
		ns.name,
		ns.momentarypressduration,
		ns.isclosedon,
		st.name AS switchtypename
		FROM nodeswitch AS ns
		INNER JOIN switchtype AS st ON st.id = ns.switchtypeid
		WHERE ns.nodeid = $1`)

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
	stmt, err := db.Prepare(`INSERT INTO node
	(name, mac)
	VALUES ($1, $2) RETURNING id`)

	if err != nil {
		log.Println("Error preparing create node sql")
		return err
	}

	lastInsertId := 0

	err = stmt.QueryRow(&item.Name,
		&item.Mac).Scan(&lastInsertId)

	if err != nil {
		log.Println("Error creating node")
		return err
	}

	defer stmt.Close()

	if err != nil {
		log.Println("Error getting the id of the inserted node")
		return err
	}

	item.Id = int(lastInsertId)

	return nil
}

func UpdateNode(node *models.Node, db *sql.DB) error {
	stmt, err := db.Prepare(`UPDATE node
	set name = $1, mac = $2
	WHERE id = $3`)

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
