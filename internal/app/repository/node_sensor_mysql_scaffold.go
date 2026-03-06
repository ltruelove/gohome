package repository

import (
	"database/sql"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data/statements"
	"github.com/ltruelove/gohome/internal/app/models"
)

type mysqlNodeSensorRepository struct {
	db   *sql.DB
	stmt statements.CrudStatement
}

func NewMySQLNodeSensorRepository(db *sql.DB, cfg *config.Configuration) CrudRepository {
	return &mysqlNodeSensorRepository{db: db, stmt: statements.NewNodeSensorDataStatements(cfg)}
}

func (r *mysqlNodeSensorRepository) SelectAll() ([]models.Model, error) {
	rows, err := r.db.Query(r.stmt.SelectAll())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.Model
	for rows.Next() {
		var s models.NodeSensor
		if err := rows.Scan(&s.Id, &s.NodeId, &s.SensorTypeId, &s.Name, &s.Pin, &s.DHTType); err != nil {
			return nil, err
		}
		list = append(list, s)
	}
	return list, nil
}

func (r *mysqlNodeSensorRepository) SelectByParentId(id int) ([]models.Model, error) {
	rows, err := r.db.Query(r.stmt.SelectByParentId(), id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []models.Model
	for rows.Next() {
		var s models.NodeSensor
		if err := rows.Scan(&s.Id, &s.NodeId, &s.SensorTypeId, &s.Name, &s.Pin, &s.DHTType); err != nil {
			return nil, err
		}
		list = append(list, s)
	}
	return list, nil
}

func (r *mysqlNodeSensorRepository) SelectById(id int) (models.Model, error) {
	var s models.NodeSensor
	err := r.db.QueryRow(r.stmt.SelectById(), id).Scan(&s.Id, &s.NodeId, &s.SensorTypeId, &s.Name, &s.Pin, &s.DHTType)
	return s, err
}

func (r *mysqlNodeSensorRepository) Insert(data models.Model) (models.Model, error) {
	item, ok := data.(*models.NodeSensor)
	if !ok {
		return nil, nil
	}
	var last int
	err := r.db.QueryRow(r.stmt.Insert(), item.NodeId, item.SensorTypeId, item.Name, item.Pin, item.DHTType).Scan(&last)
	if err != nil {
		return nil, err
	}
	item.Id = last
	return item, nil
}

func (r *mysqlNodeSensorRepository) Update(data models.Model) error {
	item, ok := data.(*models.NodeSensor)
	if !ok {
		return nil
	}
	_, err := r.db.Exec(r.stmt.Update(), item.NodeId, item.SensorTypeId, item.Name, item.Pin, item.DHTType, item.Id)
	return err
}

func (r *mysqlNodeSensorRepository) Delete(id int) error {
	_, err := r.db.Exec(r.stmt.Delete(), id)
	return err
}

func (r *mysqlNodeSensorRepository) DeleteAll() error {
	_, err := r.db.Exec(r.stmt.DeleteAll())
	return err
}

func (r *mysqlNodeSensorRepository) DeleteByParentId(id int) error {
	_, err := r.db.Exec(r.stmt.DeleteByParentId(), id)
	return err
}

func (r *mysqlNodeSensorRepository) DeleteBySecondParentId(id int) error {
	_, err := r.db.Exec(r.stmt.DeleteBySecondParentId(), id)
	return err
}
