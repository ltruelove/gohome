package data

import (
	"database/sql"
	"fmt"

	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/ltruelove/gohome/internal/app/setup"
)

const defaultViewNodeSwitchDataSelect string = `SELECT
	Id,
	NodeId,
	ViewId,
	NodeSwitchId,
	Name
	FROM View`

func FetchAllViewNodeSwitchData(db *sql.DB) []models.ViewNodeSwitchData {
	stmt, err := db.Prepare(defaultViewNodeSwitchDataSelect)

	setup.CheckErr(err)
	var listData []models.ViewNodeSwitchData

	rows, err := stmt.Query()
	setup.CheckErr(err)

	for rows.Next() {
		var item models.ViewNodeSwitchData
		rows.Scan(&item.Id,
			&item.NodeId,
			&item.ViewId,
			&item.NodeSwitchId,
			&item.Name)
		listData = append(listData, item)
	}
	defer stmt.Close()

	return listData
}

func FetchViewNodeSwitchData(id int, db *sql.DB) models.ViewNodeSwitchData {
	query := fmt.Sprintf("%s %s", defaultViewNodeSwitchDataSelect, "WHERE id = ?")
	stmt, err := db.Prepare(query)

	setup.CheckErr(err)
	defer stmt.Close()

	var item models.ViewNodeSwitchData

	err = stmt.QueryRow(id).Scan(&item.Id,
		&item.NodeId,
		&item.ViewId,
		&item.NodeSwitchId,
		&item.Name)

	if err != sql.ErrNoRows {
		setup.CheckErr(err)
	}

	return item
}

func CreateViewNodeSwitchData(item *models.ViewNodeSwitchData, db *sql.DB) {
	stmt, err := db.Prepare(`INSERT INTO ViewNodeSwitchData
	(NodeId, ViewId, NodeSwitchId, Name)
	VALUES (?, ?, ?, ?, ?)`)

	setup.CheckErr(err)

	_, err = stmt.Exec(item.NodeId,
		item.ViewId,
		item.NodeSwitchId,
		item.Name)

	defer stmt.Close()

	setup.CheckErr(err)
}

func UpdateViewNodeSwitchData(item *models.ViewNodeSwitchData, db *sql.DB) {
	stmt, err := db.Prepare(`UPDATE ViewNodeSwitchData
	SET NodeId = ?, ViewId = ?, NodeSwitchId = ?, Name = ?
	WHERE id = ?`)

	setup.CheckErr(err)

	_, err = stmt.Exec(item.NodeId,
		item.ViewId,
		item.NodeSwitchId,
		item.Name,
		item.Id)

	defer stmt.Close()

	setup.CheckErr(err)
}

func DeleteViewNodeSwitchData(id int, db *sql.DB) {
	stmt, err := db.Prepare(`DELETE FROM ViewNodeSwitchData
	WHERE id = ?`)

	setup.CheckErr(err)

	_, err = stmt.Exec(id)

	defer stmt.Close()

	setup.CheckErr(err)
}
