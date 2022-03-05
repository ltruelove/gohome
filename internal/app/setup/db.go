package setup

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Creates the database if it doesn't exist and then opens it. Checks if tables exist, and if not
// creates them.
func InitDb() *sql.DB {
	//open/create db
	db, sqlErr := sql.Open("sqlite3", "gohomedb.s3db")
	checkErr(sqlErr)

	checkTables(db)

	return db
}

func checkTables(db *sql.DB) {
	//query to see if the TemeratureSensors table exists
	var rowCount = 0
	rows, sqlErr := db.Query(`SELECT
	name
	FROM sqlite_master
	WHERE type='table' AND name='TemperatureSensors';`)

	checkErr(sqlErr)

	//determine if any rows have been returned
	defer rows.Close()
	for rows.Next() {
		rowCount++
	}

	//if no rows exist, create the table
	if rowCount < 1 {
		_, sqlErr := db.Exec(CreateTemperatureSensors)
		checkErr(sqlErr)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
