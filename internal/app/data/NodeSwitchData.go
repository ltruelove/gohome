package data

import (
	"database/sql"
	"errors"
	"log"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data/statements"
	"github.com/ltruelove/gohome/internal/app/models"
)

type NodeSwitchData struct {
	db   *sql.DB
	stmt statements.CrudStatement
}

func NewNodeSwitchData(db *sql.DB, config *config.Configuration) CrudDataInterface {
	return &NodeSwitchData{
		db:   db,
		stmt: statements.NewNodeSwitchDataStatements(config),
	}
}

func (d NodeSwitchData) DB() *sql.DB {
	return d.db
}

func (d NodeSwitchData) Stmt() statements.CrudStatement {
	return d.stmt
}

func (d NodeSwitchData) SelectAll() ([]models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectAll())
	if err != nil {
		log.Println("Error preparing fetch all node switches sql")
		return nil, err
	}

	var nodeSwitches []models.Model

	rows, err := stmt.Query()
	if err != nil {
		log.Println("Error querying for all node switches")
		return nil, err
	}
	defer stmt.Close()

	for rows.Next() {
		var nodeSwitch models.NodeSwitch

		err := rows.Scan(&nodeSwitch.Id,
			&nodeSwitch.NodeId,
			&nodeSwitch.SwitchTypeId,
			&nodeSwitch.Name,
			&nodeSwitch.Pin,
			&nodeSwitch.MomentaryPressDuration,
			&nodeSwitch.IsClosedOn)

		if err != nil {
			log.Println("Error scanning node switch")
			return nil, err
		}

		nodeSwitches = append(nodeSwitches, nodeSwitch)
	}

	return nodeSwitches, nil
}

func (d NodeSwitchData) SelectById(id int) (models.Model, error) {
	var nodeSwitch models.NodeSwitch

	stmt, err := d.DB().Prepare(d.Stmt().SelectById())
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

// SelectByParentId retrieves NodeSwitch records by NodeId
func (d NodeSwitchData) SelectByParentId(id int) ([]models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectByParentId())
	if err != nil {
		log.Println("Error preparing fetch node switches by parent id sql")
		return nil, err
	}

	var nodeSwitches []models.Model

	rows, err := stmt.Query(id)
	if err != nil {
		log.Println("Error querying for node switches by parent id")
		return nil, err
	}
	defer stmt.Close()

	for rows.Next() {
		var nodeSwitch models.NodeSwitch

		err := rows.Scan(&nodeSwitch.Id,
			&nodeSwitch.NodeId,
			&nodeSwitch.SwitchTypeId,
			&nodeSwitch.Name,
			&nodeSwitch.Pin,
			&nodeSwitch.MomentaryPressDuration,
			&nodeSwitch.IsClosedOn)

		if err != nil {
			log.Println("Error scanning node switch")
			return nil, err
		}

		nodeSwitches = append(nodeSwitches, nodeSwitch)
	}

	return nodeSwitches, nil
}

// SelectBySecondParentId retrieves NodeSwitch records by SwitchTypeId
func (d NodeSwitchData) SelectBySecondParentId(id int) ([]models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectBySecondParentId())
	if err != nil {
		log.Println("Error preparing fetch node switches by second parent id sql")
		return nil, err
	}

	var nodeSwitches []models.Model

	rows, err := stmt.Query(id)
	if err != nil {
		log.Println("Error querying for node switches by second parent id")
		return nil, err
	}
	defer stmt.Close()

	for rows.Next() {
		var nodeSwitch models.NodeSwitch

		err := rows.Scan(&nodeSwitch.Id,
			&nodeSwitch.NodeId,
			&nodeSwitch.SwitchTypeId,
			&nodeSwitch.Name,
			&nodeSwitch.Pin,
			&nodeSwitch.MomentaryPressDuration,
			&nodeSwitch.IsClosedOn)

		if err != nil {
			log.Println("Error scanning node switch")
			return nil, err
		}

		nodeSwitches = append(nodeSwitches, nodeSwitch)
	}

	return nodeSwitches, nil
}

