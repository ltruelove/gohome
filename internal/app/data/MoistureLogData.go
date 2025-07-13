package data

import (
	"database/sql"
	"errors"
	"log"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data/statements"
	"github.com/ltruelove/gohome/internal/app/models"
)

type MoistureLogData struct {
	db   *sql.DB
	stmt statements.CrudStatement
}

func NewMoistureLogData(DB *sql.DB, config *config.Configuration) CrudDataInterface {
	return &MoistureLogData{
		db:   DB,
		stmt: statements.NewMoistureLogDataStatements(config),
	}
}

func (d *MoistureLogData) DB() *sql.DB {
	return d.db
}

func (d *MoistureLogData) Stmt() statements.CrudStatement {
	return d.stmt
}

func (d *MoistureLogData) SelectAll() ([]models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectAll())
	if err != nil {
		log.Println("Error preparing fetch all moisture logs sql")
		return nil, err
	}

	var logs []models.Model

	rows, err := stmt.Query()
	if err != nil {
		log.Println("Error querying for all moisture logs")
		return nil, err
	}
	defer stmt.Close()

	for rows.Next() {
		var logEntry models.MoistureLogData

		err := rows.Scan(&logEntry.Id,
			&logEntry.NodeSensorLogId,
			&logEntry.Moisture)

		if err != nil {
			log.Println("Error scanning moisture log")
			return nil, err
		}

		logs = append(logs, &logEntry)
	}

	return logs, nil
}

func (d *MoistureLogData) SelectById(id int) (models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectById())
	if err != nil {
		log.Println("Error preparing fetch moisture log by id sql")
		return nil, err
	}

	var logEntry models.MoistureLogData

	err = stmt.QueryRow(id).Scan(&logEntry.Id,
		&logEntry.NodeSensorLogId,
		&logEntry.Moisture)

	if err != nil {
		log.Println("Error querying for moisture log by id")
		return nil, err
	}

	return &logEntry, nil
}

func (d *MoistureLogData) SelectByParentId(id int) ([]models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectByParentId())
	if err != nil {
		log.Println("Error preparing fetch moisture logs by parent id sql")
		return nil, err
	}

	var logs []models.Model

	rows, err := stmt.Query(id)
	if err != nil {
		log.Println("Error querying for moisture logs by parent id")
		return nil, err
	}
	defer stmt.Close()

	for rows.Next() {
		var logEntry models.MoistureLogData

		err := rows.Scan(&logEntry.Id,
			&logEntry.NodeSensorLogId,
			&logEntry.Moisture)

		if err != nil {
			log.Println("Error scanning moisture log by parent id")
			return nil, err
		}

		logs = append(logs, &logEntry)
	}

	return logs, nil
}

func (d *MoistureLogData) SelectBySecondParentId(id int) ([]models.Model, error) {
	return nil, errors.New("MoistureLogData table has no second parent ID")
}

func (d *MoistureLogData) Insert(data models.Model) (models.Model, error) {
	logEntry, ok := data.(*models.MoistureLogData)
	if !ok {
		log.Println("Error casting model to MoistureLogData")
		return nil, errors.New("Invalid model type")
	}

	stmt, err := d.DB().Prepare(d.Stmt().Insert())
	if err != nil {
		log.Println("Error preparing insert moisture log sql")
		return nil, err
	}

	defer stmt.Close()

	var id int
	err = stmt.QueryRow(logEntry.NodeSensorLogId, logEntry.Moisture).Scan(&id)
	if err != nil {
		log.Println("Error inserting moisture log")
		return nil, err
	}

	logEntry.Id = id

	return logEntry, nil
}

func (d *MoistureLogData) Update(data models.Model) error {
	logEntry, ok := data.(*models.MoistureLogData)
	if !ok {
		log.Println("Error casting model to MoistureLogData")
		return errors.New("Invalid model type")
	}

	stmt, err := d.DB().Prepare(d.Stmt().Update())
	if err != nil {
		log.Println("Error preparing update moisture log sql")
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(logEntry.NodeSensorLogId, logEntry.Moisture, logEntry.Id)
	if err != nil {
		log.Println("Error updating moisture log")
		return err
	}

	return nil
}

func (d *MoistureLogData) Delete(id int) error {
	stmt, err := d.DB().Prepare(d.Stmt().Delete())
	if err != nil {
		log.Println("Error preparing delete moisture log sql")
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Println("Error deleting moisture log")
		return err
	}

	return nil
}

func (d *MoistureLogData) DeleteAll() error {
	stmt, err := d.DB().Prepare(d.Stmt().DeleteAll())
	if err != nil {
		log.Println("Error preparing delete all moisture logs sql")
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		log.Println("Error deleting all moisture logs")
		return err
	}

	return nil
}

func (d *MoistureLogData) DeleteByParentId(parentId int) error {
	stmt, err := d.DB().Prepare(d.Stmt().DeleteByParentId())
	if err != nil {
		log.Println("Error preparing delete moisture logs by parent id sql")
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(parentId)
	if err != nil {
		log.Println("Error deleting moisture logs by parent id")
		return err
	}

	return nil
}

func (d *MoistureLogData) DeleteBySecondParentId(secondParentId int) error {
	return errors.New("MoistureLogData table has no second parent ID")
}
