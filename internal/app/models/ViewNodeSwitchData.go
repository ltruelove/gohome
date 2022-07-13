package models

type ViewNodeSwitchData struct {
	Id           int    `json:"Id"`
	NodeId       int    `json:"NodeId"`
	ViewId       int    `json:"ViewId"`
	NodeSwitchId int    `json:"NodeSwitchId"`
	Name         string `json:"Name"`
}
