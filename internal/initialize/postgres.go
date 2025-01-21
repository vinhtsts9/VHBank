package initialize

import (
	"Golang-Masterclass/simplebank/global"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:123456@localhost:5432/simple_bank?sslmode=disable"
)

func InitPostgres() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot  connect to db", err)
	}
	global.Postgres = conn
}
