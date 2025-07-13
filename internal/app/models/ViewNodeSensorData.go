package models

import "errors"

type ViewNodeSensorData struct {
	Id           int    `json:"Id"`
	NodeId       int    `json:"NodeId"`
	ViewId       int    `json:"ViewId"`
	NodeSensorId int    `json:"NodeSensorId"`
	Name         string `json:"Name"`
}

func (v ViewNodeSensorData) IsValid(checkId bool) (bool, error) {
	isValid := true
	errMsg := ""

	if checkId && v.Id <= 0 {
		isValid = false
		errMsg = "the Id cannot be less than or equal to 0"
	}
	if v.NodeId <= 0 {
		isValid = false
		errMsg = "the NodeId cannot be less than or equal to 0"
	}
	if v.ViewId <= 0 {
		isValid = false
		errMsg = "the ViewId cannot be less than or equal to 0"
	}
	if v.NodeSensorId <= 0 {
		isValid = false
		errMsg = "the NodeSensorId cannot be less than or equal to 0"
	}
	if v.Name == "" {
		isValid = false
		errMsg = "the Name cannot be empty"
	}

	if !isValid {
		if errMsg != "" {
			return false, errors.New(errMsg)
		}
		return false, nil
	}
	return true, nil
}
