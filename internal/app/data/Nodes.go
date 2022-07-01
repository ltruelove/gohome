package data

import (
	"database/sql"

	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/ltruelove/gohome/internal/app/setup"
)

func VerifyNodeIdIsNew(nodeId string, db *sql.DB) bool {
	node := FetchNode(nodeId, db)
	return node.Id > 0
}

func FetchAllNodes(db *sql.DB) []models.Node {
	stmt, err := db.Prepare(`SELECT Id, Name FROM Nodes`)

	setup.CheckErr(err)
	var nodes []models.Node

	rows, err := stmt.Query()
	setup.CheckErr(err)

	for rows.Next() {
		var node models.Node
		rows.Scan(&node.Id,
			&node.Name)
		nodes = append(nodes, node)
	}
	defer stmt.Close()

	return nodes
}

func FetchNode(nodeId string, db *sql.DB) models.Node {
	stmt, err := db.Prepare("SELECT Id, Name FROM Nodes WHERE id = ?")
	setup.CheckErr(err)
	defer stmt.Close()

	var node models.Node

	err = stmt.QueryRow(nodeId).Scan(&node.Id,
		&node.Name)

	if err != sql.ErrNoRows {
		setup.CheckErr(err)
	}

	return node
}

func AddNewNode(node *models.Node, db *sql.DB) {
	stmt, err := db.Prepare(`INSERT INTO Nodes
	(Id, Name, Mac)
	VALUES (?, ?, ?)`)

	setup.CheckErr(err)

	_, err = stmt.Exec(node.Id,
		node.Name,
		node.Mac)

	defer stmt.Close()

	setup.CheckErr(err)
}

func UpdateNode(node *models.Node, db *sql.DB) {
	stmt, err := db.Prepare(`UPDATE Nodes
	SET Name = ?, Mac = ?
	WHERE id = ?`)

	setup.CheckErr(err)

	_, err = stmt.Exec(node.Name,
		node.Mac,
		node.Id)

	defer stmt.Close()

	setup.CheckErr(err)
}

func DeleteNode(nodeId string, db *sql.DB) {
	stmt, err := db.Prepare(`DELETE FROM Nodes
	WHERE id = ?`)

	setup.CheckErr(err)

	_, err = stmt.Exec(nodeId)

	defer stmt.Close()

	setup.CheckErr(err)
}
