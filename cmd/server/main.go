package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hibiken/asynq"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"Golang-Masterclass/simplebank/gapi"
	"Golang-Masterclass/simplebank/global"
	"Golang-Masterclass/simplebank/internal/database"
	"Golang-Masterclass/simplebank/internal/initialize"
	pb "Golang-Masterclass/simplebank/pb"
	"Golang-Masterclass/simplebank/worker"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// Khởi tạo cấu hình
	initialize.Run()
	taskDistributor := worker.NewRedisTaskDistributor(*global.Redis.GetRedisOpt())
	r := database.New(global.Postgres)
	log.Println("Server configuration:", global.Config.GRPCServerAddress)
	log.Println("DB_SOURCE:", global.Config.DBSource)

	RunDbMigrations(global.Config.MigrationURL, global.Config.DBSource)
	// Tạo gRPC server
	go runTaskProcessor(*global.Redis.GetRedisOpt(), r)
	go RunGateway(r, taskDistributor)
	RunGrpc(r, taskDistributor)

}

func RunDbMigrations(migrationUrl string, dbSource string) {
	migration, err := migrate.New(migrationUrl, dbSource)
	if err != nil {
		log.Fatalf("cannot create migration: %v", err)
	}
	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("cannot migrate: %v", err)
	}
	log.Println("migrate database successfull")
}
func runTaskProcessor(redisOpt asynq.RedisClientOpt, r *database.Queries) {
	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, r)
	log.Println("starting task processor")
	err := taskProcessor.Start()
	if err != nil {
		log.Fatalf("cannot start task processor: %v", err)
	}
}
func RunGateway(r *database.Queries, taskDistributor worker.TaskDistributor) {
	server, err := gapi.NewServer(r, taskDistributor)
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
func RunGrpc(r *database.Queries, taskDistributor worker.TaskDistributor) {
	grpcServer := grpc.NewServer()

	// Tạo và đăng ký server
	server, err := gapi.NewServer(r, taskDistributor)
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
