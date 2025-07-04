package store

import (
	"context"

	"github.com/arjunksofficial/kart-challenge/internal/rediscli"
	"github.com/redis/go-redis/v9"
)

type Cache interface {
	AddToSet(ctx context.Context, key string, code string) error
	IsPresentInSet(ctx context.Context, key string, code string) (bool, error)
}

type redisCache struct {
	redisClient *redis.Client
}

var cache Cache

func New() Cache {
	cache = &redisCache{
		redisClient: rediscli.GetRedisClient(),
	}
	return cache
}

func Get() Cache {
	if cache == nil {
		cache = New()
	}
	return cache
}

func (c *redisCache) AddToSet(ctx context.Context, key string, code string) error {
	return c.redisClient.SAdd(ctx, key, code).Err()
}

func (c *redisCache) IsPresentInSet(ctx context.Context, key string, code string) (bool, error) {
	found, err := c.redisClient.SIsMember(ctx, key, code).Result()
	if err != nil {
		return false, err
	}
	return found, nil
}
