package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data"
	"github.com/ltruelove/gohome/internal/pkg/routing"
)

type SensorTypeController struct {
	SensorTypeData *data.SensorTypeData
}

func NewSensorTypeController(db *sql.DB, config *config.Configuration) *SensorTypeController {
	return &SensorTypeController{
		SensorTypeData: data.NewSensorTypeData(db, config),
	}

}

func (controller *SensorTypeController) RegisterSensorTypeEndpoints() {
	routing.AddRouteWithMethod("/sensorType", "GET", controller.GetAll)
	routing.AddRouteWithMethod("/sensorType/{id}", "GET", controller.GetById)
	routing.AddRouteWithMethod("/sensorType/data/{id}", "GET", controller.DataById)
}

func (controller *SensorTypeController) GetAll(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")

	log.Println("Fetch all sensor types")

	allTypes, fetchErr := controller.SensorTypeData.FetchAllSensorTypes()

	if fetchErr != nil {
		log.Printf("Error fetching sensor types from the db: %v", fetchErr)
		http.Error(writer, "Data error", http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(allTypes)
	if err != nil {
		log.Printf("An error occurred marshalling sensor data: %v", err)
		http.Error(writer, "Data error", http.StatusInternalServerError)
		return
	}

	log.Printf("%d sensor types found", len(allTypes))
	writeResponse(writer, result)
}

func (controller *SensorTypeController) GetById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Error getting sensor type id from request: %v", err)
		http.Error(writer, "Request error", http.StatusBadRequest)
		return
	}

	log.Printf("Fetch sensor type by id: %d", id)

	item, err := controller.SensorTypeData.FetchSensorType(id)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Printf("Error getting sensor type: %v", err)
			http.Error(writer, "Data error", http.StatusInternalServerError)
		} else {
			log.Println("Sensor type not found")
			http.Error(writer, "Sensor type not found", http.StatusNotFound)
		}
		return
	}

	result, err := json.Marshal(item)
	if err != nil {
		log.Printf("Error marshalling json for sensor type: %v", err)
		http.Error(writer, "Data error", http.StatusInternalServerError)
		return
	}

	writeResponse(writer, result)
}

func (controller *SensorTypeController) DataById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Error getting sensor type id from request: %v", err)
		http.Error(writer, "Request error", http.StatusBadRequest)
		return
	}

	log.Printf("Fetch all sensor type data for a sensor with the id: %d", id)

	item, err := controller.SensorTypeData.FetchSensorTypeData(id)

	if err != nil {
		if err != sql.ErrNoRows {
			log.Printf("Error getting sensor type data: %v", err)
			http.Error(writer, "Data error", http.StatusInternalServerError)
		} else {
			log.Println("Sensor type data not found")
			http.Error(writer, "Sensor type data not found", http.StatusNotFound)
		}
		return
	}

	result, err := json.Marshal(item)
	if err != nil {
		log.Printf("Error marshalling json for sensor type data: %v", err)
		http.Error(writer, "Data error", http.StatusInternalServerError)
		return
	}

	writeResponse(writer, result)
}
