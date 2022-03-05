package data

type RoomTemperature struct {
	Fahrenheit   float32 `json:"fahrenheit"`
	Celcius      float32 `json:"celcius"`
	Humidity     float32 `json:"humidity"`
	Name         string  `json:"name"`
	ErrorMessage string  `json:"errorMessage"`
}
