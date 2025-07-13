package data

import (
	"database/sql"
	"errors"
	"log"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data/statements"
	"github.com/ltruelove/gohome/internal/app/models"
)

type ControlPointNodeData struct {
	db   *sql.DB
	stmt statements.CrudStatement
}

func NewControlPointNodeData(db *sql.DB, config *config.Configuration) CrudDataInterface {
	return &ControlPointNodeData{
		db:   db,
		stmt: statements.NewControlPointNodeDataStatements(config),
	}
}

func (d *ControlPointNodeData) DB() *sql.DB {
	return d.db
}

func (d *ControlPointNodeData) Stmt() statements.CrudStatement {
	return d.stmt
}

func (d *ControlPointNodeData) SelectAll() ([]models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectAll())
	if err != nil {
		log.Println("Error preparing fetch all control point node sql")
		return nil, err
	}
	defer stmt.Close()

	var controlPointNodes []models.Model

	rows, err := stmt.Query()
	if err != nil {
		log.Println("Error querying for all control point nodes")
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var controlPointNode models.ControlPointNode
		err := rows.Scan(&controlPointNode.Id,
			&controlPointNode.ControlPointId,
			&controlPointNode.NodeId)

		if err != nil {
			log.Println("Error scanning control point node")
			return nil, err
		}
		controlPointNodes = append(controlPointNodes, &controlPointNode)
	}

	return controlPointNodes, nil
}

func (d *ControlPointNodeData) SelectByParentId(id int) ([]models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectByParentId())
	if err != nil {
		log.Println("Error preparing fetch control point nodes by ControlPointId sql")
		return nil, err
	}
	defer stmt.Close()

	var controlPointNodes []models.Model

	rows, err := stmt.Query(id)
	if err != nil {
		log.Println("Error querying for control point nodes by ControlPointId")
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var controlPointNode models.ControlPointNode
		err := rows.Scan(&controlPointNode.Id,
			&controlPointNode.ControlPointId,
			&controlPointNode.NodeId)

		if err != nil {
			log.Println("Error scanning control point node")
			return nil, err
		}
		controlPointNodes = append(controlPointNodes, &controlPointNode)
	}

	return controlPointNodes, nil
}

func (d *ControlPointNodeData) SelectBySecondParentId(id int) ([]models.Model, error) {
	return nil, errors.New("ControlPointNodeData table has no second parent ID")
}

func (d *ControlPointNodeData) SelectById(id int) (models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectById())
	if err != nil {
		log.Println("Error preparing fetch control point node by id sql")
		return nil, err
	}
	defer stmt.Close()

	var controlPointNode models.ControlPointNode

	err = stmt.QueryRow(id).Scan(&controlPointNode.Id,
		&controlPointNode.ControlPointId,
		&controlPointNode.NodeId)

	if err != nil {
		log.Println("Error querying for control point node by id")
		return nil, err
	}

	return &controlPointNode, nil
}

func (d *ControlPointNodeData) Insert(data models.Model) (models.Model, error) {
	controlPointNode, ok := data.(*models.ControlPointNode)
	if !ok {
		log.Println("Error casting model to ControlPointNode")
		return nil, errors.New("Invalid model type")
	}

	stmt, err := d.DB().Prepare(d.Stmt().Insert())
	if err != nil {
		log.Println("Error preparing insert control point node sql")
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(controlPointNode.ControlPointId, controlPointNode.NodeId)
	if err != nil {
		log.Println("Error executing insert for control point node")
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Println("Error getting last insert id for control point node")
		return nil, err
	}
	controlPointNode.Id = int(id)

	return controlPointNode, nil
}

func (d *ControlPointNodeData) Update(data models.Model) error {
	controlPointNode, ok := data.(*models.ControlPointNode)
	if !ok {
		log.Println("Error casting model to ControlPointNode")
		return errors.New("Invalid model type")
	}

	stmt, err := d.DB().Prepare(d.Stmt().Update())
	if err != nil {
		log.Println("Error preparing update control point node sql")
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(controlPointNode.ControlPointId,
		controlPointNode.NodeId,
		controlPointNode.Id)

	if err != nil {
		log.Println("Error executing update for control point node")
		return err
	}

	return nil
}

func (d *ControlPointNodeData) Delete(id int) error {
	stmt, err := d.DB().Prepare(d.Stmt().Delete())
	if err != nil {
		log.Println("Error preparing delete control point node sql")
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Println("Error executing delete for control point node")
		return err
	}

	return nil
}

func (d *ControlPointNodeData) DeleteByParentId(id int) error {
	stmt, err := d.DB().Prepare(d.Stmt().DeleteByParentId())
	if err != nil {
		log.Println("Error preparing delete control point nodes by ControlPointId sql")
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Println("Error executing delete for control point nodes by ControlPointId")
		return err
	}

	return nil
}

func (d *ControlPointNodeData) DeleteAll() error {
	stmt, err := d.DB().Prepare(d.Stmt().DeleteAll())
	if err != nil {
		log.Println("Error preparing delete all control point nodes sql")
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		log.Println("Error executing delete all control point nodes")
		return err
	}

	return nil
}

func (d *ControlPointNodeData) DeleteBySecondParentId(id int) error {
	return errors.New("ControlPointNodeData does not support delete by second parent ID")
}
