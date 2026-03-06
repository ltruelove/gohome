package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/ltruelove/gohome/internal/app/dto"
	"github.com/ltruelove/gohome/internal/app/models"
)

// fake ControlPointRepository for in-package tests
type fakeControlPointRepoInpkg struct{}

func (f *fakeControlPointRepoInpkg) FetchAll() ([]models.ControlPoint, error) {
	return []models.ControlPoint{{Id: 1, Name: "CP1", IpAddress: "192.168.1.5", Mac: "aa:bb:cc"}}, nil
}
func (f *fakeControlPointRepoInpkg) FetchAllAvailable() ([]models.ControlPoint, error) {
	return []models.ControlPoint{{Id: 2, Name: "CP2", IpAddress: "192.168.1.6", Mac: "cc:dd:ee"}}, nil
}
func (f *fakeControlPointRepoInpkg) FetchAllNodes(controlPointId int) ([]dto.ControlPointNode, error) {
	return []dto.ControlPointNode{}, nil
}
func (f *fakeControlPointRepoInpkg) FetchById(id int) (models.ControlPoint, error) {
	return models.ControlPoint{Id: id, Name: "CP1", IpAddress: "192.168.1.5", Mac: "aa:bb:cc"}, nil
}
func (f *fakeControlPointRepoInpkg) FetchByMac(mac string) (models.ControlPoint, error) {
	return models.ControlPoint{}, fmt.Errorf("not found")
}
func (f *fakeControlPointRepoInpkg) Create(cp *models.ControlPoint) error   { return nil }
func (f *fakeControlPointRepoInpkg) VerifyIdIsNew(id int) (bool, error)     { return false, nil }
func (f *fakeControlPointRepoInpkg) UpdateIp(cp *models.ControlPoint) error { return nil }
func (f *fakeControlPointRepoInpkg) Update(cp *models.ControlPoint) error   { return nil }
func (f *fakeControlPointRepoInpkg) Delete(id int) error                    { return nil }
func (f *fakeControlPointRepoInpkg) AddNodeToControlPoint(cpnode *models.ControlPointNode) error {
	return nil
}
func (f *fakeControlPointRepoInpkg) FetchControlPointByNode(nodeId int) (models.ControlPoint, error) {
	return models.ControlPoint{}, nil
}

func TestControlPointController_GetAllAndCreate_InPkg(t *testing.T) {
	repo := &fakeControlPointRepoInpkg{}
	ctrl := NewControlPointControllerWithDeps(repo)

	// GetAll
	rw := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/controlPoint", nil)
	ctrl.GetAll(rw, req)
	res := rw.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 from GetAll, got %d", res.StatusCode)
	}
	var list []models.ControlPoint
	_ = json.NewDecoder(res.Body).Decode(&list)
	if len(list) != 1 {
		t.Fatalf("expected 1 controlPoint, got %d", len(list))
	}

	// GetById
	rw = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/controlPoint/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	ctrl.GetById(rw, req)
	res = rw.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 from GetById, got %d", res.StatusCode)
	}

	// Create
	rw = httptest.NewRecorder()
	payload := strings.NewReader("{\"Name\":\"CPX\",\"IpAddress\":\"192.168.1.7\",\"Mac\":\"ff:ee:dd\"}")
	req = httptest.NewRequest(http.MethodPost, "/controlPoint", payload)
	ctrl.Create(rw, req)
	res = rw.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 from Create, got %d", res.StatusCode)
	}
}
