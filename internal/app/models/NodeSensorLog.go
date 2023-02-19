package models

import "time"

// swagger:model Node
type NodeSensorLog struct {
	Id                 int               `json:"id"`
	NodeId             int               `json:"nodeId"`
	DateLogged         time.Time         `json:"DateLogged"`
	TemperatureEntries []TempLogData     `json:"TemperatureEntries"`
	MoistureEntries    []MoistureLogData `json:"MoistureEntries"`
	MagneticEntries    []MagneticLogData `json:"MagneticEntries"`
	ResistorEntries    []ResistorLogData `json:"ResistorEntries"`
}
