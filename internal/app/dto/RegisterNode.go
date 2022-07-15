package dto

import "github.com/ltruelove/gohome/internal/app/models"

type RegsiterNode struct {
	Node     models.Node         `json:"Node"`
	Sensors  []models.NodeSensor `json:"Sensors"`
	Switches []models.NodeSwitch `json:"Switches"`
}
