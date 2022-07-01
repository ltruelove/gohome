package models

type SensorTypeData struct {
	Id           int    `json:"Id"`
	SensorTypeId int    `json:"SensorTypeId"`
	Name         string `json:"Name"`
	ValueType    string `json:"ValueType"`
}
