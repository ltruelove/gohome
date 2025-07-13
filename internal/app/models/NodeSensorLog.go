package models

import (
	"errors"
	"time"
)

// swagger:model Node
type NodeSensorLog struct {
	Id                 int               `json:"id"`
	NodeId             int               `json:"nodeId"`
	DateLogged         time.Time         `json:"DateLogged"`
	TemperatureEntries []TempLogData     `json:"TemperatureEntries"`
	MoistureEntries    []MoistureLogData `json:"MoistureEntries"`
	MagneticEntries    []MagneticLogData `json:"MagneticEntries"`
	ResistorEntries    []ResistorLogData `json:"ResistorEntries"`
}

func (n NodeSensorLog) IsValid(checkId bool) (bool, error) {
	isValid := false
	errMsg := ""

	if checkId && n.Id <= 0 {
		isValid = false
		errMsg = "Id cannot be less than or equal to 0"
	}
	if n.NodeId <= 0 {
		isValid = false
		errMsg = "NodeId cannot be less than or equal to 0"
	}
	if n.DateLogged.IsZero() {
		isValid = false
		errMsg = "DateLogged cannot be zero"
	}

	if !isValid {
		if errMsg != "" {
			return false, errors.New(errMsg)
		}
		return false, nil
	}
	return true, nil
}
