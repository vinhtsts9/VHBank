package initialize

import (
	routers "Golang-Masterclass/simplebank/internal/routes"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if "dev" == " dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	// middlewares
	manageRouter := routers.RouterGroupApp.Manage
	userRouter := routers.RouterGroupApp.User
	MainGroup := r.Group("/v1/2024")
	{
		userRouter.InitUserRouter(MainGroup)
	}
	{
		manageRouter.InitAdminRouter(MainGroup)
	}
	return r
}
