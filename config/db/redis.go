package db

import (
	"context"
	"errors"
	"fmt"
	"ginSample/config"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

var Redis *redis.Client
var ctx = context.Background()

func InitRedis(toml *config.Config) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     toml.Local.RedisHost,
		Password: toml.Local.RedisPassword,
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal(errors.New("redis connection error"))
	}
	fmt.Println("⭐️redis connected successfully !!")

	// redis set test
	if err := rdb.Set(context.TODO(), "test", 1, 30*time.Second).Err(); err != nil {
		fmt.Println("===== REDIS SET ERR ====")
		fmt.Println(err)
		fmt.Println("========================")
		panic(err)
	}

	// redis get test
	_, err = rdb.Get(context.TODO(), "test").Result()
	if err != nil {
		fmt.Println("===== REDIS GET ERR ====")
		fmt.Println(err)
		fmt.Println("========================")
		panic(err)
	}

	Redis = rdb
}
