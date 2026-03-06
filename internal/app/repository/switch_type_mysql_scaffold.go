package repository

import (
	"database/sql"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data/statements"
	"github.com/ltruelove/gohome/internal/app/models"
)

type mysqlSwitchTypeRepository struct {
	db   *sql.DB
	stmt statements.CrudStatement
}

func NewMySQLSwitchTypeRepository(db *sql.DB, cfg *config.Configuration) CrudRepository {
	return &mysqlSwitchTypeRepository{db: db, stmt: statements.NewSwitchTypeStatements(cfg)}
}

func (r *mysqlSwitchTypeRepository) SelectAll() ([]models.Model, error) {
	rows, err := r.db.Query(r.stmt.SelectAll())
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []models.Model
	for rows.Next() {
		var s models.SwitchType
		if err := rows.Scan(&s.Id, &s.Name); err != nil {
			return nil, err
		}
		list = append(list, s)
	}
	return list, nil
}

func (r *mysqlSwitchTypeRepository) SelectByParentId(id int) ([]models.Model, error) { return nil, nil }

func (r *mysqlSwitchTypeRepository) SelectById(id int) (models.Model, error) {
	var s models.SwitchType
	err := r.db.QueryRow(r.stmt.SelectById(), id).Scan(&s.Id, &s.Name)
	return s, err
}

func (r *mysqlSwitchTypeRepository) Insert(data models.Model) (models.Model, error) {
	s, ok := data.(*models.SwitchType)
	if !ok {
		return nil, nil
	}
	var last int
	err := r.db.QueryRow(r.stmt.Insert(), s.Name).Scan(&last)
	if err != nil {
		return nil, err
	}
	s.Id = last
	return s, nil
}

func (r *mysqlSwitchTypeRepository) Update(data models.Model) error {
	s, ok := data.(*models.SwitchType)
	if !ok {
		return nil
	}
	_, err := r.db.Exec(r.stmt.Update(), s.Name, s.Id)
	return err
}

func (r *mysqlSwitchTypeRepository) Delete(id int) error {
	_, err := r.db.Exec(r.stmt.Delete(), id)
	return err
}

func (r *mysqlSwitchTypeRepository) DeleteAll() error {
	_, err := r.db.Exec(r.stmt.DeleteAll())
	return err
}

func (r *mysqlSwitchTypeRepository) DeleteByParentId(id int) error       { return nil }
func (r *mysqlSwitchTypeRepository) DeleteBySecondParentId(id int) error { return nil }
