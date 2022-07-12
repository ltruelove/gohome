package models

type ControlPoint struct {
	Id        int    `json:"Id"`
	Name      string `json:"Name"`
	IpAddress string `json:"IpAddress"`
	Mac       string `json:"Mac"`
}
