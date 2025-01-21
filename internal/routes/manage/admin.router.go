package manage

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminRouter struct {
}

func (at *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	adminRouterPublic := Router.Group("/admin")
	{
		adminRouterPublic.POST("/login", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
		})
	}
}
