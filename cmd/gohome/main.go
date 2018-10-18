package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/electrical"
	"github.com/ltruelove/gohome/internal/app/garage"
	"github.com/ltruelove/gohome/internal/app/garden"
	"github.com/ltruelove/gohome/internal/app/home"
	"github.com/ltruelove/gohome/internal/app/pin"
	"github.com/ltruelove/gohome/internal/pkg/routing"
	//"database/sql"
	//_ "github.com/mattn/go-sqlite3"
)

var Config config.Configuration

func main() {
	// parse config settings
	file, err := os.Open("config/config.json")
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)

	if err != nil {
		panic(err)
	}

	//register application routes
	//each app section should have its own handlers to register with the
	//routing package which now lives outside main in internal/pkg
	home.RegisterHandlers(Config)
	pin.RegisterHandlers(Config)
	garden.RegisterHandlers(Config)
	garage.RegisterHandlers(Config)
	electrical.RegisterHandlers()

	//handle file system requests
	routing.AppRouter.PathPrefix("/").Handler(http.FileServer(http.Dir("./" + Config.WebDir + "/")))

	//use mux to handle http requests
	http.Handle("/", routing.AppRouter)

	log.Fatal(http.ListenAndServe(":8080", nil))

}

/*** Add this to main() when we're ready to do sql stuff ///
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
/*
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
*/