func (d NodeSwitchData) Insert(data models.Model) (models.Model, error) {
	item, ok := data.(*models.NodeSwitch)
	if !ok {
		log.Println("Error casting data to NodeSwitch")
		return nil, errors.New("invalid data type for NodeSwitch")
	}

	stmt, err := d.DB().Prepare(d.Stmt().Insert())
	if err != nil {
		log.Println("Error preparing create node switch sql")
		return nil, err
	}

	lastInsertId := 0
	isClosedOnInt := 0
	if item.IsClosedOn {
		isClosedOnInt = 1
	}

	err = stmt.QueryRow(&item.NodeId,
		&item.SwitchTypeId,
		&item.Name,
		&item.Pin,
		&item.MomentaryPressDuration,
		&isClosedOnInt).Scan(&lastInsertId)

	if err != nil {
		log.Println("Error creating node switch")
		return nil, err
	}

	defer stmt.Close()

	item.Id = int(lastInsertId)

	return item, nil
}

func (d NodeSwitchData) Update(data models.Model) error {
	item, ok := data.(*models.NodeSwitch)
	if !ok {
		log.Println("Error casting data to NodeSwitch")
		return errors.New("invalid data type for NodeSwitch")
	}

	stmt, err := d.DB().Prepare(d.Stmt().Update())
	if err != nil {
		log.Println("Error preparing update node switch sql")
		return err
	}

	_, err = stmt.Exec(item.NodeId,
		item.SwitchTypeId,
		item.Name,
		item.Pin,
		item.MomentaryPressDuration,
		item.IsClosedOn,
		item.Id)

	if err != nil {
		log.Println("Error updating node switch")
		return err
	}

	defer stmt.Close()

	return nil
}

func (d NodeSwitchData) Delete(id int) error {
	stmt, err := d.DB().Prepare(d.Stmt().Delete())
	if err != nil {
		log.Println("Error preparing delete node switch sql")
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		log.Println("Error deleting node switch")
		return err
	}

	defer stmt.Close()

	return nil
}

func (d NodeSwitchData) DeleteAll() error {
	stmt, err := d.DB().Prepare(d.Stmt().DeleteAll())
	if err != nil {
		log.Println("Error preparing delete all node switches sql")
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Println("Error deleting all node switches")
		return err
	}

	defer stmt.Close()

	return nil
}

func (d NodeSwitchData) DeleteByParentId(id int) error {
	stmt, err := d.DB().Prepare(d.Stmt().DeleteByParentId())
	if err != nil {
		log.Println("Error preparing delete node switches by parent id sql")
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		log.Println("Error deleting node switches by parent id")
		return err
	}

	defer stmt.Close()

	return nil
}

func (d NodeSwitchData) DeleteBySecondParentId(id int) error {
	stmt, err := d.DB().Prepare(d.Stmt().DeleteBySecondParentId())
	if err != nil {
		log.Println("Error preparing delete node switches by second parent id sql")
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		log.Println("Error deleting node switches by second parent id")
		return err
	}

	defer stmt.Close()

	return nil
}

func FetchNodeSwitch(id int, db *sql.DB) (models.NodeSwitch, error) {
	var ns models.NodeSwitch
	stmt, err := db.Prepare(`SELECT id, nodeid, switchtypeid, name, pin, momentarypressduration, isclosedon FROM nodeswitch WHERE id = $1`)
	if err != nil {
		log.Println("Error preparing fetch node switch sql")
		return ns, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&ns.Id, &ns.NodeId, &ns.SwitchTypeId, &ns.Name, &ns.Pin, &ns.MomentaryPressDuration, &ns.IsClosedOn)
	if err != nil {
		log.Println("Error querying for node switch")
		return ns, err
	}

	return ns, nil
}

func FetchNodeSwitches(nodeId int, db *sql.DB) ([]models.NodeSwitch, error) {
	stmt, err := db.Prepare(`SELECT id, nodeid, switchtypeid, name, pin, momentarypressduration, isclosedon FROM nodeswitch WHERE nodeid = $1`)
	if err != nil {
		log.Println("Error preparing fetch node switches sql")
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(nodeId)
	if err != nil {
		log.Println("Error querying for node switches")
		return nil, err
	}

	var list []models.NodeSwitch
	for rows.Next() {
		var ns models.NodeSwitch
		if err := rows.Scan(&ns.Id, &ns.NodeId, &ns.SwitchTypeId, &ns.Name, &ns.Pin, &ns.MomentaryPressDuration, &ns.IsClosedOn); err != nil {
			log.Println("Error scanning node switch")
			return nil, err
		}
		list = append(list, ns)
	}

	return list, nil
}
