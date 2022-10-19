package viewModels

type NodeVM struct {
	Id               int            `json:"Id"`
	Name             string         `json:"Name"`
	Mac              string         `json:"Mac"`
	ControlPointId   int            `json:"controlPointId"`
	ControlPointIP   string         `json:"controlPointIp"`
	ControlPointName string         `json:"controlPointName"`
	Sensors          []NodeSensorVM `json:"sensors"`
	Switches         []NodeSwitchVM `json:"switches"`
}
