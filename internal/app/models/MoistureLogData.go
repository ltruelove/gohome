package models

import "errors"

// swagger:model Node
type MoistureLogData struct {
	// The ID of the Node
	Id int `json:"id"`
	// The ID of the Node
	NodeSensorLogId int `json:"nodeSensorLogId"`
	// The ID of the Node
	Moisture int `json:"Moisture"`
}

func (m MoistureLogData) IsValid(checkId bool) (bool, error) {
	if checkId && m.Id <= 0 {
		return false, errors.New("id cannot be less than or equal to 0")
	}
	if m.NodeSensorLogId <= 0 {
		return false, errors.New("nodeSensorLogId cannot be less than or equal to 0")
	}
	if m.Moisture < 0 {
		return false, errors.New("moisture cannot be negative")
	}
	return true, nil
}
