package models

type ControllerNode struct {
	Id               int `json:"Id"`
	NodeId           int `json:"NodeId"`
	NodeControllerId int `json:"NodeControllerId"`
}
