package config

import (
	"strings"
)

type Configuration struct {
	Pin     string `json:"pin"`
	WebDir  string `json:"webDir"`
	LogFile string `json:"logFile"`
	Port    string `json:"port"`
	DbHost  string `json:"dbhost"`
	DbPort  int    `json:"dbport"`
	DbUser  string `json:"dbuser"`
	DbPass  string `json:"dbpass"`
	DbName  string `json:"dbname"`
}

func (c Configuration) ValidatePin(pin string) bool {
	return strings.Compare(pin, c.Pin) == 0
}
