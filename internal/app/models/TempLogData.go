package models

// swagger:model Node
type TempLogData struct {
	// The ID of the Node
	Id int `json:"id"`
	// The ID of the Node
	NodeSensorLogId int `json:"nodeSensorLogId"`
	// The F temp reading
	TemperatureF float32 `json:"TemperatureF"`
	// The C temp reading
	TemperatureC float32 `json:"TemperatureC"`
	// The humidity reading
	Humidity float32 `json:"Humidity"`
}
