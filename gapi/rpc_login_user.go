package gapi

import (
	"Golang-Masterclass/simplebank/global"
	"Golang-Masterclass/simplebank/internal/database"
	"Golang-Masterclass/simplebank/pb"
	"Golang-Masterclass/simplebank/response"
	"Golang-Masterclass/simplebank/util/password"
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {

	user, err := s.r.GetUser(ctx, req.GetUsername())
	if err != nil {
		if errors.Is(err, response.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	err = password.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "incorrect password")
	}

	accessToken, accessPayload, err := global.TokenMaker.CreateToken(
		user.Username,
		user.Role,
		global.Config.AccessTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create accesstoken")
	}

	refreshToken, refreshPayload, err := global.TokenMaker.CreateToken(
		user.Username,
		user.Role,
		global.Config.RefreshTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to create refreshtoken")
	}
	mtdt := s.extractMetadata(ctx)
	session, err := s.r.CreateSession(ctx, database.CreateSessionParams{
		ID:           refreshPayload.ID,
		Username:     user.Username,
		RefreshToken: refreshToken,
		UserAgent:    mtdt.UserAgent,
		ClientIp:     mtdt.ClientIp,
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to create session")
	}

	rsp := &pb.LoginUserResponse{
		User:                  convertUser(user),
		SessionId:             session.ID.String(),
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiredAt),
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiredAt),
	}
	return rsp, nil

}
