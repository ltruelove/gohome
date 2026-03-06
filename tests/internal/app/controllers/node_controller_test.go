package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	controllers "github.com/ltruelove/gohome/internal/app/controllers"
	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/stretchr/testify/assert"
)

// fakeNodeRepo implements NodeRepository for testing
type fakeNodeRepo struct {
	FetchAllFunc                  func() ([]models.Node, error)
	FetchByIdFunc                 func(int) (models.Node, error)
	CreateFunc                    func(*models.Node) error
	VerifyIdIsNewFunc             func(int) (bool, error)
	UpdateFunc                    func(*models.Node) error
	DeleteFunc                    func(int) error
	FetchIndividualFunc           func(int) (models.Node, error)
	FetchControlPointByNodeFunc   func(int) (models.ControlPoint, error)
	FetchNodeSwitchesFunc         func(int) ([]models.NodeSwitch, error)
	FetchNodeSwitchFunc           func(int) (models.NodeSwitch, error)
	GetSensorLogDataFunc          func(int, time.Time, time.Time) ([]models.NodeSensorLog, error)
	GetTempLogDataByLogIdFunc     func(int) ([]models.TempLogData, error)
	GetMoistureLogDataByLogIdFunc func(int) ([]models.MoistureLogData, error)
	GetResistorLogDataByLogIdFunc func(int) ([]models.ResistorLogData, error)
	GetMagneticLogDataByLogIdFunc func(int) ([]models.MagneticLogData, error)
	CreateNewLogFunc              func(models.NodeData) error
	CreateNodeSensorFunc          func(*models.NodeSensor) error
	CreateNodeSwitchFunc          func(*models.NodeSwitch) error
}

func (f *fakeNodeRepo) FetchAll() ([]models.Node, error) {
	if f.FetchAllFunc != nil {
		return f.FetchAllFunc()
	}
	return nil, nil
}

func (f *fakeNodeRepo) FetchById(id int) (models.Node, error) {
	if f.FetchByIdFunc != nil {
		return f.FetchByIdFunc(id)
	}
	return models.Node{}, errors.New("not implemented")
}

func (f *fakeNodeRepo) Create(n *models.Node) error {
	if f.CreateFunc != nil {
		return f.CreateFunc(n)
	}
	return nil
}

func (f *fakeNodeRepo) VerifyIdIsNew(id int) (bool, error) {
	if f.VerifyIdIsNewFunc != nil {
		return f.VerifyIdIsNewFunc(id)
	}
	return false, nil
}

func (f *fakeNodeRepo) Update(n *models.Node) error {
	if f.UpdateFunc != nil {
		return f.UpdateFunc(n)
	}
	return nil
}

func (f *fakeNodeRepo) Delete(id int) error {
	if f.DeleteFunc != nil {
		return f.DeleteFunc(id)
	}
	return nil
}

func (f *fakeNodeRepo) FetchIndividual(id int) (models.Node, error) {
	if f.FetchIndividualFunc != nil {
		return f.FetchIndividualFunc(id)
	}
	return models.Node{}, nil
}

func (f *fakeNodeRepo) FetchControlPointByNode(nodeId int) (models.ControlPoint, error) {
	if f.FetchControlPointByNodeFunc != nil {
		return f.FetchControlPointByNodeFunc(nodeId)
	}
	return models.ControlPoint{}, nil
}

func (f *fakeNodeRepo) FetchNodeSwitches(nodeId int) ([]models.NodeSwitch, error) {
	if f.FetchNodeSwitchesFunc != nil {
		return f.FetchNodeSwitchesFunc(nodeId)
	}
	return nil, nil
}

func (f *fakeNodeRepo) FetchNodeSwitch(id int) (models.NodeSwitch, error) {
	if f.FetchNodeSwitchFunc != nil {
		return f.FetchNodeSwitchFunc(id)
	}
	return models.NodeSwitch{}, nil
}

func (f *fakeNodeRepo) GetSensorLogData(nodeId int, start time.Time, end time.Time) ([]models.NodeSensorLog, error) {
	if f.GetSensorLogDataFunc != nil {
		return f.GetSensorLogDataFunc(nodeId, start, end)
	}
	return nil, nil
}

func (f *fakeNodeRepo) GetTempLogDataByLogId(logId int) ([]models.TempLogData, error) {
	if f.GetTempLogDataByLogIdFunc != nil {
		return f.GetTempLogDataByLogIdFunc(logId)
	}
	return nil, nil
}

