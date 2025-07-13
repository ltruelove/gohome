package data

import (
	"database/sql"
	"errors"
	"log"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data/statements"
	"github.com/ltruelove/gohome/internal/app/models"
)

type NodeSensorData struct {
	db   *sql.DB
	stmt statements.CrudStatement
}

func NewNodeSensorData(db *sql.DB, config *config.Configuration) CrudDataInterface {
	return &NodeSensorData{
		db:   db,
		stmt: statements.NewNodeSensorDataStatements(config),
	}
}

func (d *NodeSensorData) DB() *sql.DB {
	return d.db
}

func (d *NodeSensorData) Stmt() statements.CrudStatement {
	return d.stmt
}

func (d *NodeSensorData) SelectAll() ([]models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectAll())
	if err != nil {
		log.Println("Error preparing fetch all node sensors sql")
		return nil, err
	}

	var sensors []models.Model

	rows, err := stmt.Query()
	if err != nil {
		log.Println("Error querying for all node sensors")
		return nil, err
	}
	defer stmt.Close()

	for rows.Next() {
		var sensor models.NodeSensor

		err := rows.Scan(&sensor.Id,
			&sensor.NodeId,
			&sensor.SensorTypeId,
			&sensor.Name,
			&sensor.Pin,
			&sensor.DHTType)

		if err != nil {
			log.Println("Error scanning node sensor")
			return nil, err
		}

		sensors = append(sensors, sensor)
	}

	return sensors, nil
}

func (d *NodeSensorData) SelectById(id int) (models.Model, error) {
	var sensor models.NodeSensor

	stmt, err := d.DB().Prepare(d.Stmt().SelectById())
	if err != nil {
		log.Println("Error preparing fetch node sensor sql")
		return sensor, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&sensor.Id,
		&sensor.NodeId,
		&sensor.SensorTypeId,
		&sensor.Name,
		&sensor.Pin,
		&sensor.DHTType)
	if err != nil {
		log.Println("Error querying for node sensor")
		return sensor, err
	}

	return sensor, nil
}

// SelectByParentId retrieves NodeSensor records by NodeId
func (d *NodeSensorData) SelectByParentId(id int) ([]models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectByParentId())
	if err != nil {
		log.Println("Error preparing fetch all node sensors by parent id sql")
		return nil, err
	}

	var sensors []models.Model

	rows, err := stmt.Query(id)
	if err != nil {
		log.Println("Error querying for all node sensors by parent id")
		return nil, err
	}
	defer stmt.Close()

	for rows.Next() {
		var sensor models.NodeSensor

		err := rows.Scan(&sensor.Id,
			&sensor.NodeId,
			&sensor.SensorTypeId,
			&sensor.Name,
			&sensor.Pin,
			&sensor.DHTType)

		if err != nil {
			log.Println("Error scanning node sensor")
			return nil, err
		}

		sensors = append(sensors, sensor)
	}

	return sensors, nil
}

// SelectBySecondParentId retrieves NodeSensor records by SensorTypeId
func (d *NodeSensorData) SelectBySecondParentId(id int) ([]models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectBySecondParentId())
	if err != nil {
		log.Println("Error preparing fetch node sensors by second parent id sql")
		return nil, err
	}

	var sensors []models.Model

	rows, err := stmt.Query(id)
	if err != nil {
		log.Println("Error querying for node sensors by second parent id")
		return nil, err
	}
	defer stmt.Close()

	for rows.Next() {
		var sensor models.NodeSensor

		err := rows.Scan(&sensor.Id,
			&sensor.NodeId,
			&sensor.SensorTypeId,
			&sensor.Name,
			&sensor.Pin,
			&sensor.DHTType)

		if err != nil {
			log.Println("Error scanning node sensor")
			return nil, err
		}

		sensors = append(sensors, sensor)
	}

	return sensors, nil
}

func (d *NodeSensorData) Insert(data models.Model) (models.Model, error) {
	item, ok := data.(*models.NodeSensor)
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

	err = stmt.QueryRow(&item.NodeId,
		&item.SensorTypeId,
		&item.Name,
		&item.Pin,
		&item.DHTType).Scan(&lastInsertId)

	if err != nil {
		log.Println("Error creating node sensor")
		return item, err
	}

	defer stmt.Close()

	item.Id = int(lastInsertId)

	return item, nil
}

func (d *NodeSensorData) Update(data models.Model) error {
	item, ok := data.(*models.NodeSensor)
	if !ok {
		log.Println("Error casting model to NodeSensor")
		return errors.New("Invalid model type")
	}

	stmt, err := d.DB().Prepare(d.Stmt().Update())
	if err != nil {
		log.Println("Error preparing update node sensor sql")
		return err
	}

	_, err = stmt.Exec(item.NodeId,
		item.SensorTypeId,
		item.Name,
		item.Pin,
		item.DHTType,
		item.Id)

	if err != nil {
		log.Println("Error updating node sensor")
		return err
	}

	defer stmt.Close()

	return nil
}

func (d *NodeSensorData) Delete(id int) error {
	stmt, err := d.DB().Prepare(d.Stmt().Delete())
	if err != nil {
		log.Println("Error preparing delete node sensor sql")
		return err
	}

	_, err = stmt.Exec(id)

	if err != nil {
		log.Println("Error deleting node sensor")
		return err
	}

	defer stmt.Close()

	return nil
}

func (d *NodeSensorData) DeleteAll() error {
	stmt, err := d.DB().Prepare(d.Stmt().DeleteAll())
	if err != nil {
		log.Println("Error preparing delete all node sensors sql")
		return err
	}

	_, err = stmt.Exec()

	if err != nil {
		log.Println("Error deleting all node sensors")
		return err
	}

	defer stmt.Close()

	return nil
}

func (d *NodeSensorData) DeleteByParentId(id int) error {
	stmt, err := d.DB().Prepare(d.Stmt().DeleteByParentId())
	if err != nil {
		log.Println("Error preparing delete node sensor by parent id sql")
		return err
	}

	_, err = stmt.Exec(id)

	if err != nil {
		log.Println("Error deleting node sensor by parent id")
		return err
	}

	defer stmt.Close()

	return nil
}

func (d *NodeSensorData) DeleteBySecondParentId(id int) error {
	stmt, err := d.DB().Prepare(d.Stmt().DeleteBySecondParentId())
	if err != nil {
		log.Println("Error preparing delete node sensor by second parent id sql")
		return err
	}

	_, err = stmt.Exec(id)

	if err != nil {
		log.Println("Error deleting node sensor by second parent id")
		return err
	}

	defer stmt.Close()
	return nil
}
