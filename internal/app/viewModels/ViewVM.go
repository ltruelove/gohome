package viewModels

type ViewVM struct {
	Id       int                `json:"Id"`
	Name     string             `json:"Name"`
	Sensors  []ViewNodeSensorVM `json:"sensors"`
	Switches []ViewNodeSwitchVM `json:"switches"`
}
