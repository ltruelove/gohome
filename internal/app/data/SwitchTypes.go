package data

import (
	"database/sql"

	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/ltruelove/gohome/internal/app/setup"
)

func FetchAllSwitchTypes(db *sql.DB) []models.SwitchType {
	stmt, err := db.Prepare(`SELECT Id, Name
	FROM SwitchType`)

	setup.CheckErr(err)
	var nodeSwitches []models.SwitchType

	rows, err := stmt.Query()
	setup.CheckErr(err)

	for rows.Next() {
		var nodeSwitch models.SwitchType
		rows.Scan(&nodeSwitch.Id,
			&nodeSwitch.Name)
		nodeSwitches = append(nodeSwitches, nodeSwitch)
	}
	defer stmt.Close()

	return nodeSwitches
}

func FetchSwitchType(nodeSwitchTypeId int, db *sql.DB) models.SwitchType {
	stmt, err := db.Prepare("SELECT Id, Name FROM SwitchType WHERE id = ?")
	setup.CheckErr(err)
	defer stmt.Close()

	var nodeSwitch models.SwitchType

	err = stmt.QueryRow(nodeSwitchTypeId).Scan(&nodeSwitch.Id,
		&nodeSwitch.Name)

	if err != sql.ErrNoRows {
		setup.CheckErr(err)
	}

	return nodeSwitch
}
