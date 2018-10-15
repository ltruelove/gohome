package main

import (
	//"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux" //_ "github.com/mattn/go-sqlite3"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

type page struct {
	Title string
	Body  []byte
}

type ipAddress struct {
	IP string `json:"ip"`
}

type garden struct {
	SoilReading int `json:"soilReading"`
}

type Water struct {
	Status string `json:"status"`
}

type Configuration struct {
	Pin           string `json:"pin"`
	DoorIp        string `json:"doorIp"`
	GardenIp      string `json:"gardenIp"`
	WaterIp       string `json:"waterIp"`
	SoilThreshold int    `json:"soilThreshold"`
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

	go func() {
		ticker := time.NewTicker(time.Minute * 5)
		for range ticker.C {
			t := time.Now()
			fmt.Println(t.Format("2006-01-02 15:04:05"), "Ticker ticked")
			GetSoilReading()
		}
	}()

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
		return
	}

	reading := fmt.Sprintf("Soil Reading: %d", soilResponse.SoilReading)
	fmt.Println(reading)

	if soilResponse.SoilReading < config.SoilThreshold {
		StartWater()
	}
}

func StartWater() {
	waterStatusAddress := fmt.Sprintf("http://%s/status", config.WaterIp)
	waterResp, waterErr := http.Get(waterStatusAddress)
	if waterErr != nil {
		// handle error
		fmt.Println("Timeout?")
		return
	}

	defer waterResp.Body.Close()
	waterBody, waterErr := ioutil.ReadAll(waterResp.Body)
	waterString := string(waterBody)

	//clear out that annoying line ending
	re := regexp.MustCompile(`\r?\n`)
	waterString = re.ReplaceAllString(waterString, " ")

	waterResponse := &Water{}
	if err := json.Unmarshal(waterBody, &waterResponse); err != nil {
		fmt.Println("Probably got a bad reading")
		return
	}

	if waterResponse.Status == "on" {
		return
	}

	waterOnAddress := fmt.Sprintf("http://%s/on", config.WaterIp)
	_, waterOnErr := http.Get(waterOnAddress)
	if waterOnErr != nil {
		// handle error
		fmt.Println("Water On Timeout?")
		return
	}

	timeChan := time.NewTimer(time.Minute * 5).C
	for {
		select {
		case <-timeChan:
			fmt.Println("Timer expired")
			waterOffAddress := fmt.Sprintf("http://%s/off", config.WaterIp)
			_, waterOffErr := http.Get(waterOffAddress)
			if waterOffErr != nil {
				// handle error
				fmt.Println("Water Off Timeout?")
				return
			}

			waitChan := time.NewTimer(time.Minute * 5).C
			for {
				select {
				case <-waitChan:
					return
				}
			}
		}
	}
}

func homeHandler(writer http.ResponseWriter, request *http.Request) {
	p := &page{Title: "This is the GoHome Home Page"}
	t, _ := template.ParseFiles("home.html")
	t.Execute(writer, p)
}

func doorHandler(writer http.ResponseWriter, request *http.Request) {
	p := &page{Title: "This is the GoHome Door Page"}
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

	soilResponse := &garden{}
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
		fmt.Println("Water started remotely")
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
		fmt.Println("Water stopped remotely")
	}

	pinresponse, _ := json.Marshal(v)
	writer.Write(pinresponse)
}

func lightIPHandler(writer http.ResponseWriter, request *http.Request) {
	ip := &ipAddress{}
	ip.IP = "127.0.0.1"

	response, _ := json.Marshal(ip)
	fmt.Fprintf(writer, "%s", string(response[:]))
}

func loadPage(title string) (*page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &page{Title: title, Body: body}, nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
