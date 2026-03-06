package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/ltruelove/gohome/internal/app/repository"
	"github.com/ltruelove/gohome/internal/app/viewModels"
	"github.com/ltruelove/gohome/internal/pkg/routing"
)

type ViewController struct {
	ViewData           repository.CrudRepository
	NodeSensorData     repository.CrudRepository
	ViewNodeSensorData repository.CrudRepository
	NodeSwitchData     repository.CrudRepository
	ViewNodeSwitchData repository.CrudRepository
	CompoundData       repository.CompoundRepository
}

func (controller *ViewController) RegisterViewEndpoints() {
	routing.AddRouteWithMethod("/view", "GET", controller.GetAll)
	routing.AddRouteWithMethod("/view/{id}", "GET", controller.GetById)
	routing.AddRouteWithMethod("/view", "POST", controller.Create)
	routing.AddRouteWithMethod("/view", "PUT", controller.Update)
	routing.AddRouteWithMethod("/view/{id}", "DELETE", controller.Delete)
	routing.AddRouteWithMethod("/view/node/sensor", "POST", controller.AddNodeSensorToView)
	routing.AddRouteWithMethod("/view/node/sensor/{id}", "DELETE", controller.RemoveNodeSensorFromView)
	routing.AddRouteWithMethod("/view/node/switch", "POST", controller.AddNodeSwitchToView)
	routing.AddRouteWithMethod("/view/node/switch/{id}", "DELETE", controller.RemoveNodeSwitchFromView)
}

// convenience wrapper removed — use NewViewControllerWithDeps for DI

// NewViewControllerWithDeps constructs a ViewController with pre-built repositories.
func NewViewControllerWithDeps(viewRepo repository.CrudRepository, nodeSensorRepo repository.CrudRepository, viewNodeSensorRepo repository.CrudRepository, nodeSwitchRepo repository.CrudRepository, viewNodeSwitchRepo repository.CrudRepository, compoundRepo repository.CompoundRepository) *ViewController {
	return &ViewController{
		ViewData:           viewRepo,
		NodeSensorData:     nodeSensorRepo,
		ViewNodeSensorData: viewNodeSensorRepo,
		NodeSwitchData:     nodeSwitchRepo,
		ViewNodeSwitchData: viewNodeSwitchRepo,
		CompoundData:       compoundRepo,
	}
}

