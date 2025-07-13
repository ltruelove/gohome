package models

import "errors"

// swagger:model Node
type MagneticLogData struct {
	// The ID of the Node
	Id int `json:"id"`
	// The ID of the Node
	NodeSensorLogId int `json:"nodeSensorLogId"`
	// The ID of the Node
	IsClosed bool `json:"IsClosed"`
}

func (m MagneticLogData) IsValid(checkId bool) (bool, error) {
	if checkId && m.Id <= 0 {
		return false, errors.New("id cannot be less than or equal to 0")
	}
	if m.NodeSensorLogId <= 0 {
		return false, errors.New("nodeSensorLogId cannot be less than or equal to 0")
	}

	return true, nil
}
