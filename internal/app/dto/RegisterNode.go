package dto

import "github.com/ltruelove/gohome/internal/app/models"

type RegsiterNode struct {
	Node         models.Node         `json:"Node"`
	ControlPoint models.ControlPoint `json:"ControlPoint"`
	Sensors      []models.NodeSensor `json:"Sensors"`
	Switches     []models.NodeSwitch `json:"Switches"`
}
