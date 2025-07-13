package models

import "errors"

// swagger:model Node
type ResistorLogData struct {
	// The ID of the Node
	Id int `json:"id"`
	// The ID of the Node
	NodeSensorLogId int `json:"nodeSensorLogId"`
	// The ID of the Node
	ResistorValue int `json:"ResistorValue"`
}

func (r ResistorLogData) IsValid(checkId bool) (bool, error) {
	if checkId && r.Id <= 0 {
		return false, errors.New("the Id cannot be less than or equal to 0")
	}
	if r.NodeSensorLogId <= 0 {
		return false, errors.New("the NodeSensorLogId cannot be less than or equal to 0")
	}
	if r.ResistorValue < 0 {
		return false, errors.New("the ResistorValue cannot be negative")
	}
	return true, nil
}
