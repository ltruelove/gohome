package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/ltruelove/gohome/internal/app/services"
	"github.com/ltruelove/gohome/internal/pkg/routing"
)

type TempController struct {
	DB *sql.DB
}

func (tempController *TempController) RegisterTempControllers(mainConfig config.Configuration) {
	Config = mainConfig
	routing.AddGenericRoute("/temps", TempsController)
	routing.AddGenericRoute("/temps/garage", tempController.GarageTempRequest)
	routing.AddGenericRoute("/temps/all", tempController.AllTempsRequest)
	routing.AddRouteWithMethod("/temps/registersensor", "POST", tempController.RegisterTempSensor)
	routing.AddRouteWithMethod("/temps/updatesensor", "PUT", tempController.UpdateTempSensor)
	routing.AddRouteWithMethod("/temps/deletesensor", "DELETE", tempController.DeleteTempSensor)
}

func TempsController(writer http.ResponseWriter, request *http.Request) {
	p := &models.Page{
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

func fetchTemperature(address string, name string) models.RoomTemperature {
	status, err := http.Get(address)

	var t models.RoomTemperature
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

func (tempHandler *TempController) RegisterTempSensor(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var sensor models.TemperatureSensor

	err := decoder.Decode(&sensor)
	if err != nil {
		panic(err)
	}

	sensor.SensorId = uuid.NewString()

	if services.VerifyTemperatureSensorIdIsNew(sensor.SensorId, tempHandler.DB) {
		services.AddNewTemperatureSensor(&sensor, tempHandler.DB)

		result, err := json.Marshal(sensor)
		if err != nil {
			panic(err)
		}

		writeResponse(writer, result)
	}
}

func (tempHandler *TempController) UpdateTempSensor(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var sensor models.TemperatureSensor

	err := decoder.Decode(&sensor)
	if err != nil {
		panic(err)
	}

	services.UpdateTemperatureSensor(&sensor, tempHandler.DB)
}

func (tempHandler *TempController) DeleteTempSensor(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var sensor models.TemperatureSensor

	err := decoder.Decode(&sensor)
	if err != nil {
		panic(err)
	}

	services.DeleteTemperatureSensor(sensor.SensorId, tempHandler.DB)
}

func (tempHandler *TempController) AllTempsRequest(writer http.ResponseWriter, request *http.Request) {
	allSensors := services.FetchAllTemperatureSensors(tempHandler.DB)
	allTemps := make([]models.RoomTemperature, 0)

	for _, sensor := range allSensors {
		temp := fetchTemperature(fmt.Sprintf("http://%s", sensor.IpAddress), sensor.Name)
		allTemps = append(allTemps, temp)
	}

	result, err := json.Marshal(allTemps)

	if err != nil {
		panic(err)
	}

	writeResponse(writer, result)
}

func (tempHandler *TempController) GarageTempRequest(writer http.ResponseWriter, request *http.Request) {
	garageSensor := services.FetchGarageTemperatureSensor(tempHandler.DB)
	var response models.RoomTemperature

	//no garage sensor found return 404
	if garageSensor.SensorId == "" {
		writer.WriteHeader(http.StatusNotFound)
		response.ErrorMessage = "Garage sensor not found"
		jsonResponse, err := json.Marshal(response)

		if err != nil {
			panic(err)
		}

		writeResponse(writer, jsonResponse)
		return
	}

	response = fetchTemperature(fmt.Sprintf("http://%s", garageSensor.IpAddress), garageSensor.Name)

	result, err := json.Marshal(response)

	if err != nil {
		panic(err)
	}

	writeResponse(writer, result)
}
