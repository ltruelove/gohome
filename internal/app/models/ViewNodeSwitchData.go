package models

import "errors"

type ViewNodeSwitchData struct {
	Id           int    `json:"Id"`
	NodeId       int    `json:"NodeId"`
	ViewId       int    `json:"ViewId"`
	NodeSwitchId int    `json:"NodeSwitchId"`
	Name         string `json:"Name"`
}

func (v ViewNodeSwitchData) IsValid(checkId bool) (bool, error) {
	if checkId && v.Id <= 0 {
		return false, errors.New("invalid id")
	}
	if v.NodeId <= 0 {
		return false, errors.New("invalid node id")
	}
	if v.ViewId <= 0 {
		return false, errors.New("invalid view id")
	}
	if v.NodeSwitchId <= 0 {
		return false, errors.New("invalid node switch id")
	}
	if v.Name == "" {
		return false, errors.New("invalid name")
	}
	return true, nil
}
