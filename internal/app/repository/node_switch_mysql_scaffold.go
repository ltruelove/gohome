package repository

import (
	"database/sql"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data/statements"
	"github.com/ltruelove/gohome/internal/app/models"
)

type mysqlNodeSwitchRepository struct {
	db   *sql.DB
	stmt statements.CrudStatement
}

func NewMySQLNodeSwitchRepository(db *sql.DB, cfg *config.Configuration) CrudRepository {
	return &mysqlNodeSwitchRepository{db: db, stmt: statements.NewNodeSwitchDataStatements(cfg)}
}

func (r *mysqlNodeSwitchRepository) SelectAll() ([]models.Model, error) {
	rows, err := r.db.Query(r.stmt.SelectAll())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.Model
	for rows.Next() {
		var s models.NodeSwitch
		if err := rows.Scan(&s.Id, &s.NodeId, &s.SwitchTypeId, &s.Name, &s.Pin, &s.MomentaryPressDuration, &s.IsClosedOn); err != nil {
			return nil, err
		}
		list = append(list, s)
	}
	return list, nil
}

func (r *mysqlNodeSwitchRepository) SelectByParentId(id int) ([]models.Model, error) {
	rows, err := r.db.Query(r.stmt.SelectByParentId(), id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []models.Model
	for rows.Next() {
		var s models.NodeSwitch
		if err := rows.Scan(&s.Id, &s.NodeId, &s.SwitchTypeId, &s.Name, &s.Pin, &s.MomentaryPressDuration, &s.IsClosedOn); err != nil {
			return nil, err
		}
		list = append(list, s)
	}
	return list, nil
}

func (r *mysqlNodeSwitchRepository) SelectById(id int) (models.Model, error) {
	var s models.NodeSwitch
	err := r.db.QueryRow(r.stmt.SelectById(), id).Scan(&s.Id, &s.NodeId, &s.SwitchTypeId, &s.Name, &s.Pin, &s.MomentaryPressDuration, &s.IsClosedOn)
	return s, err
}

func (r *mysqlNodeSwitchRepository) Insert(data models.Model) (models.Model, error) {
	item, ok := data.(*models.NodeSwitch)
	if !ok {
		return nil, nil
	}
	var last int
	err := r.db.QueryRow(r.stmt.Insert(), item.NodeId, item.SwitchTypeId, item.Name, item.Pin, item.MomentaryPressDuration, item.IsClosedOn).Scan(&last)
	if err != nil {
		return nil, err
	}
	item.Id = last
	return item, nil
}

func (r *mysqlNodeSwitchRepository) Update(data models.Model) error {
	item, ok := data.(*models.NodeSwitch)
	if !ok {
		return nil
	}
	_, err := r.db.Exec(r.stmt.Update(), item.NodeId, item.SwitchTypeId, item.Name, item.Pin, item.MomentaryPressDuration, item.IsClosedOn, item.Id)
	return err
}

func (r *mysqlNodeSwitchRepository) Delete(id int) error {
	_, err := r.db.Exec(r.stmt.Delete(), id)
	return err
}

func (r *mysqlNodeSwitchRepository) DeleteAll() error {
	_, err := r.db.Exec(r.stmt.DeleteAll())
	return err
}

func (r *mysqlNodeSwitchRepository) DeleteByParentId(id int) error {
	_, err := r.db.Exec(r.stmt.DeleteByParentId(), id)
	return err
}

func (r *mysqlNodeSwitchRepository) DeleteBySecondParentId(id int) error {
	_, err := r.db.Exec(r.stmt.DeleteBySecondParentId(), id)
	return err
}
