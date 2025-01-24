package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"Golang-Masterclass/simplebank/gapi"
	"Golang-Masterclass/simplebank/global"
	"Golang-Masterclass/simplebank/internal/database"
	"Golang-Masterclass/simplebank/internal/initialize"
	pb "Golang-Masterclass/simplebank/pb"
)

func main() {
	// Khởi tạo cấu hình
	initialize.Run()
	r := database.New(global.Postgres)
	log.Println("Server configuration:", global.Config.GRPCServerAddress)

	// Tạo gRPC server
	go RunGateway(r)
	RunGrpc(r)

}
func RunGateway(r *database.Queries) {
	server, err := gapi.NewServer(r)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	grpcMux := runtime.NewServeMux()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = pb.RegisterSimpleBankHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal("cannot register handler server:", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	listener, err := net.Listen("tcp", global.Config.HTTPServerAddress)
	if err != nil {
		log.Fatalf("cannot create listener: %v", err)
	}
	log.Println("starting HTTP/REST gateway on %s", global.Config.HTTPServerAddress)
	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatalf("cannot start HTTP server: %v", err)
	}
}
func RunGrpc(r *database.Queries) {
	grpcServer := grpc.NewServer()

	// Tạo và đăng ký server
	server, err := gapi.NewServer(r)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	// Đăng ký service với gRPC server
	pb.RegisterSimpleBankServer(grpcServer, server)

	// Đăng ký reflection
	reflection.Register(grpcServer)

	// Lắng nghe trên cổng
	listener, err := net.Listen("tcp", global.Config.GRPCServerAddress)
	if err != nil {
		log.Fatalf("cannot create listener: %v", err)
	}
	log.Println("starting listen on : %v", global.Config.GRPCServerAddress)

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("cannot start grpc server: %v", err)
	}
}
