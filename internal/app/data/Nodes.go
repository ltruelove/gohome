package data

import (
	"database/sql"
	"log"

	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/ltruelove/gohome/internal/app/setup"
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

func FetchNode(nodeId int, db *sql.DB) (models.Node, error) {
	stmt, err := db.Prepare("SELECT Id, Name FROM Node WHERE id = ?")
	setup.CheckErr(err)
	defer stmt.Close()

	var node models.Node

	err = stmt.QueryRow(nodeId).Scan(&node.Id,
		&node.Name)

	if err != nil {
		log.Println(err)
		return node, err
	}

	return node, nil
}

func FetchNodeSensors(nodeId int, db *sql.DB) ([]models.NodeSensor, error) {
	stmt, err := db.Prepare(`SELECT
		Id,
		SensorTypeId,
		Pin,
		Name FROM NodeSenor WHERE NodeId = ?`)

	if err != nil {
		log.Printf("Error preparing select node sensors sql: %v", err)
		return nil, err
	}

	var sensors []models.NodeSensor

	rows, err := stmt.Query()

	if err != nil {
		log.Println("Error querying for node sensors")
		return nil, err
	}

	defer stmt.Close()

	for rows.Next() {
		var sensor models.NodeSensor
		sensor.NodeId = nodeId

		err = rows.Scan(&sensor.Id,
			&sensor.SensorTypeId,
			&sensor.Pin,
			&sensor.Name)

		if err != nil {
			log.Println("Error scanning node sensor")
			return nil, err
		}

		sensors = append(sensors, sensor)
	}

	return sensors, nil
}

func FetchNodeSwitches(nodeId int, db *sql.DB) ([]models.NodeSwitch, error) {
	stmt, err := db.Prepare(`SELECT
		Id,
		SwitchTypeId,
		Pin,
		Name FROM NodeSwitch WHERE NodeId = ?`)

	if err != nil {
		log.Printf("Error preparing select node switches sql: %v", err)
		return nil, err
	}

	var nodeSwitches []models.NodeSwitch

	rows, err := stmt.Query(nodeId)

	if err != nil {
		log.Println("Error querying for node switches")
		return nil, err
	}

	defer stmt.Close()

	for rows.Next() {
		var nodeSwitch models.NodeSwitch
		nodeSwitch.NodeId = nodeId

		err = rows.Scan(&nodeSwitch.Id,
			&nodeSwitch.SwitchTypeId,
			&nodeSwitch.Pin,
			&nodeSwitch.Name)

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
	stmt, err := db.Prepare(`DELETE FROM Node
	WHERE id = ?`)

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
