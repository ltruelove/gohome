package viewModels

type ViewNodeSensorVM struct {
	Id             int    `json:"Id"`
	NodeId         int    `json:"NodeId"`
	ViewId         int    `json:"ViewId"`
	NodeSensorId   int    `json:"NodeSensorId"`
	Name           string `json:"Name"`
	NodeName       string `json:"NodeName"`
	SensorName     string `json:"SensorName"`
	SensorTypeName string `json:"SensorTypeName"`
}
