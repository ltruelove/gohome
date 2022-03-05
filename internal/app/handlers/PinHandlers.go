package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data"
	"github.com/ltruelove/gohome/internal/pkg/routing"
)

func RegisterPinHandlers(mainConfig config.Configuration) {
	Config = mainConfig
	routing.AddRouteWithMethod("/pinValid", "POST", PinValid)
}

func PinValid(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var t data.PinRequest

	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	var v = new(data.Validator)
	v.IsValid = Config.ValidatePin(t.PinCode)

	if !v.IsValid {
		http.Error(writer, "Not valid", 401)
		return
	}

	pinresponse, _ := json.Marshal(v)
	writer.Write(pinresponse)
}
