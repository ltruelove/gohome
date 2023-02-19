package models

// swagger:model Node
type MagneticLogData struct {
	// The ID of the Node
	Id int `json:"id"`
	// The ID of the Node
	NodeSensorLogId int `json:"nodeSensorLogId"`
	// The ID of the Node
	IsClosed bool `json:"IsClosed"`
}
