package models

type ViewNodeSensorData struct {
	Id               int    `json:"Id"`
	NodeId           int    `json:"NodeId"`
	ViewId           int    `json:"ViewId"`
	NodeSensorId     int    `json:"NodeSensorId"`
	SensorTypeDataId int    `json:"SensorTypeDataId"`
	Name             string `json:"Name"`
}
