package internal

import "context"

// Todo : seperate interface for service for better modularity and testing
// type RateLimiter interface {
// }

type RateLimiterService struct {
	redisClient RateLimiterClient
}

func NewRateLimiterService(redisClient RateLimiterClient) *RateLimiterService {
	return &RateLimiterService{
		redisClient: redisClient,
	}
}

func (rls *RateLimiterService) IsRequestAllowed(ctx context.Context, service, ip string) bool {
	return false
}
