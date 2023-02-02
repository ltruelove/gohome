package models

// swagger:model ControlPointNode
type ControlPointNode struct {
	// Id of the Control Point Node
	Id int `json:"Id"`
	// Id of the Control Point
	ControlPointId int `json:"ControlPointId"`
	// Id of the Node
	NodeId int `json:"NodeId"`
}
