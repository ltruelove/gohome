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

type SwitchTypeController struct {
	DB       *sql.DB
	AllTypes []models.SwitchType
}

func (controller *SwitchTypeController) RegisterSwitchTypeEndpoints() {
	routing.AddRouteWithMethod("/switchType", "GET", controller.AllSwitchTypes)
	routing.AddRouteWithMethod("/switchType/{id}", "GET", controller.SwitchTypeById)
}

func (controller *SwitchTypeController) AllSwitchTypes(writer http.ResponseWriter, request *http.Request) {
	if len(controller.AllTypes) == 0 {
		controller.AllTypes = data.FetchAllSwitchTypes(controller.DB)
	}

	result, err := json.Marshal(controller.AllTypes)

	setup.CheckErr(err)

	writeResponse(writer, result)
}

func (controller *SwitchTypeController) SwitchTypeById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

	setup.CheckErr(err)

	item := data.FetchSwitchType(id, controller.DB)

	result, err := json.Marshal(item)

	setup.CheckErr(err)

	writeResponse(writer, result)
}
