package data

import (
	"database/sql"
	"errors"
	"log"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data/statements"
	"github.com/ltruelove/gohome/internal/app/models"
)

type SensorTypeData struct {
	db   *sql.DB
	stmt statements.CrudStatement
}

func NewSensorTypeData(db *sql.DB, config *config.Configuration) CrudDataInterface {
	return &SensorTypeData{
		db:   db,
		stmt: statements.NewSensorTypeDataStatements(config),
	}
}

func (d *SensorTypeData) DB() *sql.DB {
	return d.db
}

func (d *SensorTypeData) Stmt() statements.CrudStatement {
	return d.stmt
}

func (d *SensorTypeData) SelectAll() ([]models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectAll())
	if err != nil {
		log.Println("Error preparing fetch all sensor types sql")
		return nil, err
	}

	var sensorTypes []models.Model

	rows, err := stmt.Query()
	if err != nil {
		log.Println("Error querying for all sensor types")
		return nil, err
	}
	defer stmt.Close()

	for rows.Next() {
		var sensorType models.SensorTypeData
		rows.Scan(&sensorType.Id,
			&sensorType.SensorTypeId,
			&sensorType.Name,
			&sensorType.ValueType)
		sensorTypes = append(sensorTypes, sensorType)
	}

	return sensorTypes, nil
}

func (d *SensorTypeData) SelectById(id int) (models.Model, error) {
	var sensorTypeData models.SensorTypeData

	stmt, err := d.DB().Prepare(d.Stmt().SelectById())
	if err != nil {
		log.Printf("Error preparing SelectById statement: %v", err)
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&sensorTypeData.Id,
		&sensorTypeData.SensorTypeId,
		&sensorTypeData.Name,
		&sensorTypeData.ValueType)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("no sensor type data record found with the given id")
		}
		log.Printf("Error executing SelectById query: %v", err)
		return nil, err
	}

	return sensorTypeData, nil
}

func (d *SensorTypeData) SelectByParentId(id int) ([]models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectByParentId())
	if err != nil {
		log.Println("Error preparing SelectByParentId statement")
		return nil, err
	}
	defer stmt.Close()

	var sensorTypes []models.Model
	rows, err := stmt.Query(id)
	if err != nil {
		log.Println("Error querying for sensor types by parent ID")
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var sensorType models.SensorTypeData
		err = rows.Scan(&sensorType.Id,
			&sensorType.SensorTypeId,
			&sensorType.Name,
			&sensorType.ValueType)
		if err != nil {
			log.Println("Error scanning row for sensor type data")
			return nil, err
		}
		sensorTypes = append(sensorTypes, sensorType)
	}

	return sensorTypes, nil
}

func (d *SensorTypeData) SelectBySecondParentId(id int) ([]models.Model, error) {
	return nil, errors.New("SelectBySecondParentId not implemented for SensorTypeData")
}

func (d *SensorTypeData) Insert(data models.Model) (models.Model, error) {
	sensorTypeData, ok := data.(models.SensorTypeData)
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

	res, err := stmt.Exec(sensorTypeData.SensorTypeId, sensorTypeData.Name, sensorTypeData.ValueType)
	if err != nil {
		log.Println("Error executing insert statement")
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Println("Error getting last insert id")
		return nil, err
	}

	sensorTypeData.Id = int(id)
	return sensorTypeData, nil
}

func (d *SensorTypeData) Update(data models.Model) error {
	sensorTypeData, ok := data.(models.SensorTypeData)
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

	_, err = stmt.Exec(sensorTypeData.SensorTypeId, sensorTypeData.Name, sensorTypeData.ValueType, sensorTypeData.Id)
	if err != nil {
		log.Println("Error executing update statement")
		return err
	}

	return nil
}

func (d *SensorTypeData) Delete(id int) error {
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

func (d *SensorTypeData) DeleteByParentId(id int) error {
	stmt, err := d.DB().Prepare(d.Stmt().DeleteByParentId())
	if err != nil {
		log.Println("Error preparing delete by parent ID statement")
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Println("Error executing delete by parent ID statement")
		return err
	}

	return nil
}

func (d *SensorTypeData) DeleteAll() error {
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

func (d *SensorTypeData) DeleteBySecondParentId(id int) error {
	return errors.New("no second parent exists for sensor types, nothing to delete")
}
