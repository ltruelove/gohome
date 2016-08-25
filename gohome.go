package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

type Page struct {
	Title string
	Body  []byte
}

type IpAddress struct {
	Ip string `json:"ip"`
}

var pin *string
var doorIp *string

func main() {
	pin = flag.String("pin", "", "PIN for entering the garage door")
	doorIp = flag.String("doorIp", "", "IP address of the garage door module")
	flag.Parse()

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

	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/door", DoorHandler)
	router.HandleFunc("/lightIp", LightIpHandler)
	router.HandleFunc("/pinValid", PinValid).Methods("POST")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./")))
	http.Handle("/", router)
	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err.Error())
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

type pinRequest struct {
	PinCode string `json:"pinCode"`
}

type validator struct {
	IsValid bool
}

func PinValid(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var t pinRequest

	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	var v = new(validator)
	if strings.Compare(t.PinCode, *pin) == 0 {
		v.IsValid = true
	}

	// Moved the click functionality to here so the IP of the module wouldn't be publicly
	// available
	if v.IsValid {
		address := fmt.Sprintf("http://%s/click", *doorIp)
		http.Get(address)
	}

	pinResponse, _ := json.Marshal(v)
	writer.Write(pinResponse)
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
