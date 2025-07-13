package data

import (
	"database/sql"
	"log"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data/statements"
	"github.com/ltruelove/gohome/internal/app/models"
)

type SwitchTypeData struct {
	db   *sql.DB
	stmt statements.CrudStatement
}

func NewSwitchTypeData(db *sql.DB, config *config.Configuration) *SwitchTypeData {
	return &SwitchTypeData{
		db:   db,
		stmt: statements.NewSwitchTypeStatements(config),
	}
}

func (d *SwitchTypeData) DB() *sql.DB {
	return d.db
}

func (d *SwitchTypeData) Stmt() statements.CrudStatement {
	return d.stmt
}

func (d *SwitchTypeData) SelectAll() ([]models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectAll())
	if err != nil {
		log.Println("Error preparing fetch all switch types sql")
		return nil, err
	}

	var nodeSwitches []models.Model

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

func (d *SwitchTypeData) SelectById(id int) (models.Model, error) {
	var nodeSwitch models.SwitchType

	stmt, err := d.DB().Prepare(d.Stmt().SelectById())
	if err != nil {
		log.Println("Error preparing fetch switch type sql")
		return nodeSwitch, err
	}

	err = stmt.QueryRow(id).Scan(&nodeSwitch.Id,
		&nodeSwitch.Name)

	if err != nil {
		log.Println("Error querying for the switch type")
		return nodeSwitch, err
	}

	defer stmt.Close()

	return nodeSwitch, nil
}

func (d *SwitchTypeData) SelectByParentId(parentId int) ([]models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectByParentId())
	if err != nil {
		log.Println("Error preparing fetch switch types by parent id sql")
		return nil, err
	}

	var nodeSwitches []models.Model

	rows, err := stmt.Query(parentId)
	if err != nil {
		log.Println("Error querying for switch types by parent id")
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

func (d *SwitchTypeData) Insert(switchType models.SwitchType) (int64, error) {
	stmt, err := d.DB().Prepare(d.Stmt().Insert())
	if err != nil {
		log.Println("Error preparing insert switch type sql")
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(switchType.Name)
	if err != nil {
		log.Println("Error inserting switch type")
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error getting last insert id for switch type")
		return 0, err
	}

	return id, nil
}

func (d *SwitchTypeData) Update(switchType models.SwitchType) error {
	stmt, err := d.DB().Prepare(d.Stmt().Update())
	if err != nil {
		log.Println("Error preparing update switch type sql")
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(switchType.Name, switchType.Id)
	if err != nil {
		log.Println("Error updating switch type")
		return err
	}

	return nil
}

func (d *SwitchTypeData) Delete(id int) error {
	stmt, err := d.DB().Prepare(d.Stmt().Delete())
	if err != nil {
		log.Println("Error preparing delete switch type sql")
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Println("Error deleting switch type")
		return err
	}

	return nil
}

func (d *SwitchTypeData) DeleteAll() error {
	stmt, err := d.DB().Prepare(d.Stmt().DeleteAll())
	if err != nil {
		log.Println("Error preparing delete all switch types sql")
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		log.Println("Error deleting all switch types")
		return err
	}

	return nil
}

func (d *SwitchTypeData) DeleteByParentId(parentId int) error {
	stmt, err := d.DB().Prepare(d.Stmt().DeleteByParentId())
	if err != nil {
		log.Println("Error preparing delete switch types by parent id sql")
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(parentId)
	if err != nil {
		log.Println("Error deleting switch types by parent id")
		return err
	}

	return nil
}
