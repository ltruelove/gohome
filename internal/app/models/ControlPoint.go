package models

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

// swagger:model ControlPoint
type ControlPoint struct {
	// Id of the Control Point
	Id int `json:"Id"`
	// Name of the Control Point
	Name string `json:"Name"`
	// IP Address of the Control Point
	IpAddress string `json:"IpAddress"`
	// MAC Address of the Control Point
	Mac string `json:"Mac"`
}

func (item *ControlPoint) IsValid(checkId bool) error {
	var isValid = true
	var validationMessage = ""
	var err error = nil

	if checkId {
		if item.Id < 1 {
			validationMessage = fmt.Sprintf("%s", "Id cannot be less than 1")
			isValid = false
		}
	}

	if strings.TrimSpace(item.Name) == "" {
		if len(validationMessage) > 0 {
			validationMessage = fmt.Sprintf("%s, %s", validationMessage, "Name cannot be empty")
		} else {
			validationMessage = fmt.Sprintf("%s", "Name cannot be empty")
		}
		isValid = false
	}

	if strings.TrimSpace(item.IpAddress) == "" {
		if len(validationMessage) > 0 {
			validationMessage = fmt.Sprintf("%s, %s", validationMessage, "IpAddress cannot be empty")
		} else {
			validationMessage = fmt.Sprintf("%s", "IpAddress cannot be empty")
		}
		isValid = false
	}

	if strings.TrimSpace(item.Mac) == "" {
		if len(validationMessage) > 0 {
			validationMessage = fmt.Sprintf("%s, %s", validationMessage, "Mac cannot be empty")
		} else {
			validationMessage = fmt.Sprintf("%s", "Mac cannot be empty")
		}
		isValid = false
	}

	if !isValid {
		log.Println(validationMessage)
		err = errors.New(validationMessage)
	}

	return err
}

func (item *ControlPoint) IsIpAddressValid() error {
	var isValid = true
	var validationMessage = ""
	var err error = nil

	if strings.TrimSpace(item.IpAddress) == "" {
		if len(validationMessage) > 0 {
			validationMessage = fmt.Sprintf("%s, %s", validationMessage, "IpAddress cannot be empty")
		} else {
			validationMessage = fmt.Sprintf("%s", "IpAddress cannot be empty")
		}
		isValid = false
	}

	if !isValid {
		log.Println(validationMessage)
		err = errors.New(validationMessage)
	}

	return err
}
