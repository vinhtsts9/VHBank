package impl

import (
	"Golang-Masterclass/simplebank/global"
	"Golang-Masterclass/simplebank/internal/database"
	"Golang-Masterclass/simplebank/internal/models"
	"Golang-Masterclass/simplebank/response"
	"Golang-Masterclass/simplebank/util/password"
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type sUserLogin struct {
	r  *database.Queries
	db *sql.DB
}

func NewUserLoginImpl(r *database.Queries, db *sql.DB) *sUserLogin {
	return &sUserLogin{
		r:  r,
		db: db,
	}
}
func newUserResponse(user database.User) models.UserResponse {
	return models.UserResponse{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
}

func (s *sUserLogin) CreateUser(ctx *gin.Context, req *models.CreateUserRequest) (database.User, error) {
	hashedPassword, err := password.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return database.User{}, err
	}

	arg := database.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
	}

	user, err := s.r.CreateUser(ctx, arg)
	if err != nil {
		if response.ErrorCode(err) == response.UniqueViolation {
			ctx.JSON(http.StatusForbidden, response.ErrorResponse(err))
			return database.User{}, err
		}
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return database.User{}, err
	}

	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
	return user, nil
}

func (s *sUserLogin) LoginUser(ctx *gin.Context, req models.LoginUserRequest) {

	user, err := s.r.GetUser(ctx, req.Username)
	if err != nil {
		if errors.Is(err, response.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, response.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}

	err = password.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse(err))
		return
	}

	accessToken, accessPayload, err := global.TokenMaker.CreateToken(
		user.Username,
		user.Role,
		global.Config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}

	refreshToken, refreshPayload, err := global.TokenMaker.CreateToken(
		user.Username,
		user.Role,
		global.Config.RefreshTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}

	session, err := s.r.CreateSession(ctx, database.CreateSessionParams{
		ID:           refreshPayload.ID,
		Username:     user.Username,
		RefreshToken: refreshToken,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIp:     ctx.ClientIP(),
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}

	rsp := models.LoginUserResponse{
		SessionID:             session.ID,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		User:                  newUserResponse(user),
	}
	ctx.JSON(http.StatusOK, rsp)

}

func (s *sUserLogin) CreateUserTx(ctx context.Context, arg *models.CreateUserTxParams) (rs models.CreateUserTxResult, err error) {
	var result models.CreateUserTxResult

	err = ExecTx(ctx, s.db, func(q *database.Queries) error {
		var err error
		result.User, err = q.CreateUser(ctx, arg.CreateUserParams)
		if err != nil {
			return err
		}
		err = arg.AfterCreate(result.User)
		return err
	})
	return result, err
}
