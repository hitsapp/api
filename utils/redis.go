package utils

import (
	"context"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func ConnectRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}

func GetKey(key string) string {
	return RedisClient.Get(context.Background(), key).Val()
}

func AddKey(key string, value string, expiration int) {
	RedisClient.Set(context.Background(), key, value, 0).Val()
	RedisClient.Expire(context.Background(), key, time.Duration(time.Second*time.Duration(expiration))).Val()
}

func GetHit(url string) float64 {
	return RedisClient.ZScore(context.Background(), "hits", url).Val()
}

func GetHits(limit int64) []redis.Z {
	return RedisClient.ZRevRangeWithScores(context.Background(), "hits", 0, limit).Val()
}

func AddHit(url string) {
	RedisClient.ZAdd(context.Background(), "hits", redis.Z{
		Score:  1,
		Member: url,
	})
}

func IncrementHit(url string) {
	RedisClient.ZIncrBy(context.Background(), "hits", 1, url)
}
