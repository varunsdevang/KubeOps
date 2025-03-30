package internal

import (
	"context"
	"errors"
	"ratelimitapi/constants"
)

type RateService struct {
	redisClient RateClient
	defaultRPM  int
}

func NewRateService(redisClient RateClient, defaultRPM int) *RateService {
	return &RateService{
		redisClient: redisClient,
		defaultRPM:  defaultRPM,
	}
}

func (rs *RateService) GetRate(ctx context.Context, service string) (int, error) {
	rpm, err := rs.redisClient.GetRate(ctx, service)
	if err != nil {
		return 0, err
	}
	if rpm == 0 {
		return rs.defaultRPM, nil
	}
	return rpm, nil
}

func (rs *RateService) SetRate(ctx context.Context, service string, rpm int) error {
	if rpm < 0 {
		return errors.New(constants.InvalidRPMValueError)
	}
	return rs.redisClient.SetRate(ctx, service, rpm)
}
