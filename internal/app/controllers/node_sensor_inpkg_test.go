package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/ltruelove/gohome/internal/app/viewModels"
)

// fake repos implementing CrudRepository used by controller in-package test
type fakeNodeSensorRepoInpkg struct{}

func (f *fakeNodeSensorRepoInpkg) SelectAll() ([]models.Model, error) {
	ns := models.NodeSensor{Id: 1, NodeId: 2, SensorTypeId: 11, Name: "Sensor1", Pin: 4, DHTType: 11}
	return []models.Model{ns}, nil
}
func (f *fakeNodeSensorRepoInpkg) SelectById(id int) (models.Model, error) {
	return models.NodeSensor{Id: id, NodeId: 2, SensorTypeId: 11, Name: "Sensor1", Pin: 4, DHTType: 11}, nil
}
func (f *fakeNodeSensorRepoInpkg) SelectByParentId(id int) ([]models.Model, error) {
	return []models.Model{}, nil
}
func (f *fakeNodeSensorRepoInpkg) Insert(data models.Model) (models.Model, error) { return data, nil }
func (f *fakeNodeSensorRepoInpkg) Update(data models.Model) error                 { return nil }
func (f *fakeNodeSensorRepoInpkg) Delete(id int) error                            { return nil }
func (f *fakeNodeSensorRepoInpkg) DeleteByParentId(id int) error                  { return nil }
func (f *fakeNodeSensorRepoInpkg) DeleteBySecondParentId(id int) error            { return nil }

type fakeSensorTypeRepoInpkg struct{}

func (f *fakeSensorTypeRepoInpkg) SelectAll() ([]models.Model, error) {
	t := models.SensorType{Id: 11, TypeName: "DHT11"}
	return []models.Model{t}, nil
}
func (f *fakeSensorTypeRepoInpkg) SelectById(id int) (models.Model, error) {
	return models.SensorType{Id: id, TypeName: "Type"}, nil
}
func (f *fakeSensorTypeRepoInpkg) SelectByParentId(id int) ([]models.Model, error) {
	return []models.Model{}, nil
}
func (f *fakeSensorTypeRepoInpkg) Insert(data models.Model) (models.Model, error) { return data, nil }
func (f *fakeSensorTypeRepoInpkg) Update(data models.Model) error                 { return nil }
func (f *fakeSensorTypeRepoInpkg) Delete(id int) error                            { return nil }
func (f *fakeSensorTypeRepoInpkg) DeleteByParentId(id int) error                  { return nil }
func (f *fakeSensorTypeRepoInpkg) DeleteBySecondParentId(id int) error            { return nil }

func TestNodeSensorController_GetAll_InPkg(t *testing.T) {
	nsRepo := &fakeNodeSensorRepoInpkg{}
	tsRepo := &fakeSensorTypeRepoInpkg{}
	ctrl := NewNodeSensorControllerWithDeps(nsRepo, tsRepo)

	rw := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/sensor", nil)
	ctrl.GetAll(rw, req)
	res := rw.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200 from GetAll, got %d", res.StatusCode)
	}

	var list []viewModels.NodeSensorVM
	_ = json.NewDecoder(res.Body).Decode(&list)
	if len(list) != 1 {
		t.Fatalf("expected 1 node sensor, got %d", len(list))
	}
}
