package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/electrical"
	"github.com/ltruelove/gohome/internal/garage"
	"github.com/ltruelove/gohome/internal/garden"
	"github.com/ltruelove/gohome/internal/page"
	"github.com/ltruelove/gohome/internal/pin"
	//"database/sql"
	//_ "github.com/mattn/go-sqlite3"
)

var Config config.Configuration

func main() {
	file, err := os.Open("config/config.json")
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)

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
	router.HandleFunc("/", homeHandler)

	electrical.RegisterHandlers(router)
	pin.RegisterHandlers(router, Config)
	garage.RegisterHandlers(router)
	garden.RegisterHandlers(router, Config)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./web/")))
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
	//fmt.Println(err.Error())

}

func homeHandler(writer http.ResponseWriter, request *http.Request) {
	p := &page.Page{Title: "This is the GoHome Home Page"}
	t, _ := template.ParseFiles("web/html/home.html")
	t.Execute(writer, p)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
