package data

import (
	"database/sql"
	"log"

	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/ltruelove/gohome/internal/app/viewModels"
)

func VerifyNodeSwitchIdIsNew(nodeId int, db *sql.DB) (bool, error) {
	item, err := FetchNodeSwitch(nodeId, db)
	if err != nil {
		log.Println("Error fetching node switch")
		return false, err
	}

	return item.Id > 0, nil
}

func FetchAllNodeSwitches(db *sql.DB) ([]viewModels.NodeSwitchVM, error) {
	stmt, err := db.Prepare(`SELECT
		ns.id,
		ns.nodeid,
		ns.switchtypeid,
		ns.name,
		ns.pin,
		ns.momentarypressduration,
		ns.isclosedon,
		st.name AS switchtypename
		FROM nodeswitch AS ns
		INNER JOIN switchtype AS st ON st.id = ns.switchtypeid`)
	if err != nil {
		log.Println("Error preparing all node switches sql")
		return nil, err
	}

	var nodeSwitches []viewModels.NodeSwitchVM

	rows, err := stmt.Query()
	if err != nil {
		log.Println("Error querying for all node switches")
		return nil, err
	}

	defer stmt.Close()

	for rows.Next() {
		var nodeSwitch viewModels.NodeSwitchVM

		err := rows.Scan(&nodeSwitch.Id,
			&nodeSwitch.NodeId,
			&nodeSwitch.SwitchTypeId,
			&nodeSwitch.Name,
			&nodeSwitch.Pin,
			&nodeSwitch.MomentaryPressDuration,
			&nodeSwitch.IsClosedOn,
			&nodeSwitch.SwitchTypeName)

		if err != nil {
			log.Println("Error scanning node switch")
			return nil, err
		}

		nodeSwitches = append(nodeSwitches, nodeSwitch)
	}

	return nodeSwitches, nil
}

func FetchNodeSwitch(id int, db *sql.DB) (models.NodeSwitch, error) {
	var nodeSwitch models.NodeSwitch

	stmt, err := db.Prepare(`SELECT
		id,
		nodeid,
		switchtypeid,
		name,
		pin,
		momentarypressduration,
		isclosedon FROM nodeswitch WHERE id = $1`)
	if err != nil {
		log.Println("Error preparing fetch node switch sql")
		return nodeSwitch, err
	}

	err = stmt.QueryRow(id).Scan(&nodeSwitch.Id,
		&nodeSwitch.NodeId,
		&nodeSwitch.SwitchTypeId,
		&nodeSwitch.Name,
		&nodeSwitch.Pin,
		&nodeSwitch.MomentaryPressDuration,
		&nodeSwitch.IsClosedOn)

	if err != nil {
		log.Println("Error querying for node switch")
		return nodeSwitch, err
	}

	defer stmt.Close()

	return nodeSwitch, nil
}

func CreateNodeSwitch(item *models.NodeSwitch, db *sql.DB) error {
	stmt, err := db.Prepare(`INSERT INTO nodeswitch
	(nodeid, switchtypeid, name, pin, momentarypressduration, isclosedon)
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`)

	if err != nil {
		log.Println("Error preparing create node switch sql")
		return err
	}

	lastInsertId := 0

	err = stmt.QueryRow(&item.NodeId,
		&item.SwitchTypeId,
		&item.Name,
		&item.Pin,
		&item.MomentaryPressDuration,
		&item.IsClosedOn).Scan(&lastInsertId)

	if err != nil {
		log.Println("Error creating node sensor")
		return err
	}

	defer stmt.Close()

	if err != nil {
		log.Println("Error getting the id of the inserted node sensor")
		return err
	}

	item.Id = int(lastInsertId)

	return nil
}

func UpdateNodeSwitch(nodeSwitch *models.NodeSwitch, db *sql.DB) error {
	stmt, err := db.Prepare(`UPDATE nodeswitch
	SET nodeid = $1, switchtypeid = $2, name = $3, pin = $4, momentarypressduration = $5, isclosedon = $6
	WHERE id = $7`)

	if err != nil {
		log.Println("Error preparing update node switch sql")
		return err
	}

	_, err = stmt.Exec(nodeSwitch.NodeId,
		nodeSwitch.SwitchTypeId,
		nodeSwitch.Name,
		nodeSwitch.Pin,
		nodeSwitch.Id,
		nodeSwitch.MomentaryPressDuration,
		nodeSwitch.IsClosedOn)

	if err != nil {
		log.Println("Error updating node switch")
		return err
	}

	defer stmt.Close()

	return nil
}

func DeleteNodeSwitch(nodeSwitchId int, db *sql.DB) error {
	stmt, err := db.Prepare(`DELETE FROM nodeswitch
	WHERE id = $1`)

	if err != nil {
		log.Println("Error preparing delete node switch sql")
		return err
	}

	_, err = stmt.Exec(nodeSwitchId)

	if err != nil {
		log.Println("Error deleting node switch")
		return err
	}

	defer stmt.Close()

	return nil
}

func DeleteAllNodeSwitches(nodeId int, db *sql.DB) error {
	stmt, err := db.Prepare(`DELETE FROM nodeswitch
	WHERE nodeid = $1`)

	if err != nil {
		log.Println("Error preparing delete node switch sql")
		return err
	}

	_, err = stmt.Exec(nodeId)

	if err != nil {
		log.Println("Error deleting node switch")
		return err
	}

	defer stmt.Close()

	return nil
}
