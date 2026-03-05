package data

import (
	"database/sql"
	"errors"
	"log"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data/statements"
	"github.com/ltruelove/gohome/internal/app/models"
)

type NodeData struct {
	db   *sql.DB
	stmt statements.CrudStatement
}

func NewNodeData(DB *sql.DB, config *config.Configuration) CrudDataInterface {
	return &NodeData{
		db:   DB,
		stmt: statements.NewNodeDataStatements(config),
	}
}

func (d *NodeData) DB() *sql.DB {
	return d.db
}

func (d *NodeData) Stmt() statements.CrudStatement {
	return d.stmt
}

func (d *NodeData) SelectAll() ([]models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectAll())
	if err != nil {
		log.Println("Error preparing fetch all nodes sql")
		return nil, err
	}

	var nodes []models.Model

	rows, err := stmt.Query()
	if err != nil {
		log.Println("Error querying for all nodes")
		return nil, err
	}
	defer stmt.Close()

	for rows.Next() {
		var node models.Node

		err := rows.Scan(&node.Id,
			&node.Mac,
			&node.Name,
			&node.IpAddress)

		if err != nil {
			log.Println("Error scanning node")
			return nil, err
		}

		nodes = append(nodes, node)
	}

	return nodes, nil

}

func (d *NodeData) SelectById(id int) (models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectById())
	if err != nil {
		log.Println("Error preparing fetch node by id sql")
		return nil, err
	}

	var node models.Node

	err = stmt.QueryRow(id).Scan(&node.Id,
		&node.Mac,
		&node.Name,
		&node.IpAddress)

	if err != nil {
		log.Println("Error querying for node by id")
		return nil, err
	}

	defer stmt.Close()

	return node, nil
}

func (d *NodeData) SelectByParentId(id int) ([]models.Model, error) {
	return nil, errors.New("Node table has no parent ID")
}

func (d *NodeData) SelectBySecondParentId(id int) ([]models.Model, error) {
	return nil, errors.New("Node table has no second parent ID")
}

func (d *NodeData) Insert(data models.Model) (models.Model, error) {
	item, ok := data.(*models.Node)
	if !ok {
		log.Println("Error casting model to NodeSensor")
		return nil, errors.New("Invalid model type")
	}

	stmt, err := d.DB().Prepare(d.Stmt().Insert())
	if err != nil {
		log.Println("Error preparing create node sensor sql")
		return item, err
	}

	lastInsertId := 0
	err = stmt.QueryRow(&item.Name,
		&item.Mac,
		&item.IpAddress).Scan(&lastInsertId)

	if err != nil {
		log.Println("Error creating node")
		return nil, err
	}

	defer stmt.Close()

	item.Id = int(lastInsertId)

	return item, nil
}

func (d *NodeData) Update(data models.Model) error {
	item, ok := data.(*models.Node)
	if !ok {
		log.Println("Error casting model to Node")
		return errors.New("Invalid model type")
	}

	stmt, err := d.DB().Prepare(d.Stmt().Update())
	if err != nil {
		log.Println("Error preparing update node sql")
		return err
	}

	_, err = stmt.Exec(&item.Name,
		&item.Mac,
		&item.IpAddress,
		&item.Id)

	if err != nil {
		log.Println("Error updating node")
		return err
	}

	defer stmt.Close()

	return nil
}

func (d *NodeData) Delete(id int) error {
	stmt, err := d.DB().Prepare(d.Stmt().Delete())
	if err != nil {
		log.Println("Error preparing delete node sql")
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		log.Println("Error deleting node")
		return err
	}

	defer stmt.Close()

	return nil
}

func (d *NodeData) DeleteAll() error {
	stmt, err := d.DB().Prepare(d.Stmt().DeleteAll())
	if err != nil {
		log.Println("Error preparing delete all nodes sql")
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Println("Error deleting all nodes")
		return err
	}

	defer stmt.Close()

	return nil
}

func (d *NodeData) DeleteByParentId(id int) error {
	return errors.New("Node table has no parent ID")
}

func (d *NodeData) DeleteBySecondParentId(id int) error {
	return errors.New("Node table has no second parent ID")
}

