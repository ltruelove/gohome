package data

import (
	"database/sql"
	"errors"
	"log"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data/statements"
	"github.com/ltruelove/gohome/internal/app/models"
)

type NodeSensorLogData struct {
	db   *sql.DB
	stmt statements.CrudStatement
}

func NewNodeSensorLogData(DB *sql.DB, config *config.Configuration) CrudDataInterface {
	return &NodeSensorLogData{
		db:   DB,
		stmt: statements.NewNodeSensorLogDataStatements(config),
	}
}

func (d *NodeSensorLogData) DB() *sql.DB {
	return d.db
}

func (d *NodeSensorLogData) Stmt() statements.CrudStatement {
	return d.stmt
}

func (d *NodeSensorLogData) SelectAll() ([]models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectAll())
	if err != nil {
		log.Println("Error preparing fetch all node sensor logs sql")
		return nil, err
	}

	var logs []models.Model

	rows, err := stmt.Query()
	if err != nil {
		log.Println("Error querying for all node sensor logs")
		return nil, err
	}
	defer stmt.Close()

	for rows.Next() {
		var logEntry models.NodeSensorLog

		err := rows.Scan(&logEntry.Id,
			&logEntry.NodeId,
			&logEntry.DateLogged)

		if err != nil {
			log.Println("Error scanning node sensor log")
			return nil, err
		}

		logs = append(logs, &logEntry)
	}

	return logs, nil
}

func (d *NodeSensorLogData) SelectById(id int) (models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectById())
	if err != nil {
		log.Println("Error preparing fetch node sensor log by id sql")
		return nil, err
	}

	var logEntry models.NodeSensorLog

	row := stmt.QueryRow(id)

	err = row.Scan(&logEntry.Id,
		&logEntry.NodeId,
		&logEntry.DateLogged)

	if err != nil {
		log.Println("Error querying for node sensor log by id")
		return nil, err
	}

	return logEntry, nil
}

func (d *NodeSensorLogData) SelectByParentId(nodeId int) ([]models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectByParentId())
	if err != nil {
		log.Println("Error preparing fetch node sensor logs by parent id sql")
		return nil, err
	}

	var logs []models.Model

	rows, err := stmt.Query(nodeId)
	if err != nil {
		log.Println("Error querying for node sensor logs by parent id")
		return nil, err
	}
	defer stmt.Close()

	for rows.Next() {
		var logEntry models.NodeSensorLog

		err := rows.Scan(&logEntry.Id,
			&logEntry.NodeId,
			&logEntry.DateLogged)

		if err != nil {
			log.Println("Error scanning node sensor log by parent id")
			return nil, err
		}

		logs = append(logs, &logEntry)
	}

	return logs, nil
}

func (d *NodeSensorLogData) SelectBySecondParentId(id int) ([]models.Model, error) {
	return nil, errors.New("NodeSensorLog table has no second parent ID")
}

func (d *NodeSensorLogData) Insert(data models.Model) (models.Model, error) {
	item, ok := data.(*models.NodeSensorLog)
	if !ok {
		log.Println("Error casting model to NodeSensorLog")
		return nil, errors.New("Invalid model type")
	}

	stmt, err := d.DB().Prepare(d.Stmt().Insert())
	if err != nil {
		log.Println("Error preparing insert node sensor log sql")
		return nil, err
	}

	var id int
	err = stmt.QueryRow(&item.NodeId, &item.DateLogged).Scan(&id)
	if err != nil {
		log.Println("Error inserting node sensor log")
		return nil, err
	}

	item.Id = id

	defer stmt.Close()

	return item, nil
}

func (d *NodeSensorLogData) Update(data models.Model) error {
	item, ok := data.(*models.NodeSensorLog)
	if !ok {
		log.Println("Error casting model to NodeSensorLog")
		return errors.New("Invalid model type")
	}

	stmt, err := d.DB().Prepare(d.Stmt().Update())
	if err != nil {
		log.Println("Error preparing update node sensor log sql")
		return err
	}

	_, err = stmt.Exec(&item.NodeId, &item.DateLogged, &item.Id)
	if err != nil {
		log.Println("Error updating node sensor log")
		return err
	}

	defer stmt.Close()

	return nil
}

func (d *NodeSensorLogData) Delete(id int) error {
	stmt, err := d.DB().Prepare(d.Stmt().Delete())
	if err != nil {
		log.Println("Error preparing delete node sensor log sql")
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		log.Println("Error deleting node sensor log")
		return err
	}

	defer stmt.Close()

	return nil
}

func (d *NodeSensorLogData) DeleteAll() error {
	stmt, err := d.DB().Prepare(d.Stmt().DeleteAll())
	if err != nil {
		log.Println("Error preparing delete all node sensor logs sql")
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Println("Error deleting all node sensor logs")
		return err
	}

	defer stmt.Close()

	return nil
}

func (d *NodeSensorLogData) DeleteByParentId(id int) error {
	stmt, err := d.DB().Prepare(d.Stmt().DeleteByParentId())
	if err != nil {
		log.Println("Error preparing delete node sensor logs by parent id sql")
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		log.Println("Error deleting node sensor logs by parent id")
		return err
	}

	defer stmt.Close()

	return nil
}

func (d *NodeSensorLogData) DeleteBySecondParentId(id int) error {
	return errors.New("NodeSensorLog table has no second parent ID")
}
