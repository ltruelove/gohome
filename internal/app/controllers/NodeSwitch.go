package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data"
	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/ltruelove/gohome/internal/app/viewModels"
	"github.com/ltruelove/gohome/internal/pkg/routing"
)

type NodeSwitchController struct {
	DB             *sql.DB
	NodeSwitchData data.CrudDataInterface
	SwitchTypeData data.CrudDataInterface
}

func NewNodeSwitchController(db *sql.DB, config *config.Configuration) *NodeSwitchController {
	return &NodeSwitchController{
		DB:             db,
		NodeSwitchData: data.NewNodeSwitchData(db, config),
		SwitchTypeData: data.NewSwitchTypeData(db, config),
	}
}

func (controller *NodeSwitchController) RegisterNodeSwitchEndpoints() {
	routing.AddRouteWithMethod("/switch", "GET", controller.GetAll)
}

func (controller *NodeSwitchController) GetAll(writer http.ResponseWriter, request *http.Request) {
	log.Println("Fetch all node sensors request initiated")

	allTypes, err := controller.SwitchTypeData.SelectAll()

	if err != nil {
		log.Printf("An error occurred fetching all sensor types: %v", err)
		http.Error(writer, "Unknown error has occured", http.StatusInternalServerError)
		return
	}

	allItems, err := controller.NodeSwitchData.SelectAll()

	if err != nil {
		log.Printf("An error occurred fetching all nodes: %v", err)
		http.Error(writer, "Unknown error has occured", http.StatusInternalServerError)
		return
	}

	allResults := []viewModels.NodeSwitchVM{}

	for _, switchData := range allItems {
		value := switchData.(models.NodeSwitch)
		selectedType := models.SwitchType{}

		for _, typeValue := range allTypes {
			typeValue := typeValue.(models.SwitchType)
			if typeValue.Id == value.SwitchTypeId {
				selectedType = typeValue
			}
		}

		sensorViewModel := viewModels.NodeSwitchVM{}
		sensorViewModel.ImportModel(&value)
		sensorViewModel.SwitchTypeName = selectedType.Name

		allResults = append(allResults, sensorViewModel)
	}

	result, err := json.Marshal(allResults)

	if err != nil {
		log.Printf("An error occurred marshalling node data: %v", err)
		http.Error(writer, "Data error", http.StatusInternalServerError)
		return
	}

	log.Printf("Found %d nodes", len(allResults))
	writeResponse(writer, result)
}
