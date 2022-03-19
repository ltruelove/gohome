package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/data"
	"github.com/ltruelove/gohome/internal/app/providers"
	"github.com/ltruelove/gohome/internal/pkg/routing"
)

type TempHandler struct {
	DB *sql.DB
}

func (tempHandler *TempHandler) RegisterTempHandlers(mainConfig config.Configuration) {
	Config = mainConfig
	routing.AddGenericRoute("/temps", TempsHandler)
	routing.AddGenericRoute("/temps/kids", HandleKidsSettingsRequest)
	routing.AddGenericRoute("/temps/garage", HandleGarageSettingsRequest)
	routing.AddGenericRoute("/temps/master", HandleMasterSettingsRequest)
	routing.AddGenericRoute("/temps/main", HandleMainSettingsRequest)
	routing.AddGenericRoute("/temps/all", HandleAllSettingsRequest)
	routing.AddRouteWithMethod("/temps/registersensor", "POST", tempHandler.RegisterTempSensor)
	routing.AddRouteWithMethod("/temps/updatesensor", "PUT", tempHandler.UpdateTempSensor)
	routing.AddRouteWithMethod("/temps/deletesensor", "DELETE", tempHandler.DeleteTempSensor)
}

func TempsHandler(writer http.ResponseWriter, request *http.Request) {
	p := &data.Page{
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

func fetchTemperature(address string, name string) data.RoomTemperature {
	status, err := http.Get(address)

	var t data.RoomTemperature
	t.Name = name

	if err != nil {
		t.ErrorMessage = "There was an error fetching the temperature"
		return t
	}

	defer status.Body.Close()
	responseData, rErr := ioutil.ReadAll(status.Body)

	if rErr != nil {
		t.ErrorMessage = "There was an error fetching the temperature"
		return t
	}

	err = json.Unmarshal(responseData, &t)
	if err != nil {
		t.ErrorMessage = "There was an error fetching the temperature"
		return t
	}

	return t
}

func (tempHandler *TempHandler) RegisterTempSensor(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var sensor data.TemperatureSensor

	err := decoder.Decode(&sensor)
	if err != nil {
		panic(err)
	}

	sensor.SensorId = uuid.NewString()

	fmt.Println((sensor.SensorId))
	if providers.VerifyTemperatureSensorIdIsNew(sensor.SensorId, tempHandler.DB) {
		fmt.Println("Add a new sensor")
		providers.AddNewTemperatureSensor(&sensor, tempHandler.DB)

		result, err := json.Marshal(sensor)
		if err != nil {
			panic(err)
		}

		fmt.Println(result)
		writeResponse(writer, result)
	}
}

func (tempHandler *TempHandler) UpdateTempSensor(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var sensor data.TemperatureSensor

	err := decoder.Decode(&sensor)
	if err != nil {
		panic(err)
	}

	providers.UpdateTemperatureSensor(&sensor, tempHandler.DB)
}

func (tempHandler *TempHandler) DeleteTempSensor(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var sensor data.TemperatureSensor

	err := decoder.Decode(&sensor)
	if err != nil {
		panic(err)
	}

	providers.DeleteTemperatureSensor(sensor.SensorId, tempHandler.DB)
}

func HandleAllSettingsRequest(writer http.ResponseWriter, request *http.Request) {
	allTemps := make([]data.RoomTemperature, 0)
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
