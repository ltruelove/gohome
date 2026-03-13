package controllers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/stretchr/testify/assert"
)

type fakeNodeSensorErrRepo struct{}

func (f *fakeNodeSensorErrRepo) SelectAll() ([]models.Model, error) {
	return nil, errors.New("db fail")
}
func (f *fakeNodeSensorErrRepo) SelectById(id int) (models.Model, error)         { return nil, nil }
func (f *fakeNodeSensorErrRepo) SelectByParentId(id int) ([]models.Model, error) { return nil, nil }
func (f *fakeNodeSensorErrRepo) Insert(data models.Model) (models.Model, error)  { return nil, nil }
func (f *fakeNodeSensorErrRepo) Update(data models.Model) error                  { return nil }
func (f *fakeNodeSensorErrRepo) Delete(id int) error                             { return nil }
func (f *fakeNodeSensorErrRepo) DeleteByParentId(id int) error                   { return nil }
func (f *fakeNodeSensorErrRepo) DeleteBySecondParentId(id int) error             { return nil }

func TestNodeSensor_GetAll_Error_InPkg(t *testing.T) {
	nsRepo := &fakeNodeSensorErrRepo{}
	tsRepo := &fakeSensorTypeRepoInpkg{}
	ctrl := NewNodeSensorControllerWithDeps(nsRepo, tsRepo)

	rw := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/sensor", nil)
	ctrl.GetAll(rw, req)
	res := rw.Result()
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
}
