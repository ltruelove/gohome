package repository

import (
	"database/sql"
	"log"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data/statements"
	"github.com/ltruelove/gohome/internal/app/viewModels"
)

type mysqlCompoundRepository struct {
	db   *sql.DB
	stmt *statements.CompoundStatements
}

func NewMySQLCompoundRepository(db *sql.DB, config *config.Configuration) CompoundRepository {
	return &mysqlCompoundRepository{db: db, stmt: statements.NewCompoundStatements(config)}
}

func (r *mysqlCompoundRepository) FetchViewNodeSensorDataByViewId(viewId int) ([]viewModels.ViewNodeSensorVM, error) {
	q := r.stmt.SelectViewNodeSensorDataByViewId()
	rows, err := r.db.Query(q, viewId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []viewModels.ViewNodeSensorVM
	for rows.Next() {
		var item viewModels.ViewNodeSensorVM
		if err := rows.Scan(&item.Id, &item.NodeId, &item.ViewId, &item.NodeSensorId, &item.Name, &item.NodeName, &item.SensorName, &item.SensorTypeName); err != nil {
			log.Println("error scanning view node sensor row:", err)
			return nil, err
		}
		list = append(list, item)
	}
	return list, nil
}

func (r *mysqlCompoundRepository) FetchViewNodeSwitchDataByViewId(viewId int) ([]viewModels.ViewNodeSwitchVM, error) {
	q := r.stmt.SelectViewNodeSwitchDataByViewId()
	rows, err := r.db.Query(q, viewId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []viewModels.ViewNodeSwitchVM
	for rows.Next() {
		var item viewModels.ViewNodeSwitchVM
		if err := rows.Scan(&item.Id, &item.NodeId, &item.ViewId, &item.NodeSwitchId, &item.Name, &item.NodeName, &item.SwitchName, &item.SwitchTypeName); err != nil {
			log.Println("error scanning view node switch row:", err)
			return nil, err
		}
		list = append(list, item)
	}
	return list, nil
}
