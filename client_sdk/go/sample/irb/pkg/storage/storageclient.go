package storage

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"os"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type DataHandle struct {
	key string
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func newRedisClient() *redis.Client {
	var (
		host     = getEnv("REDIS_HOST", "localhost")
		port     = string(getEnv("REDIS_PORT", "62482"))
		password = getEnv("REDIS_PASSWORD", "")
	)

	return redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password, // no password set
		DB:       0,        // use default DB
	})
}

func Load(key string) string {
	rdb := newRedisClient()

	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}

	return val
}

func Store(key string, encodedData string) {
	rdb := newRedisClient()

	err := rdb.Set(ctx, key, encodedData, 0).Err()
	if err != nil {
		panic(err)
	}
}

func Upload(data []byte) string {
	hashedContent := sha256.Sum256(data)
	encodedContent := base64.StdEncoding.EncodeToString(data)
	key := base64.StdEncoding.EncodeToString(hashedContent[:])

	Store(key, encodedContent)
	return key
}

func Download(key string) []byte {
	val := Load(key)
	bytes, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		panic(err)
	}
	return bytes
}
