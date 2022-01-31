package models

import (
	"database/sql"
	"log"
	"number-guessing-game/config"
	"os"

	"github.com/lib/pq"
)

var Db *sql.DB

var err error

func init() { //データベースを作成する

	url := os.Getenv("DATABASE_URL")
	connection, _ := pq.ParseURL(url)
	connection += "sslmode=require"
	Db, err = sql.Open(config.Config.SQLDriver, connection)
	if err != nil {
		log.Fatalln(err)
	}
}
