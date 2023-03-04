package models

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

// swagger:model Node
type Node struct {
	// The ID of the Node
	Id int `json:"Id"`
	// The name of the Node
	Name string `json:"Name"`
	// The MAC address of the Node
	Mac string `json:"Mac"`
	// The IP address of the Node
	IpAddress string `json:"IpAddress"`
}

func (item *Node) IsIpAddressValid() error {
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

func (item *Node) IsValid(checkId bool) error {
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
