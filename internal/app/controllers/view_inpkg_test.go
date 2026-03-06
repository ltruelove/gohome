package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/ltruelove/gohome/internal/app/viewModels"
	"github.com/stretchr/testify/assert"
)

// fakeCrudRepo implements repository.CrudRepository methods used by ViewController
type fakeCrudRepo struct{}

func (f *fakeCrudRepo) SelectAll() ([]models.Model, error) {
	v := &models.View{Id: 1, Name: "v1"}
	return []models.Model{v}, nil
}
func (f *fakeCrudRepo) SelectByParentId(id int) ([]models.Model, error) { return nil, nil }
func (f *fakeCrudRepo) SelectById(id int) (models.Model, error) {
	return &models.View{Id: id, Name: "v1"}, nil
}
func (f *fakeCrudRepo) Insert(data models.Model) (models.Model, error) {
	if v, ok := data.(models.View); ok {
		vv := v
		vv.Id = 42
		return &vv, nil
	}
	if vp, ok := data.(*models.View); ok {
		vp.Id = 42
		return vp, nil
	}
	return &models.View{Id: 42, Name: "created"}, nil
}
func (f *fakeCrudRepo) Update(data models.Model) error      { return nil }
func (f *fakeCrudRepo) Delete(id int) error                 { return nil }
func (f *fakeCrudRepo) DeleteByParentId(id int) error       { return nil }
func (f *fakeCrudRepo) DeleteBySecondParentId(id int) error { return nil }

// fakeCompoundRepo implements repository.CompoundRepository used by ViewController
type fakeCompoundRepo struct{}

func (f *fakeCompoundRepo) FetchViewNodeSensorDataByViewId(viewId int) ([]viewModels.ViewNodeSensorVM, error) {
	return []viewModels.ViewNodeSensorVM{{Id: 7, Name: "s1"}}, nil
}
func (f *fakeCompoundRepo) FetchViewNodeSwitchDataByViewId(viewId int) ([]viewModels.ViewNodeSwitchVM, error) {
	return []viewModels.ViewNodeSwitchVM{{Id: 9, Name: "sw1"}}, nil
}

// minimal fake for node sensor/switch/view node repos used by Add/Remove flows
type fakeSimpleCrud struct{}

func (f *fakeSimpleCrud) SelectAll() ([]models.Model, error)              { return nil, nil }
func (f *fakeSimpleCrud) SelectByParentId(id int) ([]models.Model, error) { return nil, nil }
func (f *fakeSimpleCrud) SelectById(id int) (models.Model, error) {
	return &models.View{Id: id}, nil
}
func (f *fakeSimpleCrud) Insert(data models.Model) (models.Model, error) {
	switch v := data.(type) {
	case models.ViewNodeSensorData:
		vv := v
		return &vv, nil
	case *models.ViewNodeSensorData:
		return v, nil
	case models.ViewNodeSwitchData:
		vv := v
		return &vv, nil
	case *models.ViewNodeSwitchData:
		return v, nil
	default:
		return data, nil
	}
}
func (f *fakeSimpleCrud) Update(data models.Model) error      { return nil }
func (f *fakeSimpleCrud) Delete(id int) error                 { return nil }
func (f *fakeSimpleCrud) DeleteByParentId(id int) error       { return nil }
func (f *fakeSimpleCrud) DeleteBySecondParentId(id int) error { return nil }

