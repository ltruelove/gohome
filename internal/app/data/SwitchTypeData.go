package data

import (
	"database/sql"
	"errors"
	"log"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data/statements"
	"github.com/ltruelove/gohome/internal/app/models"
)

type SwitchTypeData struct {
	db   *sql.DB
	stmt statements.CrudStatement
}

func NewSwitchTypeData(db *sql.DB, config *config.Configuration) CrudDataInterface {
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
	return nil, errors.New("SelectByParentId not implemented for SwitchTypeData")
}

func (d *SwitchTypeData) SelectBySecondParentId(parentId int) ([]models.Model, error) {
	return nil, errors.New("SelectBySecondParentId not implemented for SwitchTypeData")
}

func (d *SwitchTypeData) Insert(data models.Model) (models.Model, error) {
	switchType, ok := data.(models.SwitchType)
	if !ok {
		log.Println("Error asserting data to SwitchType")
		return nil, errors.New("invalid data type for Insert")
	}
	stmt, err := d.DB().Prepare(d.Stmt().Insert())
	if err != nil {
		log.Println("Error preparing insert switch type sql")
		return nil, err
	}

	defer stmt.Close()

	lastInsertId := 0
	err = stmt.QueryRow(switchType.Name).Scan(&lastInsertId)
	if err != nil {
		log.Println("Error inserting switch type")
		return nil, err
	}
	switchType.Id = lastInsertId

	return switchType, nil
}

func (d *SwitchTypeData) Update(data models.Model) error {
	switchType, ok := data.(models.SwitchType)
	if !ok {
		log.Println("Error asserting data to SwitchType")
		return errors.New("invalid data type for Update")
	}
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
	return errors.New("DeleteByParentId not implemented for SwitchTypeData")
}

func (d *SwitchTypeData) DeleteBySecondParentId(secondParentId int) error {
	return errors.New("DeleteBySecondParentId not implemented for SwitchTypeData")
}
