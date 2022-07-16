package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ltruelove/gohome/internal/app/data"
	"github.com/ltruelove/gohome/internal/app/dto"
	"github.com/ltruelove/gohome/internal/app/handler"
	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/ltruelove/gohome/internal/pkg/routing"
)

type NodeController struct {
	DB *sql.DB
}

func (controller *NodeController) RegisterNodeEndpoints() {
	routing.AddRouteWithMethod("/node", "GET", controller.GetAll)
	routing.AddRouteWithMethod("/node/{id}", "GET", controller.GetById)
	routing.AddRouteWithMethod("/node", "POST", controller.Create)
	routing.AddRouteWithMethod("/node", "PUT", controller.Update)
	routing.AddRouteWithMethod("/node", "DELETE", controller.Delete)
	routing.AddRouteWithMethod("/node/register", "POST", controller.Register)
}

func (controller *NodeController) GetAll(writer http.ResponseWriter, request *http.Request) {
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

func (controller *NodeController) GetById(writer http.ResponseWriter, request *http.Request) {
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

func (controller *NodeController) Create(writer http.ResponseWriter, request *http.Request) {
	log.Println("Create node request made")

	decoder := json.NewDecoder(request.Body)
	var item models.Node

	err := decoder.Decode(&item)
	if err != nil {
		log.Printf("Error decoding the node data: %v", err)
		http.Error(writer, "Error decoding the request", http.StatusBadRequest)
		return
	}

	err = item.IsValid(false)

	if err != nil {
		vError := fmt.Sprintf("Validation error: %v", err)
		log.Println(vError)
		http.Error(writer, vError, http.StatusBadRequest)
		return
	}

	err = data.CreateNode(&item, controller.DB)

	if err != nil {
		log.Printf("Error creating a node: %v", err)
		http.Error(writer, "There was an error creating the record", http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(item)

	if err != nil {
		log.Printf("An error occurred marshalling node data: %v", err)
		http.Error(writer, "Data error", http.StatusInternalServerError)
		return
	}

	writeResponse(writer, result)
}

func (controller *NodeController) Update(writer http.ResponseWriter, request *http.Request) {
	log.Println("Update node request made")

	decoder := json.NewDecoder(request.Body)
	var item models.Node

	err := decoder.Decode(&item)
	if err != nil {
		log.Printf("Error decoding the node data: %v", err)
		http.Error(writer, "Error decoding the request", http.StatusBadRequest)
		return
	}

	err = item.IsValid(true)

	if err != nil {
		vError := fmt.Sprintf("Validation error: %v", err)
		log.Println(vError)
		http.Error(writer, vError, http.StatusBadRequest)
		return
	}

	isNew, err := data.VerifyNodeIdIsNew(item.Id, controller.DB)

	if err != nil {
		log.Printf("Error checking node id: %v", err)
		http.Error(writer, "Error checking id", http.StatusInternalServerError)
		return
	}

	if isNew {
		log.Printf("Node for id %d doesn't exist", item.Id)
		http.Error(writer, "Node not found", http.StatusNotFound)
		return
	}

	err = data.UpdateNode(&item, controller.DB)

	if err != nil {
		log.Printf("Error updating a node: %v", err)
		http.Error(writer, "There was an error updating the record", http.StatusInternalServerError)
		return
	}
}

func (controller *NodeController) Delete(writer http.ResponseWriter, request *http.Request) {
	log.Println("Delete a node")

	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Println("Could not get node id for delete")
		http.Error(writer, "Could not resolve node id", http.StatusBadRequest)
		return
	}

	isNew, err := data.VerifyNodeIdIsNew(id, controller.DB)

	if err != nil {
		log.Printf("Error checking node id: %v", err)
		http.Error(writer, "Error checking id", http.StatusInternalServerError)
		return
	}

	if isNew {
		log.Printf("Node for id %d doesn't exist", id)
		http.Error(writer, "Node not found", http.StatusNotFound)
		return
	}

	err = data.DeleteNode(id, controller.DB)

	if err != nil {
		log.Printf("There was an error attempting to delete a node: %v", err)
		http.Error(writer, "There was an error attempting to delete a node", http.StatusInternalServerError)
	}
}

func (controller *NodeController) Register(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token, Authorization, X-Requested-With")

	log.Println("Register node request made")

	decoder := json.NewDecoder(request.Body)
	var item dto.RegsiterNode

	err := decoder.Decode(&item)
	if err != nil {
		log.Printf("Error decoding the node data: %v", err)
		http.Error(writer, "Error decoding the request", http.StatusBadRequest)
		return
	}

	err = handler.RegisterNode(&item, controller.DB)

	if err != nil {
		log.Printf("Error registering a node: %v", err)
		http.Error(writer, "There was an error registering the node", http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(item)

	if err != nil {
		log.Printf("An error occurred marshalling node data: %v", err)
		http.Error(writer, "Data error", http.StatusInternalServerError)
		return
	}

	log.Println("Register node request succeeded")
	writeResponse(writer, result)
}
