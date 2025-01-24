package gapi

import (
	"Golang-Masterclass/simplebank/internal/database"
	"Golang-Masterclass/simplebank/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user database.User) *pb.User {
	return &pb.User{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}
