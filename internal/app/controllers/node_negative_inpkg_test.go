package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/dto"
	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/stretchr/testify/assert"
)

// Negative-path tests for Node controller

func TestNodeController_Toggle_InvalidPin_InPkg(t *testing.T) {
	// fake repo won't be called because PIN is invalid
	repo := &fakeNodeRepoInpkg{}
	ctrl := NewNodeControllerWithDeps(repo, nil)

	// configure server PIN
	Config = config.Configuration{Pin: "0000"}

	pin := models.PinRequest{PinCode: "1234"}
	pb, _ := json.Marshal(pin)
	req := httptest.NewRequest("POST", "/node/switch/toggle/10", bytes.NewReader(pb))
	req = mux.SetURLVars(req, map[string]string{"id": "10"})
	rr := httptest.NewRecorder()

	ctrl.ToggleNodeSwitch(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestNodeController_Toggle_FetchNodeSwitchError_InPkg(t *testing.T) {
	repo := &fakeNodeRepoInpkg{FetchNodeSwitchFunc: func(id int) (models.NodeSwitch, error) { return models.NodeSwitch{}, errors.New("not found") }}
	ctrl := NewNodeControllerWithDeps(repo, nil)

	Config = config.Configuration{Pin: "1234"}
	pin := models.PinRequest{PinCode: "1234"}
	pb, _ := json.Marshal(pin)
	req := httptest.NewRequest("POST", "/node/switch/toggle/10", bytes.NewReader(pb))
	req = mux.SetURLVars(req, map[string]string{"id": "10"})
	rr := httptest.NewRecorder()

	ctrl.ToggleNodeSwitch(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestNodeController_Toggle_ExternalFailure_InPkg(t *testing.T) {
	// create a server and close it to force http.Get connection error
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	addr := srv.Listener.Addr().String()
	srv.Close()

	// use fakeNodeRepo that points to the closed server address
	repo := &fakeNodeRepo{serverAddr: addr}
	ctrl := NewNodeControllerWithDeps(repo, nil)

	Config = config.Configuration{Pin: "1234"}
	pin := models.PinRequest{PinCode: "1234"}
	pb, _ := json.Marshal(pin)
	req := httptest.NewRequest("POST", "/node/switch/toggle/10", bytes.NewReader(pb))
	req = mux.SetURLVars(req, map[string]string{"id": "10"})
	rr := httptest.NewRecorder()

	ctrl.ToggleNodeSwitch(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestNodeController_GetNodeData_ExternalFailure_InPkg(t *testing.T) {
	// server closed to simulate connection error
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	addr := srv.Listener.Addr().String()
	srv.Close()

	repo := &fakeNodeRepo{serverAddr: addr}
	ctrl := NewNodeControllerWithDeps(repo, nil)

	req := httptest.NewRequest("GET", "/node/data/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	rr := httptest.NewRecorder()

	ctrl.GetNodeData(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

// fake control-point repo that returns error for FetchById
type fakeCPErr struct{}

func (f *fakeCPErr) FetchAll() ([]models.ControlPoint, error)          { return nil, nil }
func (f *fakeCPErr) FetchAllAvailable() ([]models.ControlPoint, error) { return nil, nil }
func (f *fakeCPErr) FetchAllNodes(controlPointId int) ([]dto.ControlPointNode, error) {
	return nil, nil
}
func (f *fakeCPErr) FetchById(id int) (models.ControlPoint, error) {
	return models.ControlPoint{}, errors.New("not found")
}
func (f *fakeCPErr) FetchByMac(mac string) (models.ControlPoint, error) {
	return models.ControlPoint{}, nil
}
func (f *fakeCPErr) Create(cp *models.ControlPoint) error                        { return nil }
func (f *fakeCPErr) VerifyIdIsNew(id int) (bool, error)                          { return false, nil }
func (f *fakeCPErr) UpdateIp(cp *models.ControlPoint) error                      { return nil }
func (f *fakeCPErr) Update(cp *models.ControlPoint) error                        { return nil }
func (f *fakeCPErr) Delete(id int) error                                         { return nil }
func (f *fakeCPErr) AddNodeToControlPoint(cpnode *models.ControlPointNode) error { return nil }
func (f *fakeCPErr) FetchControlPointByNode(nodeId int) (models.ControlPoint, error) {
	return models.ControlPoint{}, nil
}

func TestNodeController_Register_ControlPointNotFound_InPkg(t *testing.T) {
	nodeRepo := &fakeNodeRepoInpkg{CreateFunc: func(n *models.Node) error { n.Id = 77; return nil }, CreateNodeSensorFunc: func(s *models.NodeSensor) error { return nil }, CreateNodeSwitchFunc: func(sw *models.NodeSwitch) error { return nil }}
	cpRepo := &fakeCPErr{}

	ctrl := NewNodeControllerWithDeps(nodeRepo, cpRepo)

	payload := dto.RegsiterNode{Node: models.Node{Name: "nreg", Mac: "mm"}, ControlPoint: models.ControlPoint{Id: 1}, Sensors: []models.NodeSensor{}, Switches: []models.NodeSwitch{}}
	b, _ := json.Marshal(payload)
	req := httptest.NewRequest("POST", "/node/register", bytes.NewReader(b))
	rr := httptest.NewRecorder()
	ctrl.Register(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestNodeController_Delete_FetchIndividualError_InPkg(t *testing.T) {
	repo := &fakeNodeRepoInpkg{FetchIndividualFunc: func(id int) (models.Node, error) { return models.Node{}, errors.New("not found") }}
	ctrl := NewNodeControllerWithDeps(repo, nil)

	req := httptest.NewRequest("DELETE", "/node/1/delete", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	rr := httptest.NewRecorder()

	ctrl.Delete(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
