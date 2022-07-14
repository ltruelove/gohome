package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ltruelove/gohome/internal/app/data"
	"github.com/ltruelove/gohome/internal/app/models"
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
	log.Println("Fetch all sensor types")

	if len(controller.AllTypes) == 0 {
		var fetchErr error
		controller.AllTypes, fetchErr = data.FetchAllSensorTypes(controller.DB)

		if fetchErr != nil {
			log.Printf("Error fetching sensor types from the db: %v", fetchErr)
			http.Error(writer, "Data error", http.StatusInternalServerError)
			return
		}
	}

	result, err := json.Marshal(controller.AllTypes)
	if err != nil {
		log.Printf("An error occurred marshalling sensor data: %v", err)
		http.Error(writer, "Data error", http.StatusInternalServerError)
		return
	}

	log.Printf("%d sensor types found", len(controller.AllTypes))
	writeResponse(writer, result)
}

func (controller *SensorTypeController) SensorTypeById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Error getting sensor type id from request: %v", err)
		http.Error(writer, "Request error", http.StatusBadRequest)
		return
	}

	log.Printf("Fetch sensor type by id: %d", id)

	item, err := data.FetchSensorType(id, controller.DB)
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

func (controller *SensorTypeController) SensorTypeDataById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Error getting sensor type id from request: %v", err)
		http.Error(writer, "Request error", http.StatusBadRequest)
		return
	}

	log.Printf("Fetch all sensor type data for a sensor with the id: %d", id)

	item, err := data.FetchSensorTypeData(id, controller.DB)

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
