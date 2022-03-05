package data

type TemperatureSensor struct {
	SensorId  string `json:"sensorId"`
	Name      string `json:"name"`
	IsGarage  int    `json:"isGarage"`
	IpAddress string `json:"ipAddress"`
}
