package models

import "errors"

type SwitchType struct {
	Id   int    `json:"Id"`
	Name string `json:"Name"`
}

func (s SwitchType) IsValid(checkId bool) (bool, error) {
	if checkId && s.Id <= 0 {
		return false, errors.New("invalid id")
	}
	if s.Name == "" {
		return false, errors.New("invalid name")
	}
	return true, nil
}
