package setup

import (
	"database/sql"
	"log"

	"github.com/ltruelove/gohome/config"
)

func InitDb(config config.Configuration) *sql.DB {
	if config.DbType == "mysql" {
		return initMySqlDb(config)
	} else {
		return InitPostgresDb(config)
	}
}

func CheckErr(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
