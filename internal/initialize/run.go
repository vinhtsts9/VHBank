package initialize

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Run() *gin.Engine {
	LoadConfig()

	InitPostgres()
	InitTokenMaker()
	InitRedis()
	InitServiceInterface()
	log.Println("Config ok")
	r := InitRouter()

	return r
}
