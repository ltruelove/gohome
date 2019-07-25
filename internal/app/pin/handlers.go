package pin

import (
	"encoding/json"
	"net/http"
	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/pkg/routing"
)

var Config config.Configuration

type PinRequest struct {
	PinCode string `json:"pinCode"`
}

type Validator struct {
	IsValid bool
}

func RegisterHandlers(mainConfig config.Configuration) {
	Config = mainConfig
	routing.AddRouteWithMethod("/pinValid", "POST", PinValid)
}

func PinValid(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var t PinRequest

	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	var v = new(Validator)
	v.IsValid = Config.ValidatePin(t.PinCode)

	if !v.IsValid {
		http.Error(writer, "Not valid", 401)
		return
	}

	pinresponse, _ := json.Marshal(v)
	writer.Write(pinresponse)
}
