package viewModels

import "github.com/ltruelove/gohome/internal/app/models"

type NodeSensorVM struct {
	Id             int    `json:"Id"`
	NodeId         int    `json:"NodeId"`
	SensorTypeId   int    `json:"SensorTypeId"`
	Name           string `json:"Name"`
	SensorTypeName string `json:"SensorTypeName"`
	Pin            int    `json:"Pin"`
	DHTType        int    `json:"DHTType"`
}

func (nodeSensor *NodeSensorVM) ImportModel(model *models.NodeSensor) {
	nodeSensor.Id = model.Id
	nodeSensor.NodeId = model.NodeId
	nodeSensor.SensorTypeId = model.SensorTypeId
	nodeSensor.Name = model.Name
	nodeSensor.Pin = model.Pin
	nodeSensor.DHTType = model.DHTType
}
