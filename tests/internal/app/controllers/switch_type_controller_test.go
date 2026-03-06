package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/ltruelove/gohome/internal/app/controllers"
	"github.com/ltruelove/gohome/internal/app/models"
)

// minimal fake CRUD repo for switch types
type fakeCrudRepoSwitch struct{}

func (f *fakeCrudRepoSwitch) SelectAll() ([]models.Model, error) {
	v := models.SwitchType{Id: 99, Name: "Relay"}
	return []models.Model{v}, nil
}
func (f *fakeCrudRepoSwitch) SelectById(id int) (models.Model, error) {
	return models.SwitchType{Id: id, Name: "FakeSwitch"}, nil
}
func (f *fakeCrudRepoSwitch) SelectByParentId(id int) ([]models.Model, error) {
	return []models.Model{}, nil
}
func (f *fakeCrudRepoSwitch) Insert(data models.Model) (models.Model, error) { return data, nil }
func (f *fakeCrudRepoSwitch) Update(data models.Model) error                 { return nil }
func (f *fakeCrudRepoSwitch) Delete(id int) error                            { return nil }
func (f *fakeCrudRepoSwitch) DeleteByParentId(id int) error                  { return nil }
func (f *fakeCrudRepoSwitch) DeleteBySecondParentId(id int) error            { return nil }

func TestSwitchTypeController_GetEndpoints(t *testing.T) {
	repo := &fakeCrudRepoSwitch{}
	ctrl := controllers.NewSwitchTypeControllerWithDeps(repo)

	// GetAll
	rw := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/switchType", nil)
	ctrl.GetAll(rw, req)
	res := rw.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200 from GetAll, got %d", res.StatusCode)
	}

	// GetById
	rw = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/switchType/99", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "99"})
	ctrl.GetById(rw, req)
	res = rw.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200 from GetById, got %d", res.StatusCode)
	}
}
