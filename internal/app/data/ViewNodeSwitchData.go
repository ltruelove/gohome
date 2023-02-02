package data

import (
	"database/sql"
	"fmt"
	"log"

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

func FetchAllViewNodeSwitchData(db *sql.DB) ([]models.ViewNodeSwitchData, error) {
	stmt, err := db.Prepare(defaultViewNodeSwitchDataSelect)

	if err != nil {
		log.Println("Error preparing all node switch data sql")
		return nil, err
	}
	var listData []models.ViewNodeSwitchData

	rows, err := stmt.Query()
	if err != nil {
		log.Println("Error querying for all node switch data")
		return nil, err
	}

	defer stmt.Close()

	for rows.Next() {
		var item models.ViewNodeSwitchData
		err := rows.Scan(&item.Id,
			&item.NodeId,
			&item.ViewId,
			&item.NodeSwitchId,
			&item.Name)

		if err != nil {
			log.Println("Error scanning node switch data")
			return nil, err
		}

		listData = append(listData, item)
	}

	return listData, nil
}

func FetchViewNodeSwitchData(id int, db *sql.DB) (models.ViewNodeSwitchData, error) {
	var item models.ViewNodeSwitchData

	query := fmt.Sprintf("%s %s", defaultViewNodeSwitchDataSelect, "WHERE id = ?")
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println("Error preparing fetch node switch data sql")
		return item, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&item.Id,
		&item.NodeId,
		&item.ViewId,
		&item.NodeSwitchId,
		&item.Name)

	if err != nil {
		log.Println("Error querying for node switch data")
		return item, err
	}

	return item, nil
}

func CreateViewNodeSwitchData(item *models.ViewNodeSwitchData, db *sql.DB) error {
	stmt, err := db.Prepare(`INSERT INTO ViewNodeSwitchData
	(NodeId, ViewId, NodeSwitchId, Name)
	VALUES (?, ?, ?, ?)`)

	if err != nil {
		log.Println("Error preparing create node switch data sql")
		return err
	}

	result, err := stmt.Exec(item.NodeId,
		item.ViewId,
		item.NodeSwitchId,
		item.Name)

	if err != nil {
		log.Println("Error creating node switch data")
		return err
	}

	lastInsertId, err := result.LastInsertId()

	if err != nil {
		log.Println("Error getting the id of the inserted view node sensor")
		return err
	}

	item.Id = int(lastInsertId)

	defer stmt.Close()

	return nil
}

func UpdateViewNodeSwitchData(item *models.ViewNodeSwitchData, db *sql.DB) error {
	stmt, err := db.Prepare(`UPDATE ViewNodeSwitchData
	SET NodeId = ?, ViewId = ?, NodeSwitchId = ?, Name = ?
	WHERE id = ?`)

	if err != nil {
		log.Println("Error preparing update node switch data sql")
		return err
	}

	setup.CheckErr(err)

	_, err = stmt.Exec(item.NodeId,
		item.ViewId,
		item.NodeSwitchId,
		item.Name,
		item.Id)

	if err != nil {
		log.Println("Error updating node switch data")
		return err
	}

	defer stmt.Close()

	return nil
}

func DeleteViewNodeSwitchData(id int, db *sql.DB) error {
	stmt, err := db.Prepare(`DELETE FROM ViewNodeSwitchData
	WHERE id = ?`)

	if err != nil {
		log.Println("Error preparing delete node switch data sql")
		return err
	}

	_, err = stmt.Exec(id)

	if err != nil {
		log.Println("Error deleting node switch data")
		return err
	}

	defer stmt.Close()

	return nil
}
