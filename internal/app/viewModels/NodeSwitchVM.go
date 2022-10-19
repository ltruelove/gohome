package viewModels

import "github.com/ltruelove/gohome/internal/app/models"

type NodeSwitchVM struct {
	Id                     int    `json:"Id"`
	NodeId                 int    `json:"NodeId"`
	SwitchTypeId           int    `json:"SwitchTypeId"`
	Name                   string `json:"Name"`
	SwitchTypeName         string `json:"SwitchTypeName"`
	Pin                    int    `json:"Pin"`
	MomentaryPressDuration int    `json:"MomentaryPressDuration"`
	IsClosedOn             bool   `json:"IsClosedOn"`
}

func (nodeSwitch *NodeSwitchVM) ImportModel(model *models.NodeSwitch) {
	nodeSwitch.Id = model.Id
	nodeSwitch.NodeId = model.NodeId
	nodeSwitch.SwitchTypeId = model.SwitchTypeId
	nodeSwitch.Name = model.Name
	nodeSwitch.Pin = model.Pin
	nodeSwitch.MomentaryPressDuration = model.MomentaryPressDuration
	nodeSwitch.IsClosedOn = model.IsClosedOn
}
