package repository

import (
	"database/sql"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data/statements"
	"github.com/ltruelove/gohome/internal/app/models"
)

type mysqlSensorTypeRepository struct {
	db   *sql.DB
	stmt statements.CrudStatement
}

func NewMySQLSensorTypeRepository(db *sql.DB, cfg *config.Configuration) CrudRepository {
	return &mysqlSensorTypeRepository{db: db, stmt: statements.NewSensorTypeStatements(cfg)}
}

func (r *mysqlSensorTypeRepository) SelectAll() ([]models.Model, error) {
	rows, err := r.db.Query(r.stmt.SelectAll())
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []models.Model
	for rows.Next() {
		var s models.SensorType
		if err := rows.Scan(&s.Id, &s.SensorTypeId, &s.Name, &s.ValueType); err != nil {
			return nil, err
		}
		list = append(list, s)
	}
	return list, nil
}

func (r *mysqlSensorTypeRepository) SelectByParentId(id int) ([]models.Model, error) { return nil, nil }

func (r *mysqlSensorTypeRepository) SelectById(id int) (models.Model, error) {
	var s models.SensorType
	err := r.db.QueryRow(r.stmt.SelectById(), id).Scan(&s.Id, &s.SensorTypeId, &s.Name, &s.ValueType)
	return s, err
}

func (r *mysqlSensorTypeRepository) Insert(data models.Model) (models.Model, error) {
	s, ok := data.(*models.SensorType)
	if !ok {
		return nil, nil
	}
	var last int
	err := r.db.QueryRow(r.stmt.Insert(), s.SensorTypeId, s.Name, s.ValueType).Scan(&last)
	if err != nil {
		return nil, err
	}
	s.Id = last
	return s, nil
}

func (r *mysqlSensorTypeRepository) Update(data models.Model) error {
	s, ok := data.(*models.SensorType)
	if !ok {
		return nil
	}
	_, err := r.db.Exec(r.stmt.Update(), s.SensorTypeId, s.Name, s.ValueType, s.Id)
	return err
}

func (r *mysqlSensorTypeRepository) Delete(id int) error {
	_, err := r.db.Exec(r.stmt.Delete(), id)
	return err
}

func (r *mysqlSensorTypeRepository) DeleteAll() error {
	_, err := r.db.Exec(r.stmt.DeleteAll())
	return err
}

func (r *mysqlSensorTypeRepository) DeleteByParentId(id int) error       { return nil }
func (r *mysqlSensorTypeRepository) DeleteBySecondParentId(id int) error { return nil }
