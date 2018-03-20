package main

import (
	"github.com/go-redis/redis"
)

// GetClient
func GetClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.56.101:6379",
		Password: "", // no password
		DB:       0,  // use default DB
	})
	return client
}