package controllers

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/ltruelove/gohome/internal/app/viewModels"
	"github.com/stretchr/testify/assert"
)

// fake view repo that returns error for SelectById
type fakeViewErrRepo struct{}

func (f *fakeViewErrRepo) SelectAll() ([]models.Model, error) { return nil, nil }
func (f *fakeViewErrRepo) SelectById(id int) (models.Model, error) {
	return nil, errors.New("not found")
}
func (f *fakeViewErrRepo) SelectByParentId(id int) ([]models.Model, error) { return nil, nil }
func (f *fakeViewErrRepo) Insert(data models.Model) (models.Model, error)  { return nil, nil }
func (f *fakeViewErrRepo) Update(data models.Model) error                  { return nil }
func (f *fakeViewErrRepo) Delete(id int) error                             { return nil }
func (f *fakeViewErrRepo) DeleteByParentId(id int) error                   { return nil }
func (f *fakeViewErrRepo) DeleteBySecondParentId(id int) error             { return nil }

type fakeCompoundEmpty struct{}

func (f *fakeCompoundEmpty) FetchViewNodeSensorDataByViewId(id int) ([]viewModels.ViewNodeSensorVM, error) {
	return nil, nil
}
func (f *fakeCompoundEmpty) FetchViewNodeSwitchDataByViewId(id int) ([]viewModels.ViewNodeSwitchVM, error) {
	return nil, nil
}

func TestViewController_GetById_NotFound_InPkg(t *testing.T) {
	viewRepo := &fakeViewErrRepo{}
	compoundRepo := &fakeCompoundEmpty{}
	ctrl := NewViewControllerWithDeps(viewRepo, nil, nil, nil, nil, compoundRepo)

	req := httptest.NewRequest("GET", "/view/99", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "99"})
	rr := httptest.NewRecorder()
	ctrl.GetById(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestViewController_AddNodeSensor_BadJSON_InPkg(t *testing.T) {
	viewRepo := &fakeCrudRepoInpkg{}
	ctrl := NewViewControllerWithDeps(viewRepo, nil, nil, nil, nil, nil)

	req := httptest.NewRequest("POST", "/view/node/sensor", bytes.NewReader([]byte("{bad")))
	rr := httptest.NewRecorder()
	ctrl.AddNodeSensorToView(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestViewController_RemoveNodeSensor_BadId_InPkg(t *testing.T) {
	viewRepo := &fakeCrudRepoInpkg{}
	ctrl := NewViewControllerWithDeps(viewRepo, nil, nil, nil, nil, nil)

	req := httptest.NewRequest("DELETE", "/view/node/sensor/bad", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "bad"})
	rr := httptest.NewRecorder()
	ctrl.RemoveNodeSensorFromView(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
