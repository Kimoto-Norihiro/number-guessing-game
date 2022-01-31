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

/*
const (
	tableNameUser = "users"
)
*/

func init() { //データベースを作成する

	url := os.Getenv("DATABASE_URL")
	connection, _ := pq.ParseURL(url)
	connection += "sslmode=require"
	Db, err = sql.Open(config.Config.SQLDriver, connection)
	if err != nil {
		log.Fatalln(err)
	}

	/*
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name STRING,
		score INTEGER,
		created_at DATETIME)`, tableNameUser)

	_, err = Db.Exec(cmdU)
	if err != nil {
		log.Println(err)
	}
	*/
}
