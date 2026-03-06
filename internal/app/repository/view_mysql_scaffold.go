package repository

import (
	"database/sql"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data/statements"
	"github.com/ltruelove/gohome/internal/app/models"
)

type mysqlViewRepository struct {
	db   *sql.DB
	stmt *statements.ViewDataStatements
}

func NewMySQLViewRepository(db *sql.DB, cfg *config.Configuration) CrudRepository {
	return &mysqlViewRepository{db: db, stmt: statements.NewViewDataStatements(cfg)}
}

func (r *mysqlViewRepository) SelectAll() ([]models.Model, error) {
	rows, err := r.db.Query(r.stmt.SelectAll())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.Model
	for rows.Next() {
		var v models.View
		if err := rows.Scan(&v.Id, &v.Name); err != nil {
			return nil, err
		}
		list = append(list, v)
	}
	return list, nil
}

func (r *mysqlViewRepository) SelectByParentId(id int) ([]models.Model, error) {
	return nil, nil
}

func (r *mysqlViewRepository) SelectById(id int) (models.Model, error) {
	var v models.View
	err := r.db.QueryRow(r.stmt.SelectById(), id).Scan(&v.Id, &v.Name)
	return v, err
}

func (r *mysqlViewRepository) Insert(data models.Model) (models.Model, error) {
	v, ok := data.(*models.View)
	if !ok {
		return nil, nil
	}
	var lastId int
	err := r.db.QueryRow(r.stmt.Insert(), v.Name).Scan(&lastId)
	if err != nil {
		return nil, err
	}
	v.Id = lastId
	return v, nil
}

func (r *mysqlViewRepository) Update(data models.Model) error {
	v, ok := data.(*models.View)
	if !ok {
		return nil
	}
	_, err := r.db.Exec(r.stmt.Update(), v.Name, v.Id)
	return err
}

func (r *mysqlViewRepository) Delete(id int) error {
	_, err := r.db.Exec(r.stmt.Delete(), id)
	return err
}

func (r *mysqlViewRepository) DeleteAll() error {
	_, err := r.db.Exec(r.stmt.DeleteAll())
	return err
}

func (r *mysqlViewRepository) DeleteByParentId(id int) error       { return nil }
func (r *mysqlViewRepository) DeleteBySecondParentId(id int) error { return nil }
