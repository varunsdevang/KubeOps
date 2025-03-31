


clean:
	go mod tidy

grpc:
    protoc --go_out=. --go-grpc_out=. ./proto/ratelimit.proto

run:
	go run main.go

dockerbuild:
	docker build -t ratelimiter:latest .

dockerrun:
	docker run -p 3001:3001 ratelimiter:latest