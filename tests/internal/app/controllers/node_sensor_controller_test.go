package controllers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ltruelove/gohome/internal/app/controllers"
	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/ltruelove/gohome/internal/app/viewModels"
)

type fakeNodeSensorRepo struct{}

func (f *fakeNodeSensorRepo) SelectAll() ([]models.Model, error) {
	ns := models.NodeSensor{Id: 1, NodeId: 2, SensorTypeId: 11, Name: "Sensor1", Pin: 4, DHTType: 11}
	return []models.Model{ns}, nil
}
func (f *fakeNodeSensorRepo) SelectById(id int) (models.Model, error) {
	return models.NodeSensor{Id: id, NodeId: 2, SensorTypeId: 11, Name: "Sensor1", Pin: 4, DHTType: 11}, nil
}
func (f *fakeNodeSensorRepo) SelectByParentId(id int) ([]models.Model, error) {
	return []models.Model{}, nil
}
func (f *fakeNodeSensorRepo) Insert(data models.Model) (models.Model, error) { return data, nil }
func (f *fakeNodeSensorRepo) Update(data models.Model) error                 { return nil }
func (f *fakeNodeSensorRepo) Delete(id int) error                            { return nil }
func (f *fakeNodeSensorRepo) DeleteByParentId(id int) error                  { return nil }
func (f *fakeNodeSensorRepo) DeleteBySecondParentId(id int) error            { return nil }

type fakeSensorTypeRepo struct{}

func (f *fakeSensorTypeRepo) SelectAll() ([]models.Model, error) {
	t := models.SensorType{Id: 11, TypeName: "DHT11"}
	return []models.Model{t}, nil
}
func (f *fakeSensorTypeRepo) SelectById(id int) (models.Model, error) {
	return models.SensorType{Id: id, TypeName: "Type"}, nil
}
func (f *fakeSensorTypeRepo) SelectByParentId(id int) ([]models.Model, error) {
	return []models.Model{}, nil
}
func (f *fakeSensorTypeRepo) Insert(data models.Model) (models.Model, error) { return data, nil }
func (f *fakeSensorTypeRepo) Update(data models.Model) error                 { return nil }
func (f *fakeSensorTypeRepo) Delete(id int) error                            { return nil }
func (f *fakeSensorTypeRepo) DeleteByParentId(id int) error                  { return nil }
func (f *fakeSensorTypeRepo) DeleteBySecondParentId(id int) error            { return nil }

func TestNodeSensorController_GetAll(t *testing.T) {
	nsRepo := &fakeNodeSensorRepo{}
	tsRepo := &fakeSensorTypeRepo{}
	ctrl := controllers.NewNodeSensorControllerWithDeps(nsRepo, tsRepo)

	rw := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/sensor", nil)
	ctrl.GetAll(rw, req)
	res := rw.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200 from GetAll, got %d", res.StatusCode)
	}

	// basic decode to ensure payload shape
	var list []viewModels.NodeSensorVM
	_ = json.NewDecoder(res.Body).Decode(&list)
	if len(list) != 1 {
		t.Fatalf("expected 1 node sensor, got %d", len(list))
	}
}
