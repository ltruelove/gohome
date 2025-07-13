package models

import "errors"

type View struct {
	Id   int    `json:"Id"`
	Name string `json:"Name"`
}

func (v View) IsValid(checkId bool) (bool, error) {
	isValid := true
	errMsg := ""

	if checkId && v.Id <= 0 {
		isValid = false
		errMsg = "the Id cannot be less than or equal to 0"
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
