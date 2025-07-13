package data

import (
	"database/sql"
	"errors"
	"log"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data/statements"
	"github.com/ltruelove/gohome/internal/app/models"
)

type TempLogData struct {
	db   *sql.DB
	stmt statements.CrudStatement
}

func NewTempLogData(db *sql.DB, config *config.Configuration) CrudDataInterface {
	return &TempLogData{
		db:   db,
		stmt: statements.NewTempLogDataStatements(config),
	}
}

func (d *TempLogData) DB() *sql.DB {
	return d.db
}

func (d *TempLogData) Stmt() statements.CrudStatement {
	return d.stmt
}

func (d *TempLogData) SelectAll() ([]models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectAll())
	if err != nil {
		log.Println("Error preparing fetch all temperature logs sql")
		return nil, err
	}

	var logs []models.Model

	rows, err := stmt.Query()
	if err != nil {
		log.Println("Error querying for all temperature logs")
		return nil, err
	}
	defer stmt.Close()

	for rows.Next() {
		var logEntry models.TempLogData

		err := rows.Scan(&logEntry.Id,
			&logEntry.NodeSensorLogId,
			&logEntry.TemperatureF,
			&logEntry.TemperatureC,
			&logEntry.Humidity)
		if err != nil {
			log.Println("Error scanning temperature log")
			return nil, err
		}
		logs = append(logs, &logEntry)
	}

	return logs, nil
}

func (d *TempLogData) SelectById(id int) (models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectById())
	if err != nil {
		log.Println("Error preparing fetch temperature log by ID sql")
		return nil, err
	}

	var logEntry models.TempLogData

	err = stmt.QueryRow(id).Scan(&logEntry.Id,
		&logEntry.NodeSensorLogId,
		&logEntry.TemperatureF,
		&logEntry.TemperatureC,
		&logEntry.Humidity)

	if err != nil {
		log.Println("Error querying for temperature log by ID")
		return nil, err
	}

	return &logEntry, nil
}

func (d *TempLogData) SelectByParentId(id int) ([]models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectByParentId())
	if err != nil {
		log.Println("Error preparing fetch temperature logs by NodeSensorLogId sql")
		return nil, err
	}

	var logData []models.Model

	rows, err := stmt.Query(id)
	if err != nil {
		log.Println("Error querying for temperature logs by NodeSensorLogId")
		return nil, err
	}
	defer stmt.Close()

	for rows.Next() {
		var logRecord models.TempLogData
		err := rows.Scan(&logRecord.Id,
			&logRecord.NodeSensorLogId,
			&logRecord.TemperatureF,
			&logRecord.TemperatureC,
			&logRecord.Humidity)
		if err != nil {
			log.Println("Error scanning temperature log")
			return nil, err
		}
		logData = append(logData, logRecord)
	}

	return logData, nil
}

func (d *TempLogData) SelectBySecondParentId(id int) ([]models.Model, error) {
	// This method is not applicable for TempLogData as it does not have a second parent ID.
	return nil, errors.New("TempLogData does not support second parent ID queries")
}

func (d *TempLogData) Insert(data models.Model) (models.Model, error) {
	item, ok := data.(*models.TempLogData)
	if !ok {
		log.Println("Error casting model to TempLogData")
		return nil, errors.New("Invalid model type")
	}

	stmt, err := d.DB().Prepare(d.Stmt().Insert())
	if err != nil {
		log.Println("Error preparing insert temperature log sql")
		return nil, err
	}

	res, err := stmt.Exec(item.NodeSensorLogId, item.TemperatureF, item.TemperatureC, item.Humidity)
	if err != nil {
		log.Println("Error inserting temperature log")
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Println("Error getting last insert ID for temperature log")
		return nil, err
	}
	item.Id = int(id)

	defer stmt.Close()

	return item, nil
}

func (d *TempLogData) Update(data models.Model) error {
	item, ok := data.(*models.TempLogData)
	if !ok {
		log.Println("Error casting model to TempLogData")
		return errors.New("Invalid model type")
	}

	stmt, err := d.DB().Prepare(d.Stmt().Update())
	if err != nil {
		log.Println("Error preparing update temperature log sql")
		return err
	}

	_, err = stmt.Exec(item.NodeSensorLogId, item.TemperatureF, item.TemperatureC, item.Humidity, item.Id)
	if err != nil {
		log.Println("Error updating temperature log")
		return err
	}

	defer stmt.Close()

	return nil
}

func (d *TempLogData) Delete(id int) error {
	stmt, err := d.DB().Prepare(d.Stmt().Delete())
	if err != nil {
		log.Println("Error preparing delete temperature log sql")
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		log.Println("Error deleting temperature log")
		return err
	}

	defer stmt.Close()

	return nil
}

func (d *TempLogData) DeleteAll() error {
	stmt, err := d.DB().Prepare(d.Stmt().DeleteAll())
	if err != nil {
		log.Println("Error preparing delete all temperature logs sql")
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Println("Error deleting all temperature logs")
		return err
	}

	defer stmt.Close()

	return nil
}

func (d *TempLogData) DeleteByParentId(id int) error {
	stmt, err := d.DB().Prepare(d.Stmt().DeleteByParentId())
	if err != nil {
		log.Println("Error preparing delete temperature logs by NodeSensorLogId sql")
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		log.Println("Error deleting temperature logs by NodeSensorLogId")
		return err
	}

	defer stmt.Close()

	return nil
}

func (d *TempLogData) DeleteBySecondParentId(id int) error {
	// This method is not applicable for TempLogData as it does not have a second parent ID.
	return errors.New("TempLogData does not support second parent ID deletions")
}
