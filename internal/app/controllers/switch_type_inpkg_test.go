package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/ltruelove/gohome/internal/app/models"
)

type fakeSwitchTypeRepoInpkg2 struct{}

func (f *fakeSwitchTypeRepoInpkg2) SelectAll() ([]models.Model, error) {
	v := models.SwitchType{Id: 99, Name: "Relay"}
	return []models.Model{v}, nil
}
func (f *fakeSwitchTypeRepoInpkg2) SelectById(id int) (models.Model, error) {
	return models.SwitchType{Id: id, Name: "FakeSwitch"}, nil
}
func (f *fakeSwitchTypeRepoInpkg2) SelectByParentId(id int) ([]models.Model, error) {
	return []models.Model{}, nil
}
func (f *fakeSwitchTypeRepoInpkg2) Insert(data models.Model) (models.Model, error) { return data, nil }
func (f *fakeSwitchTypeRepoInpkg2) Update(data models.Model) error                 { return nil }
func (f *fakeSwitchTypeRepoInpkg2) Delete(id int) error                            { return nil }
func (f *fakeSwitchTypeRepoInpkg2) DeleteByParentId(id int) error                  { return nil }
func (f *fakeSwitchTypeRepoInpkg2) DeleteBySecondParentId(id int) error            { return nil }

func TestSwitchTypeController_GetEndpoints_InPkg(t *testing.T) {
	repo := &fakeSwitchTypeRepoInpkg2{}
	ctrl := NewSwitchTypeControllerWithDeps(repo)

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
