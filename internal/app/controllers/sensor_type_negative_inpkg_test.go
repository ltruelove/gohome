package controllers

import (
	"database/sql"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/stretchr/testify/assert"
)

type fakeSensorTypeErrRepo2 struct{}

func (f *fakeSensorTypeErrRepo2) SelectAll() ([]models.Model, error) {
	return nil, errors.New("db fail")
}
func (f *fakeSensorTypeErrRepo2) SelectById(id int) (models.Model, error) {
	return nil, errors.New("fail")
}
func (f *fakeSensorTypeErrRepo2) SelectByParentId(id int) ([]models.Model, error) {
	return nil, errors.New("fail")
}
func (f *fakeSensorTypeErrRepo2) Insert(data models.Model) (models.Model, error) { return nil, nil }
func (f *fakeSensorTypeErrRepo2) Update(data models.Model) error                 { return nil }
func (f *fakeSensorTypeErrRepo2) Delete(id int) error                            { return nil }
func (f *fakeSensorTypeErrRepo2) DeleteByParentId(id int) error                  { return nil }
func (f *fakeSensorTypeErrRepo2) DeleteBySecondParentId(id int) error            { return nil }

type fakeSensorTypeNoRowsRepo struct{}

func (f *fakeSensorTypeNoRowsRepo) SelectAll() ([]models.Model, error) { return nil, nil }
func (f *fakeSensorTypeNoRowsRepo) SelectById(id int) (models.Model, error) {
	return nil, sql.ErrNoRows
}
func (f *fakeSensorTypeNoRowsRepo) SelectByParentId(id int) ([]models.Model, error) {
	return nil, sql.ErrNoRows
}
func (f *fakeSensorTypeNoRowsRepo) Insert(data models.Model) (models.Model, error) { return nil, nil }
func (f *fakeSensorTypeNoRowsRepo) Update(data models.Model) error                 { return nil }
func (f *fakeSensorTypeNoRowsRepo) Delete(id int) error                            { return nil }
func (f *fakeSensorTypeNoRowsRepo) DeleteByParentId(id int) error                  { return nil }
func (f *fakeSensorTypeNoRowsRepo) DeleteBySecondParentId(id int) error            { return nil }

func TestSensorType_GetAll_Error_InPkg(t *testing.T) {
	repo := &fakeSensorTypeErrRepo2{}
	ctrl := NewSensorTypeControllerWithDeps(repo, nil)

	rw := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/sensorType", nil)
	ctrl.GetAll(rw, req)
	res := rw.Result()
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
}

func TestSensorType_GetById_NotFoundAndError_InPkg(t *testing.T) {
	// not found
	repoNoRows := &fakeSensorTypeNoRowsRepo{}
	ctrl1 := NewSensorTypeControllerWithDeps(repoNoRows, repoNoRows)
	rw := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/sensorType/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	ctrl1.GetById(rw, req)
	assert.Equal(t, http.StatusNotFound, rw.Result().StatusCode)

	// other error
	repoErr := &fakeSensorTypeErrRepo2{}
	ctrl2 := NewSensorTypeControllerWithDeps(repoErr, repoErr)
	rw2 := httptest.NewRecorder()
	req2 := httptest.NewRequest("GET", "/sensorType/1", nil)
	req2 = mux.SetURLVars(req2, map[string]string{"id": "1"})
	ctrl2.GetById(rw2, req2)
	assert.Equal(t, http.StatusInternalServerError, rw2.Result().StatusCode)
}

func TestSensorType_DataById_Errors_InPkg(t *testing.T) {
	// not found
	repoNoRows := &fakeSensorTypeNoRowsRepo{}
	ctrl1 := NewSensorTypeControllerWithDeps(nil, repoNoRows)
	rw := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/sensorType/data/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	ctrl1.DataById(rw, req)
	assert.Equal(t, http.StatusNotFound, rw.Result().StatusCode)

	// other error
	repoErr := &fakeSensorTypeErrRepo2{}
	ctrl2 := NewSensorTypeControllerWithDeps(nil, repoErr)
	rw2 := httptest.NewRecorder()
	req2 := httptest.NewRequest("GET", "/sensorType/data/1", nil)
	req2 = mux.SetURLVars(req2, map[string]string{"id": "1"})
	ctrl2.DataById(rw2, req2)
	assert.Equal(t, http.StatusInternalServerError, rw2.Result().StatusCode)
}
