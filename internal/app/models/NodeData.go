package models

import "errors"

// swagger:model Node
type NodeData struct {
	// The ID of the Node
	NodeId int `json:"nodeId"`
	// The F temp reading
	TemperatureF float32 `json:"TemperatureF"`
	// The C temp reading
	TemperatureC float32 `json:"TemperatureC"`
	// The humidity reading
	Humidity float32 `json:"Humidity"`
	// The moisture reading
	Moisture int `json:"Moisture"`
	// The resistor reading
	ResistorValue int `json:"ResistorValue"`
	// The resistor reading
	IsClosed bool `json:"IsClosed"`
	// The magnetic reading
	MagneticValue bool `json:"MagneticValue"`
}

func (n NodeData) IsValid(checkId bool) (bool, error) {
	isValid := false
	errMsg := ""

	if checkId && n.NodeId <= 0 {
		isValid = false
		errMsg = "NodeId cannot be less than or equal to 0"
	}
	if n.TemperatureF < -459.67 {
		isValid = false
		errMsg = "TemperatureF cannot be less than absolute zero"
	}
	if n.TemperatureC < -273.15 {
		isValid = false
		errMsg = "TemperatureC cannot be less than absolute zero"
	}
	if n.Humidity < 0 || n.Humidity > 100 {
		isValid = false
		errMsg = "Humidity must be between 0 and 100"
	}
	if n.Moisture < 0 || n.Moisture > 100 {
		isValid = false
		errMsg = "Moisture must be between 0 and 100"
	}
	if n.ResistorValue < 0 {
		isValid = false
		errMsg = "ResistorValue cannot be negative"
	}

	if !isValid {
		if errMsg != "" {
			return false, errors.New(errMsg)
		}
		return false, nil
	}

	return true, nil
}
