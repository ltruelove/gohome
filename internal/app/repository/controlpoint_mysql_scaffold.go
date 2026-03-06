package repository

import (
	"database/sql"
	"log"

	"github.com/ltruelove/gohome/internal/app/dto"
	"github.com/ltruelove/gohome/internal/app/models"
)

type mysqlControlPointRepository struct {
	db *sql.DB
}

func NewMySQLControlPointRepository(db *sql.DB) ControlPointRepository {
	return &mysqlControlPointRepository{db: db}
}

func (r *mysqlControlPointRepository) FetchAll() ([]models.ControlPoint, error) {
	rows, err := r.db.Query("SELECT id, name, ipaddress, mac FROM controlpoint")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cps []models.ControlPoint
	for rows.Next() {
		var cp models.ControlPoint
		if err := rows.Scan(&cp.Id, &cp.Name, &cp.IpAddress, &cp.Mac); err != nil {
			return nil, err
		}
		cps = append(cps, cp)
	}
	return cps, nil
}

func (r *mysqlControlPointRepository) FetchAllAvailable() ([]models.ControlPoint, error) {
	query := `SELECT id, name, ipaddress, mac FROM controlpoint AS c WHERE (SELECT COUNT(id) FROM controlpointnodes WHERE controlpointid = c.id) < 20`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cps []models.ControlPoint
	for rows.Next() {
		var cp models.ControlPoint
		if err := rows.Scan(&cp.Id, &cp.Name, &cp.IpAddress, &cp.Mac); err != nil {
			return nil, err
		}
		cps = append(cps, cp)
	}
	return cps, nil
}

func (r *mysqlControlPointRepository) FetchAllNodes(controlPointId int) ([]dto.ControlPointNode, error) {
	rows, err := r.db.Query(`SELECT node.id, node.name, node.mac, cpn.id AS relationid FROM controlpointnodes AS cpn INNER JOIN node ON node.id = cpn.nodeid WHERE cpn.controlpointid = ?`, controlPointId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []dto.ControlPointNode
	for rows.Next() {
		var rec dto.ControlPointNode
		if err := rows.Scan(&rec.Id, &rec.Name, &rec.Mac, &rec.RelationId); err != nil {
			return nil, err
		}
		list = append(list, rec)
	}
	return list, nil
}

func (r *mysqlControlPointRepository) FetchById(id int) (models.ControlPoint, error) {
	var cp models.ControlPoint
	err := r.db.QueryRow("SELECT id, name, ipaddress, mac FROM controlpoint WHERE id = ?", id).Scan(&cp.Id, &cp.Name, &cp.IpAddress, &cp.Mac)
	if err != nil {
		return cp, err
	}
	return cp, nil
}

func (r *mysqlControlPointRepository) FetchByMac(mac string) (models.ControlPoint, error) {
	var cp models.ControlPoint
	err := r.db.QueryRow("SELECT id, name, ipaddress, mac FROM controlpoint WHERE mac = ?", mac).Scan(&cp.Id, &cp.Name, &cp.IpAddress, &cp.Mac)
	if err != nil {
		return cp, err
	}
	return cp, nil
}

func (r *mysqlControlPointRepository) Create(cp *models.ControlPoint) error {
	res, err := r.db.Exec("INSERT INTO controlpoint (name, ipaddress, mac) VALUES (?, ?, ?)", cp.Name, cp.IpAddress, cp.Mac)
	if err != nil {
		log.Println("Error creating control point:", err)
		return err
	}
	id, _ := res.LastInsertId()
	cp.Id = int(id)
	return nil
}

func (r *mysqlControlPointRepository) VerifyIdIsNew(id int) (bool, error) {
	_, err := r.FetchById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func (r *mysqlControlPointRepository) UpdateIp(cp *models.ControlPoint) error {
	_, err := r.db.Exec("UPDATE controlpoint SET ipaddress = ? WHERE id = ?", cp.IpAddress, cp.Id)
	return err
}

func (r *mysqlControlPointRepository) Update(cp *models.ControlPoint) error {
	_, err := r.db.Exec("UPDATE controlpoint SET name = ?, ipaddress = ?, mac = ? WHERE id = ?", cp.Name, cp.IpAddress, cp.Mac, cp.Id)
	return err
}

func (r *mysqlControlPointRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM controlpoint WHERE id = ?", id)
	return err
}

func (r *mysqlControlPointRepository) AddNodeToControlPoint(cpnode *models.ControlPointNode) error {
	_, err := r.db.Exec("INSERT INTO controlpointnodes (controlpointid, nodeid) VALUES (?, ?)", cpnode.ControlPointId, cpnode.NodeId)
	return err
}
