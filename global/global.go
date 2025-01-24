package global

import (
	"Golang-Masterclass/simplebank/package/setting"
	"Golang-Masterclass/simplebank/util/token"
	"database/sql"
)

var (
	Postgres   *sql.DB
	Config     setting.Config
	TokenMaker token.Maker
)
