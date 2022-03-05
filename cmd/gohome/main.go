package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/handlers"
	"github.com/ltruelove/gohome/internal/app/setup"
	"github.com/ltruelove/gohome/internal/pkg/routing"
)

var Config config.Configuration
var db *sql.DB

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

	db = setup.InitDb()
	defer db.Close()

	//register application routes
	//each app section should have its own handlers to register with the
	//routing package which now lives outside main in internal/pkg
	handlers.RegisterHomeHandlers(Config)
	handlers.RegisterPinHandlers(Config)
	handlers.RegisterGardenHandlers(Config)
	handlers.RegisterGarageHandlers(Config)
	handlers.RegisterTempHandlers(Config)
	handlers.RegisterElectricHandlers()

	//handle file system requests
	routing.AppRouter.PathPrefix("/").Handler(http.FileServer(http.Dir("./" + Config.WebDir + "/")))

	//use mux to handle http requests
	http.Handle("/", routing.AppRouter)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
