package repository

import (
	"database/sql"

	"github.com/ltruelove/gohome/internal/app/data"
	"github.com/ltruelove/gohome/internal/app/dto"
	"github.com/ltruelove/gohome/internal/app/models"
)

type controlPointRepositoryAdapter struct {
	db *sql.DB
}

func NewControlPointRepository(db *sql.DB, dbType string) ControlPointRepository {
	if dbType == "mysql" {
		return NewMySQLControlPointRepository(db)
	}
	return &controlPointRepositoryAdapter{db: db}
}

func (r *controlPointRepositoryAdapter) FetchAll() ([]models.ControlPoint, error) {
	return data.FetchAllControlPoints(r.db)
}

func (r *controlPointRepositoryAdapter) FetchAllAvailable() ([]models.ControlPoint, error) {
	return data.FetchAllAvailableControlPoints(r.db)
}

func (r *controlPointRepositoryAdapter) FetchAllNodes(controlPointId int) ([]dto.ControlPointNode, error) {
	return data.FetchAllControlPointNodes(controlPointId, r.db)
}

func (r *controlPointRepositoryAdapter) FetchById(id int) (models.ControlPoint, error) {
	return data.FetchControlPoint(id, r.db)
}

func (r *controlPointRepositoryAdapter) FetchByMac(mac string) (models.ControlPoint, error) {
	return data.FetchControlPointByMac(mac, r.db)
}

func (r *controlPointRepositoryAdapter) Create(cp *models.ControlPoint) error {
	return data.CreateControlPoint(cp, r.db)
}

func (r *controlPointRepositoryAdapter) VerifyIdIsNew(id int) (bool, error) {
	return data.VerifyControlPointIdIsNew(id, r.db)
}

func (r *controlPointRepositoryAdapter) UpdateIp(cp *models.ControlPoint) error {
	return data.UpdateControlPointIp(cp, r.db)
}

func (r *controlPointRepositoryAdapter) Update(cp *models.ControlPoint) error {
	return data.UpdateControlPoint(cp, r.db)
}

func (r *controlPointRepositoryAdapter) Delete(id int) error {
	return data.DeleteControlPoint(id, r.db)
}

func (r *controlPointRepositoryAdapter) AddNodeToControlPoint(cpnode *models.ControlPointNode) error {
	return data.AddNodeToControlPoint(cpnode, r.db)
}
