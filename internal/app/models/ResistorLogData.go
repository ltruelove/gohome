package models

// swagger:model Node
type ResistorLogData struct {
	// The ID of the Node
	Id int `json:"id"`
	// The ID of the Node
	NodeSensorLogId int `json:"nodeSensorLogId"`
	// The ID of the Node
	ResistorValue int `json:"ResistorValue"`
}
