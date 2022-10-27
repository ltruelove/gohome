package models

type ViewNodeSensorData struct {
	Id           int    `json:"Id"`
	NodeId       int    `json:"NodeId"`
	ViewId       int    `json:"ViewId"`
	NodeSensorId int    `json:"NodeSensorId"`
	Name         string `json:"Name"`
}
