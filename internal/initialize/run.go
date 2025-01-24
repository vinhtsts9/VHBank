package initialize

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Run() *gin.Engine {
	LoadConfig()

	InitPostgres()
	InitTokenMaker()
	log.Println("Config ok")
	r := InitRouter()

	return r
}
