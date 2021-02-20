package garden

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/page"
	"github.com/ltruelove/gohome/internal/app/pin"
	"github.com/ltruelove/gohome/internal/pkg/routing"
)

func RegisterHandlers(mainConfig config.Configuration) {
	Config = mainConfig

	if Config.TickerActive {
		Init()
	}

	routing.AddGenericRoute("/garden", GardenHandler)
	routing.AddRouteWithMethod("/waterOn", "POST", WaterOn)
	routing.AddRouteWithMethod("/waterOff", "POST", WaterOff)
	routing.AddRouteWithMethod("/soil", "GET", SoilHandle)
	routing.AddRouteWithMethod("/waterStatus", "GET", WaterStatus)
}

func GardenHandler(writer http.ResponseWriter, request *http.Request) {
	p := &page.Page{Title: "This is the GoHome Garden Page"}
	t, _ := template.ParseFiles(Config.WebDir + "/html/garden.html")
	t.Execute(writer, p)
}

func SoilHandle(writer http.ResponseWriter, request *http.Request) {
	address := fmt.Sprintf("http://%s/status", Config.GardenIp)
	resp, err := http.Get(address)
	if err != nil {
		// handle error
		writer.Write([]byte("Timeout?"))
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	jsonString := string(body)

	//clear out that annoying line ending
	re := regexp.MustCompile(`\r?\n`)
	jsonString = re.ReplaceAllString(jsonString, " ")

	soilResponse := &Garden{}
	if err := json.Unmarshal(body, &soilResponse); err != nil {
		errorResponse := "Probably got a bad soil reading"
		writer.Write([]byte(errorResponse))
		fmt.Println(errorResponse)
	} else {
		reading := fmt.Sprintf("%d", soilResponse.SoilReading)
		writer.Write([]byte(reading))
	}
}

func WaterStatus(writer http.ResponseWriter, request *http.Request) {
	address := fmt.Sprintf("http://%s/status", Config.WaterIp)
	resp, err := http.Get(address)
	if err != nil {
		// handle error
		writer.Write([]byte("Timeout?"))
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	jsonString := string(body)

	//clear out that annoying line ending
	re := regexp.MustCompile(`\r?\n`)
	jsonString = re.ReplaceAllString(jsonString, " ")

	waterResponse := &Water{}
	if err := json.Unmarshal(body, &waterResponse); err != nil {
		errorResponse := "Probably got a bad water reading"
		writer.Write([]byte(errorResponse))
		fmt.Println(errorResponse)
	} else {
		reading := fmt.Sprintf("%s", waterResponse.Status)
		writer.Write([]byte(reading))
	}
}

func WaterOn(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var t pin.PinRequest

	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	var v = new(pin.Validator)
	v.IsValid = Config.ValidatePin(t.PinCode)

	// Moved the click functionality to here so the IP of the module wouldn't be publicly
	// available
	if v.IsValid {
		address := fmt.Sprintf("http://%s/on", Config.WaterIp)
		http.Get(address)
		fmt.Println("Water started remotely")
	}

	pinresponse, _ := json.Marshal(v)
	writer.Write(pinresponse)
}

func WaterOff(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var t pin.PinRequest

	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	var v = new(pin.Validator)
	v.IsValid = Config.ValidatePin(t.PinCode)

	// Moved the click functionality to here so the IP of the module wouldn't be publicly
	// available
	if v.IsValid {
		address := fmt.Sprintf("http://%s/off", Config.WaterIp)
		http.Get(address)
		fmt.Println("Water stopped remotely")
	}

	pinresponse, _ := json.Marshal(v)
	writer.Write(pinresponse)
}
