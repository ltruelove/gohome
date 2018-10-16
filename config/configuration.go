package config

type Configuration struct {
	Pin           string `json:"pin"`
	DoorIp        string `json:"doorIp"`
	GardenIp      string `json:"gardenIp"`
	WaterIp       string `json:"waterIp"`
	SoilThreshold int    `json:"soilThreshold"`
	TickerActive  bool   `json:"tickerActive"`
}