func (f *fakeNodeRepo) GetMoistureLogDataByLogId(logId int) ([]models.MoistureLogData, error) {
	if f.GetMoistureLogDataByLogIdFunc != nil {
		return f.GetMoistureLogDataByLogIdFunc(logId)
	}
	return nil, nil
}

func (f *fakeNodeRepo) GetResistorLogDataByLogId(logId int) ([]models.ResistorLogData, error) {
	if f.GetResistorLogDataByLogIdFunc != nil {
		return f.GetResistorLogDataByLogIdFunc(logId)
	}
	return nil, nil
}

func (f *fakeNodeRepo) GetMagneticLogDataByLogId(logId int) ([]models.MagneticLogData, error) {
	if f.GetMagneticLogDataByLogIdFunc != nil {
		return f.GetMagneticLogDataByLogIdFunc(logId)
	}
	return nil, nil
}

func (f *fakeNodeRepo) CreateNewLog(item models.NodeData) error {
	if f.CreateNewLogFunc != nil {
		return f.CreateNewLogFunc(item)
	}
	return nil
}

func (f *fakeNodeRepo) CreateNodeSensor(sensor *models.NodeSensor) error {
	if f.CreateNodeSensorFunc != nil {
		return f.CreateNodeSensorFunc(sensor)
	}
	return nil
}

func (f *fakeNodeRepo) CreateNodeSwitch(sw *models.NodeSwitch) error {
	if f.CreateNodeSwitchFunc != nil {
		return f.CreateNodeSwitchFunc(sw)
	}
	return nil
}

func TestNodeController_GetAll(t *testing.T) {
	fetchAll := func() ([]models.Node, error) {
		return []models.Node{{Id: 1, Name: "n1", Mac: "m1"}, {Id: 2, Name: "n2", Mac: "m2"}}, nil
	}
	repo := &fakeNodeRepo{FetchAllFunc: fetchAll}

	ctrl := controllers.NewNodeControllerWithDeps(repo, nil)

	req := httptest.NewRequest("GET", "/node", nil)
	rr := httptest.NewRecorder()

	ctrl.GetAll(rr, req)

	assert.Equal(t, 200, rr.Code)

	var got []models.Node
	err := json.Unmarshal(rr.Body.Bytes(), &got)
	assert.NoError(t, err)
	assert.Len(t, got, 2)
}

func TestNodeController_GetById_SuccessAndNotFound(t *testing.T) {
	fetchById := func(id int) (models.Node, error) {
		if id == 1 {
			return models.Node{Id: 1, Name: "n1", Mac: "m1"}, nil
		}
		return models.Node{}, errors.New("not found")
	}
	repo := &fakeNodeRepo{FetchByIdFunc: fetchById}

	ctrl := controllers.NewNodeControllerWithDeps(repo, nil)

	// success
	req := httptest.NewRequest("GET", "/node/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	rr := httptest.NewRecorder()
	ctrl.GetById(rr, req)
	assert.Equal(t, 200, rr.Code)
	var n models.Node
	err := json.Unmarshal(rr.Body.Bytes(), &n)
	assert.NoError(t, err)
	assert.Equal(t, 1, n.Id)

	// not found
	req2 := httptest.NewRequest("GET", "/node/999", nil)
	req2 = mux.SetURLVars(req2, map[string]string{"id": "999"})
	rr2 := httptest.NewRecorder()
	ctrl.GetById(rr2, req2)
	assert.Equal(t, 404, rr2.Code)
}

func TestNodeController_Create_Success(t *testing.T) {
	createFn := func(n *models.Node) error {
		n.Id = 55
		return nil
	}
	repo := &fakeNodeRepo{CreateFunc: createFn}

	ctrl := controllers.NewNodeControllerWithDeps(repo, nil)

	payload := models.Node{Name: "nn", Mac: "mm", IpAddress: "1.2.3.4"}
	b, _ := json.Marshal(payload)
	req := httptest.NewRequest("POST", "/node", bytes.NewReader(b))
	rr := httptest.NewRecorder()

	ctrl.Create(rr, req)

	assert.Equal(t, 200, rr.Code)
	var out models.Node
	err := json.Unmarshal(rr.Body.Bytes(), &out)
	assert.NoError(t, err)
	assert.Equal(t, 55, out.Id)
}
