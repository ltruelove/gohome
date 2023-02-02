package models

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

// swagger:model NodeSensor
type NodeSensor struct {
	// The ID of the node sensor
	Id int `json:"Id"`
	// The ID of the node the sensor is attached to
	NodeId int `json:"NodeId"`
	// The ID of the type of sensor the node sensor is
	SensorTypeId int `json:"SensorTypeId"`
	// The name of the node sensor
	Name string `json:"Name"`
	// The I/O pin the sensor is attached to
	Pin int `json:"Pin"`
	// The type of DHT sensor if it's a DHT sensor (11 for 11, 22 for 22)
	DHTType int `json:"DHTType"`
}

func (item *NodeSensor) IsValid(checkId bool) error {
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

	if item.SensorTypeId < 1 {
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
