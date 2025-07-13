package models

import "errors"

type RoomTemperature struct {
	Fahrenheit   float32 `json:"fahrenheit"`
	Celcius      float32 `json:"celcius"`
	Humidity     float32 `json:"humidity"`
	Name         string  `json:"name"`
	ErrorMessage string  `json:"errorMessage"`
}

func (r RoomTemperature) IsValid() (bool, error) {
	if r.Fahrenheit < -459.67 {
		return false, errors.New("the value for Fahrenheit cannot be less than absolute zero")
	}
	if r.Celcius < -273.15 {
		return false, errors.New("the value for Celcius cannot be less than absolute zero")
	}
	if r.Humidity < 0 || r.Humidity > 100 {
		return false, errors.New("the value for Humidity must be between 0 and 100")
	}
	if r.Name == "" {
		return false, errors.New("the value for Name cannot be empty")
	}

	return true, nil
}
