package temps

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/page"
	"github.com/ltruelove/gohome/internal/pkg/routing"
)

var Config config.Configuration

func RegisterHandlers(mainConfig config.Configuration) {
	Config = mainConfig
	routing.AddGenericRoute("/temps", TempsHandler)
	routing.AddGenericRoute("/kids", HandleKidsSettingsRequest)
	routing.AddGenericRoute("/garage", HandleGarageSettingsRequest)
	routing.AddGenericRoute("/master", HandleMasterSettingsRequest)
	routing.AddGenericRoute("/main", HandleMainSettingsRequest)
}

func TempsHandler(writer http.ResponseWriter, request *http.Request) {
	p := &page.Page{
		Title: "This is the GoHome Room Temperatures Page",
	}
	t, _ := template.ParseFiles(Config.WebDir + "/html/temps.html")
	t.Execute(writer, p)
}

func makeRequest(writer http.ResponseWriter, request *http.Request, address string) {
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

func HandleKidsSettingsRequest(writer http.ResponseWriter, request *http.Request) {
	address := fmt.Sprintf("http://%s", Config.KidsStatusIP)
	makeRequest(writer, request, address)
}

func HandleGarageSettingsRequest(writer http.ResponseWriter, request *http.Request) {
	address := fmt.Sprintf("http://%s", Config.GarageStatusIP)
	makeRequest(writer, request, address)
}

func HandleMasterSettingsRequest(writer http.ResponseWriter, request *http.Request) {
	address := fmt.Sprintf("http://%s", Config.MasterStatusIP)
	makeRequest(writer, request, address)
}

func HandleMainSettingsRequest(writer http.ResponseWriter, request *http.Request) {
	address := fmt.Sprintf("http://%s", Config.MainStatusIP)
	makeRequest(writer, request, address)
}
