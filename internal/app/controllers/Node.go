package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ltruelove/gohome/internal/app/data"
	"github.com/ltruelove/gohome/internal/pkg/routing"
)

type NodeController struct {
	DB *sql.DB
}

func (controller *NodeController) RegisterNodeEndpoints() {
	routing.AddRouteWithMethod("/node", "GET", controller.AllNodes)
	routing.AddRouteWithMethod("/node/{id}", "GET", controller.NodeById)
}

func (controller *NodeController) AllNodes(writer http.ResponseWriter, request *http.Request) {
	log.Println("Fetch all nodes request initiated")

	allItems, err := data.FetchAllNodes(controller.DB)

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

func (controller *NodeController) NodeById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Printf("An error occurred fetching a node by id: %v", err)
		http.Error(writer, "Unknown error has occured", http.StatusInternalServerError)
		return
	}

	log.Printf("Fetch node by id: %d", id)

	item, err := data.FetchNode(id, controller.DB)

	if err != nil {
		log.Println("Node not found")
		http.Error(writer, "Node not found", http.StatusNotFound)
		return
	}

	result, err := json.Marshal(item)

	if err != nil {
		log.Printf("An error occurred marshalling node data: %v", err)
		http.Error(writer, "Data error", http.StatusInternalServerError)
		return
	}

	log.Println("Node found")

	writeResponse(writer, result)
}
