package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/dto"
	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/stretchr/testify/assert"
)

// fakeNodeRepo implements repository.NodeRepository (minimal subset used by controller)
type fakeNodeRepo struct {
	serverAddr string
}

// fakeNodeRepoInpkg is a configurable fake used by multiple node tests
type fakeNodeRepoInpkg struct {
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
	CreateNewLogFunc              func(models.NodeData) error
	GetMoistureLogDataByLogIdFunc func(int) ([]models.MoistureLogData, error)
	GetResistorLogDataByLogIdFunc func(int) ([]models.ResistorLogData, error)
	GetMagneticLogDataByLogIdFunc func(int) ([]models.MagneticLogData, error)
	CreateNodeSensorFunc          func(*models.NodeSensor) error
	CreateNodeSwitchFunc          func(*models.NodeSwitch) error
}

func (f *fakeNodeRepoInpkg) FetchAll() ([]models.Node, error) {
	if f.FetchAllFunc != nil {
		return f.FetchAllFunc()
	}
	return nil, nil
}

func (f *fakeNodeRepoInpkg) FetchById(id int) (models.Node, error) {
	if f.FetchByIdFunc != nil {
		return f.FetchByIdFunc(id)
	}
	return models.Node{}, nil
}

func (f *fakeNodeRepoInpkg) Create(n *models.Node) error {
	if f.CreateFunc != nil {
		return f.CreateFunc(n)
	}
	return nil
}

func (f *fakeNodeRepoInpkg) VerifyIdIsNew(id int) (bool, error) {
	if f.VerifyIdIsNewFunc != nil {
		return f.VerifyIdIsNewFunc(id)
	}
	return false, nil
}

func (f *fakeNodeRepoInpkg) Update(n *models.Node) error {
	if f.UpdateFunc != nil {
		return f.UpdateFunc(n)
	}
	return nil
}

func (f *fakeNodeRepoInpkg) Delete(id int) error {
	if f.DeleteFunc != nil {
		return f.DeleteFunc(id)
	}
	return nil
}

func (f *fakeNodeRepoInpkg) FetchIndividual(id int) (models.Node, error) {
	if f.FetchIndividualFunc != nil {
		return f.FetchIndividualFunc(id)
	}
	return models.Node{}, nil
}

func (f *fakeNodeRepoInpkg) FetchControlPointByNode(nodeId int) (models.ControlPoint, error) {
	if f.FetchControlPointByNodeFunc != nil {
		return f.FetchControlPointByNodeFunc(nodeId)
	}
	return models.ControlPoint{}, nil
}

func (f *fakeNodeRepoInpkg) FetchNodeSwitches(nodeId int) ([]models.NodeSwitch, error) {
	if f.FetchNodeSwitchesFunc != nil {
		return f.FetchNodeSwitchesFunc(nodeId)
	}
	return nil, nil
}

func (f *fakeNodeRepoInpkg) FetchNodeSwitch(id int) (models.NodeSwitch, error) {
	if f.FetchNodeSwitchFunc != nil {
		return f.FetchNodeSwitchFunc(id)
	}
	return models.NodeSwitch{}, nil
}

