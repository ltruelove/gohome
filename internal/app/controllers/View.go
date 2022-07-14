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

type ViewController struct {
	DB *sql.DB
}

func (controller *ViewController) RegisterViewEndpoints() {
	routing.AddRouteWithMethod("/view", "GET", controller.AllViews)
	routing.AddRouteWithMethod("/view/{id}", "GET", controller.ViewById)
	routing.AddRouteWithMethod("/view", "POST", controller.CreateView)
	routing.AddRouteWithMethod("/view", "PUT", controller.UpdateView)
	routing.AddRouteWithMethod("/view/{id}", "DELETE", controller.DeleteView)
}

func (controller *ViewController) AllViews(writer http.ResponseWriter, request *http.Request) {
	log.Println("Fetch all views request initiated")

	allItems, err := data.FetchAllViews(controller.DB)

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

func (controller *ViewController) ViewById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Printf("An error occurred fetching a view by id: %v", err)
		http.Error(writer, "Unknown error has occured", http.StatusInternalServerError)
		return
	}

	log.Printf("Fetch view by id: %d", id)

	item, err := data.FetchView(id, controller.DB)

	if err != nil {
		log.Println("view not found")
		http.Error(writer, "view not found", http.StatusNotFound)
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

func (controller *ViewController) CreateView(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var item models.View

	err := decoder.Decode(&item)
	if err != nil {
		log.Printf("Error decoding the view data: %v", err)
		http.Error(writer, "Error decoding the request", http.StatusBadRequest)
		return
	}

	err = data.CreateView(&item, controller.DB)

	if err != nil {
		log.Printf("Error creating a view: %v", err)
		http.Error(writer, "There was an error creating the record", http.StatusInternalServerError)
		return
	}
}

func (controller *ViewController) UpdateView(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var item models.View

	err := decoder.Decode(&item)
	if err != nil {
		log.Printf("Error decoding the view data: %v", err)
		http.Error(writer, "Error decoding the request", http.StatusBadRequest)
		return
	}

	err = data.UpdateView(&item, controller.DB)

	if err != nil {
		log.Printf("Error updating a view: %v", err)
		http.Error(writer, "There was an error updating the record", http.StatusInternalServerError)
		return
	}
}

func (controller *ViewController) DeleteView(writer http.ResponseWriter, request *http.Request) {
	log.Println("Delete a view")

	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Println("Could not get view id for delete")
		http.Error(writer, "Could not resolve view id", http.StatusBadRequest)
		return
	}

	isNew, err := data.VerifyViewIdIsNew(id, controller.DB)

	if err != nil {
		log.Printf("Error checking view id: %v", err)
		http.Error(writer, "Error checking id", http.StatusInternalServerError)
		return
	}

	if isNew {
		log.Printf("View for id %d doesn't exist")
		http.Error(writer, "View not found", http.StatusNotFound)
		return
	}

	err = data.DeleteView(id, controller.DB)

	if err != nil {
		log.Printf("There was an error attempting to delete a view: %v", err)
		http.Error(writer, "There was an error attempting to delete a view", http.StatusInternalServerError)
	}
}
