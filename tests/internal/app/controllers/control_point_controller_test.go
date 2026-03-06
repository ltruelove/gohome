package controllers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/ltruelove/gohome/internal/app/controllers"
	"github.com/ltruelove/gohome/internal/app/dto"
	"github.com/ltruelove/gohome/internal/app/models"
)

type fakeControlPointRepo struct{}

func (f *fakeControlPointRepo) FetchAll() ([]models.ControlPoint, error) {
	return []models.ControlPoint{{Id: 1, Mac: "aa:bb:cc"}}, nil
}
func (f *fakeControlPointRepo) FetchAllAvailable() ([]models.ControlPoint, error) {
	return []models.ControlPoint{{Id: 2, Mac: "cc:dd:ee"}}, nil
}
func (f *fakeControlPointRepo) FetchAllNodes(controlPointId int) ([]dto.ControlPointNode, error) {
	return []dto.ControlPointNode{}, nil
}
func (f *fakeControlPointRepo) FetchById(id int) (models.ControlPoint, error) {
	return models.ControlPoint{Id: id, Mac: "aa:bb"}, nil
}
func (f *fakeControlPointRepo) FetchByMac(mac string) (models.ControlPoint, error) {
	return models.ControlPoint{}, fmt.Errorf("not found")
}
func (f *fakeControlPointRepo) Create(cp *models.ControlPoint) error   { return nil }
func (f *fakeControlPointRepo) VerifyIdIsNew(id int) (bool, error)     { return false, nil }
func (f *fakeControlPointRepo) UpdateIp(cp *models.ControlPoint) error { return nil }
func (f *fakeControlPointRepo) Update(cp *models.ControlPoint) error   { return nil }
func (f *fakeControlPointRepo) Delete(id int) error                    { return nil }
func (f *fakeControlPointRepo) AddNodeToControlPoint(cpnode *models.ControlPointNode) error {
	return nil
}
func (f *fakeControlPointRepo) FetchControlPointByNode(nodeId int) (models.ControlPoint, error) {
	return models.ControlPoint{}, nil
}

func TestControlPointController_GetAndCreate(t *testing.T) {
	repo := &fakeControlPointRepo{}
	ctrl := controllers.NewControlPointControllerWithDeps(repo)

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

	// Create (happy path)
	rw = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodPost, "/controlPoint", strings.NewReader("{\"Name\":\"CP1\", \"IpAddress\":\"192.168.1.5\", \"Mac\":\"aa:bb:cc\"}"))
	ctrl.Create(rw, req)
	res = rw.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 from Create, got %d", res.StatusCode)
	}
}
