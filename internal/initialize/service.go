package initialize

import (
	"Golang-Masterclass/simplebank/global"
	"Golang-Masterclass/simplebank/internal/database"
	"Golang-Masterclass/simplebank/internal/service"
	"Golang-Masterclass/simplebank/internal/service/impl"
)

func InitServiceInterface() {
	queries := database.New(global.Postgres)
	service.InitUserLogin(impl.NewUserLoginImpl(queries, global.Postgres))
}
