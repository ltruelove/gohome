package repository

import (
	"database/sql"
	"time"

	"github.com/ltruelove/gohome/internal/app/models"
)

type nodeRepositoryAdapter struct {
	db     *sql.DB
	cpRepo ControlPointRepository
}

func NewNodeRepository(db *sql.DB, dbType string) NodeRepository {
	if dbType == "mysql" {
		return NewMySQLNodeRepository(db)
	}
	return &nodeRepositoryAdapter{db: db, cpRepo: NewControlPointRepository(db, dbType)}
}

func (r *nodeRepositoryAdapter) FetchAll() ([]models.Node, error) {
	return data.FetchAllNodes(r.db)
}

func (r *nodeRepositoryAdapter) FetchById(id int) (models.Node, error) {
	return data.FetchNode(id, r.db)
}

func (r *nodeRepositoryAdapter) Create(node *models.Node) error {
	return data.CreateNode(node, r.db)
}

func (r *nodeRepositoryAdapter) VerifyIdIsNew(id int) (bool, error) {
	return data.VerifyNodeIdIsNew(id, r.db)
}

func (r *nodeRepositoryAdapter) Update(node *models.Node) error {
	return data.UpdateNode(node, r.db)
}

func (r *nodeRepositoryAdapter) Delete(id int) error {
	return data.DeleteNode(id, r.db)
}

func (r *nodeRepositoryAdapter) FetchIndividual(id int) (models.Node, error) {
	return data.FetchIndividualNode(id, r.db)
}

func (r *nodeRepositoryAdapter) FetchControlPointByNode(nodeId int) (models.ControlPoint, error) {
	return r.cpRepo.FetchControlPointByNode(nodeId)
}

func (r *nodeRepositoryAdapter) FetchNodeSwitches(nodeId int) ([]models.NodeSwitch, error) {
	return data.FetchNodeSwitches(nodeId, r.db)
}

func (r *nodeRepositoryAdapter) FetchNodeSwitch(id int) (models.NodeSwitch, error) {
	return data.FetchNodeSwitch(id, r.db)
}

func (r *nodeRepositoryAdapter) GetSensorLogData(nodeId int, start time.Time, end time.Time) ([]models.NodeSensorLog, error) {
	return data.GetSensorLogData(nodeId, r.db, start, end)
}

func (r *nodeRepositoryAdapter) GetTempLogDataByLogId(logId int) ([]models.TempLogData, error) {
	return data.GetTempLogDataByLogId(logId, r.db)
}

func (r *nodeRepositoryAdapter) CreateNewLog(item models.NodeData) error {
	return data.CreateNewLog(item, r.db)
}

func (r *nodeRepositoryAdapter) GetMoistureLogDataByLogId(logId int) ([]models.MoistureLogData, error) {
	return data.GetMoistureLogDataByLogId(logId, r.db)
}

func (r *nodeRepositoryAdapter) GetResistorLogDataByLogId(logId int) ([]models.ResistorLogData, error) {
	return data.GetResistorLogDataByLogId(logId, r.db)
}

func (r *nodeRepositoryAdapter) GetMagneticLogDataByLogId(logId int) ([]models.MagneticLogData, error) {
	return data.GetMagneticLogDataByLogId(logId, r.db)
}

func (r *nodeRepositoryAdapter) CreateNodeSensor(sensor *models.NodeSensor) error {
	return data.CreateNodeSensor(sensor, r.db)
}

func (r *nodeRepositoryAdapter) CreateNodeSwitch(sw *models.NodeSwitch) error {
	return data.CreateNodeSwitch(sw, r.db)
}
