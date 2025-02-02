postgres: 
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123456 -d postgres:alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres dropdb simple_bank

migrateup:
	migrate -path migration -database "postgresql://root:123456@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path migration -database "postgresql://root:123456@localhost:5432/simple_bank?sslmode=disable" -verbose down

migrateup1:
	migrate -path migration -database "postgresql://root:123456@localhost:5432/simple_bank?sslmode=disable" -verbose up 1
migratedown1:
	migrate -path migration -database "postgresql://root:123456@localhost:5432/simple_bank?sslmode=disable" -verbose down 1
sqlc:
	sqlc generate
test: 
	go test -v -cover ./...
server:
	go run ./cmd/server
mock: 
	mockgen -package mockdb -destination db/mock/store.go Golang-Masterclass/simplebank/db/sqlc Store
proto:
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=simple_bank \
	proto/*.proto
evans: 
	evans --host localhost --port 9099 -r repl
redis:
	docker run --name spbank_redis -p 6379:6380 -d redis:latest
.PHONY: createdb dropdb postgres migratedown migrateup sqlc server mock migratedown1 migrateup1 proto evans 
