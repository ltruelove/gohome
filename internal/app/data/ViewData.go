package data

import (
	"database/sql"
	"errors"
	"log"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data/statements"
	"github.com/ltruelove/gohome/internal/app/models"
)

// ViewData is a struct that holds the database connection and statements for view data operations
// It implements the CrudData interface for managing views in the database.
type ViewData struct {
	db   *sql.DB
	stmt statements.CrudStatement
}

func NewViewData(db *sql.DB, config *config.Configuration) CrudDataInterface {
	return &ViewData{
		db:   db,
		stmt: statements.NewViewDataStatements(config),
	}
}

func (d *ViewData) DB() *sql.DB {
	return d.db
}

func (d *ViewData) Stmt() statements.CrudStatement {
	return d.stmt
}

func (d *ViewData) SelectAll() ([]models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectAll())

	if err != nil {
		log.Println("Error preparing all views sql")
		return nil, err
	}

	var views []models.Model

	rows, err := stmt.Query()
	if err != nil {
		log.Println("Error querying for all views")
		return nil, err
	}

	defer stmt.Close()

	for rows.Next() {
		var view models.View
		err := rows.Scan(&view.Id,
			&view.Name)

		if err != nil {
			log.Println("Error scanning view")
			return nil, err
		}

		views = append(views, view)
	}

	return views, nil
}

func (d *ViewData) SelectById(id int) (models.Model, error) {
	var item models.View
	log.Printf("Fetching view for id: %d", id)
	stmt, err := d.DB().Prepare(d.Stmt().SelectById())

	if err != nil {
		log.Println("Error preparing fetch view sql")
		return item, err
	}

	err = stmt.QueryRow(id).Scan(&item.Id,
		&item.Name)

	if err != nil {
		log.Println("Error querying for view")
		return item, err
	}

	defer stmt.Close()

	return item, nil
}

func (d *ViewData) SelectByParentId(id int) ([]models.Model, error) {
	return nil, errors.New("SelectByParentId not implemented for ViewData")
}

func (d *ViewData) SelectBySecondParentId(id int) ([]models.Model, error) {
	return nil, errors.New("SelectBySecondParentId not implemented for ViewData")
}

func (d *ViewData) Insert(data models.Model) (models.Model, error) {
	view, ok := data.(*models.View)
	if !ok {
		log.Println("Error: data is not of type *models.View")
		return nil, errors.New("data is not of type *models.View")
	}

	stmt, err := d.DB().Prepare(d.Stmt().Insert())

	if err != nil {
		log.Println("Error preparing create view sql")
		return nil, err
	}

	lastInsertId := 0

	err = stmt.QueryRow(view.Name).Scan(&lastInsertId)

	if err != nil {
		log.Println("Error creating view")
		return nil, err
	}

	view.Id = int(lastInsertId)
	log.Printf("Created a view with the id: %d", view.Id)
	defer stmt.Close()

	return view, nil
}

func (d *ViewData) Update(data models.Model) error {
	view, ok := data.(*models.View)
	if !ok {
		log.Println("Error: data is not of type *models.View")
		return errors.New("data is not of type *models.View")
	}
	stmt, err := d.DB().Prepare(d.Stmt().Update())

	if err != nil {
		log.Println("Error preparing update view sql")
		return err
	}

	_, err = stmt.Exec(view.Name,
		view.Id)

	if err != nil {
		log.Println("Error updating view")
		return err
	}

	defer stmt.Close()

	return nil
}

func (d *ViewData) Delete(id int) error {
	stmt, err := d.DB().Prepare(d.Stmt().Delete())

	if err != nil {
		log.Println("Error preparing delete view sql")
		return err
	}

	_, err = stmt.Exec(id)

	if err != nil {
		log.Println("Error deleting view")
		return err
	}

	defer stmt.Close()

	return nil
}

func (d *ViewData) DeleteAll() error {
	stmt, err := d.DB().Prepare(d.Stmt().DeleteAll())

	if err != nil {
		log.Println("Error preparing delete all views sql")
		return err
	}

	_, err = stmt.Exec()

	if err != nil {
		log.Println("Error deleting all views")
		return err
	}

	defer stmt.Close()

	return nil
}

func (d *ViewData) DeleteByParentId(id int) error {
	return errors.New("DeleteByParentId not implemented for ViewData")
}

func (d *ViewData) DeleteBySecondParentId(id int) error {
	return errors.New("DeleteBySecondParentId not implemented for ViewData")
}