func (f *fakeNodeRepo) FetchAll() ([]models.Node, error) {
	return []models.Node{{Id: 1, Name: "n1", Mac: "MAC1"}}, nil
}
func (f *fakeNodeRepo) FetchById(id int) (models.Node, error) {
	return models.Node{Id: id, Name: "n1", Mac: "AA:BB:CC", IpAddress: "127.0.0.1"}, nil
}
func (f *fakeNodeRepo) Create(node *models.Node) error     { return nil }
func (f *fakeNodeRepo) VerifyIdIsNew(id int) (bool, error) { return false, nil }
func (f *fakeNodeRepo) Update(node *models.Node) error     { return nil }
func (f *fakeNodeRepo) Delete(id int) error                { return nil }
func (f *fakeNodeRepo) FetchIndividual(id int) (models.Node, error) {
	return models.Node{Id: id, Mac: "AA:BB:CC"}, nil
}
func (f *fakeNodeRepo) FetchControlPointByNode(nodeId int) (models.ControlPoint, error) {
	return models.ControlPoint{Id: 1, IpAddress: f.serverAddr, Name: "cp1", Mac: "CPMAC"}, nil
}
func (f *fakeNodeRepo) FetchNodeSwitches(nodeId int) ([]models.NodeSwitch, error) {
	return []models.NodeSwitch{{Id: 10, NodeId: nodeId, Name: "sw1"}}, nil
}
func (f *fakeNodeRepo) FetchNodeSwitch(id int) (models.NodeSwitch, error) {
	return models.NodeSwitch{Id: id, NodeId: 1, MomentaryPressDuration: 50}, nil
}
func (f *fakeNodeRepo) GetSensorLogData(nodeId int, start time.Time, end time.Time) ([]models.NodeSensorLog, error) {
	return nil, nil
}
func (f *fakeNodeRepo) GetTempLogDataByLogId(logId int) ([]models.TempLogData, error) {
	return nil, nil
}
func (f *fakeNodeRepo) CreateNewLog(item models.NodeData) error { return nil }
func (f *fakeNodeRepo) GetMoistureLogDataByLogId(logId int) ([]models.MoistureLogData, error) {
	return nil, nil
}
func (f *fakeNodeRepo) GetResistorLogDataByLogId(logId int) ([]models.ResistorLogData, error) {
	return nil, nil
}
func (f *fakeNodeRepo) GetMagneticLogDataByLogId(logId int) ([]models.MagneticLogData, error) {
	return nil, nil
}
func (f *fakeNodeRepo) CreateNodeSensor(sensor *models.NodeSensor) error { return nil }
func (f *fakeNodeRepo) CreateNodeSwitch(sw *models.NodeSwitch) error     { return nil }

// fakeCP is a top-level fake implementing control-point repo methods used by tests
type fakeCP struct{}

func (f *fakeCP) FetchAll() ([]models.ControlPoint, error)                         { return nil, nil }
func (f *fakeCP) FetchAllAvailable() ([]models.ControlPoint, error)                { return nil, nil }
func (f *fakeCP) FetchAllNodes(controlPointId int) ([]dto.ControlPointNode, error) { return nil, nil }
func (f *fakeCP) FetchById(id int) (models.ControlPoint, error) {
	return models.ControlPoint{Id: id, IpAddress: "127.0.0.1"}, nil
}
func (f *fakeCP) FetchByMac(mac string) (models.ControlPoint, error) {
	return models.ControlPoint{}, nil
}
func (f *fakeCP) Create(cp *models.ControlPoint) error                        { return nil }
func (f *fakeCP) VerifyIdIsNew(id int) (bool, error)                          { return false, nil }
func (f *fakeCP) UpdateIp(cp *models.ControlPoint) error                      { return nil }
func (f *fakeCP) Update(cp *models.ControlPoint) error                        { return nil }
func (f *fakeCP) Delete(id int) error                                         { return nil }
func (f *fakeCP) AddNodeToControlPoint(cpnode *models.ControlPointNode) error { return nil }
func (f *fakeCP) FetchControlPointByNode(nodeId int) (models.ControlPoint, error) {
	return models.ControlPoint{}, nil
}

