package setup

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/ltruelove/gohome/config"
)

// Creates the database if it doesn't exist and then opens it. Checks if tables exist, and if not
// creates them. Adds static data if it doesn't exist.
func InitDb(config config.Configuration) *sql.DB {
	//open/create db
	//psqlconn := fmt.Sprintf("postgres://%s:%s@%s/gohome?sslmode=disable", config.DbUser, config.DbPass, config.DbHost)
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.DbHost, config.DbPort, config.DbUser, config.DbPass, config.DbName)
	db, sqlErr := sql.Open("postgres", psqlconn)
	CheckErr(sqlErr)

	db.Exec(`set search_path='public'`)
	//checkTables(db)
	// TODO make sure data is not duped
	//populateStaticData(db)

	return db
}

func checkTables(db *sql.DB) {
	log.Println("Creating db tables if they don't exist")
	_, err := db.Exec(CreateTables)

	CheckErr(err)
}

func populateStaticData(db *sql.DB) {
	log.Println("Populating static data if it doesn't exist")
	_, err := db.Exec(StaticData)

	CheckErr(err)
}

func CheckErr(err error) {
	if err != nil {
		log.Println(err)
		//panic(err)
	}
}
