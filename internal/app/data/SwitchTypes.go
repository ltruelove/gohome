package data

import (
	"database/sql"
	"log"

	"github.com/ltruelove/gohome/internal/app/models"
)

func FetchAllSwitchTypes(db *sql.DB) ([]models.SwitchType, error) {
	stmt, err := db.Prepare(`SELECT Id, Name
	FROM SwitchType`)
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

func FetchSwitchType(nodeSwitchTypeId int, db *sql.DB) (models.SwitchType, error) {
	var nodeSwitch models.SwitchType

	stmt, err := db.Prepare("SELECT Id, Name FROM SwitchType WHERE id = ?")
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
