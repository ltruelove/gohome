package data

import (
	"database/sql"
	"log"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/models"
)

type SwitchTypeData struct {
	DB         *sql.DB
	Statements *SwitchTypeStatements
}

func NewSwitchTypeData(db *sql.DB, config *config.Configuration) *SwitchTypeData {
	return &SwitchTypeData{
		DB:         db,
		Statements: NewSwitchTypeStatements(config),
	}
}

func (switchType *SwitchTypeData) FetchAllSwitchTypes() ([]models.SwitchType, error) {
	stmt, err := switchType.DB.Prepare(switchType.Statements.SelectAllSwitchTypes())
	if err != nil {
		log.Println("Error preparing fetch all switch types sql")
		return nil, err
	}

	var nodeSwitches []models.SwitchType

	rows, err := stmt.Query()
	if err != nil {
		log.Println("Error querying for all switch types")
		return nil, err
	}
	defer stmt.Close()

	for rows.Next() {
		var nodeSwitch models.SwitchType
		rows.Scan(&nodeSwitch.Id,
			&nodeSwitch.Name)
		nodeSwitches = append(nodeSwitches, nodeSwitch)
	}

	return nodeSwitches, nil
}

func (switchType *SwitchTypeData) FetchSwitchType(nodeSwitchTypeId int) (models.SwitchType, error) {
	var nodeSwitch models.SwitchType

	stmt, err := switchType.DB.Prepare(switchType.Statements.SelectSwitchTypeById())
	if err != nil {
		log.Println("Error preparing fetch switch type sql")
		return nodeSwitch, err
	}

	err = stmt.QueryRow(nodeSwitchTypeId).Scan(&nodeSwitch.Id,
		&nodeSwitch.Name)

	if err != nil {
		log.Println("Error querying for the switch type")
		return nodeSwitch, err
	}

	defer stmt.Close()

	return nodeSwitch, nil
}
