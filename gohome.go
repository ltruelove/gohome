package main

import (
	//"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	//_ "github.com/mattn/go-sqlite3"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

type Page struct {
	Title string
	Body  []byte
}

type IpAddress struct {
	Ip string `json:"ip"`
}

type Garden struct {
	SoilReading int `json:"soilReading"`
}

type Water struct {
	Status string `json:"status"`
}

type Configuration struct {
	Pin      string `json:"pin"`
	DoorIp   string `json:"doorIp"`
	GardenIp string `json:"gardenIp"`
	WaterIp  string `json:"waterIp"`
}

var config Configuration

func main() {
	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)

	if err != nil {
		panic(err)
	}

	/*
		db, sqlErr := sql.Open("sqlite3", "gohomedb.s3db")
		checkErr(sqlErr)

		rows, sqlErr := db.Query("SELECT * FROM user;")
		checkErr(sqlErr)

		for rows.Next() {
			var uid int
			var username string
			var password string
			var isDisabled string
			sqlErr = rows.Scan(&uid, &username, &password, &isDisabled)
			checkErr(sqlErr)
		}
	*/

	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/door", DoorHandler)
	router.HandleFunc("/garden", GardenHandler)
	router.HandleFunc("/lightIp", LightIpHandler)
	router.HandleFunc("/pinValid", PinValid).Methods("POST")
	router.HandleFunc("/waterOn", WaterOn).Methods("POST")
	router.HandleFunc("/waterOff", WaterOff).Methods("POST")
	router.HandleFunc("/soil", SoilHandle).Methods("GET")
	router.HandleFunc("/waterStatus", WaterStatus).Methods("GET")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./")))
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
	//fmt.Println(err.Error())

	ticker := time.NewTicker(time.Minute * 5)

	go func() {
		for range ticker.C {
			GetSoilReading()
		}
	}()
}

func GetSoilReading() {
	address := fmt.Sprintf("http://%s/status", config.GardenIp)
	resp, err := http.Get(address)
	if err != nil {
		// handle error
		fmt.Println("Timeout?")
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	jsonString := string(body)

	//clear out that annoying line ending
	re := regexp.MustCompile(`\r?\n`)
	jsonString = re.ReplaceAllString(jsonString, " ")

	soilResponse := &Garden{}
	soilErr := json.Unmarshal(body, &soilResponse)
	if soilErr != nil {
		errorResponse := "Probably got a bad reading"
		fmt.Println(errorResponse)
	} else {
		reading := fmt.Sprintf("Soil Reading: %d", soilResponse.SoilReading)
		fmt.Println(reading)
	}
}

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	p := &Page{Title: "This is the GoHome Home Page"}
	t, _ := template.ParseFiles("home.html")
	t.Execute(writer, p)
}

func DoorHandler(writer http.ResponseWriter, request *http.Request) {
	p := &Page{Title: "This is the GoHome Door Page"}
	t, _ := template.ParseFiles("door.html")
	t.Execute(writer, p)
}

func GardenHandler(writer http.ResponseWriter, request *http.Request) {
	p := &Page{Title: "This is the GoHome Garden Page"}
	t, _ := template.ParseFiles("garden.html")
	t.Execute(writer, p)
}

type pinRequest struct {
	PinCode string `json:"pinCode"`
}

type validator struct {
	IsValid bool
}

func SoilHandle(writer http.ResponseWriter, request *http.Request) {
	address := fmt.Sprintf("http://%s/status", config.GardenIp)
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
		errorResponse := "Probably got a bad reading"
		writer.Write([]byte(errorResponse))
		fmt.Println(errorResponse)
	} else {
		reading := fmt.Sprintf("%d", soilResponse.SoilReading)
		writer.Write([]byte(reading))
	}
}

func WaterStatus(writer http.ResponseWriter, request *http.Request) {
	address := fmt.Sprintf("http://%s/status", config.WaterIp)
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
		errorResponse := "Probably got a bad reading"
		writer.Write([]byte(errorResponse))
		fmt.Println(errorResponse)
	} else {
		reading := fmt.Sprintf("%s", waterResponse.Status)
		writer.Write([]byte(reading))
	}
}

func PinValid(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var t pinRequest

	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	var v = new(validator)
	if strings.Compare(t.PinCode, config.Pin) == 0 {
		v.IsValid = true
	}

	// Moved the click functionality to here so the IP of the module wouldn't be publicly
	// available
	if v.IsValid {
		address := fmt.Sprintf("http://%s/click", config.DoorIp)
		http.Get(address)
	}

	pinresponse, _ := json.Marshal(v)
	writer.Write(pinresponse)
}

func WaterOn(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var t pinRequest

	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	var v = new(validator)
	if strings.Compare(t.PinCode, config.Pin) == 0 {
		v.IsValid = true
	}

	// Moved the click functionality to here so the IP of the module wouldn't be publicly
	// available
	if v.IsValid {
		address := fmt.Sprintf("http://%s/on", config.WaterIp)
		http.Get(address)
	}

	pinresponse, _ := json.Marshal(v)
	writer.Write(pinresponse)
}

func WaterOff(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var t pinRequest

	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	var v = new(validator)
	if strings.Compare(t.PinCode, config.Pin) == 0 {
		v.IsValid = true
	}

	// Moved the click functionality to here so the IP of the module wouldn't be publicly
	// available
	if v.IsValid {
		address := fmt.Sprintf("http://%s/off", config.WaterIp)
		http.Get(address)
	}

	pinresponse, _ := json.Marshal(v)
	writer.Write(pinresponse)
}

func LightIpHandler(writer http.ResponseWriter, request *http.Request) {
	ip := &IpAddress{}
	ip.Ip = "127.0.0.1"

	response, _ := json.Marshal(ip)
	fmt.Fprintf(writer, "%s", string(response[:]))
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
