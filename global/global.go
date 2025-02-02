package global

import (
	"Golang-Masterclass/simplebank/package/setting"
	"Golang-Masterclass/simplebank/util/token"
	"database/sql"

	"github.com/hibiken/asynq"
)

type RedisClientInterface interface {
	GetRedisOpt() *asynq.RedisClientOpt
}

var (
	Postgres   *sql.DB
	Config     setting.Config
	TokenMaker token.Maker
	Redis      RedisClientInterface
)
