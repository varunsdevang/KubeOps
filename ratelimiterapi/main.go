package main

import (
	"net"
	"ratelimitapi/config"
	"ratelimitapi/handlers"
	"ratelimitapi/internal"
	"time"

	pb "ratelimitapi/proto/grpc"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	// Note the import path
)

var configPath = "./config/config.yml"

func main() {

	config, err := config.GetConfig(configPath)
	if err != nil {
		log.Error("error loading config", err)
		return
	}

	log.Info(config)

	redisClient := internal.NewRedisClient(config.Redis.Host, config.Redis.Password, config.Redis.DB)
	rateLimiterSrvc := internal.NewRateLimiterService(redisClient, time.Second*time.Duration(config.RateDuration))
	rateSrvc := internal.NewRateService(redisClient, config.DefaultRPM)
	rateLimitHnldr := handlers.RateLimitHandler{Service: rateLimiterSrvc}
	rateHandler := handlers.RateHandler{Service: rateSrvc}

	lis, err := net.Listen("tcp", ":"+config.Server.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterRateLimitServiceServer(server, &rateLimitHnldr)
	pb.RegisterRateServiceServer(server, &rateHandler)
	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
