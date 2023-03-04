package data

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/ltruelove/gohome/internal/app/setup"
)

const defaultViewNodeSwitchDataSelect string = `SELECT
	id,
	nodeid,
	viewid,
	nodeswitchid,
	name
	FROM view`

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

	query := fmt.Sprintf("%s %s", defaultViewNodeSwitchDataSelect, "WHERE id = $1")
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
	stmt, err := db.Prepare(`INSERT INTO viewnodeswitchdata
	(nodeid, viewid, nodeswitchid, name)
	VALUES ($1, $2, $3, $4) RETURNING id`)

	if err != nil {
		log.Println("Error preparing create node switch data sql")
		return err
	}

	lastInsertId := 0

	err = stmt.QueryRow(item.NodeId,
		item.ViewId,
		item.NodeSwitchId,
		item.Name).Scan(&lastInsertId)

	if err != nil {
		log.Println("Error creating node switch data")
		return err
	}

	if err != nil {
		log.Println("Error getting the id of the inserted view node sensor")
		return err
	}

	item.Id = int(lastInsertId)

	defer stmt.Close()

	return nil
}

func UpdateViewNodeSwitchData(item *models.ViewNodeSwitchData, db *sql.DB) error {
	stmt, err := db.Prepare(`UPDATE viewnodeswitchdata
	SET nodeid = $1, viewid = $2, nodeswitchid = $3, name = $4
	WHERE id = $5`)

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
	stmt, err := db.Prepare(`DELETE FROM viewnodeswitchdata
	WHERE id = $1`)

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