func TestViewController_CreateUpdateDelete_AddRemoveNodes(t *testing.T) {
	viewRepo := &fakeCrudRepo{}
	nodeSensorRepo := &fakeSimpleCrud{}
	viewNodeSensorRepo := &fakeSimpleCrud{}
	nodeSwitchRepo := &fakeSimpleCrud{}
	viewNodeSwitchRepo := &fakeSimpleCrud{}
	compoundRepo := &fakeCompoundRepo{}

	ctrl := NewViewControllerWithDeps(viewRepo, nodeSensorRepo, viewNodeSensorRepo, nodeSwitchRepo, viewNodeSwitchRepo, compoundRepo)

	// Create
	body := models.View{Name: "newview"}
	b, _ := json.Marshal(body)
	req := httptest.NewRequest("POST", "/view", bytes.NewReader(b))
	rr := httptest.NewRecorder()
	ctrl.Create(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	var created models.View
	err := json.Unmarshal(rr.Body.Bytes(), &created)
	assert.NoError(t, err)
	assert.Equal(t, 42, created.Id)

	// Update
	body2 := models.View{Id: created.Id, Name: "updated"}
	b2, _ := json.Marshal(body2)
	req2 := httptest.NewRequest("PUT", "/view", bytes.NewReader(b2))
	rr2 := httptest.NewRecorder()
	ctrl.Update(rr2, req2)
	assert.Equal(t, http.StatusOK, rr2.Code)

	// Delete - controller expects id from mux vars; simulate via URL and vars
	req3 := httptest.NewRequest("DELETE", "/view/42", nil)
	req3 = mux.SetURLVars(req3, map[string]string{"id": strconv.Itoa(created.Id)})
	rr3 := httptest.NewRecorder()
	ctrl.Delete(rr3, req3)
	// controller may return 200 or 204 on success
	if rr3.Code != http.StatusNoContent {
		assert.Equal(t, http.StatusOK, rr3.Code)
	}

	// Add NodeSensor
	addSensor := models.ViewNodeSensorData{ViewId: created.Id, NodeSensorId: 7}
	b4, _ := json.Marshal(addSensor)
	req4 := httptest.NewRequest("POST", "/view/node/sensor", bytes.NewReader(b4))
	rr4 := httptest.NewRecorder()
	ctrl.AddNodeSensorToView(rr4, req4)
	assert.Equal(t, http.StatusOK, rr4.Code)

	// Remove NodeSensor (expects id in mux vars)
	req5 := httptest.NewRequest("DELETE", "/view/node/sensor/7", nil)
	req5 = mux.SetURLVars(req5, map[string]string{"id": "7"})
	rr5 := httptest.NewRecorder()
	ctrl.RemoveNodeSensorFromView(rr5, req5)
	if rr5.Code != http.StatusNoContent {
		assert.Equal(t, http.StatusOK, rr5.Code)
	}

	// Add NodeSwitch
	addSwitch := models.ViewNodeSwitchData{ViewId: created.Id, NodeSwitchId: 9}
	b6, _ := json.Marshal(addSwitch)
	req6 := httptest.NewRequest("POST", "/view/node/switch", bytes.NewReader(b6))
	rr6 := httptest.NewRecorder()
	ctrl.AddNodeSwitchToView(rr6, req6)
	assert.Equal(t, http.StatusOK, rr6.Code)

	// Remove NodeSwitch
	req7 := httptest.NewRequest("DELETE", "/view/node/switch/9", nil)
	req7 = mux.SetURLVars(req7, map[string]string{"id": "9"})
	rr7 := httptest.NewRecorder()
	ctrl.RemoveNodeSwitchFromView(rr7, req7)
	if rr7.Code != http.StatusNoContent {
		assert.Equal(t, http.StatusOK, rr7.Code)
	}
}

// fake repo + test for GetAll and GetById
type fakeCrudRepoInpkg struct {
	SelectAllFunc            func() ([]models.Model, error)
	SelectByIdFunc           func(int) (models.Model, error)
	InsertFunc               func(models.Model) (models.Model, error)
	UpdateFunc               func(models.Model) error
	DeleteFunc               func(int) error
	DeleteByParentIdFunc     func(int) error
	DeleteBySecondParentIdFn func(int) error
}

func (f *fakeCrudRepoInpkg) SelectAll() ([]models.Model, error) {
	if f.SelectAllFunc != nil {
		return f.SelectAllFunc()
	}
	return nil, nil
}
func (f *fakeCrudRepoInpkg) SelectByParentId(id int) ([]models.Model, error) { return nil, nil }
func (f *fakeCrudRepoInpkg) SelectById(id int) (models.Model, error) {
	if f.SelectByIdFunc != nil {
		return f.SelectByIdFunc(id)
	}
	return nil, nil
}
func (f *fakeCrudRepoInpkg) Insert(data models.Model) (models.Model, error) {
	if f.InsertFunc != nil {
		return f.InsertFunc(data)
	}
	return nil, nil
}
func (f *fakeCrudRepoInpkg) Update(data models.Model) error {
	if f.UpdateFunc != nil {
		return f.UpdateFunc(data)
	}
	return nil
}
func (f *fakeCrudRepoInpkg) Delete(id int) error {
	if f.DeleteFunc != nil {
		return f.DeleteFunc(id)
	}
	return nil
}
func (f *fakeCrudRepoInpkg) DeleteByParentId(id int) error {
	if f.DeleteByParentIdFunc != nil {
		return f.DeleteByParentIdFunc(id)
	}
	return nil
}
func (f *fakeCrudRepoInpkg) DeleteBySecondParentId(id int) error {
	if f.DeleteBySecondParentIdFn != nil {
		return f.DeleteBySecondParentIdFn(id)
	}
	return nil
}

// fake CompoundRepository
type fakeCompoundRepoInpkg struct {
	FetchViewNodeSensorDataByViewIdFunc func(int) ([]viewModels.ViewNodeSensorVM, error)
	FetchViewNodeSwitchDataByViewIdFunc func(int) ([]viewModels.ViewNodeSwitchVM, error)
}

func (f *fakeCompoundRepoInpkg) FetchViewNodeSensorDataByViewId(id int) ([]viewModels.ViewNodeSensorVM, error) {
	if f.FetchViewNodeSensorDataByViewIdFunc != nil {
		return f.FetchViewNodeSensorDataByViewIdFunc(id)
	}
	return nil, nil
}
func (f *fakeCompoundRepoInpkg) FetchViewNodeSwitchDataByViewId(id int) ([]viewModels.ViewNodeSwitchVM, error) {
	if f.FetchViewNodeSwitchDataByViewIdFunc != nil {
		return f.FetchViewNodeSwitchDataByViewIdFunc(id)
	}
	return nil, nil
}

func TestViewController_GetAllAndGetById_InPkg(t *testing.T) {
	view := models.View{Id: 5, Name: "Main"}
	viewModelSensors := []viewModels.ViewNodeSensorVM{{Id: 1, NodeId: 2, ViewId: 5, NodeSensorId: 3, Name: "S"}}
	viewModelSwitches := []viewModels.ViewNodeSwitchVM{{Id: 2, NodeId: 4, ViewId: 5, NodeSwitchId: 6, Name: "SW"}}

	viewRepo := &fakeCrudRepoInpkg{SelectAllFunc: func() ([]models.Model, error) { return []models.Model{view}, nil }, SelectByIdFunc: func(id int) (models.Model, error) { return &models.View{Id: 5, Name: "Main"}, nil }, InsertFunc: nil}
	compoundRepo := &fakeCompoundRepoInpkg{FetchViewNodeSensorDataByViewIdFunc: func(id int) ([]viewModels.ViewNodeSensorVM, error) { return viewModelSensors, nil }, FetchViewNodeSwitchDataByViewIdFunc: func(id int) ([]viewModels.ViewNodeSwitchVM, error) { return viewModelSwitches, nil }}

	ctrl := NewViewControllerWithDeps(viewRepo, nil, nil, nil, nil, compoundRepo)

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
