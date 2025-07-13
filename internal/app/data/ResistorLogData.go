package data

import (
	"database/sql"
	"errors"
	"log"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data/statements"
	"github.com/ltruelove/gohome/internal/app/models"
)

type ResistorLogData struct {
	db   *sql.DB
	stmt statements.CrudStatement
}

func NewResistorLogData(DB *sql.DB, config *config.Configuration) CrudDataInterface {
	return &ResistorLogData{
		db:   DB,
		stmt: statements.NewResistorLogDataStatements(config),
	}
}

func (d *ResistorLogData) DB() *sql.DB {
	return d.db
}

func (d *ResistorLogData) Stmt() statements.CrudStatement {
	return d.stmt
}

func (d *ResistorLogData) SelectAll() ([]models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectAll())
	if err != nil {
		log.Println("Error preparing fetch all resistor logs sql")
		return nil, err
	}

	var logs []models.Model

	rows, err := stmt.Query()
	if err != nil {
		log.Println("Error querying for all resistor logs")
		return nil, err
	}
	defer stmt.Close()

	for rows.Next() {
		var logEntry models.ResistorLogData

		err := rows.Scan(&logEntry.Id,
			&logEntry.NodeSensorLogId,
			&logEntry.ResistorValue)

		if err != nil {
			log.Println("Error scanning resistor log")
			return nil, err
		}

		logs = append(logs, &logEntry)
	}

	return logs, nil
}

func (d *ResistorLogData) SelectById(id int) (models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectById())
	if err != nil {
		log.Println("Error preparing fetch resistor log by id sql")
		return nil, err
	}

	var logEntry models.ResistorLogData

	err = stmt.QueryRow(id).Scan(&logEntry.Id,
		&logEntry.NodeSensorLogId,
		&logEntry.ResistorValue)

	if err != nil {
		log.Println("Error querying for resistor log by id")
		return nil, err
	}

	return &logEntry, nil
}

func (d *ResistorLogData) SelectByParentId(id int) ([]models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectByParentId())
	if err != nil {
		log.Println("Error preparing fetch resistor logs by parent id sql")
		return nil, err
	}

	var logs []models.Model

	rows, err := stmt.Query(id)
	if err != nil {
		log.Println("Error querying for resistor logs by parent id")
		return nil, err
	}
	defer stmt.Close()

	for rows.Next() {
		var logEntry models.ResistorLogData

		err := rows.Scan(&logEntry.Id,
			&logEntry.NodeSensorLogId,
			&logEntry.ResistorValue)

		if err != nil {
			log.Println("Error scanning resistor log by parent id")
			return nil, err
		}

		logs = append(logs, &logEntry)
	}

	return logs, nil
}

func (d *ResistorLogData) SelectBySecondParentId(id int) ([]models.Model, error) {
	return nil, errors.New("ResistorLog table has no second parent ID")
}

func (d *ResistorLogData) Insert(data models.Model) (models.Model, error) {
	item, ok := data.(*models.ResistorLogData)
	if !ok {
		log.Println("Error casting model to ResistorLogData")
		return nil, errors.New("Invalid model type")
	}

	stmt, err := d.DB().Prepare(d.Stmt().Insert())
	if err != nil {
		log.Println("Error preparing insert resistor log sql")
		return nil, err
	}

	defer stmt.Close()

	var id int
	err = stmt.QueryRow(item.NodeSensorLogId, item.ResistorValue).Scan(&id)
	if err != nil {
		log.Println("Error inserting resistor log")
		return nil, err
	}

	item.Id = id
	return item, nil
}

func (d *ResistorLogData) Update(data models.Model) error {
	item, ok := data.(*models.ResistorLogData)
	if !ok {
		log.Println("Error casting model to ResistorLogData")
		return errors.New("Invalid model type")
	}

	stmt, err := d.DB().Prepare(d.Stmt().Update())
	if err != nil {
		log.Println("Error preparing update resistor log sql")
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(item.NodeSensorLogId, item.ResistorValue, item.Id)
	if err != nil {
		log.Println("Error updating resistor log")
		return err
	}

	return nil
}

func (d *ResistorLogData) Delete(id int) error {
	stmt, err := d.DB().Prepare(d.Stmt().Delete())
	if err != nil {
		log.Println("Error preparing delete resistor log sql")
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Println("Error deleting resistor log")
		return err
	}

	return nil
}

func (d *ResistorLogData) DeleteAll() error {
	stmt, err := d.DB().Prepare(d.Stmt().DeleteAll())
	if err != nil {
		log.Println("Error preparing delete all resistor logs sql")
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		log.Println("Error deleting all resistor logs")
		return err
	}

	return nil
}

func (d *ResistorLogData) DeleteByParentId(id int) error {
	stmt, err := d.DB().Prepare(d.Stmt().DeleteByParentId())
	if err != nil {
		log.Println("Error preparing delete resistor logs by parent id sql")
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Println("Error deleting resistor logs by parent id")
		return err
	}

	return nil
}

func (d *ResistorLogData) DeleteBySecondParentId(id int) error {
	return errors.New("ResistorLog table has no second parent ID")
}
