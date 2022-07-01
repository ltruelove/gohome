package setup

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Creates the database if it doesn't exist and then opens it. Checks if tables exist, and if not
// creates them.
func InitDb() *sql.DB {
	//open/create db
	db, sqlErr := sql.Open("sqlite3", "gohomedb.s3db")
	CheckErr(sqlErr)

	checkTables(db)

	return db
}

func checkTables(db *sql.DB) {
	fmt.Println("Creating db tables if they don't exist")
	db.Exec(CreateTables)
}

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
