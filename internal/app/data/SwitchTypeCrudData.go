package data

import (
	"database/sql"
	"errors"
	"log"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data/statements"
	"github.com/ltruelove/gohome/internal/app/models"
)

// SwitchTypeCrudData is a specific implementation of CrudDataInterface for SwitchType models
type SwitchTypeCrudData struct {
	db   *sql.DB
	stmt statements.CrudStatement
}

// NewSwitchTypeCrudData creates a new instance of SwitchTypeCrudData
// This function returns an instance that implements CrudDataInterface
func NewSwitchTypeCrudData(db *sql.DB, config *config.Configuration) CrudDataInterface {
	return &SwitchTypeCrudData{
		db:   db,
		stmt: statements.NewSwitchTypeStatements(config),
	}
}

func (d *SwitchTypeCrudData) DB() *sql.DB {
	return d.db
}

func (d *SwitchTypeCrudData) Stmt() statements.CrudStatement {
	return d.stmt
}

func (d *SwitchTypeCrudData) SelectAll() ([]models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectAll())
	if err != nil {
		log.Printf("Error preparing SelectAll statement: %v", err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Printf("Error executing SelectAll query: %v", err)
		return nil, err
	}
	defer rows.Close()

	var results []models.Model
	for rows.Next() {
		var switchType models.SwitchType
		err := rows.Scan(&switchType.Id, &switchType.Name)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}
		results = append(results, switchType)
	}

	return results, nil
}

func (d *SwitchTypeCrudData) SelectByParentId(id int) ([]models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectByParentId())
	if err != nil {
		log.Printf("Error preparing SelectByParentId statement: %v", err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		log.Printf("Error executing SelectByParentId query: %v", err)
		return nil, err
	}
	defer rows.Close()

	var results []models.Model
	for rows.Next() {
		var switchType models.SwitchType
		err := rows.Scan(&switchType.Id, &switchType.Name)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}
		results = append(results, switchType)
	}

	return results, nil
}

func (d *SwitchTypeCrudData) SelectById(id int) (models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectById())
	if err != nil {
		log.Printf("Error preparing SelectById statement: %v", err)
		return nil, err
	}
	defer stmt.Close()

	var switchType models.SwitchType
	err = stmt.QueryRow(id).Scan(&switchType.Id, &switchType.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No switch type found with id %d", id)
		} else {
			log.Printf("Error scanning row: %v", err)
		}
		return nil, err
	}

	return switchType, nil
}

func (d *SwitchTypeCrudData) Insert(data models.Model) (models.Model, error) {
	switchType, ok := data.(models.SwitchType)
	if !ok {
		log.Printf("Invalid model type, expected SwitchType")
		return nil, errors.New("invalid model type")
	}

	// Validate the model
	if valid, err := switchType.IsValid(false); !valid {
		log.Printf("Model validation failed: %v", err)
		return nil, err
	}

	stmt, err := d.DB().Prepare(d.Stmt().Insert())
	if err != nil {
		log.Printf("Error preparing Insert statement: %v", err)
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(switchType.Name)
	if err != nil {
		log.Printf("Error executing Insert statement: %v", err)
		return nil, err
	}

	// Get the last inserted ID
	lastID, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert ID: %v", err)
		return nil, err
	}

	switchType.Id = int(lastID)
	return switchType, nil
}

func (d *SwitchTypeCrudData) Update(data models.Model) error {
	switchType, ok := data.(models.SwitchType)
	if !ok {
		log.Printf("Invalid model type, expected SwitchType")
		return errors.New("invalid model type")
	}

	// Validate the model
	if valid, err := switchType.IsValid(true); !valid {
		log.Printf("Model validation failed: %v", err)
		return err
	}

	stmt, err := d.DB().Prepare(d.Stmt().Update())
	if err != nil {
		log.Printf("Error preparing Update statement: %v", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(switchType.Name, switchType.Id)
	if err != nil {
		log.Printf("Error executing Update statement: %v", err)
		return err
	}

	return nil
}

func (d *SwitchTypeCrudData) Delete(id int) error {
	stmt, err := d.DB().Prepare(d.Stmt().Delete())
	if err != nil {
		log.Printf("Error preparing Delete statement: %v", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Printf("Error executing Delete statement: %v", err)
		return err
	}

	return nil
}

func (d *SwitchTypeCrudData) DeleteByParentId(id int) error {
	stmt, err := d.DB().Prepare(d.Stmt().DeleteByParentId())
	if err != nil {
		log.Printf("Error preparing DeleteByParentId statement: %v", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Printf("Error executing DeleteByParentId statement: %v", err)
		return err
	}

	return nil
}

func (d *SwitchTypeCrudData) DeleteAll() error {
	stmt, err := d.DB().Prepare(d.Stmt().DeleteAll())
	if err != nil {
		log.Printf("Error preparing DeleteAll statement: %v", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		log.Printf("Error executing DeleteAll statement: %v", err)
		return err
	}

	return nil
}
