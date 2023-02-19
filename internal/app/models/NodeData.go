package models

// swagger:model Node
type NodeData struct {
	// The ID of the Node
	NodeId int `json:"nodeId"`
	// The F temp reading
	TemperatureF float32 `json:"TemperatureF"`
	// The C temp reading
	TemperatureC float32 `json:"TemperatureC"`
	// The humidity reading
	Humidity float32 `json:"Humidity"`
	// The moisture reading
	Moisture int `json:"Moisture"`
	// The resistor reading
	ResistorValue int `json:"ResistorValue"`
	// The resistor reading
	IsClosed bool `json:"IsClosed"`
	// The magnetic reading
	MagneticValue bool `json:"MagneticValue"`
}
