package repository

import (
	"github.com/ltruelove/gohome/internal/app/dto"
	"github.com/ltruelove/gohome/internal/app/models"
)

type ControlPointRepository interface {
	FetchAll() ([]models.ControlPoint, error)
	FetchAllAvailable() ([]models.ControlPoint, error)
	FetchAllNodes(controlPointId int) ([]dto.ControlPointNode, error)
	FetchById(id int) (models.ControlPoint, error)
	FetchByMac(mac string) (models.ControlPoint, error)
	Create(cp *models.ControlPoint) error
	VerifyIdIsNew(id int) (bool, error)
	UpdateIp(cp *models.ControlPoint) error
	Update(cp *models.ControlPoint) error
	Delete(id int) error
	AddNodeToControlPoint(cpnode *models.ControlPointNode) error
}