func TestNodeController_BasicFlows(t *testing.T) {
	// start a local server to accept controller outgoing HTTP requests
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/toggleNodeSwitch":
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("toggled"))
		case "/pressMomentary":
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("pressed"))
		case "/triggerUpdate":
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("update"))
		case "/nodeData":
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("{\"node\":\"data\"}"))
		case "/triggerRestart":
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("restart"))
		case "/nodeUpdateMode":
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("updatemode"))
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer srv.Close()

	// server URL contains scheme and host; controller builds URLs like http://<ip>/...
	// use srv.Listener.Addr
	addr := srv.Listener.Addr().String()

	repo := &fakeNodeRepo{serverAddr: addr}
	ctrl := NewNodeControllerWithDeps(repo, nil)

	// set valid PIN
	Config = config.Configuration{Pin: "1234"}

	// ToggleNodeSwitch
	pin := models.PinRequest{PinCode: "1234"}
	pb, _ := json.Marshal(pin)
	req := httptest.NewRequest("POST", "/node/switch/toggle/10", bytes.NewReader(pb))
	req = mux.SetURLVars(req, map[string]string{"id": "10"})
	rr := httptest.NewRecorder()
	ctrl.ToggleNodeSwitch(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	// PressNodeSwitch
	req2 := httptest.NewRequest("POST", "/node/switch/press/10", bytes.NewReader(pb))
	req2 = mux.SetURLVars(req2, map[string]string{"id": "10"})
	rr2 := httptest.NewRecorder()
	ctrl.PressNodeSwitch(rr2, req2)
	assert.Equal(t, http.StatusOK, rr2.Code)

	// TriggerUpdate
	req3 := httptest.NewRequest("POST", "/node/update/1", nil)
	req3 = mux.SetURLVars(req3, map[string]string{"id": "1"})
	rr3 := httptest.NewRecorder()
	ctrl.TriggerUpdate(rr3, req3)
	assert.Equal(t, http.StatusOK, rr3.Code)

	// GetNodeData (server returns JSON payload)
	req4 := httptest.NewRequest("GET", "/node/data/1", nil)
	req4 = mux.SetURLVars(req4, map[string]string{"id": "1"})
	rr4 := httptest.NewRecorder()
	ctrl.GetNodeData(rr4, req4)
	assert.Equal(t, http.StatusOK, rr4.Code)
	body, _ := io.ReadAll(rr4.Body)
	assert.Contains(t, string(body), "node")

	// LogNodeReading
	nd := models.NodeData{NodeId: 1, TemperatureF: 72.0, TemperatureC: 22.0, Humidity: 50.0, Moisture: 0, ResistorValue: 0, IsClosed: false, MagneticValue: false}
	nb, _ := json.Marshal(nd)
	req5 := httptest.NewRequest("POST", "/node/reading", bytes.NewReader(nb))
	rr5 := httptest.NewRecorder()
	ctrl.LogNodeReading(rr5, req5)
	assert.Equal(t, http.StatusOK, rr5.Code)
}

