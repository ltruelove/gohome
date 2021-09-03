package temps

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/page"
	"github.com/ltruelove/gohome/internal/pkg/routing"
)

var Config config.Configuration

type RoomTemperature struct {
	Fahrenheit float32 `json:"fahrenheit"`
	Celcius    float32 `json:"celcius"`
	Humidity   float32 `json:"humidity"`
	Name       string  `json:"name"`
}

func RegisterHandlers(mainConfig config.Configuration) {
	Config = mainConfig
	routing.AddGenericRoute("/temps", TempsHandler)
	routing.AddGenericRoute("/temps/kids", HandleKidsSettingsRequest)
	routing.AddGenericRoute("/temps/garage", HandleGarageSettingsRequest)
	routing.AddGenericRoute("/temps/master", HandleMasterSettingsRequest)
	routing.AddGenericRoute("/temps/main", HandleMainSettingsRequest)
	routing.AddGenericRoute("/temps/all", HandleAllSettingsRequest)
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

	writeResponse(writer, responseData)
}

func writeResponse(writer http.ResponseWriter, responseData []byte) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(responseData)
}

func fetchTemperature(address string, name string) RoomTemperature {
	status, err := http.Get(address)

	if err != nil {
		panic(err)
	}

	defer status.Body.Close()
	responseData, rErr := ioutil.ReadAll(status.Body)

	if rErr != nil {
		panic(rErr)
	}

	var t RoomTemperature

	err = json.Unmarshal(responseData, &t)
	if err != nil {
		panic(err)
	}

	t.Name = name

	return t
}

func HandleAllSettingsRequest(writer http.ResponseWriter, request *http.Request) {
	allTemps := make([]RoomTemperature, 0)
	kidsTemp := fetchTemperature(fmt.Sprintf("http://%s", Config.KidsStatusIP), "Kids' Room")
	mainTemp := fetchTemperature(fmt.Sprintf("http://%s", Config.MainStatusIP), "Main Floor")
	masterTemp := fetchTemperature(fmt.Sprintf("http://%s", Config.MasterStatusIP), "Master Bedroom")
	garageTemp := fetchTemperature(fmt.Sprintf("http://%s", Config.GarageStatusIP), "Garage")

	allTemps = append(allTemps, kidsTemp)
	allTemps = append(allTemps, mainTemp)
	allTemps = append(allTemps, masterTemp)
	allTemps = append(allTemps, garageTemp)

	result, err := json.Marshal(allTemps)

	if err != nil {
		panic(err)
	}

	writeResponse(writer, result)
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
