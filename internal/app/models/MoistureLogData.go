package models

// swagger:model Node
type MoistureLogData struct {
	// The ID of the Node
	Id int `json:"id"`
	// The ID of the Node
	NodeSensorLogId int `json:"nodeSensorLogId"`
	// The ID of the Node
	Moisture int `json:"Moisture"`
}
