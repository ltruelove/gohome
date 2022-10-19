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
	routing.AddRouteWithMethod("/node/{id}/delete", "DELETE", controller.Delete)
	routing.AddRouteWithMethod("/node/register", "POST", controller.Register)
	routing.AddRouteWithMethod("/node/switchesByNode/{id}", "GET", controller.GetAllNodeSwitches)
	routing.AddRouteWithMethod("/node/switch/toggle/{id}", "POST", controller.ToggleNodeSwitch)
	routing.AddRouteWithMethod("/node/switch/press/{id}", "POST", controller.PressNodeSwitch)
}

func (controller *NodeController) GetAll(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
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
	writer.Header().Set("Access-Control-Allow-Origin", "*")

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
	writer.Header().Set("Access-Control-Allow-Origin", "*")
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
	writer.Header().Set("Access-Control-Allow-Origin", "*")
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
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Methods", "DELETE, OPTIONS")

	if request.Method == "OPTIONS" {
		log.Println("OPTIONS request")
		writer.WriteHeader(http.StatusOK)
		writeResponse(writer, []byte(""))
		return
	}

	log.Println("Delete a node")

	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Println("Could not get node id for delete")
		http.Error(writer, "Could not resolve node id", http.StatusBadRequest)
		return
	}

	/*
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
	*/

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

	controlPoint, err := data.FetchControlPoint(item.ControlPoint.Id, controller.DB)

	if err != nil {
		log.Printf("Could not find requested control point: %v", err)
		http.Error(writer, "Control point not found", http.StatusNotFound)
		return
	}

	item.ControlPoint = controlPoint

	err = handler.RegisterNode(&item, controller.DB)

	if err != nil {
		log.Printf("Error registering a node: %v", err)
		http.Error(writer, "There was an error registering the node", http.StatusInternalServerError)
		return
	}

	controlPointNode := models.ControlPointNode{}
	controlPointNode.ControlPointId = item.ControlPoint.Id
	controlPointNode.NodeId = item.Node.Id

	err = data.AddNodeToControlPoint(&controlPointNode, controller.DB)

	if err != nil {
		log.Printf("An error occurred adding the node to the control point: %v", err)
		http.Error(writer, "Data error", http.StatusInternalServerError)
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

func (controller *NodeController) GetAllNodeSwitches(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	log.Println("Fetch all node switches request initiated")

	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Println("Could not get node id for delete")
		http.Error(writer, "Could not resolve node id", http.StatusBadRequest)
		return
	}

	allItems, err := data.FetchNodeSwitches(id, controller.DB)

	if err != nil {
		log.Printf("An error occurred fetching all node switches: %v", err)
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

func (controller *NodeController) ToggleNodeSwitch(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	log.Println("Toggle node switch request initiated")
	log.Println(request.Body)

	decoder := json.NewDecoder(request.Body)
	var pinRequest models.PinRequest

	err := decoder.Decode(&pinRequest)

	if err != nil {
		log.Println("Could not decode pin for toggle request")
		http.Error(writer, "Could not decode pin for toggle request", http.StatusBadRequest)
		return
	}

	if !Config.ValidatePin(pinRequest.PinCode) {
		log.Println("Invalid PIN used in toggle request")
		http.Error(writer, "Invalid PIN given", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Println("Could not get node switch id")
		http.Error(writer, "Could not resolve node switch id", http.StatusBadRequest)
		return
	}

	nodeSwitch, err := data.FetchNodeSwitch(id, controller.DB)

	if err != nil {
		log.Println("Could not get node switch")
		http.Error(writer, "Could not resolve node switch", http.StatusBadRequest)
		return
	}

	node, err := data.FetchNode(nodeSwitch.NodeId, controller.DB)

	if err != nil {
		log.Println("Could not get node")
		http.Error(writer, "Could not resolve node", http.StatusBadRequest)
		return
	}

	nodeControlPoint, err := data.FetchControlPointByNode(node.Id, controller.DB)

	if err != nil {
		log.Println("Could not get control point by node")
		http.Error(writer, "Could not resolve control point for node switch", http.StatusBadRequest)
		return
	}

	_, err = http.Get("http://" + nodeControlPoint.IpAddress + "/toggleNodeSwitch?mac=" + node.Mac)

	if err != nil {
		log.Println("Could not complete the toggle request." + err.Error())
		http.Error(writer, "There was a error making the toggle request.", http.StatusInternalServerError)
		return
	}

	writeResponse(writer, []byte("Toggle successful"))
}

func (controller *NodeController) PressNodeSwitch(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	log.Println("Press node switch request initiated")
	log.Println(request.Body)

	decoder := json.NewDecoder(request.Body)
	var pinRequest models.PinRequest

	err := decoder.Decode(&pinRequest)

	if err != nil {
		log.Println("Could not decode pin for toggle request")
		http.Error(writer, "Could not decode pin for toggle request", http.StatusBadRequest)
		return
	}

	if !Config.ValidatePin(pinRequest.PinCode) {
		log.Println("Invalid PIN used in toggle request")
		http.Error(writer, "Invalid PIN given", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Println("Could not get node switch id")
		http.Error(writer, "Could not resolve node switch id", http.StatusBadRequest)
		return
	}

	nodeSwitch, err := data.FetchNodeSwitch(id, controller.DB)

	if err != nil {
		log.Println("Could not get node switch")
		http.Error(writer, "Could not resolve node switch", http.StatusBadRequest)
		return
	}

	node, err := data.FetchNode(nodeSwitch.NodeId, controller.DB)

	if err != nil {
		log.Println("Could not get node")
		http.Error(writer, "Could not resolve node", http.StatusBadRequest)
		return
	}

	nodeControlPoint, err := data.FetchControlPointByNode(node.Id, controller.DB)

	if err != nil {
		log.Println("Could not get control point by node")
		http.Error(writer, "Could not resolve control point for node switch", http.StatusBadRequest)
		return
	}

	url := fmt.Sprintf("http://%s/pressMomentary?mac=%s&MomentaryPressDuration=%d",
		nodeControlPoint.IpAddress, node.Mac, nodeSwitch.MomentaryPressDuration)

	_, err = http.Get(url)

	if err != nil {
		log.Println("Could not complete the toggle request." + err.Error())
		http.Error(writer, "There was a error making the toggle request.", http.StatusInternalServerError)
		return
	}

	writeResponse(writer, []byte("Toggle successful"))
}
