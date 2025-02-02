package util

import (
	"context"

	"github.com/gin-gonic/gin"
)

// Định nghĩa key để lưu trữ gin.Context vào context
type GinContextKey string

const GinContextKeyInstance GinContextKey = "GinContext"

// Hàm lưu gin.Context vào context.Context
func GinContextToContext(ginCtx *gin.Context) context.Context {
	return context.WithValue(ginCtx, GinContextKeyInstance, ginCtx)
}

// Hàm lấy gin.Context từ context.Context
func GinContextFromContext(ctx context.Context) (*gin.Context, bool) {
	ginCtx, ok := ctx.Value(GinContextKeyInstance).(*gin.Context)
	return ginCtx, ok
}
