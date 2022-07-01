package models

type NodeSwitch struct {
	Id                     int    `json:"Id"`
	NodeId                 int    `json:"NodeId"`
	SwitchTypeId           int    `json:"SwitchTypeId"`
	Name                   string `json:"Name"`
	Pin                    string `json:"Pin"`
	MomentaryPressDuration int    `json:"MomentaryPressDuration"`
	IsClosedOn             bool   `json:"IsClosedOn"`
}
