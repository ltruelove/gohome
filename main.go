package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/ltruelove/gohome/config"
	"github.com/ltruelove/gohome/internal/app/controllers"
	"github.com/ltruelove/gohome/internal/app/data"
	"github.com/ltruelove/gohome/internal/app/repository"
	"github.com/ltruelove/gohome/internal/app/setup"
	"github.com/ltruelove/gohome/internal/pkg/routing"
)

var Config config.Configuration
var version = "0.0.1"

// @title GoHome API
// @version 0.0.1
// @description API for managing GoHome Control Points and Nodes
// @BasePath /
func main() {
	fmt.Println("Starting GoHome API version " + version)

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

	file.Close()

	//data.LogFileTicker(Config)

	// set up logging
	currentTime := time.Now()
	datedLogFile := fmt.Sprintf("%s_%s", currentTime.Format("2006-01-02"), Config.LogFile)
	logFile, logErr := os.OpenFile(datedLogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if logErr != nil {
		log.Fatalf("error opening file: %v", err)
	}

	defer logFile.Close()
	log.SetOutput(logFile)

	log.Printf("setting up db for %s", Config.DbType)
	db := setup.InitDb(Config)

	viewController := controllers.NewViewControllerWithDeps(
		repository.NewCrudRepositoryFromData(data.NewViewData(db, &Config)),
		repository.NewCrudRepositoryFromData(data.NewNodeSensorData(db, &Config)),
		repository.NewCrudRepositoryFromData(data.NewViewNodeSensorData(db, &Config)),
		repository.NewCrudRepositoryFromData(data.NewNodeSwitchData(db, &Config)),
		repository.NewCrudRepositoryFromData(data.NewViewNodeSwitchData(db, &Config)),
		repository.NewCompoundRepository(db, Config.DbType, &Config),
	)
	sensorTypeController := controllers.NewSensorTypeControllerWithDeps(data.NewSensorType(db, &Config), data.NewSensorTypeData(db, &Config))
	switchTypeController := controllers.NewSwitchTypeControllerWithDeps(data.NewSwitchTypeData(db, &Config))
	nodeController := controllers.NewNodeControllerWithDeps(repository.NewNodeRepository(db, Config.DbType), repository.NewControlPointRepository(db, Config.DbType))
	controlPointController := controllers.NewControlPointControllerWithDeps(repository.NewControlPointRepository(db, Config.DbType))
	switchController := controllers.NewNodeSwitchControllerWithDeps(repository.NewNodeSwitchCrudRepository(db, Config.DbType, &Config), repository.NewCrudRepositoryFromData(data.NewSwitchTypeData(db, &Config)))
	sensorController := controllers.NewNodeSensorControllerWithDeps(repository.NewNodeSensorCrudRepository(db, Config.DbType, &Config), repository.NewCrudRepositoryFromData(data.NewSensorType(db, &Config)))

	//register application routes
	//each app section should have its own handlers to register with the
	//routing package which now lives outside main in internal/pkg
	controllers.RegisterHomeControllers(Config)
	controllers.RegisterPinControllers(Config)

	viewController.RegisterViewEndpoints()
	sensorTypeController.RegisterSensorTypeEndpoints()
	switchTypeController.RegisterSwitchTypeEndpoints()
	nodeController.RegisterNodeEndpoints()
	controlPointController.RegisterControlPointEndpoints()
	switchController.RegisterNodeSwitchEndpoints()
	sensorController.RegisterNodeSensorEndpoints()

	defer db.Close()

	//handle file system requests
	routing.AppRouter.PathPrefix("/").Handler(http.FileServer(http.Dir("./" + Config.WebDir + "/")))

	//use mux to handle http requests
	http.Handle("/", routing.AppRouter)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":"+Config.Port, handlers.CORS(originsOk, headersOk, methodsOk)(routing.AppRouter)))
}
