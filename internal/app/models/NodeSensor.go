package models

type NodeSensor struct {
	Id           int    `json:"Id"`
	NodeId       int    `json:"NodeId"`
	SensorTypeId int    `json:"SensorTypeId"`
	Name         string `json:"Name"`
	Pin          string `json:"Pin"`
}
