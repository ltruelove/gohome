package pin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/ltruelove/gohome/config"
)

var Config config.Configuration

type PinRequest struct {
	PinCode string `json:"pinCode"`
}

type Validator struct {
	IsValid bool
}

func RegisterHandlers(router *mux.Router, mainConfig config.Configuration) {
	Config = mainConfig
	router.HandleFunc("/pinValid", PinValid).Methods("POST")
}

func PinValid(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var t PinRequest

	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	var v = new(Validator)
	if strings.Compare(t.PinCode, Config.Pin) == 0 {
		v.IsValid = true
	}

	// Moved the click functionality to here so the IP of the module wouldn't be publicly
	//available
	if v.IsValid {
		address := fmt.Sprintf("http://%s/click", Config.DoorIp)
		http.Get(address)
	}

	pinresponse, _ := json.Marshal(v)
	writer.Write(pinresponse)
}
