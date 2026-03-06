package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	controllers "github.com/ltruelove/gohome/internal/app/controllers"
	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/ltruelove/gohome/internal/app/viewModels"
	"github.com/stretchr/testify/assert"
)

// simple fake CrudRepository for controller tests
type fakeCrudRepo struct {
	SelectAllFunc            func() ([]models.Model, error)
	SelectByIdFunc           func(int) (models.Model, error)
	InsertFunc               func(models.Model) (models.Model, error)
	UpdateFunc               func(models.Model) error
	DeleteFunc               func(int) error
	DeleteByParentIdFunc     func(int) error
	DeleteBySecondParentIdFn func(int) error
}

func (f *fakeCrudRepo) SelectAll() ([]models.Model, error) {
	if f.SelectAllFunc != nil {
		return f.SelectAllFunc()
	}
	return nil, nil
}
func (f *fakeCrudRepo) SelectByParentId(id int) ([]models.Model, error) { return nil, nil }
func (f *fakeCrudRepo) SelectById(id int) (models.Model, error) {
	if f.SelectByIdFunc != nil {
		return f.SelectByIdFunc(id)
	}
	return nil, nil
}
func (f *fakeCrudRepo) Insert(data models.Model) (models.Model, error) {
	if f.InsertFunc != nil {
		return f.InsertFunc(data)
	}
	return nil, nil
}
func (f *fakeCrudRepo) Update(data models.Model) error {
	if f.UpdateFunc != nil {
		return f.UpdateFunc(data)
	}
	return nil
}
func (f *fakeCrudRepo) Delete(id int) error {
	if f.DeleteFunc != nil {
		return f.DeleteFunc(id)
	}
	return nil
}
func (f *fakeCrudRepo) DeleteByParentId(id int) error {
	if f.DeleteByParentIdFunc != nil {
		return f.DeleteByParentIdFunc(id)
	}
	return nil
}
func (f *fakeCrudRepo) DeleteBySecondParentId(id int) error {
	if f.DeleteBySecondParentIdFn != nil {
		return f.DeleteBySecondParentIdFn(id)
	}
	return nil
}

// fake CompoundRepository
type fakeCompoundRepo struct {
	FetchViewNodeSensorDataByViewIdFunc func(int) ([]viewModels.ViewNodeSensorVM, error)
	FetchViewNodeSwitchDataByViewIdFunc func(int) ([]viewModels.ViewNodeSwitchVM, error)
}

func (f *fakeCompoundRepo) FetchViewNodeSensorDataByViewId(id int) ([]viewModels.ViewNodeSensorVM, error) {
	if f.FetchViewNodeSensorDataByViewIdFunc != nil {
		return f.FetchViewNodeSensorDataByViewIdFunc(id)
	}
	return nil, nil
}
func (f *fakeCompoundRepo) FetchViewNodeSwitchDataByViewId(id int) ([]viewModels.ViewNodeSwitchVM, error) {
	if f.FetchViewNodeSwitchDataByViewIdFunc != nil {
		return f.FetchViewNodeSwitchDataByViewIdFunc(id)
	}
	return nil, nil
}

func TestViewController_GetAllAndGetById(t *testing.T) {
	view := models.View{Id: 5, Name: "Main"}
	viewModelSensors := []viewModels.ViewNodeSensorVM{{Id: 1, NodeId: 2, ViewId: 5, NodeSensorId: 3, Name: "S"}}
	viewModelSwitches := []viewModels.ViewNodeSwitchVM{{Id: 2, NodeId: 4, ViewId: 5, NodeSwitchId: 6, Name: "SW"}}

	viewRepo := &fakeCrudRepo{SelectAllFunc: func() ([]models.Model, error) { return []models.Model{view}, nil }, SelectByIdFunc: func(id int) (models.Model, error) { return &models.View{Id: 5, Name: "Main"}, nil }, InsertFunc: nil}
	compoundRepo := &fakeCompoundRepo{FetchViewNodeSensorDataByViewIdFunc: func(id int) ([]viewModels.ViewNodeSensorVM, error) { return viewModelSensors, nil }, FetchViewNodeSwitchDataByViewIdFunc: func(id int) ([]viewModels.ViewNodeSwitchVM, error) { return viewModelSwitches, nil }}

	ctrl := controllers.NewViewControllerWithDeps(viewRepo, nil, nil, nil, nil, compoundRepo)

	// GetAll
	req := httptest.NewRequest("GET", "/view", nil)
	rr := httptest.NewRecorder()
	ctrl.GetAll(rr, req)
	assert.Equal(t, 200, rr.Code)
	var got []models.View
	_ = json.Unmarshal(rr.Body.Bytes(), &got)
	assert.Len(t, got, 1)

	// GetById
	req2 := httptest.NewRequest("GET", "/view/5", nil)
	req2 = mux.SetURLVars(req2, map[string]string{"id": "5"})
	rr2 := httptest.NewRecorder()
	ctrl.GetById(rr2, req2)
	assert.Equal(t, 200, rr2.Code)
	var vm viewModels.ViewVM
	_ = json.Unmarshal(rr2.Body.Bytes(), &vm)
	assert.Equal(t, 1, len(vm.Sensors))
	assert.Equal(t, 1, len(vm.Switches))
}

