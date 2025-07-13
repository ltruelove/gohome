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

type NodeSensorController struct {
	DB             *sql.DB
	NodeSensorData data.CrudDataInterface
	SensorType     data.CrudDataInterface
}

func (controller *NodeSensorController) RegisterNodeSensorEndpoints() {
	routing.AddRouteWithMethod("/sensor", "GET", controller.GetAll)
}

func NewNodeSensorController(db *sql.DB, config *config.Configuration) *NodeSensorController {
	return &NodeSensorController{
		DB:             db,
		NodeSensorData: data.NewNodeSensorData(db, config),
		SensorType:     data.NewSensorType(db, config),
	}
}

func (controller *NodeSensorController) GetAll(writer http.ResponseWriter, request *http.Request) {
	log.Println("Fetch all node sensors request initiated")

	allTypes, err := controller.SensorType.SelectAll()

	if err != nil {
		log.Printf("An error occurred fetching all sensor types: %v", err)
		http.Error(writer, "Unknown error has occured", http.StatusInternalServerError)
		return
	}

	allItems, err := controller.NodeSensorData.SelectAll()

	if err != nil {
		log.Printf("An error occurred fetching all nodes: %v", err)
		http.Error(writer, "Unknown error has occured", http.StatusInternalServerError)
		return
	}

	allResults := []viewModels.NodeSensorVM{}

	for _, value := range allItems {
		value := value.(models.NodeSensor)
		selectedType := models.SensorType{}

		for _, typeValue := range allTypes {
			typeValue := typeValue.(models.SensorType)
			if typeValue.Id == value.SensorTypeId {
				selectedType = typeValue
			}
		}

		sensorViewModel := viewModels.NodeSensorVM{}
		sensorViewModel.ImportModel(&value)
		sensorViewModel.SensorTypeName = selectedType.TypeName

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
