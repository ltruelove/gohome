package handler

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ltruelove/gohome/internal/app/data"
	"github.com/ltruelove/gohome/internal/app/dto"
	"github.com/ltruelove/gohome/internal/app/models"
)

func RegisterNode(dto *dto.RegsiterNode, db *sql.DB) error {
	isValid, err := dto.Node.IsValid(false)

	if !isValid || err != nil {
		log.Println("Node Validation error")
		return fmt.Errorf("%w: %s", err, "during node validation")
	}

	err = data.CreateNode(&dto.Node, db)

	if err != nil {
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

		err = data.CreateNodeSensor(&item, db)

		if err != nil {
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
		err = data.CreateNodeSwitch(&item, db)

		if err != nil {
			log.Println("Error creating node switch for register")
			return fmt.Errorf("%w: %s", err, "creating node switch in database")
		}

		updatedSwitches = append(updatedSwitches, item)
	}

	dto.Switches = updatedSwitches

	return nil
}
