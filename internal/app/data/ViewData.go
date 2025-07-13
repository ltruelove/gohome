package data

import (
	"database/sql"
	"log"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data/statements"
	"github.com/ltruelove/gohome/internal/app/models"
)

type ViewData struct {
	DB         *sql.DB
	Statements *statements.ViewDataStatements
}

func NewViewData(db *sql.DB, config *config.Configuration) *ViewData {
	return &ViewData{
		DB:         db,
		Statements: statements.NewViewDataStatements(config),
	}
}

func (viewData *ViewData) VerifyViewIdIsNew(viewId int) (bool, error) {
	view, err := viewData.FetchView(viewId)

	log.Printf("view found with id: %d, and name: %s", view.Id, view.Name)

	if err != nil {
		log.Println("Error fetching view")
		return false, err
	}

	return view.Id < 1, nil
}

func (viewData *ViewData) FetchAllViews() ([]models.View, error) {
	stmt, err := viewData.DB.Prepare(viewData.Statements.SelectAll())

	if err != nil {
		log.Println("Error preparing all views sql")
		return nil, err
	}

	var views []models.View

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

func (viewData *ViewData) FetchView(viewId int) (models.View, error) {
	var item models.View
	log.Printf("Fetching view for id: %d", viewId)
	stmt, err := viewData.DB.Prepare(viewData.Statements.SelectById())

	if err != nil {
		log.Println("Error preparing fetch view sql")
		return item, err
	}

	err = stmt.QueryRow(viewId).Scan(&item.Id,
		&item.Name)

	if err != nil {
		log.Println("Error querying for view")
		return item, err
	}

	defer stmt.Close()

	return item, nil
}

func (viewData *ViewData) CreateView(view *models.View) error {
	stmt, err := viewData.DB.Prepare(viewData.Statements.Insert())

	if err != nil {
		log.Println("Error preparing create view sql")
		return err
	}

	lastInsertId := 0

	err = stmt.QueryRow(view.Name).Scan(&lastInsertId)

	if err != nil {
		log.Println("Error creating view")
		return err
	}

	view.Id = int(lastInsertId)
	log.Printf("Created a view with the id: %d", view.Id)
	defer stmt.Close()

	return nil
}

func (viewData *ViewData) UpdateView(view *models.View) error {
	stmt, err := viewData.DB.Prepare(viewData.Statements.Update())

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

func (viewData *ViewData) DeleteView(viewId int) error {
	stmt, err := viewData.DB.Prepare(viewData.Statements.Delete())

	if err != nil {
		log.Println("Error preparing delete view sql")
		return err
	}

	_, err = stmt.Exec(viewId)

	if err != nil {
		log.Println("Error deleting view")
		return err
	}

	defer stmt.Close()

	return nil
}
