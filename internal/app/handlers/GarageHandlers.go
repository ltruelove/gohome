package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data"
	"github.com/ltruelove/gohome/internal/pkg/routing"
)

func RegisterGarageHandlers(mainConfig config.Configuration) {
	Config = mainConfig
	routing.AddGenericRoute("/door", DoorHandler)
	routing.AddRouteWithMethod("/clickGarageDoorButton", "POST", ClickGarageDoorButton)
}

func DoorHandler(writer http.ResponseWriter, request *http.Request) {
	p := &data.Page{
		Title: "This is the GoHome Door Page",
	}
	t, _ := template.ParseFiles(Config.WebDir + "/html/door.html")
	t.Execute(writer, p)
}

func ClickGarageDoorButton(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var t data.PinRequest

	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	var v = new(data.Validator)
	if Config.ValidatePin(t.PinCode) {
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
