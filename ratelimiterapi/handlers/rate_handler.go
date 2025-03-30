package handlers

import (
	"context"
	"ratelimitapi/internal"
	"ratelimitapi/proto/grpc"
)

type RateHandler struct {
	grpc.UnimplementedRateServiceServer
	Service *internal.RateService
}

func (rh *RateHandler) SetRate(context.Context, *grpc.ServiceRateMessage) (*grpc.Empty, error) {
	return nil, nil
}
func (rh *RateHandler) GetRate(context.Context, *grpc.ServiceMessage) (*grpc.ServiceRateMessage, error) {
	return nil, nil
}
