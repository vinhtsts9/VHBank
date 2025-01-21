package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
		})
		userRouterPublic.POST("/login", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
		})
	}
}
