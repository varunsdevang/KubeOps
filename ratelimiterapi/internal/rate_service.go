package internal

import "context"

type RateService struct {
	redisClient RateClient
}

func NewRateService(redisClient RateClient) *RateService {
	return &RateService{
		redisClient: redisClient,
	}
}

func (rs *RateService) GetRate(ctx context.Context, service string) (int, error) {
	return 0, nil
}

func (rs *RateService) SetRate(ctx context.Context, service string, rpm int) error {
	return nil
}
