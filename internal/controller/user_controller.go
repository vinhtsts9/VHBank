package controller

import (
	"Golang-Masterclass/simplebank/internal/models"
	"Golang-Masterclass/simplebank/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

var User = new(cUser)

type cUser struct{}

func (c *cUser) CreateUser(ctx *gin.Context) {
	var req models.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

}

func (c *cUser) LoginUser(ctx *gin.Context) {
	var req models.LoginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}
}
