package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/ltruelove/gohome/internal/app/viewModels"
)

type fakeNodeSwitchRepoInpkg struct{}

func (f *fakeNodeSwitchRepoInpkg) SelectAll() ([]models.Model, error) {
	ns := models.NodeSwitch{Id: 7, NodeId: 2, SwitchTypeId: 99, Name: "Sw1", Pin: 5, MomentaryPressDuration: 0, IsClosedOn: true}
	return []models.Model{ns}, nil
}
func (f *fakeNodeSwitchRepoInpkg) SelectById(id int) (models.Model, error) {
	return models.NodeSwitch{Id: id, NodeId: 2, SwitchTypeId: 99, Name: "Sw1", Pin: 5, MomentaryPressDuration: 0, IsClosedOn: true}, nil
}
func (f *fakeNodeSwitchRepoInpkg) SelectByParentId(id int) ([]models.Model, error) {
	return []models.Model{}, nil
}
func (f *fakeNodeSwitchRepoInpkg) Insert(data models.Model) (models.Model, error) { return data, nil }
func (f *fakeNodeSwitchRepoInpkg) Update(data models.Model) error                 { return nil }
func (f *fakeNodeSwitchRepoInpkg) Delete(id int) error                            { return nil }
func (f *fakeNodeSwitchRepoInpkg) DeleteByParentId(id int) error                  { return nil }
func (f *fakeNodeSwitchRepoInpkg) DeleteBySecondParentId(id int) error            { return nil }

type fakeSwitchTypeRepoInpkg struct{}

func (f *fakeSwitchTypeRepoInpkg) SelectAll() ([]models.Model, error) {
	st := models.SwitchType{Id: 99, Name: "Relay"}
	return []models.Model{st}, nil
}
func (f *fakeSwitchTypeRepoInpkg) SelectById(id int) (models.Model, error) {
	return models.SwitchType{Id: id, Name: "Relay"}, nil
}
func (f *fakeSwitchTypeRepoInpkg) SelectByParentId(id int) ([]models.Model, error) {
	return []models.Model{}, nil
}
func (f *fakeSwitchTypeRepoInpkg) Insert(data models.Model) (models.Model, error) { return data, nil }
func (f *fakeSwitchTypeRepoInpkg) Update(data models.Model) error                 { return nil }
func (f *fakeSwitchTypeRepoInpkg) Delete(id int) error                            { return nil }
func (f *fakeSwitchTypeRepoInpkg) DeleteByParentId(id int) error                  { return nil }
func (f *fakeSwitchTypeRepoInpkg) DeleteBySecondParentId(id int) error            { return nil }

func TestNodeSwitchController_GetAll_InPkg(t *testing.T) {
	nsRepo := &fakeNodeSwitchRepoInpkg{}
	tsRepo := &fakeSwitchTypeRepoInpkg{}
	ctrl := NewNodeSwitchControllerWithDeps(nsRepo, tsRepo)

	rw := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/switch", nil)
	ctrl.GetAll(rw, req)
	res := rw.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200 from GetAll, got %d", res.StatusCode)
	}

	var list []viewModels.NodeSwitchVM
	_ = json.NewDecoder(res.Body).Decode(&list)
	if len(list) != 1 {
		t.Fatalf("expected 1 node switch, got %d", len(list))
	}
}
