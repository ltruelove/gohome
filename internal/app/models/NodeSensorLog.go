package models

import (
	"errors"
	"fmt"
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
	isValid := true
	errMsg := ""

	if checkId && n.Id <= 0 {
		isValid = false
		if errMsg != "" {
			errMsg = fmt.Sprintf("%s, %s", errMsg, "Id cannot be less than or equal to 0")
		} else {
			errMsg = "Id cannot be less than or equal to 0"
		}
	}
	if n.NodeId <= 0 {
		isValid = false
		if errMsg != "" {
			errMsg = fmt.Sprintf("%s, %s", errMsg, "NodeId cannot be less than or equal to 0")
		} else {
			errMsg = "NodeId cannot be less than or equal to 0"
		}
	}
	if n.DateLogged.IsZero() {
		isValid = false
		if errMsg != "" {
			errMsg = fmt.Sprintf("%s, %s", errMsg, "DateLogged cannot be zero")
		} else {
			errMsg = "DateLogged cannot be zero"
		}
	}

	if !isValid {
		if errMsg != "" {
			return false, errors.New(errMsg)
		}
		return false, nil
	}
	return true, nil
}
