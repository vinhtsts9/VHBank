package service

import (
	"Golang-Masterclass/simplebank/internal/database"
	"Golang-Masterclass/simplebank/internal/models"
	"context"

	"github.com/gin-gonic/gin"
)

type (
	IUserLogin interface {
		CreateUserTx(ctx context.Context, arg *models.CreateUserTxParams) (rs models.CreateUserTxResult, err error)
		LoginUser(ctx *gin.Context, req models.LoginUserRequest)
		CreateUser(ctx *gin.Context, req *models.CreateUserRequest) (database.User, error)
	}
)

var (
	localUserLogin IUserLogin
)

func InitUserLogin(i IUserLogin) {
	localUserLogin = i
}

func NewUserLogin() IUserLogin {
	if localUserLogin == nil {
		panic("implement localUserLogin notfound")
	}

	return localUserLogin
}
