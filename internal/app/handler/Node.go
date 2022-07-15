package handler

import (
	"database/sql"
	"log"

	"github.com/ltruelove/gohome/internal/app/data"
	"github.com/ltruelove/gohome/internal/app/dto"
)

func RegisterNode(dto *dto.RegsiterNode, db *sql.DB) error {
	err := dto.Node.IsValid(false)

	if err != nil {
		log.Println("Node Validation error")
		return err
	}

	err = data.CreateNode(&dto.Node, db)

	if err != nil {
		log.Println("Error creating node for register")
		return err
	}

	for _, item := range dto.Sensors {
		err = item.IsValid(false)

		if err != nil {
			log.Println("Node sensor validation error")
			return err
		}

		err = data.CreateNodeSensor(&item, db)

		if err != nil {
			log.Println("Error creating node sensor for register")
			return err
		}
	}

	for _, item := range dto.Switches {
		err = item.IsValid(false)

		if err != nil {
			log.Println("Node switch validation error")
			return err
		}
		err = data.CreateNodeSwitch(&item, db)

		if err != nil {
			log.Println("Error creating node switch for register")
			return err
		}
	}

	return nil
}
