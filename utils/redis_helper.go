package utils

import "github.com/go-redis/redis"

// GetRedisClient returns redis client instance
func GetRedisClient() *redis.Client {
	redisURL := "redis://localhost:6379/0"
	option, err := redis.ParseURL(redisURL)

	if err != nil {
		panic(err)
	}

	return redis.NewClient(option)
}
