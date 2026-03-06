package handler

import (
	"fmt"
	"log"

	"github.com/ltruelove/gohome/internal/app/dto"
	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/ltruelove/gohome/internal/app/repository"
)

func RegisterNode(dto *dto.RegsiterNode, nodeRepo repository.NodeRepository) error {
	isValid, err := dto.Node.IsValid(false)

	if !isValid || err != nil {
		log.Println("Node Validation error")
		return fmt.Errorf("%w: %s", err, "during node validation")
	}

	// create the node using repository
	if err := nodeRepo.Create(&dto.Node); err != nil {
		log.Println("Error creating node for register")
		return fmt.Errorf("%w: %s", err, "creating node in database")
	}

	updatedSensors := []models.NodeSensor{}
	for _, item := range dto.Sensors {
		item.NodeId = dto.Node.Id

		isValid, err := item.IsValid(false)

		if !isValid || err != nil {
			log.Println("Node sensor validation error")
			return fmt.Errorf("%w: %s", err, "during node sensor validation")
		}

		if err := nodeRepo.CreateNodeSensor(&item); err != nil {
			log.Println("Error creating node sensor for register")
			return fmt.Errorf("%w: %s", err, "creating node sensor in database")
		}

		updatedSensors = append(updatedSensors, item)
	}

	dto.Sensors = updatedSensors

	updatedSwitches := []models.NodeSwitch{}
	for _, item := range dto.Switches {
		item.NodeId = dto.Node.Id

		isValid, err := item.IsValid(false)

		if !isValid || err != nil {
			log.Println("Node switch validation error")
			return fmt.Errorf("%w: %s", err, "during node switch validation")
		}
		if err := nodeRepo.CreateNodeSwitch(&item); err != nil {
			log.Println("Error creating node switch for register")
			return fmt.Errorf("%w: %s", err, "creating node switch in database")
		}

		updatedSwitches = append(updatedSwitches, item)
	}

	dto.Switches = updatedSwitches

	return nil
}
