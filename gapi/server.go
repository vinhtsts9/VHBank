package gapi

import (
	"Golang-Masterclass/simplebank/internal/database"
	"Golang-Masterclass/simplebank/pb"
	"Golang-Masterclass/simplebank/worker"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	r               *database.Queries
	taskDistributor worker.TaskDistributor
}

func NewServer(r *database.Queries, taskDistributor worker.TaskDistributor) (*Server, error) {
	return &Server{
		r:               r,
		taskDistributor: taskDistributor,
	}, nil
}
