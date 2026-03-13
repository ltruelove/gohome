package controllers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/stretchr/testify/assert"
)

type fakeNodeSwitchErrRepo struct{}

func (f *fakeNodeSwitchErrRepo) SelectAll() ([]models.Model, error) {
	return nil, errors.New("db fail")
}
func (f *fakeNodeSwitchErrRepo) SelectById(id int) (models.Model, error)         { return nil, nil }
func (f *fakeNodeSwitchErrRepo) SelectByParentId(id int) ([]models.Model, error) { return nil, nil }
func (f *fakeNodeSwitchErrRepo) Insert(data models.Model) (models.Model, error)  { return nil, nil }
func (f *fakeNodeSwitchErrRepo) Update(data models.Model) error                  { return nil }
func (f *fakeNodeSwitchErrRepo) Delete(id int) error                             { return nil }
func (f *fakeNodeSwitchErrRepo) DeleteByParentId(id int) error                   { return nil }
func (f *fakeNodeSwitchErrRepo) DeleteBySecondParentId(id int) error             { return nil }

func TestNodeSwitch_GetAll_Error_InPkg(t *testing.T) {
	nsRepo := &fakeNodeSwitchErrRepo{}
	tsRepo := &fakeSwitchTypeRepoInpkg{}
	ctrl := NewNodeSwitchControllerWithDeps(nsRepo, tsRepo)

	rw := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/switch", nil)
	ctrl.GetAll(rw, req)
	res := rw.Result()
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
}
