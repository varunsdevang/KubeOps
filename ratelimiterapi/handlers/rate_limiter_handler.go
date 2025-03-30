package handlers

import (
	"context"
	"errors"
	"net"
	"ratelimitapi/constants"
	"ratelimitapi/internal"
	"ratelimitapi/proto/grpc"
)

//CRUD operations for storing rate limits based on service

type RateLimitHandler struct {
	grpc.UnimplementedRateLimitServiceServer
	Service *internal.RateLimiterService
}

func (rlh *RateLimitHandler) IsRequestAllowed(ctx context.Context, req *grpc.RateLimitRequest) (*grpc.RateLimitResponse, error) {
	res := grpc.RateLimitResponse{}
	if req.Service == "" {
		return &res, errors.New(constants.InvalidServiceNameError)
	}

	if req.Ip == "" && !isValidIPv4(req.Ip) {
		return &res, errors.New(constants.InvalidIPAddressError)
	}

	isAllowed, err := rlh.Service.IsRequestAllowed(ctx, req.Service, req.Ip)
	res.IsAllowed = isAllowed
	return &res, err
}

func isValidIPv4(ip string) bool {
	parsedIP := net.ParseIP(ip)
	return parsedIP != nil && parsedIP.To4() != nil
}
