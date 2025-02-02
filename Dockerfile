FROM golang:1.23-alpine3.19 as builder
workdir /app
copy . .
run go build -o main ./cmd/server

from alpine:3.19
workdir /app 
copy --from=builder /app/main .
copy app.env .
copy start.sh .
copy wait-for.sh .
copy migrations ./migrations
expose 9090
cmd ["/app/main"]
entrypoint ["/app/start.sh"]