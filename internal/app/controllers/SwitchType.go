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

type SwitchTypeController struct {
	SwitchTypeData *data.SwitchTypeData
}

func NewSwitchTypeController(db *sql.DB, config *config.Configuration) *SwitchTypeController {
	return &SwitchTypeController{
		SwitchTypeData: data.NewSwitchTypeData(db, config),
	}
}

func (controller *SwitchTypeController) RegisterSwitchTypeEndpoints() {
	routing.AddRouteWithMethod("/switchType", "GET", controller.GetAll)
	routing.AddRouteWithMethod("/switchType/{id}", "GET", controller.GetById)
}

func (controller *SwitchTypeController) GetAll(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")

	log.Println("Fetch all switch types")

	allTypes, fetchErr := controller.SwitchTypeData.FetchAllSwitchTypes()

	if fetchErr != nil {
		log.Printf("Error fetching switch types from the db: %v", fetchErr)
		http.Error(writer, "Data error", http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(allTypes)
	if err != nil {
		log.Printf("An error occurred marshalling switch data: %v", err)
		http.Error(writer, "Data error", http.StatusInternalServerError)
		return
	}

	log.Printf("%d switch types found", len(allTypes))
	writeResponse(writer, result)
}

func (controller *SwitchTypeController) GetById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Error getting switch type id from request: %v", err)
		http.Error(writer, "Request error", http.StatusBadRequest)
		return
	}

	log.Printf("Fetch switch type by id: %d", id)

	item, err := controller.SwitchTypeData.FetchSwitchType(id)
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
