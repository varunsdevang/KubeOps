package internal

import (
	"context"
	"sync"

	redis "github.com/redis/go-redis/v9"
)

var client *redis.Client
var once sync.Once

// Get Value for service
// Set Value for service

// Get Value based on IP
// Set Value based on IP

type RateClient interface {
	GetRate(ctx context.Context, service string) int
	SetRate(ctx context.Context, service string, rpm int) error
}

type RateLimiterClient interface {
	GetTokens(ctx context.Context, service string, ip string) int
	SetTokens(ctx context.Context, service string, ip string, rpm int) error
}

type RedisClient struct {
	client redis.Client
}

func NewRedisClient(addr, passowrd string, db int) *RedisClient {

	once.Do(
		func() {
			client = redis.NewClient(&redis.Options{
				Addr:     addr,
				Password: passowrd, // No password set
				DB:       db,       // Use default DB
			})
		})
	return &RedisClient{
		client: *client,
	}
}

func (rc *RedisClient) GetRate(ctx context.Context, service string) int {
	return 0
}

func (rc *RedisClient) SetRate(ctx context.Context, service string, rpm int) error {
	return nil
}

func (rc *RedisClient) GetTokens(ctx context.Context, service string, ip string) int {
	return 0
}

func (rc *RedisClient) SetTokens(ctx context.Context, service string, ip string, rpm int) error {
	return nil
}
