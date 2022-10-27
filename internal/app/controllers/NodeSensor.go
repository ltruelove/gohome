package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/ltruelove/gohome/internal/app/data"
	"github.com/ltruelove/gohome/internal/pkg/routing"
)

type NodeSensorController struct {
	DB *sql.DB
}

func (controller *NodeSensorController) RegisterNodeSensorEndpoints() {
	routing.AddRouteWithMethod("/sensor", "GET", controller.GetAll)
}

func (controller *NodeSensorController) GetAll(writer http.ResponseWriter, request *http.Request) {
	log.Println("Fetch all node sensors request initiated")

	allItems, err := data.FetchAllNodeSensors(controller.DB)

	if err != nil {
		log.Printf("An error occurred fetching all nodes: %v", err)
		http.Error(writer, "Unknown error has occured", http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(allItems)

	if err != nil {
		log.Printf("An error occurred marshalling node data: %v", err)
		http.Error(writer, "Data error", http.StatusInternalServerError)
		return
	}

	log.Printf("Found %d nodes", len(allItems))
	writeResponse(writer, result)
}
