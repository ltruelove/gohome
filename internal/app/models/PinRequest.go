package models

import "errors"

type PinRequest struct {
	PinCode string `json:"pinCode"`
}

func (p PinRequest) IsValid(checkId bool) (bool, error) {
	if p.PinCode == "" {
		return false, errors.New("the PinCode cannot be empty")
	}
	return true, nil
}
