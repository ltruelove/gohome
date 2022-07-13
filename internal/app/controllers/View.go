package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ltruelove/gohome/internal/app/data"
	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/ltruelove/gohome/internal/app/setup"
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
	allItems := data.FetchAllViews(controller.DB)

	result, err := json.Marshal(allItems)

	setup.CheckErr(err)

	writeResponse(writer, result)
}

func (controller *ViewController) ViewById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

	setup.CheckErr(err)

	item := data.FetchView(id, controller.DB)

	result, err := json.Marshal(item)

	setup.CheckErr(err)

	writeResponse(writer, result)
}

func (controller *ViewController) CreateView(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var item models.View

	err := decoder.Decode(&item)
	setup.CheckErr(err)

	data.CreateView(&item, controller.DB)
}

func (controller *ViewController) UpdateView(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var item models.View

	err := decoder.Decode(&item)
	setup.CheckErr(err)

	data.UpdateView(&item, controller.DB)
}

func (controller *ViewController) DeleteView(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

	setup.CheckErr(err)

	data.DeleteView(id, controller.DB)
}
