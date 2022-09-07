package connections

import "github.com/go-redis/redis/v9"

func getRedisConnection() *redis.Client {

	options := redis.Options{}

	redisClient := redis.NewClient(&options)

	return redisClient
}
