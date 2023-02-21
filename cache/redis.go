package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func ConnectRedis(ctx context.Context) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ping, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println("erro", err)
		panic(err)
	}

	fmt.Println("Successfully connected!", ping)
	redisClient = client
}

func SetToRedis(ctx context.Context, key string, val []byte) {
	err := redisClient.Set(ctx, key, val, 120*time.Second).Err()
	if err != nil {
		fmt.Println(err)
	}

}

func GetFromRedis(ctx context.Context, key string) string {
	val, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
	}
	// var info model.DataInfo
	// fmt.Println(val)
	// fmt.Println(json.Unmarshal(val, &info))
	return val
}

func GetAllKeys(ctx context.Context, key string) []string {
	keys := []string{}

	iter := redisClient.Scan(ctx, 0, key, 0).Iterator()
	for iter.Next(ctx) {
		keys = append(keys, iter.Val())
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}

	return keys
}
