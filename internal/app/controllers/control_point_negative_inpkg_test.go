package controllers

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/ltruelove/gohome/internal/app/dto"
	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/stretchr/testify/assert"
)

type fakeControlPointErrRepo struct{}

func (f *fakeControlPointErrRepo) FetchAll() ([]models.ControlPoint, error) {
	return nil, errors.New("fail")
}
func (f *fakeControlPointErrRepo) FetchAllAvailable() ([]models.ControlPoint, error) { return nil, nil }
func (f *fakeControlPointErrRepo) FetchAllNodes(controlPointId int) ([]dto.ControlPointNode, error) {
	return nil, nil
}
func (f *fakeControlPointErrRepo) FetchById(id int) (models.ControlPoint, error) {
	return models.ControlPoint{}, errors.New("not found")
}
func (f *fakeControlPointErrRepo) FetchByMac(mac string) (models.ControlPoint, error) {
	return models.ControlPoint{}, nil
}
func (f *fakeControlPointErrRepo) Create(cp *models.ControlPoint) error   { return nil }
func (f *fakeControlPointErrRepo) VerifyIdIsNew(id int) (bool, error)     { return true, nil }
func (f *fakeControlPointErrRepo) UpdateIp(cp *models.ControlPoint) error { return nil }
func (f *fakeControlPointErrRepo) Update(cp *models.ControlPoint) error   { return nil }
func (f *fakeControlPointErrRepo) Delete(id int) error                    { return nil }
func (f *fakeControlPointErrRepo) AddNodeToControlPoint(cpnode *models.ControlPointNode) error {
	return nil
}
func (f *fakeControlPointErrRepo) FetchControlPointByNode(nodeId int) (models.ControlPoint, error) {
	return models.ControlPoint{}, nil
}

func TestControlPoint_Create_BadJSON_InPkg(t *testing.T) {
	repo := &fakeControlPointErrRepo{}
	ctrl := NewControlPointControllerWithDeps(repo)

	// bad JSON
	req := httptest.NewRequest("POST", "/controlPoint", bytes.NewReader([]byte("{bad json")))
	rr := httptest.NewRecorder()
	ctrl.Create(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestControlPoint_Delete_NotFound_InPkg(t *testing.T) {
	repo := &fakeControlPointErrRepo{}
	// override VerifyIdIsNew to return true -> not found
	ctrl := NewControlPointControllerWithDeps(repo)

	req := httptest.NewRequest("DELETE", "/controlPoint/1/delete", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	rr := httptest.NewRecorder()
	ctrl.Delete(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Code)
}
