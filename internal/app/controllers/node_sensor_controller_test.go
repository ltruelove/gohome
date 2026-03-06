package controllers

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ltruelove/gohome/internal/app/data/statements"
	"github.com/ltruelove/gohome/internal/app/models"
)

type mockSensorTypeData struct{}

func (m *mockSensorTypeData) DB() *sql.DB                    { return nil }
func (m *mockSensorTypeData) Stmt() statements.CrudStatement { return nil }
func (m *mockSensorTypeData) SelectAll() ([]models.Model, error) {
	types := []models.Model{models.SensorType{Id: 1, TypeName: "Temp"}}
	return types, nil
}
func (m *mockSensorTypeData) SelectByParentId(id int) ([]models.Model, error)       { return nil, nil }
func (m *mockSensorTypeData) SelectBySecondParentId(id int) ([]models.Model, error) { return nil, nil }
func (m *mockSensorTypeData) SelectById(id int) (models.Model, error)               { return nil, nil }
func (m *mockSensorTypeData) Insert(data models.Model) (models.Model, error)        { return nil, nil }
func (m *mockSensorTypeData) Update(data models.Model) error                        { return nil }
func (m *mockSensorTypeData) Delete(id int) error                                   { return nil }
func (m *mockSensorTypeData) DeleteAll() error                                      { return nil }
func (m *mockSensorTypeData) DeleteByParentId(id int) error                         { return nil }
func (m *mockSensorTypeData) DeleteBySecondParentId(id int) error                   { return nil }

type mockNodeSensorData struct{}

func (m *mockNodeSensorData) DB() *sql.DB                    { return nil }
func (m *mockNodeSensorData) Stmt() statements.CrudStatement { return nil }
func (m *mockNodeSensorData) SelectAll() ([]models.Model, error) {
	sensors := []models.Model{models.NodeSensor{Id: 1, SensorTypeId: 1, Name: "Sensor1"}}
	return sensors, nil
}
func (m *mockNodeSensorData) SelectByParentId(id int) ([]models.Model, error)       { return nil, nil }
func (m *mockNodeSensorData) SelectBySecondParentId(id int) ([]models.Model, error) { return nil, nil }
func (m *mockNodeSensorData) SelectById(id int) (models.Model, error)               { return nil, nil }
func (m *mockNodeSensorData) Insert(data models.Model) (models.Model, error)        { return nil, nil }
func (m *mockNodeSensorData) Update(data models.Model) error                        { return nil }
func (m *mockNodeSensorData) Delete(id int) error                                   { return nil }
func (m *mockNodeSensorData) DeleteAll() error                                      { return nil }
func (m *mockNodeSensorData) DeleteByParentId(id int) error                         { return nil }
func (m *mockNodeSensorData) DeleteBySecondParentId(id int) error                   { return nil }

func TestNodeSensorGetAll(t *testing.T) {
	sensorData := &mockNodeSensorData{}
	sensorTypes := &mockSensorTypeData{}

	controller := NewNodeSensorControllerWithDeps(sensorData, sensorTypes)

	req := httptest.NewRequest(http.MethodGet, "/sensor", nil)
	rr := httptest.NewRecorder()

	controller.GetAll(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200 got %d", rr.Code)
	}
}
