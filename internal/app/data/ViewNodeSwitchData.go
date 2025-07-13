package data

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data/statements"
	"github.com/ltruelove/gohome/internal/app/models"
)

type ViewNodeSwitchData struct {
	db   *sql.DB
	stmt statements.CrudStatement
}

func NewViewNodeSwitchData(db *sql.DB, config *config.Configuration) CrudDataInterface {
	return &ViewNodeSwitchData{
		db:   db,
		stmt: statements.NewViewNodeSwitchDataStatements(config),
	}
}

func (d ViewNodeSwitchData) DB() *sql.DB {
	return d.db
}

func (d ViewNodeSwitchData) Stmt() statements.CrudStatement {
	return d.stmt
}

func (d ViewNodeSwitchData) SelectAll() ([]models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectAll())
	if err != nil {
		log.Println("Error preparing fetch all view node switch data sql")
		return nil, err
	}

	var listData []models.Model

	rows, err := stmt.Query()
	if err != nil {
		log.Println("Error querying for all view node switch data")
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
			log.Println("Error scanning view node switch data")
			return nil, err
		}

		listData = append(listData, item)
	}

	return listData, nil
}

func (d ViewNodeSwitchData) SelectById(id int) (models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectById())
	if err != nil {
		log.Println("Error preparing fetch view node switch data by id sql")
		return nil, err
	}

	var item models.ViewNodeSwitchData

	err = stmt.QueryRow(id).Scan(&item.Id,
		&item.NodeId,
		&item.ViewId,
		&item.NodeSwitchId,
		&item.Name)

	if err != nil {
		log.Println("Error querying for view node switch data by id")
		return nil, err
	}

	defer stmt.Close()

	return item, nil
}

func (d ViewNodeSwitchData) SelectByParentId(id int) ([]models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectByParentId())
	if err != nil {
		log.Println("Error preparing fetch view node switch data by parent id sql")
		return nil, err
	}

	defer stmt.Close()

	var items []models.Model

	rows, err := stmt.Query(id)
	if err != nil {
		log.Println("Error querying for view node switch data by parent id")
		return nil, err
	}

	for rows.Next() {
		var item models.ViewNodeSwitchData

		err = rows.Scan(&item.Id,
			&item.NodeId,
			&item.ViewId,
			&item.NodeSwitchId,
			&item.Name)

		if err != nil {
			log.Println("Error scanning view node switch data by parent id")
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

func (d ViewNodeSwitchData) SelectBySecondParentId(id int) ([]models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectBySecondParentId())
	if err != nil {
		log.Println("Error preparing fetch view node switch data by second parent id sql")
		return nil, err
	}

	defer stmt.Close()

	var items []models.Model

	rows, err := stmt.Query(id)
	if err != nil {
		log.Println("Error querying for view node switch data by second parent id")
		return nil, err
	}

	for rows.Next() {
		var item models.ViewNodeSwitchData

		err = rows.Scan(&item.Id,
			&item.NodeId,
			&item.ViewId,
			&item.NodeSwitchId,
			&item.Name)

		if err != nil {
			log.Println("Error scanning view node switch data by second parent id")
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

func (d ViewNodeSwitchData) Insert(data models.Model) (models.Model, error) {
	item, ok := data.(*models.ViewNodeSwitchData)
	if !ok {
		log.Println("Error casting data to ViewNodeSwitchData")
		return nil, fmt.Errorf("invalid type for Insert: %T", data)
	}

	stmt, err := d.DB().Prepare(d.Stmt().Insert())
	if err != nil {
		log.Println("Error preparing insert view node switch data sql")
		return nil, err
	}

	lastInsertId := 0

	err = stmt.QueryRow(item.NodeId,
		item.ViewId,
		item.NodeSwitchId,
		item.Name).Scan((&lastInsertId))

	if err != nil {
		log.Println("Error creating node switch data")
		return nil, err
	}

	item.Id = int(lastInsertId)
	log.Printf("Created node switch data with id: %d", item.Id)
	defer stmt.Close()

	return item, nil
}

func (d ViewNodeSwitchData) Update(data models.Model) error {
	item, ok := data.(*models.ViewNodeSwitchData)
	if !ok {
		log.Println("Error casting data to ViewNodeSwitchData")
		return fmt.Errorf("invalid type for Update: %T", data)
	}

	stmt, err := d.DB().Prepare(d.Stmt().Update())
	if err != nil {
		log.Println("Error preparing update view node switch data sql")
		return err
	}

	_, err = stmt.Exec(item.Id,
		item.NodeId,
		item.ViewId,
		item.NodeSwitchId,
		item.Name)

	if err != nil {
		log.Println("Error updating view node switch data")
		return err
	}

	defer stmt.Close()

	return nil
}

func (d ViewNodeSwitchData) Delete(id int) error {
	stmt, err := d.DB().Prepare(d.Stmt().Delete())
	if err != nil {
		log.Println("Error preparing delete view node switch data sql")
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		log.Println("Error deleting view node switch data")
		return err
	}

	defer stmt.Close()

	return nil
}

func (d ViewNodeSwitchData) DeleteAll() error {
	stmt, err := d.DB().Prepare(d.Stmt().DeleteAll())
	if err != nil {
		log.Println("Error preparing delete all view node switch data sql")
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Println("Error deleting all view node switch data")
		return err
	}

	defer stmt.Close()

	return nil
}

func (d ViewNodeSwitchData) DeleteByParentId(parentId int) error {
	stmt, err := d.DB().Prepare(d.Stmt().DeleteByParentId())
	if err != nil {
		log.Println("Error preparing delete all view node switch data by parent id sql")
		return err
	}

	_, err = stmt.Exec(parentId)
	if err != nil {
		log.Println("Error deleting all view node switch data by parent id")
		return err
	}

	defer stmt.Close()

	return nil
}

func (d ViewNodeSwitchData) DeleteBySecondParentId(secondParentId int) error {
	stmt, err := d.DB().Prepare(d.Stmt().DeleteBySecondParentId())
	if err != nil {
		log.Println("Error preparing delete all view node switch data by second parent id sql")
		return err
	}

	_, err = stmt.Exec(secondParentId)
	if err != nil {
		log.Println("Error deleting all view node switch data by second parent id")
		return err
	}

	defer stmt.Close()

	return nil
}
