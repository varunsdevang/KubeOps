


clean:
	go mod tidy

grpc:
    protoc --go_out=. --go-grpc_out=. ./proto/ratelimit.proto

run:
	go run main.go