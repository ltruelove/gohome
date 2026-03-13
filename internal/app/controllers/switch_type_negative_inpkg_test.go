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

type fakeSwitchTypeErrRepo2 struct{}

func (f *fakeSwitchTypeErrRepo2) SelectAll() ([]models.Model, error) {
	return nil, errors.New("db fail")
}
func (f *fakeSwitchTypeErrRepo2) SelectById(id int) (models.Model, error) {
	return nil, errors.New("fail")
}
func (f *fakeSwitchTypeErrRepo2) SelectByParentId(id int) ([]models.Model, error) { return nil, nil }
func (f *fakeSwitchTypeErrRepo2) Insert(data models.Model) (models.Model, error)  { return nil, nil }
func (f *fakeSwitchTypeErrRepo2) Update(data models.Model) error                  { return nil }
func (f *fakeSwitchTypeErrRepo2) Delete(id int) error                             { return nil }
func (f *fakeSwitchTypeErrRepo2) DeleteByParentId(id int) error                   { return nil }
func (f *fakeSwitchTypeErrRepo2) DeleteBySecondParentId(id int) error             { return nil }

type fakeSwitchTypeNoRowsRepo struct{}

func (f *fakeSwitchTypeNoRowsRepo) SelectAll() ([]models.Model, error) { return nil, nil }
func (f *fakeSwitchTypeNoRowsRepo) SelectById(id int) (models.Model, error) {
	return nil, sql.ErrNoRows
}
func (f *fakeSwitchTypeNoRowsRepo) SelectByParentId(id int) ([]models.Model, error) { return nil, nil }
func (f *fakeSwitchTypeNoRowsRepo) Insert(data models.Model) (models.Model, error)  { return nil, nil }
func (f *fakeSwitchTypeNoRowsRepo) Update(data models.Model) error                  { return nil }
func (f *fakeSwitchTypeNoRowsRepo) Delete(id int) error                             { return nil }
func (f *fakeSwitchTypeNoRowsRepo) DeleteByParentId(id int) error                   { return nil }
func (f *fakeSwitchTypeNoRowsRepo) DeleteBySecondParentId(id int) error             { return nil }

func TestSwitchType_GetAll_Error_InPkg(t *testing.T) {
	repo := &fakeSwitchTypeErrRepo2{}
	ctrl := NewSwitchTypeControllerWithDeps(repo)

	rw := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/switchType", nil)
	ctrl.GetAll(rw, req)
	res := rw.Result()
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
}

func TestSwitchType_GetById_NotFoundAndError_InPkg(t *testing.T) {
	repoNoRows := &fakeSwitchTypeNoRowsRepo{}
	ctrl1 := NewSwitchTypeControllerWithDeps(repoNoRows)
	rw := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/switchType/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	ctrl1.GetById(rw, req)
	assert.Equal(t, http.StatusNotFound, rw.Result().StatusCode)

	repoErr := &fakeSwitchTypeErrRepo2{}
	ctrl2 := NewSwitchTypeControllerWithDeps(repoErr)
	rw2 := httptest.NewRecorder()
	req2 := httptest.NewRequest("GET", "/switchType/1", nil)
	req2 = mux.SetURLVars(req2, map[string]string{"id": "1"})
	ctrl2.GetById(rw2, req2)
	assert.Equal(t, http.StatusInternalServerError, rw2.Result().StatusCode)
}
