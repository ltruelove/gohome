package setup

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// Creates the database if it doesn't exist and then opens it. Checks if tables exist, and if not
// creates them. Adds static data if it doesn't exist.
func InitDb() *sql.DB {
	//open/create db
	db, sqlErr := sql.Open("sqlite3", "gohomedb.s3db")
	CheckErr(sqlErr)

	checkTables(db)
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
