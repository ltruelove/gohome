package data

import (
	"database/sql"
	"log"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data/statements"
	"github.com/ltruelove/gohome/internal/app/setup"
	"github.com/ltruelove/gohome/internal/app/viewModels"
)

type CompoundData struct {
	DB   *sql.DB
	Stmt *statements.CompoundStatements
}

func NewCompoundData(db *sql.DB, config *config.Configuration) *CompoundData {
	return &CompoundData{
		DB:   db,
		Stmt: statements.NewCompoundStatements(config),
	}
}

func (compoundData *CompoundData) FetchViewNodeSensorDataByViewId(viewId int) ([]viewModels.ViewNodeSensorVM, error) {
	stmt, err := compoundData.DB.Prepare(compoundData.Stmt.SelectViewNodeSensorDataByViewId())
	if err != nil {
		log.Println("Error preparing fetch node sensor data by view id sql")
		return nil, err
	}

	var listData []viewModels.ViewNodeSensorVM

	rows, err := stmt.Query(viewId)
	if err != nil {
		log.Println("Error querying for node sensor data by view id")
		return nil, err
	}

	for rows.Next() {
		var item viewModels.ViewNodeSensorVM

		err := rows.Scan(&item.Id,
			&item.NodeId,
			&item.ViewId,
			&item.NodeSensorId,
			&item.Name,
			&item.NodeName,
			&item.SensorName,
			&item.SensorTypeName)

		if err != nil {
			log.Println("Error scanning node sensor data by view id")
			return nil, err
		}
		listData = append(listData, item)
	}
	defer stmt.Close()

	return listData, nil
}

func (compoundData *CompoundData) FetchViewNodeSwitchDataByViewId(viewId int) ([]viewModels.ViewNodeSwitchVM, error) {
	log.Printf("Finding all switches for view with id %d", viewId)
	stmt, err := compoundData.DB.Prepare(compoundData.Stmt.SelectViewNodeSwitchDataByViewId())

	if err != nil {
		log.Println("Error preparing fetch view switch data sql")
		return nil, err
	}

	var switchDataList []viewModels.ViewNodeSwitchVM

	rows, err := stmt.Query(viewId)

	if err != nil {
		log.Println("Error querying for view")
		return nil, err
	}

	defer stmt.Close()
	setup.CheckErr(err)

	for rows.Next() {
		var switchDataItem viewModels.ViewNodeSwitchVM

		err := rows.Scan(&switchDataItem.Id,
			&switchDataItem.NodeId,
			&switchDataItem.ViewId,
			&switchDataItem.NodeSwitchId,
			&switchDataItem.Name,
			&switchDataItem.NodeName,
			&switchDataItem.SwitchName,
			&switchDataItem.SwitchTypeName)

		if err != nil {
			log.Println("Error scanning view")
			return nil, err
		}

		switchDataList = append(switchDataList, switchDataItem)
	}

	log.Printf("Found %d switches", len(switchDataList))
	return switchDataList, nil
}
