package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/ltruelove/gohome/internal/app/repository"
	"github.com/ltruelove/gohome/internal/app/viewModels"
	"github.com/ltruelove/gohome/internal/pkg/routing"
)

type NodeSwitchController struct {
	NodeSwitch repository.CrudRepository
	SwitchType repository.CrudRepository
}

// convenience wrapper removed — use NewNodeSwitchControllerWithDeps for DI

// NewNodeSwitchControllerWithDeps constructs a NodeSwitchController with explicit
// data layer dependencies to enable DI and testing.
func NewNodeSwitchControllerWithDeps(nodeSwitchData repository.CrudRepository, switchTypeData repository.CrudRepository) *NodeSwitchController {
	return &NodeSwitchController{
		NodeSwitch: nodeSwitchData,
		SwitchType: switchTypeData,
	}
}

func (controller *NodeSwitchController) RegisterNodeSwitchEndpoints() {
	routing.AddRouteWithMethod("/switch", "GET", controller.GetAll)
}

func (controller *NodeSwitchController) GetAll(writer http.ResponseWriter, request *http.Request) {
	log.Println("Fetch all node sensors request initiated")

	allTypes, err := controller.SwitchType.SelectAll()

	if err != nil {
		log.Printf("An error occurred fetching all sensor types: %v", err)
		http.Error(writer, "Unknown error has occured", http.StatusInternalServerError)
		return
	}

	allItems, err := controller.NodeSwitch.SelectAll()

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
