package initialize

import (
	"Golang-Masterclass/simplebank/global"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Postgres struct {
	Conn *sql.DB
}

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:123456@localhost:5432/simple_bank?sslmode=disable"
)

func InitPostgres() {
	conn, err := NewDatabse(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot  connect to db", err)
	}
	global.Postgres = conn.Conn
}

func NewDatabse(driver, source string) (*Postgres, error) {
	conn, err := sql.Open(driver, source)
	if err != nil {
		return nil, err
	}
	if err := conn.Ping(); err != nil {
		return nil, err
	}
	log.Println(("Successfully connected to db"))
	return &Postgres{Conn: conn}, nil
}
