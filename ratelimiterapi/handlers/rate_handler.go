package handlers

import (
	"context"
	"errors"
	"ratelimitapi/constants"
	"ratelimitapi/internal"
	"ratelimitapi/proto/grpc"
)

type RateHandler struct {
	grpc.UnimplementedRateServiceServer
	Service *internal.RateService
}

func (rh *RateHandler) SetRate(ctx context.Context, req *grpc.ServiceRateMessage) (*grpc.Empty, error) {
	if req.Service == "" {
		return nil, errors.New(constants.InvalidServiceNameError)
	}
	err := rh.Service.SetRate(ctx, req.Service, int(req.Rpm))
	return nil, err
}
func (rh *RateHandler) GetRate(ctx context.Context, req *grpc.ServiceMessage) (*grpc.ServiceRateMessage, error) {
	if req.Service == "" {
		return nil, errors.New(constants.InvalidServiceNameError)
	}
	res, err := rh.Service.GetRate(ctx, req.Service)
	return &grpc.ServiceRateMessage{
		Rpm:     int32(res),
		Service: req.Service,
	}, err
}
