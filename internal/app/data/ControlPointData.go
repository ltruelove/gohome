package data

import (
	"database/sql"
	"errors"
	"log"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data/statements"
	"github.com/ltruelove/gohome/internal/app/models"
)

type ControlPointData struct {
	db   *sql.DB
	stmt statements.CrudStatement
}

func NewControlPointData(db *sql.DB, config *config.Configuration) CrudDataInterface {
	return &ControlPointData{
		db:   db,
		stmt: statements.NewControlPointDataStatements(config),
	}
}

func (d *ControlPointData) DB() *sql.DB {
	return d.db
}

func (d *ControlPointData) Stmt() statements.CrudStatement {
	return d.stmt
}

func (d *ControlPointData) SelectAll() ([]models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectAll())
	if err != nil {
		log.Println("Error preparing fetch all control point sql")
		return nil, err
	}

	var controlPoints []models.Model

	rows, err := stmt.Query()
	if err != nil {
		log.Println("Error querying for all control points")
		return nil, err
	}
	defer stmt.Close()

	for rows.Next() {
		var controlPoint models.ControlPoint
		err := rows.Scan(&controlPoint.Id,
			&controlPoint.Name,
			&controlPoint.IpAddress,
			&controlPoint.Mac)

		if err != nil {
			log.Println("Error scanning magnetic log")
			return nil, err
		}
		controlPoints = append(controlPoints, &controlPoint)
	}

	return controlPoints, nil
}

func (d *ControlPointData) SelectById(id int) (models.Model, error) {
	stmt, err := d.DB().Prepare(d.Stmt().SelectById())
	if err != nil {
		log.Println("Error preparing fetch control point by id sql")
		return nil, err
	}

	var controlPoint models.ControlPoint

	err = stmt.QueryRow(id).Scan(&controlPoint.Id,
		&controlPoint.Name,
		&controlPoint.IpAddress,
		&controlPoint.Mac)

	if err != nil {
		log.Println("Error querying for control point by id")
		return nil, err
	}

	return &controlPoint, nil
}

func (d *ControlPointData) SelectByParentId(id int) ([]models.Model, error) {
	return nil, errors.New("ControlPoint does not support parent ID queries")
}

func (d *ControlPointData) SelectBySecondParentId(id int) ([]models.Model, error) {
	return nil, errors.New("ControlPoint does not support second parent ID queries")
}

func (d *ControlPointData) Insert(data models.Model) (models.Model, error) {
	controlPoint, ok := data.(*models.ControlPoint)
	if !ok {
		log.Println("Error casting model to ControlPoint")
		return nil, errors.New("Invalid model type")
	}

	stmt, err := d.DB().Prepare(d.Stmt().Insert())
	if err != nil {
		log.Println("Error preparing insert control point sql")
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(controlPoint.Name, controlPoint.IpAddress, controlPoint.Mac)
	if err != nil {
		log.Println("Error executing insert control point sql")
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Println("Error getting last insert id for control point")
		return nil, err
	}

	controlPoint.Id = int(id)
	return controlPoint, nil
}

func (d *ControlPointData) Update(data models.Model) error {
	controlPoint, ok := data.(*models.ControlPoint)
	if !ok {
		log.Println("Error casting model to ControlPoint")
		return errors.New("Invalid model type")
	}

	stmt, err := d.DB().Prepare(d.Stmt().Update())
	if err != nil {
		log.Println("Error preparing update control point sql")
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(controlPoint.Name, controlPoint.IpAddress, controlPoint.Mac, controlPoint.Id)
	if err != nil {
		log.Println("Error executing update control point sql")
		return err
	}

	return nil
}

func (d *ControlPointData) Delete(id int) error {
	stmt, err := d.DB().Prepare(d.Stmt().Delete())
	if err != nil {
		log.Println("Error preparing delete control point sql")
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Println("Error executing delete control point sql")
		return err
	}

	return nil
}

func (d *ControlPointData) DeleteAll() error {
	stmt, err := d.DB().Prepare(d.Stmt().DeleteAll())
	if err != nil {
		log.Println("Error preparing delete all control points sql")
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		log.Println("Error executing delete all control points sql")
		return err
	}

	return nil
}

func (d *ControlPointData) DeleteByParentId(id int) error {
	return errors.New("ControlPoint does not support delete by parent ID")
}

func (d *ControlPointData) DeleteBySecondParentId(id int) error {
	return errors.New("ControlPoint does not support delete by second parent ID")
}
