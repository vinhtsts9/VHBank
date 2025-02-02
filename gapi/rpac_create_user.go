package gapi

import (
	"Golang-Masterclass/simplebank/internal/database"
	"Golang-Masterclass/simplebank/internal/models"
	"Golang-Masterclass/simplebank/internal/service"
	"Golang-Masterclass/simplebank/pb"
	"Golang-Masterclass/simplebank/util/password"
	"Golang-Masterclass/simplebank/worker"
	"context"
	"time"

	"github.com/hibiken/asynq"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	hashedPassword, err := password.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	}

	arg := &models.CreateUserTxParams{
		CreateUserParams: database.CreateUserParams{
			Username:       req.GetUsername(),
			HashedPassword: hashedPassword,
			FullName:       req.GetFullName(),
			Email:          req.GetEmail(),
		},
		AfterCreate: func(user database.User) error {

			taskPayload := &worker.PayloadSendVerifyEmail{
				Username: user.Username,
			}
			opts := []asynq.Option{
				asynq.MaxRetry(10),
				asynq.ProcessIn(10 * time.Second),
				asynq.Queue(worker.QueueCritical),
			}
			return s.taskDistributor.DistributeTaskSendVerifyEmail(ctx, taskPayload, opts...)
		},
	}
	user, err := service.NewUserLogin().CreateUserTx(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "username already exist: %s", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to distribute task send verify email: %s", err)
	}
	rsp := &pb.CreateUserResponse{
		User: convertUser(user.User),
	}
	return rsp, nil
}
