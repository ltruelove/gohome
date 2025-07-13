package models

import "errors"

// swagger:model IpAddress
type IpAddress struct {
	// The IP Address string
	IP string `json:"ip"`
}

func (ip IpAddress) IsValid(checkId bool) (bool, error) {
	if ip.IP == "" {
		return false, errors.New("ip cannot be empty")
	}
	return true, nil
}
