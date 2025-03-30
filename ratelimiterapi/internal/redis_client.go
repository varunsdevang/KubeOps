package internal

import (
	"context"
	"errors"
	"fmt"
	"ratelimitapi/constants"
	"strconv"
	"sync"
	"time"

	redis "github.com/redis/go-redis/v9"
)

var client *redis.Client
var once sync.Once

// Get Value for service
// Set Value for service

// Get Value based on IP
// Set Value based on IP

type RateClient interface {
	GetRate(ctx context.Context, service string) (int, error)
	SetRate(ctx context.Context, service string, rpm int) error
	CheckRateLimit(ctx context.Context, service, ip string, limit int, window time.Duration) (bool, error)
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

func (rc *RedisClient) GetRate(ctx context.Context, service string) (int, error) {
	res, err := rc.client.HGet(ctx, "service", "rpm").Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return 0, nil // Key doesn't exist, return 0
		}
		return 0, errors.New(constants.RedisGetError + err.Error())
	}
	rpm, err := strconv.Atoi(res)
	if err != nil {
		return 0, errors.New(constants.RedisInvalidDBValueError + err.Error())
	}

	return rpm, nil
}

func (rc *RedisClient) SetRate(ctx context.Context, service string, rpm int) error {
	err := rc.client.HSet(ctx, "service", "rpm", rpm).Err()
	if err != nil {
		return errors.New(constants.RedisSetError + err.Error())
	}
	return nil
}

func (rc *RedisClient) CheckRateLimit(ctx context.Context, service, ip string, limit int, window time.Duration) (bool, error) {
	now := time.Now()
	bucket := fmt.Sprintf("rate_limit:%s:ip:%s", service, ip)
	minScore := now.Add(-window).UnixNano() / int64(time.Millisecond)

	pipe := rc.client.Pipeline()
	pipe.ZAdd(ctx, bucket, redis.Z{Score: float64(now.UnixNano() / int64(time.Millisecond)), Member: now.UnixNano() / int64(time.Millisecond)})
	pipe.ZRemRangeByScore(ctx, bucket, "0", fmt.Sprintf("%d", minScore))
	card := pipe.ZCard(ctx, bucket)
	_, err := pipe.Exec(ctx)

	if err != nil {
		return false, err
	}

	count := card.Val()
	return count <= int64(limit), nil
}
