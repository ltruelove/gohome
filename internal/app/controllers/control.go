package controllers

import (
	"net/http"

	"github.com/ltruelove/gohome/config"
)

var Config config.Configuration

func writeResponse(writer http.ResponseWriter, responseData []byte) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(responseData)
}
