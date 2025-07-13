package setup

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ltruelove/gohome/config"
)

// Creates the database if it doesn't exist and then opens it. Checks if tables exist, and if not
// creates them. Adds static data if it doesn't exist.
func initMySqlDb(config config.Configuration) *sql.DB {
	//"username:password@tcp(127.0.0.1:3306)/test"
	sqlconn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?multiStatements=true", config.DbUser, config.DbPass, config.DbHost, config.DbPort, config.DbName)
	db, sqlErr := sql.Open("mysql", sqlconn)
	CheckErr(sqlErr)

	db.Exec(`set search_path='public'`)
	dbSql, sqlErr := os.ReadFile("config/mysql_database.sql")
	CheckErr(sqlErr)
	_, sqlErr = db.Exec(string(dbSql))
	CheckErr(sqlErr)

	return db
}
