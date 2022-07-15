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
	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/ltruelove/gohome/internal/pkg/routing"
)

type ControlPointController struct {
	DB *sql.DB
}

func (controller *ControlPointController) RegisterControlPointEndpoints() {
	routing.AddRouteWithMethod("/controlPoint", "GET", controller.AllControlPoints)
	routing.AddRouteWithMethod("/controlPoint/{id}", "GET", controller.ControlPointById)
	routing.AddRouteWithMethod("/controlPoint", "POST", controller.CreateControlPoint)
	routing.AddRouteWithMethod("/controlPoint", "PUT", controller.UpdateControlPoint)
	routing.AddRouteWithMethod("/controlPoint/{id}", "DELETE", controller.DeleteControlPoint)
	routing.AddRouteWithMethod("/controlPoint/register", "POST", controller.RegisterControlPoint)
	routing.AddRouteWithMethod("/controlPoint/ipUpdate", "PUT", controller.UpdateControlPointIp)
}

func (controller *ControlPointController) AllControlPoints(writer http.ResponseWriter, request *http.Request) {
	log.Println("Fetch all controlPoints request initiated")

	allItems, err := data.FetchAllControlPoints(controller.DB)

	if err != nil {
		log.Printf("An error occurred fetching all controlPoints: %v", err)
		http.Error(writer, "Unknown error has occured", http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(allItems)

	if err != nil {
		log.Printf("An error occurred marshalling controlPoint data: %v", err)
		http.Error(writer, "Data error", http.StatusInternalServerError)
		return
	}

	log.Printf("Found %d controlPoints", len(allItems))
	writeResponse(writer, result)
}

func (controller *ControlPointController) ControlPointById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Printf("An error occurred fetching a controlPoint by id: %v", err)
		http.Error(writer, "Unknown error has occured", http.StatusInternalServerError)
		return
	}

	log.Printf("Fetch controlPoint by id: %d", id)

	item, err := data.FetchControlPoint(id, controller.DB)

	if err != nil {
		log.Println("controlPoint not found")
		http.Error(writer, "controlPoint not found", http.StatusNotFound)
		return
	}

	result, err := json.Marshal(item)

	if err != nil {
		log.Printf("An error occurred marshalling controlPoint data: %v", err)
		http.Error(writer, "Data error", http.StatusInternalServerError)
		return
	}

	writeResponse(writer, result)
}

func (controller *ControlPointController) RegisterControlPoint(writer http.ResponseWriter, request *http.Request) {
	log.Println("Register control point request made")
	decoder := json.NewDecoder(request.Body)
	var item models.ControlPoint

	err := decoder.Decode(&item)
	if err != nil {
		log.Printf("Error decoding the controlPoint data: %v", err)
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

	err = data.CreateControlPoint(&item, controller.DB)

	if err != nil {
		log.Printf("Error creating a controlPoint: %v", err)
		http.Error(writer, "There was an error creating the record", http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(item)

	if err != nil {
		log.Printf("An error occurred marshalling controlPoint data: %v", err)
		http.Error(writer, "Data error", http.StatusInternalServerError)
		return
	}

	writeResponse(writer, result)
}

func (controller *ControlPointController) CreateControlPoint(writer http.ResponseWriter, request *http.Request) {
	log.Println("Create control point request made")

	decoder := json.NewDecoder(request.Body)
	var item models.ControlPoint

	err := decoder.Decode(&item)
	if err != nil {
		log.Printf("Error decoding the controlPoint data: %v", err)
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

	err = data.CreateControlPoint(&item, controller.DB)

	if err != nil {
		log.Printf("Error creating a controlPoint: %v", err)
		http.Error(writer, "There was an error creating the record", http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(item)

	if err != nil {
		log.Printf("An error occurred marshalling controlPoint data: %v", err)
		http.Error(writer, "Data error", http.StatusInternalServerError)
		return
	}

	writeResponse(writer, result)
}

func (controller *ControlPointController) UpdateControlPointIp(writer http.ResponseWriter, request *http.Request) {
	log.Println("Update control point IP Address request made")

	decoder := json.NewDecoder(request.Body)
	var item models.ControlPoint

	err := decoder.Decode(&item)
	if err != nil {
		log.Printf("Error decoding the controlPoint data: %v", err)
		http.Error(writer, "Error decoding the request", http.StatusBadRequest)
		return
	}

	err = item.IsIpAddressValid()

	if err != nil {
		vError := fmt.Sprintf("Validation error: %v", err)
		log.Println(vError)
		http.Error(writer, vError, http.StatusBadRequest)
		return
	}

	isNew, err := data.VerifyControlPointIdIsNew(item.Id, controller.DB)

	if err != nil {
		log.Printf("Error checking controlPoint id: %v", err)
		http.Error(writer, "Error checking id", http.StatusInternalServerError)
		return
	}

	if isNew {
		log.Printf("ControlPoint for id %d doesn't exist", item.Id)
		http.Error(writer, "ControlPoint not found", http.StatusNotFound)
		return
	}

	err = data.UpdateControlPointIp(&item, controller.DB)

	if err != nil {
		log.Printf("Error updating a controlPoint: %v", err)
		http.Error(writer, "There was an error updating the record", http.StatusInternalServerError)
		return
	}
}

func (controller *ControlPointController) UpdateControlPoint(writer http.ResponseWriter, request *http.Request) {
	log.Println("Update control point request made")

	decoder := json.NewDecoder(request.Body)
	var item models.ControlPoint

	err := decoder.Decode(&item)
	if err != nil {
		log.Printf("Error decoding the controlPoint data: %v", err)
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

	isNew, err := data.VerifyControlPointIdIsNew(item.Id, controller.DB)

	if err != nil {
		log.Printf("Error checking controlPoint id: %v", err)
		http.Error(writer, "Error checking id", http.StatusInternalServerError)
		return
	}

	if isNew {
		log.Printf("ControlPoint for id %d doesn't exist", item.Id)
		http.Error(writer, "ControlPoint not found", http.StatusNotFound)
		return
	}

	err = data.UpdateControlPoint(&item, controller.DB)

	if err != nil {
		log.Printf("Error updating a controlPoint: %v", err)
		http.Error(writer, "There was an error updating the record", http.StatusInternalServerError)
		return
	}
}

func (controller *ControlPointController) DeleteControlPoint(writer http.ResponseWriter, request *http.Request) {
	log.Println("Delete a controlPoint")

	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Println("Could not get controlPoint id for delete")
		http.Error(writer, "Could not resolve controlPoint id", http.StatusBadRequest)
		return
	}

	isNew, err := data.VerifyControlPointIdIsNew(id, controller.DB)

	if err != nil {
		log.Printf("Error checking controlPoint id: %v", err)
		http.Error(writer, "Error checking id", http.StatusInternalServerError)
		return
	}

	if isNew {
		log.Printf("ControlPoint for id %d doesn't exist", id)
		http.Error(writer, "ControlPoint not found", http.StatusNotFound)
		return
	}

	err = data.DeleteControlPoint(id, controller.DB)

	if err != nil {
		log.Printf("There was an error attempting to delete a controlPoint: %v", err)
		http.Error(writer, "There was an error attempting to delete a controlPoint", http.StatusInternalServerError)
	}
}