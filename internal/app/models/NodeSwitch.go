package models

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

type NodeSwitch struct {
	Id                     int    `json:"Id"`
	NodeId                 int    `json:"NodeId"`
	SwitchTypeId           int    `json:"SwitchTypeId"`
	Name                   string `json:"Name"`
	Pin                    string `json:"Pin"`
	MomentaryPressDuration int    `json:"MomentaryPressDuration"`
	IsClosedOn             bool   `json:"IsClosedOn"`
}

func (item *NodeSwitch) IsValid(checkId bool) error {
	var isValid = false
	var validationMessage = ""
	var err error = nil

	if checkId {
		if item.Id < 1 {
			validationMessage = fmt.Sprintf("%s", "Id cannot be less than 1")
			isValid = false
		}
	}

	if item.NodeId < 1 {
		if len(validationMessage) > 0 {
			validationMessage = fmt.Sprintf("%s, %s", validationMessage, "NodeId cannot be less than 1")
		} else {
			validationMessage = fmt.Sprintf("%s", "NodeId cannot be less than 1")
		}
		isValid = false
	}

	if item.SwitchTypeId < 1 {
		if len(validationMessage) > 0 {
			validationMessage = fmt.Sprintf("%s, %s", validationMessage, "SensorTypeId cannot be less than 1")
		} else {
			validationMessage = fmt.Sprintf("%s", "SensorTypeId cannot be less than 1")
		}
		isValid = false
	}

	if strings.TrimSpace(item.Name) == "" {
		if len(validationMessage) > 0 {
			validationMessage = fmt.Sprintf("%s, %s", validationMessage, "Name cannot be empty")
		} else {
			validationMessage = fmt.Sprintf("%s", "Name cannot be empty")
		}
		isValid = false
	}

	if strings.TrimSpace(item.Pin) == "" {
		if len(validationMessage) > 0 {
			validationMessage = fmt.Sprintf("%s, %s", validationMessage, "Pin cannot be empty")
		} else {
			validationMessage = fmt.Sprintf("%s", "Pin cannot be empty")
		}
		isValid = false
	}

	if !isValid {
		log.Println(validationMessage)
		err = errors.New(validationMessage)
	}

	return err
}
