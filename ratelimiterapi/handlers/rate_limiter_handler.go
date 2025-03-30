package handlers

import (
	"context"
	"ratelimitapi/internal"
	"ratelimitapi/proto/grpc"
)

//CRUD operations for storing rate limits based on service

type RateLimitHandler struct {
	grpc.UnimplementedRateLimitServiceServer
	Service *internal.RateLimiterService
}

func (rlh *RateLimitHandler) IsRequestAllowed(context.Context, *grpc.RateLimitRequest) (*grpc.RateLimitResponse, error) {
	return nil, nil
}
