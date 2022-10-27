package viewModels

type ViewNodeSwitchVM struct {
	Id             int    `json:"Id"`
	NodeId         int    `json:"NodeId"`
	ViewId         int    `json:"ViewId"`
	NodeSwitchId   int    `json:"NodeSwitchId"`
	Name           string `json:"Name"`
	NodeName       string `json:"NodeName"`
	SwitchName     string `json:"SwitchName"`
	SwitchTypeName string `json:"SwitchTypeName"`
}
