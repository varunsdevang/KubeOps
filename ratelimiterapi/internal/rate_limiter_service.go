package internal

import (
	"context"
	"time"
)

// Todo : seperate interface for service for better modularity and testing
// type RateLimiter interface {
// }

type RateLimiterService struct {
	redisClient RateClient
	duration    time.Duration // In seconds
}

func NewRateLimiterService(redisClient RateClient, rateDuration time.Duration) *RateLimiterService {
	return &RateLimiterService{
		redisClient: redisClient,
		duration:    rateDuration,
	}
}

func (rls *RateLimiterService) IsRequestAllowed(ctx context.Context, service, ip string) (bool, error) {
	// Handle no service rpm registered case
	rpm, err := rls.redisClient.GetRate(ctx, service)
	if rpm == 0 && err == nil {
		return true, nil
	}

	return rls.redisClient.CheckRateLimit(ctx, service, ip, rpm, rls.duration)
}
