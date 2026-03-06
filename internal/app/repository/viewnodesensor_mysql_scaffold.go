package repository

import (
	"database/sql"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data/statements"
	"github.com/ltruelove/gohome/internal/app/models"
)

type mysqlViewNodeSensorRepository struct {
	db   *sql.DB
	stmt statements.CrudStatement
}

func NewMySQLViewNodeSensorRepository(db *sql.DB, cfg *config.Configuration) CrudRepository {
	return &mysqlViewNodeSensorRepository{db: db, stmt: statements.NewViewNodeSensorDataStatements(cfg)}
}

func (r *mysqlViewNodeSensorRepository) SelectAll() ([]models.Model, error) {
	rows, err := r.db.Query(r.stmt.SelectAll())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.Model
	for rows.Next() {
		var item models.ViewNodeSensorData
		if err := rows.Scan(&item.Id, &item.NodeId, &item.ViewId, &item.NodeSensorId, &item.Name); err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, nil
}

func (r *mysqlViewNodeSensorRepository) SelectByParentId(id int) ([]models.Model, error) {
	rows, err := r.db.Query(r.stmt.SelectByParentId(), id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []models.Model
	for rows.Next() {
		var item models.ViewNodeSensorData
		if err := rows.Scan(&item.Id, &item.NodeId, &item.ViewId, &item.NodeSensorId, &item.Name); err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, nil
}

func (r *mysqlViewNodeSensorRepository) SelectById(id int) (models.Model, error) {
	var item models.ViewNodeSensorData
	err := r.db.QueryRow(r.stmt.SelectById(), id).Scan(&item.Id, &item.NodeId, &item.ViewId, &item.NodeSensorId, &item.Name)
	return item, err
}

func (r *mysqlViewNodeSensorRepository) Insert(data models.Model) (models.Model, error) {
	item, ok := data.(*models.ViewNodeSensorData)
	if !ok {
		return nil, nil
	}
	var last int
	err := r.db.QueryRow(r.stmt.Insert(), item.NodeId, item.ViewId, item.NodeSensorId, item.Name).Scan(&last)
	if err != nil {
		return nil, err
	}
	item.Id = last
	return item, nil
}

func (r *mysqlViewNodeSensorRepository) Update(data models.Model) error {
	item, ok := data.(*models.ViewNodeSensorData)
	if !ok {
		return nil
	}
	_, err := r.db.Exec(r.stmt.Update(), item.NodeId, item.ViewId, item.NodeSensorId, item.Name, item.Id)
	return err
}

func (r *mysqlViewNodeSensorRepository) Delete(id int) error {
	_, err := r.db.Exec(r.stmt.Delete(), id)
	return err
}

func (r *mysqlViewNodeSensorRepository) DeleteAll() error {
	_, err := r.db.Exec(r.stmt.DeleteAll())
	return err
}

func (r *mysqlViewNodeSensorRepository) DeleteByParentId(id int) error {
	_, err := r.db.Exec(r.stmt.DeleteByParentId(), id)
	return err
}

func (r *mysqlViewNodeSensorRepository) DeleteBySecondParentId(id int) error {
	_, err := r.db.Exec(r.stmt.DeleteBySecondParentId(), id)
	return err
}
