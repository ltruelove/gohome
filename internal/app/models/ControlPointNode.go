package models

import "errors"

// swagger:model ControlPointNode
type ControlPointNode struct {
	// Id of the Control Point Node
	Id int `json:"Id"`
	// Id of the Control Point
	ControlPointId int `json:"ControlPointId"`
	// Id of the Node
	NodeId int `json:"NodeId"`
}

func (c ControlPointNode) IsValid(checkId bool) (bool, error) {
	if checkId && c.Id <= 0 {
		return false, errors.New("id cannot be less than or equal to 0")
	}
	if c.ControlPointId <= 0 {
		return false, errors.New("ControlPointId cannot be less than or equal to 0")
	}
	if c.NodeId <= 0 {
		return false, errors.New("NodeId cannot be less than or equal to 0")
	}
	return true, nil
}
