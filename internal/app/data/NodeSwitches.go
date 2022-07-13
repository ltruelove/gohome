package data

import (
	"database/sql"

	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/ltruelove/gohome/internal/app/setup"
)

func VerifyNodeSwitchIdIsNew(nodeId int, db *sql.DB) bool {
	node := FetchNodeSwitch(nodeId, db)
	return node.Id > 0
}

func FetchAllNodeSwitches(db *sql.DB) []models.NodeSwitch {
	stmt, err := db.Prepare(`SELECT
		Id,
		NodeId,
		SwitchTypeId,
		Name,
		Pin FROM NodeSwitch`)

	setup.CheckErr(err)
	var nodeSwitches []models.NodeSwitch

	rows, err := stmt.Query()
	setup.CheckErr(err)

	for rows.Next() {
		var nodeSwitch models.NodeSwitch

		rows.Scan(&nodeSwitch.Id,
			&nodeSwitch.NodeId,
			&nodeSwitch.SwitchTypeId,
			&nodeSwitch.Name,
			&nodeSwitch.Pin)

		nodeSwitches = append(nodeSwitches, nodeSwitch)
	}
	defer stmt.Close()

	return nodeSwitches
}

func FetchNodeSwitch(nodeId int, db *sql.DB) models.NodeSwitch {
	stmt, err := db.Prepare(`SELECT
		Id,
		NodeId,
		SwitchTypeId,
		Name
		Pin FROM NodeSwitch WHERE id = ?`)
	setup.CheckErr(err)
	defer stmt.Close()

	var nodeSwitch models.NodeSwitch

	err = stmt.QueryRow(nodeId).Scan(&nodeSwitch.Id,
		&nodeSwitch.NodeId,
		&nodeSwitch.SwitchTypeId,
		&nodeSwitch.Name,
		&nodeSwitch.Pin)

	if err != sql.ErrNoRows {
		setup.CheckErr(err)
	}

	return nodeSwitch
}

func CreateNodeSwitch(nodeSwitch *models.NodeSwitch, db *sql.DB) {
	stmt, err := db.Prepare(`INSERT INTO NodeSwitch
	(Id, NodeId, SwitchTypeId, Name, Pin)
	VALUES (?, ?, ?, ?, ?)`)

	setup.CheckErr(err)

	_, err = stmt.Exec(nodeSwitch.Id,
		nodeSwitch.NodeId,
		nodeSwitch.SwitchTypeId,
		nodeSwitch.Name,
		nodeSwitch.Pin)

	defer stmt.Close()

	setup.CheckErr(err)
}

func UpdateNodeSwitch(nodeSwitch *models.NodeSwitch, db *sql.DB) {
	stmt, err := db.Prepare(`UPDATE NodeSwitch
	SET NodeId = ?, SwitchTypeId = ?, Name = ?, Pin = ?
	WHERE id = ?`)

	setup.CheckErr(err)

	_, err = stmt.Exec(nodeSwitch.NodeId,
		nodeSwitch.SwitchTypeId,
		nodeSwitch.Name,
		nodeSwitch.Pin,
		nodeSwitch.Id)

	defer stmt.Close()

	setup.CheckErr(err)
}

func DeleteNodeSwitch(nodeSwitchId int, db *sql.DB) {
	stmt, err := db.Prepare(`DELETE FROM NodeSwitch
	WHERE id = ?`)

	setup.CheckErr(err)

	_, err = stmt.Exec(nodeSwitchId)

	defer stmt.Close()

	setup.CheckErr(err)
}
