package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ltruelove/gohome/internal/app/data"
	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/ltruelove/gohome/internal/app/setup"
	"github.com/ltruelove/gohome/internal/pkg/routing"
)

type SensorTypeController struct {
	DB       *sql.DB
	AllTypes []models.SensorType
}

func (controller *SensorTypeController) RegisterSensorTypeEndpoints() {
	routing.AddRouteWithMethod("/sensorType", "GET", controller.AllSensorTypes)
	routing.AddRouteWithMethod("/sensorType/{id}", "GET", controller.SensorTypeById)
	routing.AddRouteWithMethod("/sensorType/data/{id}", "GET", controller.SensorTypeDataById)
}

func (controller *SensorTypeController) AllSensorTypes(writer http.ResponseWriter, request *http.Request) {
	if len(controller.AllTypes) == 0 {
		controller.AllTypes = data.FetchAllSensorTypes(controller.DB)
	}

	result, err := json.Marshal(controller.AllTypes)

	setup.CheckErr(err)

	writeResponse(writer, result)
}

func (controller *SensorTypeController) SensorTypeById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

	setup.CheckErr(err)

	item := data.FetchSensorType(id, controller.DB)

	result, err := json.Marshal(item)

	setup.CheckErr(err)

	writeResponse(writer, result)
}

func (controller *SensorTypeController) SensorTypeDataById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

	setup.CheckErr(err)

	item := data.FetchSensorTypeData(id, controller.DB)

	result, err := json.Marshal(item)

	setup.CheckErr(err)

	writeResponse(writer, result)
}
