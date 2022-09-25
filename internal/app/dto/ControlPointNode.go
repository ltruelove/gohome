package dto

type ControlPointNode struct {
	Id                 int    `json:"Id"`
	Name               string `json:"Name"`
	Mac                string `json:"Mac"`
	ControlPointNodeId int    `json: "ControlPointNodeId"`
}