func TestViewController_Create_Update_Delete(t *testing.T) {
	created := &models.View{Id: 10, Name: "created"}
	viewRepo := &fakeCrudRepo{InsertFunc: func(data models.Model) (models.Model, error) { v := data.(*models.View); v.Id = 10; return v, nil }, SelectByIdFunc: func(id int) (models.Model, error) {
		if id == 10 {
			return created, nil
		}
		return nil, nil
	}, UpdateFunc: func(data models.Model) error { return nil }, DeleteFunc: func(id int) error { return nil }}
	ctrl := controllers.NewViewControllerWithDeps(viewRepo, nil, nil, nil, nil, nil)

	// Create
	payload := models.View{Name: "created"}
	b, _ := json.Marshal(payload)
	req := httptest.NewRequest("POST", "/view", bytes.NewReader(b))
	rr := httptest.NewRecorder()
	ctrl.Create(rr, req)
	assert.Equal(t, 200, rr.Code)
	var out models.View
	_ = json.Unmarshal(rr.Body.Bytes(), &out)
	assert.Equal(t, 10, out.Id)

	// Update
	updatePayload := models.View{Id: 10, Name: "updated"}
	ub, _ := json.Marshal(updatePayload)
	reqUp := httptest.NewRequest("PUT", "/view", bytes.NewReader(ub))
	rrUp := httptest.NewRecorder()
	ctrl.Update(rrUp, reqUp)
	assert.Equal(t, 200, rrUp.Code)

	// Delete
	reqDel := httptest.NewRequest("DELETE", "/view/10", nil)
	reqDel = mux.SetURLVars(reqDel, map[string]string{"id": "10"})
	rrDel := httptest.NewRecorder()
	ctrl.Delete(rrDel, reqDel)
	assert.Equal(t, 200, rrDel.Code)
}

func TestViewController_AddAndRemoveNodeSensor(t *testing.T) {
	// Add
	vns := &models.ViewNodeSensorData{Id: 99, NodeId: 7, ViewId: 4, NodeSensorId: 6, Name: "nsv"}
	viewNodeSensorRepo := &fakeCrudRepo{InsertFunc: func(data models.Model) (models.Model, error) {
		item := data.(*models.ViewNodeSensorData)
		item.Id = 99
		return item, nil
	}}
	nodeSensorRepo := &fakeCrudRepo{DeleteFunc: func(id int) error { return nil }}
	ctrl := controllers.NewViewControllerWithDeps(nil, nodeSensorRepo, viewNodeSensorRepo, nil, nil, nil)

	p, _ := json.Marshal(models.ViewNodeSensorData{NodeId: 7, ViewId: 4, NodeSensorId: 6, Name: "nsv"})
	req := httptest.NewRequest("POST", "/view/node/sensor", bytes.NewReader(p))
	rr := httptest.NewRecorder()
	ctrl.AddNodeSensorToView(rr, req)
	assert.Equal(t, 200, rr.Code)
	var out models.ViewNodeSensorData
	_ = json.Unmarshal(rr.Body.Bytes(), &out)
	assert.Equal(t, 99, out.Id)

	// Remove
	reqDel := httptest.NewRequest("DELETE", "/view/node/sensor/6", nil)
	reqDel = mux.SetURLVars(reqDel, map[string]string{"id": "6"})
	rrDel := httptest.NewRecorder()
	ctrl.RemoveNodeSensorFromView(rrDel, reqDel)
	assert.Equal(t, 200, rrDel.Code)
}

func TestViewController_AddAndRemoveNodeSwitch(t *testing.T) {
	// Add
	vnswRepo := &fakeCrudRepo{InsertFunc: func(data models.Model) (models.Model, error) {
		item := data.(*models.ViewNodeSwitchData)
		item.Id = 77
		return item, nil
	}}
	nodeSwitchRepo := &fakeCrudRepo{DeleteFunc: func(id int) error { return nil }}
	ctrl := controllers.NewViewControllerWithDeps(nil, nil, nil, nodeSwitchRepo, vnswRepo, nil)

	p, _ := json.Marshal(models.ViewNodeSwitchData{NodeId: 7, ViewId: 4, NodeSwitchId: 6, Name: "nsv"})
	req := httptest.NewRequest("POST", "/view/node/switch", bytes.NewReader(p))
	rr := httptest.NewRecorder()
	ctrl.AddNodeSwitchToView(rr, req)
	assert.Equal(t, 200, rr.Code)
	var out models.ViewNodeSwitchData
	_ = json.Unmarshal(rr.Body.Bytes(), &out)
	assert.Equal(t, 77, out.Id)

	// Remove
	reqDel := httptest.NewRequest("DELETE", "/view/node/switch/6", nil)
	reqDel = mux.SetURLVars(reqDel, map[string]string{"id": "6"})
	rrDel := httptest.NewRecorder()
	ctrl.RemoveNodeSwitchFromView(rrDel, reqDel)
	assert.Equal(t, 200, rrDel.Code)
}
