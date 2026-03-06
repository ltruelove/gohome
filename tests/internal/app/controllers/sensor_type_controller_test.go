package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/ltruelove/gohome/internal/app/controllers"
	"github.com/ltruelove/gohome/internal/app/models"
)

// minimal fake CRUD repo implementing required methods used by controller
type fakeCrudRepoSensor struct{}

func (f *fakeCrudRepoSensor) SelectAll() ([]models.Model, error) {
	v := models.SensorType{Id: 42, TypeName: "DHT22"}
	return []models.Model{v}, nil
}
func (f *fakeCrudRepoSensor) SelectById(id int) (models.Model, error) {
	return models.SensorType{Id: id, TypeName: "Fake"}, nil
}
func (f *fakeCrudRepoSensor) SelectByParentId(id int) ([]models.Model, error) {
	// not used for top-level sensor types
	return []models.Model{}, nil
}
func (f *fakeCrudRepoSensor) Insert(data models.Model) (models.Model, error) { return data, nil }
func (f *fakeCrudRepoSensor) Update(data models.Model) error                 { return nil }
func (f *fakeCrudRepoSensor) Delete(id int) error                            { return nil }
func (f *fakeCrudRepoSensor) DeleteByParentId(id int) error                  { return nil }
func (f *fakeCrudRepoSensor) DeleteBySecondParentId(id int) error            { return nil }

// (Create endpoint does not exist on SensorTypeController; test Get endpoints below)

func TestSensorTypeController_GetEndpoints(t *testing.T) {
	repo := &fakeCrudRepoSensor{}
	// controller expects two repos: sensorType and sensorTypeData
	ctrl := controllers.NewSensorTypeControllerWithDeps(repo, repo)

	// GetAll
	rw := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/sensorType", nil)
	ctrl.GetAll(rw, req)
	res := rw.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200 from GetAll, got %d", res.StatusCode)
	}

	// GetById
	rw = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/sensorType/42", nil)
	// set mux vars on request
	req = mux.SetURLVars(req, map[string]string{"id": "42"})
	ctrl.GetById(rw, req)
	res = rw.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200 from GetById, got %d", res.StatusCode)
	}

	// DataById
	rw = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/sensorType/data/42", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "42"})
	ctrl.DataById(rw, req)
	res = rw.Result()
	if res.StatusCode != http.StatusNotFound && res.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 or 404 from DataById, got %d", res.StatusCode)
	}
}

// no helper needed; using mux.SetURLVars directly
