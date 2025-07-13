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
	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/ltruelove/gohome/internal/app/viewModels"
	"github.com/ltruelove/gohome/internal/pkg/routing"
)

type ViewController struct {
	ViewData       *data.ViewData
	NodeSensorData data.CrudDataInterface
	SensorData     *data.ViewNodeSensorData
	NodeSwitchData data.CrudDataInterface
	SwitchData     data.CrudDataInterface
	CompoundData   *data.CompoundData
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

func NewViewController(db *sql.DB, config *config.Configuration) *ViewController {
	return &ViewController{
		ViewData:       data.NewViewData(db, config),
		NodeSensorData: data.NewNodeSensorData(db, config),
		SensorData:     data.NewViewNodeSensorData(db, config),
		NodeSwitchData: data.NewNodeSwitchData(db, config),
		SwitchData:     data.NewViewNodeSwitchData(db, config),
		CompoundData:   data.NewCompoundData(db, config),
	}
}

func (controller *ViewController) GetAll(writer http.ResponseWriter, request *http.Request) {
	log.Println("Fetch all views request initiated")

	allItems, err := controller.ViewData.FetchAllViews()

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

	item, err := controller.ViewData.FetchView(id)

	if err != nil {
		log.Println("view not found")
		http.Error(writer, "view not found", http.StatusNotFound)
		return
	}

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

	err = controller.ViewData.CreateView(&item)

	if err != nil {
		log.Printf("Error creating a view: %v", err)
		http.Error(writer, "There was an error creating the record", http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(item)

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

	_, err = controller.ViewData.FetchView(item.Id)

	if err != nil {
		log.Println("view not found")
		http.Error(writer, "view not found", http.StatusNotFound)
		return
	}

	err = controller.ViewData.UpdateView(&item)

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

	_, err = controller.ViewData.FetchView(id)

	if err != nil {
		log.Println("view not found")
		http.Error(writer, "view not found", http.StatusNotFound)
		return
	}

	err = controller.ViewData.DeleteView(id)

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

	err = controller.SensorData.CreateViewNodeSensorData(&item)

	if err != nil {
		log.Printf("Error creating a view node sensor: %v", err)
		http.Error(writer, "There was an error creating the record", http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(item)

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

	err = controller.SensorData.DeleteByParent(id)
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

	switchItem, err := controller.SwitchData.Insert(&item)

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

	err = controller.NodeSwitchData.Delete(id)

	if err != nil {
		log.Printf("There was an error attempting to delete a view node switch: %v", err)
		http.Error(writer, "There was an error attempting to delete a view node switch", http.StatusInternalServerError)
	}
}
