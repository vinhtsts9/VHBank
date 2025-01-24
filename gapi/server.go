package gapi

import (
	"Golang-Masterclass/simplebank/internal/database"
	"Golang-Masterclass/simplebank/pb"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	r *database.Queries
}

func NewServer(r *database.Queries) (*Server, error) {
	return &Server{
		r: r,
	}, nil
}