func (controller *ViewController) GetAll(writer http.ResponseWriter, request *http.Request) {
	log.Println("Fetch all views request initiated")

	allItems, err := controller.ViewData.SelectAll()

	if err != nil {
		log.Printf("An error occurred fetching all views: %v", err)
		http.Error(writer, "Unknown error has occured", http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(allItems)

	if err != nil {
		log.Printf("An error occurred marshalling view data: %v", err)
		http.Error(writer, "Data error", http.StatusInternalServerError)
		return
	}

	log.Printf("Found %d views", len(allItems))
	writeResponse(writer, result)
}

func (controller *ViewController) GetById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Printf("An error occurred fetching a view by id: %v", err)
		http.Error(writer, "Unknown error has occured", http.StatusInternalServerError)
		return
	}

	log.Printf("Fetch view by id: %d", id)

	model, err := controller.ViewData.SelectById(id)
	if err != nil {
		log.Println("view not found")
		http.Error(writer, "view not found", http.StatusNotFound)
		return
	}
	item := model.(*models.View)

	var viewModel = viewModels.ViewVM{Id: item.Id, Name: item.Name}

	viewModel.Sensors, err = controller.CompoundData.FetchViewNodeSensorDataByViewId(id)

	if err != nil {
		log.Printf("An error occurred fetching all view sensors: %v", err)
		http.Error(writer, "Unknown error has occured", http.StatusInternalServerError)
		return
	}

	viewModel.Switches, err = controller.CompoundData.FetchViewNodeSwitchDataByViewId(id)

	if err != nil {
		log.Printf("An error occurred fetching all view switches: %v", err)
		http.Error(writer, "Unknown error has occured", http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(viewModel)

	if err != nil {
		log.Printf("An error occurred marshalling view data: %v", err)
		http.Error(writer, "Data error", http.StatusInternalServerError)
		return
	}

	writeResponse(writer, result)
}

func (controller *ViewController) Create(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var item models.View

	err := decoder.Decode(&item)
	if err != nil {
		log.Printf("Error decoding the view data: %v", err)
		http.Error(writer, "Error decoding the request", http.StatusBadRequest)
		return
	}

	model, err := controller.ViewData.Insert(item)
	viewItem := model.(*models.View)

	if err != nil {
		log.Printf("Error creating a view: %v", err)
		http.Error(writer, "There was an error creating the record", http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(viewItem)

	if err != nil {
		log.Printf("An error occurred marshalling view data: %v", err)
		http.Error(writer, "Data error", http.StatusInternalServerError)
		return
	}

	writeResponse(writer, result)
}

func (controller *ViewController) Update(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var item models.View

	err := decoder.Decode(&item)
	if err != nil {
		log.Printf("Error decoding the view data: %v", err)
		http.Error(writer, "Error decoding the request", http.StatusBadRequest)
		return
	}

	_, err = controller.ViewData.SelectById(item.Id)

	if err != nil {
		log.Println("view not found")
		http.Error(writer, "view not found", http.StatusNotFound)
		return
	}

	err = controller.ViewData.Update(item)

	if err != nil {
		log.Printf("Error updating a view: %v", err)
		http.Error(writer, "There was an error updating the record", http.StatusInternalServerError)
		return
	}
}

func (controller *ViewController) Delete(writer http.ResponseWriter, request *http.Request) {
	log.Println("Delete a view")

	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Println("Could not get view id for delete")
		http.Error(writer, "Could not resolve view id", http.StatusBadRequest)
		return
	}

	_, err = controller.ViewData.SelectById(id)

	if err != nil {
		log.Println("view not found")
		http.Error(writer, "view not found", http.StatusNotFound)
		return
	}

	err = controller.ViewData.Delete(id)

	if err != nil {
		log.Printf("There was an error attempting to delete a view: %v", err)
		http.Error(writer, "There was an error attempting to delete a view", http.StatusInternalServerError)
	}
}

func (controller *ViewController) AddNodeSensorToView(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var item models.ViewNodeSensorData

	err := decoder.Decode(&item)
	if err != nil {
		log.Printf("Error decoding the node sensor data: %v", err)
		http.Error(writer, "Error decoding the request", http.StatusBadRequest)
		return
	}

	model, err := controller.ViewNodeSensorData.Insert(item)
	sensorData := model.(*models.ViewNodeSensorData)

	if err != nil {
		log.Printf("Error creating a view node sensor: %v", err)
		http.Error(writer, "There was an error creating the record", http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(sensorData)

	if err != nil {
		log.Printf("An error occurred marshalling view node sensor data: %v", err)
		http.Error(writer, "Data error", http.StatusInternalServerError)
		return
	}

	writeResponse(writer, result)
}

func (controller *ViewController) RemoveNodeSensorFromView(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "OPTIONS" {
		log.Println("OPTIONS request")
		writer.WriteHeader(http.StatusOK)
		writeResponse(writer, []byte(""))
		return
	}

	log.Println("Delete a view node sensor")

	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Println("Could not get view node sensor id for delete")
		http.Error(writer, "Could not resolve view node sensor id", http.StatusBadRequest)
		return
	}

	err = controller.ViewNodeSensorData.DeleteBySecondParentId(id)
	err = controller.NodeSensorData.Delete(id)

	if err != nil {
		log.Printf("There was an error attempting to delete a view node sensor: %v", err)
		http.Error(writer, "There was an error attempting to delete a view node sensor", http.StatusInternalServerError)
	}
}

func (controller *ViewController) AddNodeSwitchToView(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var item models.ViewNodeSwitchData

	err := decoder.Decode(&item)
	if err != nil {
		log.Printf("Error decoding the node switch data: %v", err)
		http.Error(writer, "Error decoding the request", http.StatusBadRequest)
		return
	}

	switchItem, err := controller.ViewNodeSwitchData.Insert(&item)

	if err != nil {
		log.Printf("Error creating a view node switch: %v", err)
		http.Error(writer, "There was an error creating the record", http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(switchItem)

	if err != nil {
		log.Printf("An error occurred marshalling view node switch data: %v", err)
		http.Error(writer, "Data error", http.StatusInternalServerError)
		return
	}

	writeResponse(writer, result)
}

func (controller *ViewController) RemoveNodeSwitchFromView(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "OPTIONS" {
		log.Println("OPTIONS request")
		writer.WriteHeader(http.StatusOK)
		writeResponse(writer, []byte(""))
		return
	}

	log.Println("Delete a view node switch")

	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Println("Could not get view node switch id for delete")
		http.Error(writer, "Could not resolve view node switch id", http.StatusBadRequest)
		return
	}

	err = controller.ViewNodeSwitchData.DeleteBySecondParentId(id)
	err = controller.NodeSwitchData.Delete(id)

	if err != nil {
		log.Printf("There was an error attempting to delete a view node switch: %v", err)
		http.Error(writer, "There was an error attempting to delete a view node switch", http.StatusInternalServerError)
	}
}
