package repository

import (
	"time"

	"github.com/ltruelove/gohome/internal/app/models"
)

type NodeRepository interface {
	FetchAll() ([]models.Node, error)
	FetchById(id int) (models.Node, error)
	Create(node *models.Node) error
	VerifyIdIsNew(id int) (bool, error)
	Update(node *models.Node) error
	Delete(id int) error
	FetchIndividual(id int) (models.Node, error)
	FetchControlPointByNode(nodeId int) (models.ControlPoint, error)
	FetchNodeSwitches(nodeId int) ([]models.NodeSwitch, error)
	FetchNodeSwitch(id int) (models.NodeSwitch, error)
	GetSensorLogData(nodeId int, start time.Time, end time.Time) ([]models.NodeSensorLog, error)
	GetTempLogDataByLogId(logId int) ([]models.TempLogData, error)
	CreateNewLog(item models.NodeData) error
	GetMoistureLogDataByLogId(logId int) ([]models.MoistureLogData, error)
	GetResistorLogDataByLogId(logId int) ([]models.ResistorLogData, error)
	GetMagneticLogDataByLogId(logId int) ([]models.MagneticLogData, error)
	CreateNodeSensor(sensor *models.NodeSensor) error
	CreateNodeSwitch(sw *models.NodeSwitch) error
}
