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

type SwitchTypeController struct {
	DB       *sql.DB
	AllTypes []models.SwitchType
}

func (controller *SwitchTypeController) RegisterSwitchTypeEndpoints() {
	routing.AddRouteWithMethod("/switchType", "GET", controller.GetAll)
	routing.AddRouteWithMethod("/switchType/{id}", "GET", controller.GetById)
}

func (controller *SwitchTypeController) GetAll(writer http.ResponseWriter, request *http.Request) {
	log.Println("Fetch all switch types")

	if len(controller.AllTypes) == 0 {
		var fetchErr error
		controller.AllTypes, fetchErr = data.FetchAllSwitchTypes(controller.DB)

		if fetchErr != nil {
			log.Printf("Error fetching switch types from the db: %v", fetchErr)
			http.Error(writer, "Data error", http.StatusInternalServerError)
			return
		}
	}

	result, err := json.Marshal(controller.AllTypes)
	if err != nil {
		log.Printf("An error occurred marshalling switch data: %v", err)
		http.Error(writer, "Data error", http.StatusInternalServerError)
		return
	}

	log.Printf("%d switch types found", len(controller.AllTypes))
	writeResponse(writer, result)
}

func (controller *SwitchTypeController) GetById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Error getting switch type id from request: %v", err)
		http.Error(writer, "Request error", http.StatusBadRequest)
		return
	}

	log.Printf("Fetch switch type by id: %d", id)

	item, err := data.FetchSwitchType(id, controller.DB)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Printf("Error getting switch type: %v", err)
			http.Error(writer, "Data error", http.StatusInternalServerError)
		} else {
			log.Println("Switch type not found")
			http.Error(writer, "Switch type not found", http.StatusNotFound)
		}
		return
	}

	result, err := json.Marshal(item)
	if err != nil {
		log.Printf("Error marshalling json for switch type: %v", err)
		http.Error(writer, "Data error", http.StatusInternalServerError)
		return
	}

	writeResponse(writer, result)
}
