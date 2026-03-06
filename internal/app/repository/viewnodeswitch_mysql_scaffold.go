package repository

import (
	"database/sql"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data/statements"
	"github.com/ltruelove/gohome/internal/app/models"
)

type mysqlViewNodeSwitchRepository struct {
	db   *sql.DB
	stmt statements.CrudStatement
}

func NewMySQLViewNodeSwitchRepository(db *sql.DB, cfg *config.Configuration) CrudRepository {
	return &mysqlViewNodeSwitchRepository{db: db, stmt: statements.NewViewNodeSwitchDataStatements(cfg)}
}

func (r *mysqlViewNodeSwitchRepository) SelectAll() ([]models.Model, error) {
	rows, err := r.db.Query(r.stmt.SelectAll())
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []models.Model
	for rows.Next() {
		var item models.ViewNodeSwitchData
		if err := rows.Scan(&item.Id, &item.NodeId, &item.ViewId, &item.NodeSwitchId, &item.Name); err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, nil
}

func (r *mysqlViewNodeSwitchRepository) SelectByParentId(id int) ([]models.Model, error) {
	rows, err := r.db.Query(r.stmt.SelectByParentId(), id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []models.Model
	for rows.Next() {
		var item models.ViewNodeSwitchData
		if err := rows.Scan(&item.Id, &item.NodeId, &item.ViewId, &item.NodeSwitchId, &item.Name); err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, nil
}

func (r *mysqlViewNodeSwitchRepository) SelectById(id int) (models.Model, error) {
	var item models.ViewNodeSwitchData
	err := r.db.QueryRow(r.stmt.SelectById(), id).Scan(&item.Id, &item.NodeId, &item.ViewId, &item.NodeSwitchId, &item.Name)
	return item, err
}

func (r *mysqlViewNodeSwitchRepository) Insert(data models.Model) (models.Model, error) {
	item, ok := data.(*models.ViewNodeSwitchData)
	if !ok {
		return nil, nil
	}
	var last int
	err := r.db.QueryRow(r.stmt.Insert(), item.NodeId, item.ViewId, item.NodeSwitchId, item.Name).Scan(&last)
	if err != nil {
		return nil, err
	}
	item.Id = last
	return item, nil
}

func (r *mysqlViewNodeSwitchRepository) Update(data models.Model) error {
	item, ok := data.(*models.ViewNodeSwitchData)
	if !ok {
		return nil
	}
	_, err := r.db.Exec(r.stmt.Update(), item.NodeId, item.ViewId, item.NodeSwitchId, item.Name, item.Id)
	return err
}

func (r *mysqlViewNodeSwitchRepository) Delete(id int) error {
	_, err := r.db.Exec(r.stmt.Delete(), id)
	return err
}

func (r *mysqlViewNodeSwitchRepository) DeleteAll() error {
	_, err := r.db.Exec(r.stmt.DeleteAll())
	return err
}

func (r *mysqlViewNodeSwitchRepository) DeleteByParentId(id int) error {
	_, err := r.db.Exec(r.stmt.DeleteByParentId(), id)
	return err
}

func (r *mysqlViewNodeSwitchRepository) DeleteBySecondParentId(id int) error {
	_, err := r.db.Exec(r.stmt.DeleteBySecondParentId(), id)
	return err
}
