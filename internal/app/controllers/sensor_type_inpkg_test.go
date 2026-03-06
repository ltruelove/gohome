package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/ltruelove/gohome/internal/app/models"
)

type fakeSensorTypeRepoInpkg2 struct{}

func (f *fakeSensorTypeRepoInpkg2) SelectAll() ([]models.Model, error) {
	v := models.SensorType{Id: 42, TypeName: "DHT22"}
	return []models.Model{v}, nil
}
func (f *fakeSensorTypeRepoInpkg2) SelectById(id int) (models.Model, error) {
	return models.SensorType{Id: id, TypeName: "Fake"}, nil
}
func (f *fakeSensorTypeRepoInpkg2) SelectByParentId(id int) ([]models.Model, error) {
	return []models.Model{}, nil
}
func (f *fakeSensorTypeRepoInpkg2) Insert(data models.Model) (models.Model, error) { return data, nil }
func (f *fakeSensorTypeRepoInpkg2) Update(data models.Model) error                 { return nil }
func (f *fakeSensorTypeRepoInpkg2) Delete(id int) error                            { return nil }
func (f *fakeSensorTypeRepoInpkg2) DeleteByParentId(id int) error                  { return nil }
func (f *fakeSensorTypeRepoInpkg2) DeleteBySecondParentId(id int) error            { return nil }

func TestSensorTypeController_GetEndpoints_InPkg(t *testing.T) {
	repo := &fakeSensorTypeRepoInpkg2{}
	ctrl := NewSensorTypeControllerWithDeps(repo, repo)

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
	req = mux.SetURLVars(req, map[string]string{"id": "42"})
	ctrl.GetById(rw, req)
	res = rw.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200 from GetById, got %d", res.StatusCode)
	}

	// DataById (may return 404 or 200)
	rw = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/sensorType/data/42", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "42"})
	ctrl.DataById(rw, req)
	res = rw.Result()
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNotFound {
		t.Fatalf("expected 200 or 404 from DataById, got %d", res.StatusCode)
	}
}
