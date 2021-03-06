package garage

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/page"
	"github.com/ltruelove/gohome/internal/app/pin"
	"github.com/ltruelove/gohome/internal/pkg/routing"
)

var Config config.Configuration

func RegisterHandlers(mainConfig config.Configuration) {
	Config = mainConfig
	routing.AddGenericRoute("/door", DoorHandler)
	routing.AddGenericRoute("/doorStatus", HandleSettingsRequest)
	routing.AddRouteWithMethod("/clickGarageDoorButton", "POST", ClickGarageDoorButton)
}

func DoorHandler(writer http.ResponseWriter, request *http.Request) {
	p := &page.Page{
		Title:    "This is the GoHome Door Page",
		StatusIP: Config.GarageStatusIP,
	}
	t, _ := template.ParseFiles(Config.WebDir + "/html/door.html")
	t.Execute(writer, p)
}

func HandleSettingsRequest(writer http.ResponseWriter, request *http.Request) {
	address := fmt.Sprintf("http://%s", Config.GarageStatusIP)
	status, err := http.Get(address)

	if err != nil {
		panic(err)
	}

	defer status.Body.Close()
	responseData, rErr := ioutil.ReadAll(status.Body)

	if rErr != nil {
		panic(rErr)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(responseData)
}

func ClickGarageDoorButton(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var t pin.PinRequest

	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	var v = new(pin.Validator)
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
