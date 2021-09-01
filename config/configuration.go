package config

import (
	"strings"
)

type Configuration struct {
	Pin            string `json:"pin"`
	DoorIp         string `json:"doorIp"`
	GarageStatusIP string `json:"garageStatusIp"`
	KidsStatusIP   string `json:"kidsStatusIp"`
	MasterStatusIP string `json:"masterStatusIp"`
	MainStatusIP   string `json:"mainStatusIp"`
	GardenIp       string `json:"gardenIp"`
	WaterIp        string `json:"waterIp"`
	SoilThreshold  int    `json:"soilThreshold"`
	TickerActive   bool   `json:"tickerActive"`
	WebDir         string `json:"webDir"`
}

func (c Configuration) ValidatePin(pin string) bool {
	return strings.Compare(pin, c.Pin) == 0
}
