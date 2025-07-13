package data

import (
	"database/sql"
	"errors"
	"log"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data/statements"
	"github.com/ltruelove/gohome/internal/app/models"
)

type ViewNodeSensorData struct {
	db   *sql.DB
	stmt statements.CrudStatement
}

func NewViewNodeSensorData(db *sql.DB, config *config.Configuration) CrudDataInterface {
	return &ViewNodeSensorData{
		db:   db,
		stmt: statements.NewViewNodeSensorDataStatements(config),
	}
}

func (d *ViewNodeSensorData) DB() *sql.DB {
	return d.db
}

func (d *ViewNodeSensorData) Stmt() statements.CrudStatement {
	return d.stmt
}

func (d *ViewNodeSensorData) SelectAll() ([]models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectAll())

	if err != nil {
		log.Println("Error preparing all node sensor data sql")
		return nil, err
	}
	var listData []models.Model

	rows, err := stmt.Query()
	if err != nil {
		log.Println("Error querying for all node sensor data")
		return nil, err
	}

	for rows.Next() {
		var item models.ViewNodeSensorData
		err := rows.Scan(&item.Id,
			&item.NodeId,
			&item.ViewId,
			&item.NodeSensorId,
			&item.Name)

		if err != nil {
			log.Println("Error scanning node sensor data")
			return nil, err
		}
		listData = append(listData, item)
	}
	defer stmt.Close()

	return listData, nil
}

func (d *ViewNodeSensorData) SelectById(id int) (models.Model, error) {
	var item models.ViewNodeSensorData

	stmt, err := d.DB().Prepare(d.Stmt().SelectById())
	if err != nil {
		log.Println("Error preparing fetch node sensor data sql")
		return item, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&item.Id,
		&item.NodeId,
		&item.ViewId,
		&item.NodeSensorId,
		&item.Name)

	if err != nil {
		log.Println("Error querying for node sensor data")
		return item, err
	}

	return item, nil
}

func (d *ViewNodeSensorData) SelectByParentId(nodeId int) ([]models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectByParentId())
	if err != nil {
		log.Println("Error preparing fetch node sensor data by node id sql")
		return nil, err
	}

	var listData []models.Model

	rows, err := stmt.Query(nodeId)
	if err != nil {
		log.Println("Error querying for node sensor data by node id")
		return nil, err
	}

	for rows.Next() {
		var item models.ViewNodeSensorData
		err := rows.Scan(&item.Id,
			&item.NodeId,
			&item.ViewId,
			&item.NodeSensorId,
			&item.Name)

		if err != nil {
			log.Println("Error scanning node sensor data by node id")
			return nil, err
		}
		listData = append(listData, item)
	}
	defer stmt.Close()

	return listData, nil
}

func (d *ViewNodeSensorData) SelectBySecondParentId(viewId int) ([]models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectBySecondParentId())
	if err != nil {
		log.Println("Error preparing fetch node sensor data by second parent id sql")
		return nil, err
	}

	defer stmt.Close()

	var listData []models.Model

	rows, err := stmt.Query(viewId)
	if err != nil {
		log.Println("Error querying for node sensor data by second parent id")
		return nil, err
	}

	for rows.Next() {
		var item models.ViewNodeSensorData
		err := rows.Scan(&item.Id,
			&item.NodeId,
			&item.ViewId,
			&item.NodeSensorId,
			&item.Name)

		if err != nil {
			log.Println("Error scanning node sensor data by second parent id")
			return nil, err
		}
		listData = append(listData, item)
	}
	defer stmt.Close()

	return listData, nil
}

func (d *ViewNodeSensorData) Insert(data models.Model) (models.Model, error) {
	item, ok := data.(*models.ViewNodeSensorData)
	if !ok {
		log.Println("Error asserting model type")
		return nil, errors.New("invalid type for Insert")
	}
	stmt, err := d.DB().Prepare(d.Stmt().Insert())

	if err != nil {
		log.Println("Error preparing create node sensor data sql")
		return nil, err
	}

	lastInsertId := 0

	err = stmt.QueryRow(item.NodeId,
		item.ViewId,
		item.NodeSensorId,
		item.Name).Scan((&lastInsertId))

	if err != nil {
		log.Println("Error creating node sensor data")
		return nil, err
	}

	item.Id = int(lastInsertId)
	log.Printf("Created node sensor data with id: %d", item.Id)
	defer stmt.Close()

	return item, nil
}

func (d *ViewNodeSensorData) Update(data models.Model) error {
	item, ok := data.(*models.ViewNodeSensorData)
	if !ok {
		log.Println("Error asserting model type")
		return errors.New("invalid type for Update")
	}
	stmt, err := d.DB().Prepare(d.Stmt().Update())
	if err != nil {
		log.Println("Error preparing update node sensor data sql")
		return err
	}

	_, err = stmt.Exec(item.NodeId,
		item.ViewId,
		item.NodeSensorId,
		item.Name,
		item.Id)

	if err != nil {
		log.Println("Error updating node sensor data")
		return err
	}

	defer stmt.Close()

	log.Printf("Updated node sensor data with id: %d", item.Id)
	return nil
}

func (d *ViewNodeSensorData) Delete(id int) error {
	stmt, err := d.DB().Prepare(d.Stmt().Delete())
	if err != nil {
		log.Println("Error preparing delete node sensor data sql")
		return err
	}

	_, err = stmt.Exec(id)

	if err != nil {
		log.Println("Error deleting node sensor data")
		return err
	}

	defer stmt.Close()

	log.Printf("Deleted node sensor data with id: %d", id)
	return nil
}

func (d *ViewNodeSensorData) DeleteByParentId(id int) error {
	stmt, err := d.DB().Prepare(d.Stmt().DeleteByParentId())
	if err != nil {
		log.Println("Error preparing delete node sensor data by node id sql")
		return err
	}

	_, err = stmt.Exec(id)

	if err != nil {
		log.Println("Error deleting node sensor data by node id")
		return err
	}

	defer stmt.Close()

	log.Printf("Deleted node sensor data for node id: %d", id)
	return nil
}

func (d *ViewNodeSensorData) DeleteAll() error {
	stmt, err := d.DB().Prepare(d.Stmt().DeleteAll())
	if err != nil {
		log.Println("Error preparing delete all node sensor data sql")
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Println("Error deleting all node sensor data")
		return err
	}

	defer stmt.Close()

	log.Println("Deleted all node sensor data")
	return nil
}
