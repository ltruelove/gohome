package models

import "errors"

type SensorType struct {
	Id       int    `json:"Id"`
	TypeName string `json:"TypeName"`
}

func (s SensorType) IsValid(checkId bool) (bool, error) {
	if checkId && s.Id <= 0 {
		return false, errors.New("id must be positive")
	}
	if s.TypeName == "" {
		return false, errors.New("typeName cannot be empty")
	}
	return true, nil
}
