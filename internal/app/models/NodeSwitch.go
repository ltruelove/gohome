package models

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

// swagger:model NodeSwitch
type NodeSwitch struct {
	// The ID of the node switch
	Id int `json:"Id"`
	// The ID of the node the switch belongs to
	NodeId int `json:"NodeId"`
	// The ID of the type of switch the node switch is
	SwitchTypeId int `json:"SwitchTypeId"`
	// The name of the node switch
	Name string `json:"Name"`
	// The I/O pin the switch is attached to
	Pin int `json:"Pin"`
	// The time in milliseconds to hold a momentary button down
	MomentaryPressDuration int `json:"MomentaryPressDuration"`
	// Is a closed circuit considered "on"
	IsClosedOn bool `json:"IsClosedOn"`
}

func (item *NodeSwitch) IsValid(checkId bool) error {
	var isValid = true
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

	if item.Pin < 0 {
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
