package rediscli

import (
	"context"

	"github.com/arjunksofficial/kart-challenge/internal/config"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient() (*redis.Client, error) {
	// Load configuration
	redisConfig := config.GetRedisConfig()

	// Create a new Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Host + ":" + redisConfig.Port,
		DB:       redisConfig.DB,
		Password: redisConfig.Password,
	})

	// Test the connection
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return rdb, nil
}

// GetRedisClient returns a singleton Redis client instance
var redisClient *redis.Client

func GetRedisClient() *redis.Client {
	if redisClient == nil {
		var err error
		redisClient, err = NewRedisClient()
		if err != nil {
			panic("Failed to create Redis client: " + err.Error())
		}
	}
	return redisClient
}

// CloseRedisClient closes the Redis client connection
func CloseRedisClient() error {
	if redisClient != nil {
		if err := redisClient.Close(); err != nil {
			return err
		}
		redisClient = nil
	}
	return nil
}

func SetRedisClient(client *redis.Client) {
	redisClient = client
}
