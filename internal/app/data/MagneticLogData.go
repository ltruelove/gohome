package data

import (
	"database/sql"
	"errors"
	"log"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data/statements"
	"github.com/ltruelove/gohome/internal/app/models"
)

type MagneticLogData struct {
	db   *sql.DB
	stmt statements.CrudStatement
}

func NewMagneticLogData(DB *sql.DB, config *config.Configuration) CrudDataInterface {
	return &MagneticLogData{
		db:   DB,
		stmt: statements.NewMagneticLogDataStatements(config),
	}
}

func (d *MagneticLogData) DB() *sql.DB {
	return d.db
}

func (d *MagneticLogData) Stmt() statements.CrudStatement {
	return d.stmt
}

func (d *MagneticLogData) SelectAll() ([]models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectAll())
	if err != nil {
		log.Println("Error preparing fetch all magnetic logs sql")
		return nil, err
	}

	var logs []models.Model

	rows, err := stmt.Query()
	if err != nil {
		log.Println("Error querying for all magnetic logs")
		return nil, err
	}
	defer stmt.Close()

	for rows.Next() {
		var logEntry models.MagneticLogData

		err := rows.Scan(&logEntry.Id,
			&logEntry.NodeSensorLogId,
			&logEntry.IsClosed)

		if err != nil {
			log.Println("Error scanning magnetic log")
			return nil, err
		}

		logs = append(logs, &logEntry)
	}

	return logs, nil
}

func (d *MagneticLogData) SelectById(id int) (models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectById())
	if err != nil {
		log.Println("Error preparing fetch magnetic log by id sql")
		return nil, err
	}

	var logEntry models.MagneticLogData

	err = stmt.QueryRow(id).Scan(&logEntry.Id,
		&logEntry.NodeSensorLogId,
		&logEntry.IsClosed)

	if err != nil {
		log.Println("Error querying for magnetic log by id")
		return nil, err
	}

	return &logEntry, nil
}

func (d *MagneticLogData) SelectByParentId(id int) ([]models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectByParentId())
	if err != nil {
		log.Println("Error preparing fetch magnetic logs by parent id sql")
		return nil, err
	}

	var logs []models.Model

	rows, err := stmt.Query(id)
	if err != nil {
		log.Println("Error querying for magnetic logs by parent id")
		return nil, err
	}
	defer stmt.Close()

	for rows.Next() {
		var logEntry models.MagneticLogData

		err := rows.Scan(&logEntry.Id,
			&logEntry.NodeSensorLogId,
			&logEntry.IsClosed)

		if err != nil {
			log.Println("Error scanning magnetic log by parent id")
			return nil, err
		}

		logs = append(logs, &logEntry)
	}

	return logs, nil
}

func (d *MagneticLogData) SelectBySecondParentId(id int) ([]models.Model, error) {
	return nil, errors.New("MagneticLogData table has no second parent ID")
}

func (d *MagneticLogData) Insert(data models.Model) (models.Model, error) {
	logEntry, ok := data.(*models.MagneticLogData)
	if !ok {
		log.Println("Error casting model to MagneticLogData")
		return nil, errors.New("Invalid model type")
	}

	stmt, err := d.DB().Prepare(d.Stmt().Insert())
	if err != nil {
		log.Println("Error preparing insert magnetic log sql")
		return nil, err
	}

	defer stmt.Close()

	var id int
	err = stmt.QueryRow(logEntry.NodeSensorLogId, logEntry.IsClosed).Scan(&id)
	if err != nil {
		log.Println("Error inserting magnetic log")
		return nil, err
	}

	logEntry.Id = id

	return logEntry, nil
}

func (d *MagneticLogData) Update(data models.Model) error {
	logEntry, ok := data.(*models.MagneticLogData)
	if !ok {
		log.Println("Error casting model to MagneticLogData")
		return errors.New("Invalid model type")
	}

	stmt, err := d.DB().Prepare(d.Stmt().Update())
	if err != nil {
		log.Println("Error preparing update magnetic log sql")
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(logEntry.NodeSensorLogId, logEntry.IsClosed, logEntry.Id)
	if err != nil {
		log.Println("Error updating magnetic log")
		return err
	}

	return nil
}

func (d *MagneticLogData) Delete(id int) error {
	stmt, err := d.DB().Prepare(d.Stmt().Delete())
	if err != nil {
		log.Println("Error preparing delete magnetic log sql")
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Println("Error deleting magnetic log")
		return err
	}

	return nil
}

func (d *MagneticLogData) DeleteAll() error {
	stmt, err := d.DB().Prepare(d.Stmt().DeleteAll())
	if err != nil {
		log.Println("Error preparing delete all magnetic logs sql")
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		log.Println("Error deleting all magnetic logs")
		return err
	}

	return nil
}

func (d *MagneticLogData) DeleteByParentId(parentId int) error {
	stmt, err := d.DB().Prepare(d.Stmt().DeleteByParentId())
	if err != nil {
		log.Println("Error preparing delete magnetic logs by parent id sql")
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(parentId)
	if err != nil {
		log.Println("Error deleting magnetic logs by parent id")
		return err
	}

	return nil
}

func (d *MagneticLogData) DeleteBySecondParentId(secondParentId int) error {
	return errors.New("MagneticLogData table has no second parent ID")
}