// Convenience wrappers used by older callers that expect package-level
// create functions. These call the DB directly using Postgres style SQL
// (the rest of the project primarily targets Postgres).
func CreateNode(node *models.Node, db *sql.DB) error {
	stmt, err := db.Prepare(`INSERT INTO node (mac, name, ipaddress) VALUES ($1, $2, $3) RETURNING id`)
	if err != nil {
		log.Println("Error preparing create node sql")
		return err
	}

	lastInsertId := 0
	err = stmt.QueryRow(node.Mac, node.Name, node.IpAddress).Scan(&lastInsertId)
	if err != nil {
		log.Println("Error creating node")
		return err
	}

	defer stmt.Close()

	node.Id = int(lastInsertId)

	return nil
}

func CreateNodeSensor(sensor *models.NodeSensor, db *sql.DB) error {
	stmt, err := db.Prepare(`INSERT INTO nodesensor (nodeid, sensortypeid, name, pin, dhttype) VALUES ($1, $2, $3, $4, $5) RETURNING id`)
	if err != nil {
		log.Println("Error preparing create node sensor sql")
		return err
	}

	lastInsertId := 0
	err = stmt.QueryRow(sensor.NodeId, sensor.SensorTypeId, sensor.Name, sensor.Pin, sensor.DHTType).Scan(&lastInsertId)
	if err != nil {
		log.Println("Error creating node sensor")
		return err
	}

	defer stmt.Close()

	sensor.Id = int(lastInsertId)

	return nil
}

func CreateNodeSwitch(sw *models.NodeSwitch, db *sql.DB) error {
	stmt, err := db.Prepare(`INSERT INTO nodeswitch (nodeid, switchtypeid, name, pin, momentarypressduration, isclosedon) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`)
	if err != nil {
		log.Println("Error preparing create node switch sql")
		return err
	}

	lastInsertId := 0
	isClosedOnInt := 0
	if sw.IsClosedOn {
		isClosedOnInt = 1
	}

	err = stmt.QueryRow(sw.NodeId, sw.SwitchTypeId, sw.Name, sw.Pin, sw.MomentaryPressDuration, isClosedOnInt).Scan(&lastInsertId)
	if err != nil {
		log.Println("Error creating node switch")
		return err
	}

	defer stmt.Close()

	sw.Id = int(lastInsertId)

	return nil
}

func FetchAllNodes(db *sql.DB) ([]models.Node, error) {
	stmt, err := db.Prepare(`SELECT id, mac, name, ipaddress FROM node`)
	if err != nil {
		log.Println("Error preparing fetch all nodes sql")
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Println("Error querying for all nodes")
		return nil, err
	}
	defer stmt.Close()

	var nodes []models.Node
	for rows.Next() {
		var n models.Node
		if err := rows.Scan(&n.Id, &n.Mac, &n.Name, &n.IpAddress); err != nil {
			log.Println("Error scanning node")
			return nil, err
		}
		nodes = append(nodes, n)
	}

	return nodes, nil
}

func FetchNode(id int, db *sql.DB) (models.Node, error) {
	var n models.Node
	stmt, err := db.Prepare(`SELECT id, mac, name, ipaddress FROM node WHERE id = $1`)
	if err != nil {
		log.Println("Error preparing fetch node sql")
		return n, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&n.Id, &n.Mac, &n.Name, &n.IpAddress)
	if err != nil {
		log.Println("Error querying for node")
		return n, err
	}

	return n, nil
}

func VerifyNodeIdIsNew(nodeId int, db *sql.DB) (bool, error) {
	n, err := FetchNode(nodeId, db)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}
		return false, err
	}

	return n.Id < 1, nil
}

func UpdateNode(node *models.Node, db *sql.DB) error {
	stmt, err := db.Prepare(`UPDATE node SET mac = $1, name = $2, ipaddress = $3 WHERE id = $4`)
	if err != nil {
		log.Println("Error preparing update node sql")
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(node.Mac, node.Name, node.IpAddress, node.Id)
	if err != nil {
		log.Println("Error updating node")
		return err
	}

	return nil
}

func FetchIndividualNode(id int, db *sql.DB) (models.Node, error) {
	return FetchNode(id, db)
}

func DeleteNode(id int, db *sql.DB) error {
	stmt, err := db.Prepare(`DELETE FROM node WHERE id = $1`)
	if err != nil {
		log.Println("Error preparing delete node sql")
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Println("Error deleting node")
		return err
	}

	return nil
}

func UpdateNodeIp(node *models.Node, db *sql.DB) error {
	stmt, err := db.Prepare(`UPDATE node SET ipaddress = $1 WHERE id = $2`)
	if err != nil {
		log.Println("Error preparing update node IP sql")
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(node.IpAddress, node.Id)
	if err != nil {
		log.Println("Error updating node IP")
		return err
	}

	return nil
}
