package data

import (
	"database/sql"
	"log"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data/statements"
	"github.com/ltruelove/gohome/internal/app/models"
)

// GenericCrudData is a generic implementation of CrudDataInterface
// that can be used for any model that implements models.Model
type GenericCrudData struct {
	db   *sql.DB
	stmt statements.CrudStatement
}

// NewGenericCrudData creates a new instance of GenericCrudData
// This is the factory function that returns an instance that can be passed
// to functions expecting a CrudDataInterface parameter
func NewGenericCrudData(db *sql.DB, stmt statements.CrudStatement) CrudDataInterface {
	return &GenericCrudData{
		db:   db,
		stmt: stmt,
	}
}

// NewGenericCrudDataWithConfig creates a new instance using a statement factory function
// Example usage: NewGenericCrudDataWithConfig(db, config, statements.NewSwitchTypeStatements)
func NewGenericCrudDataWithConfig(db *sql.DB, config *config.Configuration,
	stmtFactory func(*config.Configuration) statements.CrudStatement) CrudDataInterface {
	return &GenericCrudData{
		db:   db,
		stmt: stmtFactory(config),
	}
}

// DB returns the database connection
func (d *GenericCrudData) DB() *sql.DB {
	return d.db
}

// Stmt returns the CRUD statement interface
func (d *GenericCrudData) Stmt() statements.CrudStatement {
	return d.stmt
}

// SelectAll retrieves all records from the database
func (d *GenericCrudData) SelectAll() ([]models.Model, error) {
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
		// Note: You'll need to implement the specific model scanning logic
		// based on your actual model type. This is a placeholder.
		// In practice, you might want to make this generic or have
		// model-specific implementations.
	}

	return results, nil
}

// SelectByParentId retrieves records by parent ID
func (d *GenericCrudData) SelectByParentId(id int) ([]models.Model, error) {
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
		// Implement model-specific scanning logic here
	}

	return results, nil
}

// SelectById retrieves a single record by ID
func (d *GenericCrudData) SelectById(id int) (models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectById())
	if err != nil {
		log.Printf("Error preparing SelectById statement: %v", err)
		return nil, err
	}
	defer stmt.Close()

	// Implement model-specific scanning logic here
	// For now, returning nil as placeholder
	return nil, nil
}

// Insert inserts a new record
func (d *GenericCrudData) Insert(data models.Model) (models.Model, error) {
	// Validate the model first
	if valid, err := data.IsValid(false); !valid {
		log.Printf("Model validation failed: %v", err)
		return nil, err
	}

	stmt, err := d.DB().Prepare(d.Stmt().Insert())
	if err != nil {
		log.Printf("Error preparing Insert statement: %v", err)
		return nil, err
	}
	defer stmt.Close()

	// Implement model-specific insertion logic here
	// This would typically involve extracting values from the model
	// and passing them to stmt.Exec()

	return data, nil
}

// Update updates an existing record
func (d *GenericCrudData) Update(data models.Model) error {
	// Validate the model first
	if valid, err := data.IsValid(true); !valid {
		log.Printf("Model validation failed: %v", err)
		return err
	}

	stmt, err := d.DB().Prepare(d.Stmt().Update())
	if err != nil {
		log.Printf("Error preparing Update statement: %v", err)
		return err
	}
	defer stmt.Close()

	// Implement model-specific update logic here

	return nil
}

// Delete deletes a record by ID
func (d *GenericCrudData) Delete(id int) error {
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

// DeleteByParentId deletes records by parent ID
func (d *GenericCrudData) DeleteByParentId(id int) error {
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

// DeleteAll deletes all records
func (d *GenericCrudData) DeleteAll() error {
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
