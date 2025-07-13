package models

import "errors"

type SensorTypeData struct {
	Id           int    `json:"Id"`
	SensorTypeId int    `json:"SensorTypeId"`
	Name         string `json:"Name"`
	ValueType    string `json:"ValueType"`
}

func (s SensorTypeData) IsValid(checkId bool) (bool, error) {
	if checkId && s.Id <= 0 {
		return false, errors.New("id must be positive")
	}
	if s.SensorTypeId <= 0 {
		return false, errors.New("SensorTypeId must be positive")
	}
	if s.Name == "" {
		return false, errors.New("name cannot be empty")
	}
	if s.ValueType == "" {
		return false, errors.New("ValueType cannot be empty")
	}
	return true, nil
}
