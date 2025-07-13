package models

import "errors"

// swagger:model Node
type TempLogData struct {
	// The ID of the Node
	Id int `json:"id"`
	// The ID of the Node
	NodeSensorLogId int `json:"nodeSensorLogId"`
	// The F temp reading
	TemperatureF float32 `json:"TemperatureF"`
	// The C temp reading
	TemperatureC float32 `json:"TemperatureC"`
	// The humidity reading
	Humidity float32 `json:"Humidity"`
}

func (t TempLogData) IsValid(checkId bool) (bool, error) {
	isValid := true
	errMsg := ""

	if checkId && t.Id <= 0 {
		isValid = false
		errMsg = "the Id cannot be less than or equal to 0"
	}
	if t.NodeSensorLogId <= 0 {
		isValid = false
		errMsg = "the NodeSensorLogId cannot be less than or equal to 0"
	}
	if t.TemperatureF < -459.67 {
		isValid = false
		errMsg = "the TemperatureF cannot be less than absolute zero"
	}
	if t.TemperatureC < -273.15 {
		isValid = false
		errMsg = "the TemperatureC cannot be less than absolute zero"
	}
	if t.Humidity < 0 || t.Humidity > 100 {
		isValid = false
		errMsg = "the Humidity must be between 0 and 100"
	}

	if !isValid {
		if errMsg != "" {
			return false, errors.New(errMsg)
		}
		return false, nil
	}
	return true, nil
}
