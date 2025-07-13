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
