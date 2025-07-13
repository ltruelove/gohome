package data

import (
	"database/sql"
	"errors"
	"log"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data/statements"
	"github.com/ltruelove/gohome/internal/app/models"
)

type SensorType struct {
	db   *sql.DB
	stmt statements.CrudStatement
}

func NewSensorType(db *sql.DB, config *config.Configuration) CrudDataInterface {
	return &SensorType{
		db:   db,
		stmt: statements.NewSensorTypeStatements(config),
	}
}

func (d *SensorType) DB() *sql.DB {
	return d.db
}

func (d *SensorType) Stmt() statements.CrudStatement {
	return d.stmt
}

func (d *SensorType) SelectAll() ([]models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectAll())
	if err != nil {
		log.Println("Error preparing all sensor types sql")
		return nil, err
	}

	var sensors []models.Model

	rows, err := stmt.Query()
	if err != nil {
		log.Println("Error querying for all sensor types")
		return nil, err
	}

	for rows.Next() {
		var sensor models.SensorType
		rows.Scan(&sensor.Id,
			&sensor.TypeName)
		sensors = append(sensors, sensor)
	}
	defer stmt.Close()

	return sensors, nil
}

func (d *SensorType) SelectById(id int) (models.Model, error) {
	var sensor models.SensorType

	stmt, err := d.DB().Prepare(d.Stmt().SelectById())
	if err != nil {
		log.Println("Error preparing the fetch sensor type sql")
		return sensor, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&sensor.Id,
		&sensor.TypeName)

	if err != nil {
		log.Println("Error querying for the sensor type")
		return sensor, err
	}

	return sensor, nil
}

func (d *SensorType) SelectByParentId(id int) ([]models.Model, error) {
	return nil, errors.New("no parent exists for sensor types")
}

func (d *SensorType) SelectBySecondParentId(id int) ([]models.Model, error) {
	return nil, errors.New("no second parent exists for sensor types")
}

func (d *SensorType) Insert(data models.Model) (models.Model, error) {
	sensorData, ok := data.(models.SensorType)
	if !ok {
		log.Println("Invalid data type for insert")
		return nil, sql.ErrNoRows
	}

	stmt, err := d.DB().Prepare(d.Stmt().Insert())
	if err != nil {
		log.Println("Error preparing insert statement")
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(sensorData.TypeName)
	if err != nil {
		log.Println("Error executing insert statement")
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Println("Error getting last insert id")
		return nil, err
	}

	sensorData.Id = int(id)
	return sensorData, nil
}

func (d *SensorType) Update(data models.Model) error {
	sensorData, ok := data.(models.SensorType)
	if !ok {
		log.Println("Invalid data type for update")
		return sql.ErrNoRows
	}

	stmt, err := d.DB().Prepare(d.Stmt().Update())
	if err != nil {
		log.Println("Error preparing update statement")
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(sensorData.TypeName, sensorData.Id)
	if err != nil {
		log.Println("Error executing update statement")
		return err
	}

	return nil
}

func (d *SensorType) Delete(id int) error {
	stmt, err := d.DB().Prepare(d.Stmt().Delete())
	if err != nil {
		log.Println("Error preparing delete statement")
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Println("Error executing delete statement")
		return err
	}

	return nil
}

func (d *SensorType) DeleteAll() error {
	stmt, err := d.DB().Prepare(d.Stmt().DeleteAll())
	if err != nil {
		log.Println("Error preparing delete all statement")
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		log.Println("Error executing delete all statement")
		return err
	}

	return nil
}

func (d *SensorType) DeleteByParentId(id int) error {
	log.Println("No parent exists for sensor types, nothing to delete")
	return nil
}

func (d *SensorType) DeleteBySecondParentId(id int) error {
	log.Println("No second parent exists for sensor types, nothing to delete")
	return nil
}
