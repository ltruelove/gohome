package dto

type ControlPointNode struct {
	Id         int    `json:"Id"`
	Name       string `json:"Name"`
	Mac        string `json:"Mac"`
	RelationId int    `json:"RelationId"`
}
