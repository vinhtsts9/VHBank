package initialize

import "github.com/gin-gonic/gin"

func Run() *gin.Engine {
	r := InitRouter()
	InitPostgres()
	return r
}
