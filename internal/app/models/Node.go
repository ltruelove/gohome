package models

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

type Node struct {
	Id   int    `json:"Id"`
	Name string `json:"Name"`
	Mac  string `json:"Mac"`
}

func (item *Node) IsValid(checkId bool) error {
	var isValid = false
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