func TestNodeController_Delete_InPkg(t *testing.T) {
	// server to accept eraseNodeSettings
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/eraseNodeSettings" {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("erased"))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer srv.Close()

	repo := &fakeNodeRepo{serverAddr: srv.Listener.Addr().String()}
	ctrl := NewNodeControllerWithDeps(repo, nil)

	req := httptest.NewRequest("DELETE", "/node/1/delete", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	rr := httptest.NewRecorder()
	ctrl.Delete(rr, req)
	// controller may respond 200 or 204
	if rr.Code != http.StatusNoContent {
		assert.Equal(t, http.StatusOK, rr.Code)
	}
}

func TestNodeController_Register_InPkg(t *testing.T) {
	// fake control point repo

	nodeCreateCalled := false
	nodeRepo := &fakeNodeRepoInpkg{CreateFunc: func(n *models.Node) error { n.Id = 77; nodeCreateCalled = true; return nil }, CreateNodeSensorFunc: func(s *models.NodeSensor) error { return nil }, CreateNodeSwitchFunc: func(sw *models.NodeSwitch) error { return nil }}
	cpRepo := &fakeCP{}

	ctrl := NewNodeControllerWithDeps(nodeRepo, cpRepo)

	payload := dto.RegsiterNode{Node: models.Node{Name: "nreg", Mac: "mm"}, ControlPoint: models.ControlPoint{Id: 1}, Sensors: []models.NodeSensor{}, Switches: []models.NodeSwitch{}}
	b, _ := json.Marshal(payload)
	req := httptest.NewRequest("POST", "/node/register", bytes.NewReader(b))
	rr := httptest.NewRecorder()
	ctrl.Register(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.True(t, nodeCreateCalled)
}

func TestNodeController_GetAllNodeSwitches_InPkg(t *testing.T) {
	fetchFn := func(nodeId int) ([]models.NodeSwitch, error) {
		return []models.NodeSwitch{{Id: 5, NodeId: nodeId, Name: "s"}}, nil
	}
	repo := &fakeNodeRepoInpkg{FetchNodeSwitchesFunc: fetchFn}
	ctrl := NewNodeControllerWithDeps(repo, nil)

	req := httptest.NewRequest("GET", "/node/switchesByNode/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	rr := httptest.NewRecorder()
	ctrl.GetAllNodeSwitches(rr, req)
	assert.Equal(t, 200, rr.Code)
}

func TestNodeController_GetNodeLogData_InPkg(t *testing.T) {
	// prepare a log entry
	now := time.Now()
	logEntry := models.NodeSensorLog{Id: 9, NodeId: 1, DateLogged: now}
	fetchLogs := func(nodeId int, start time.Time, end time.Time) ([]models.NodeSensorLog, error) {
		return []models.NodeSensorLog{logEntry}, nil
	}
	tempFn := func(logId int) ([]models.TempLogData, error) {
		return []models.TempLogData{{Id: 1, NodeSensorLogId: logId, TemperatureF: 22.0}}, nil
	}
	repo := &fakeNodeRepoInpkg{FetchByIdFunc: func(id int) (models.Node, error) { return models.Node{Id: id}, nil }, GetSensorLogDataFunc: fetchLogs, GetTempLogDataByLogIdFunc: tempFn, GetMoistureLogDataByLogIdFunc: func(int) ([]models.MoistureLogData, error) { return nil, nil }, GetResistorLogDataByLogIdFunc: func(int) ([]models.ResistorLogData, error) { return nil, nil }, GetMagneticLogDataByLogIdFunc: func(int) ([]models.MagneticLogData, error) { return nil, nil }}

	ctrl := NewNodeControllerWithDeps(repo, nil)
	req := httptest.NewRequest("GET", "/node/logs/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	rr := httptest.NewRecorder()
	ctrl.GetNodeLogData(rr, req)
	assert.Equal(t, 200, rr.Code)
	var out []models.NodeSensorLog
	_ = json.Unmarshal(rr.Body.Bytes(), &out)
	assert.Len(t, out, 1)
	assert.Len(t, out[0].TemperatureEntries, 1)
}

func TestNodeController_TriggerRestart_UpdateIp_EnterUpdateMode_InPkg(t *testing.T) {
	// server to accept requests
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	}))
	defer srv.Close()

	addr := srv.Listener.Addr().String()
	repo := &fakeNodeRepo{serverAddr: addr}
	ctrl := NewNodeControllerWithDeps(repo, nil)

	// TriggerRestart
	req := httptest.NewRequest("GET", "/node/restart", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	rr := httptest.NewRecorder()
	ctrl.TriggerRestart(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	// UpdateNodeIp
	node := models.Node{Id: 1, IpAddress: "1.2.3.4"}
	b, _ := json.Marshal(node)
	req2 := httptest.NewRequest("POST", "/node/ipUpdate", bytes.NewReader(b))
	rr2 := httptest.NewRecorder()
	// use fake repo that verifies id exists
	repo2 := &fakeNodeRepoInpkg{VerifyIdIsNewFunc: func(id int) (bool, error) { return false, nil }, UpdateFunc: func(n *models.Node) error { return nil }}
	ctrl2 := NewNodeControllerWithDeps(repo2, nil)
	ctrl2.UpdateNodeIp(rr2, req2)
	assert.Equal(t, http.StatusOK, rr2.Code)

	// EnterUpdateMode (requires PIN)
	Config = config.Configuration{Pin: "0000"}
	pin := models.PinRequest{PinCode: "0000"}
	pb, _ := json.Marshal(pin)
	req3 := httptest.NewRequest("POST", "/node/updateMode/1", bytes.NewReader(pb))
	req3 = mux.SetURLVars(req3, map[string]string{"id": "1"})
	rr3 := httptest.NewRecorder()
	ctrl.EnterUpdateMode(rr3, req3)
	assert.Equal(t, http.StatusOK, rr3.Code)
}

func (f *fakeNodeRepoInpkg) GetSensorLogData(nodeId int, start time.Time, end time.Time) ([]models.NodeSensorLog, error) {
	if f.GetSensorLogDataFunc != nil {
		return f.GetSensorLogDataFunc(nodeId, start, end)
	}
	return nil, nil
}

func (f *fakeNodeRepoInpkg) GetTempLogDataByLogId(logId int) ([]models.TempLogData, error) {
	if f.GetTempLogDataByLogIdFunc != nil {
		return f.GetTempLogDataByLogIdFunc(logId)
	}
	return nil, nil
}

func (f *fakeNodeRepoInpkg) GetMoistureLogDataByLogId(logId int) ([]models.MoistureLogData, error) {
	if f.GetMoistureLogDataByLogIdFunc != nil {
		return f.GetMoistureLogDataByLogIdFunc(logId)
	}
	return nil, nil
}

func (f *fakeNodeRepoInpkg) GetResistorLogDataByLogId(logId int) ([]models.ResistorLogData, error) {
	if f.GetResistorLogDataByLogIdFunc != nil {
		return f.GetResistorLogDataByLogIdFunc(logId)
	}
	return nil, nil
}

func (f *fakeNodeRepoInpkg) GetMagneticLogDataByLogId(logId int) ([]models.MagneticLogData, error) {
	if f.GetMagneticLogDataByLogIdFunc != nil {
		return f.GetMagneticLogDataByLogIdFunc(logId)
	}
	return nil, nil
}

func (f *fakeNodeRepoInpkg) CreateNewLog(item models.NodeData) error {
	if f.CreateNewLogFunc != nil {
		return f.CreateNewLogFunc(item)
	}
	return nil
}

func (f *fakeNodeRepoInpkg) CreateNodeSensor(sensor *models.NodeSensor) error {
	if f.CreateNodeSensorFunc != nil {
		return f.CreateNodeSensorFunc(sensor)
	}
	return nil
}

func (f *fakeNodeRepoInpkg) CreateNodeSwitch(sw *models.NodeSwitch) error {
	if f.CreateNodeSwitchFunc != nil {
		return f.CreateNodeSwitchFunc(sw)
	}
	return nil
}

func TestNodeController_GetAll_InPkg(t *testing.T) {
	fetchAll := func() ([]models.Node, error) {
		return []models.Node{{Id: 1, Name: "n1", Mac: "m1"}, {Id: 2, Name: "n2", Mac: "m2"}}, nil
	}
	repo := &fakeNodeRepoInpkg{FetchAllFunc: fetchAll}

	ctrl := NewNodeControllerWithDeps(repo, nil)

	req := httptest.NewRequest("GET", "/node", nil)
	rr := httptest.NewRecorder()

	ctrl.GetAll(rr, req)

	assert.Equal(t, 200, rr.Code)

	var got []models.Node
	err := json.Unmarshal(rr.Body.Bytes(), &got)
	assert.NoError(t, err)
	assert.Len(t, got, 2)
}

func TestNodeController_GetById_InPkg(t *testing.T) {
	fetchById := func(id int) (models.Node, error) {
		if id == 1 {
			return models.Node{Id: 1, Name: "n1", Mac: "m1"}, nil
		}
		return models.Node{}, errors.New("not found")
	}
	repo := &fakeNodeRepoInpkg{FetchByIdFunc: fetchById}

	ctrl := NewNodeControllerWithDeps(repo, nil)

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

func TestNodeController_Create_InPkg(t *testing.T) {
	createFn := func(n *models.Node) error {
		n.Id = 55
		return nil
	}
	repo := &fakeNodeRepoInpkg{CreateFunc: createFn}

	ctrl := NewNodeControllerWithDeps(repo, nil)

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
