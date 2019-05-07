package main

import (
	"log"
	"github.com/go-redis/redis"
)

var (
	redisClient *redis.Client
)

func connectRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := redisClient.Ping().Result()
	log.Println(pong, err)
}

func exampleClient() {
	err := redisClient.Set("key", "value", 0).Err()
	if err != nil {
		log.Fatal(err)
	} 
	val, err := redisClient.Get("key").Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("key", val)
	
	val2, err := redisClient.Get("key2").Result()
	if err == redis.Nil {
		log.Println("key2 does not exist")
	} else if err != nil {
		log.Fatal(err)
	} else {
		log.Println("key2", val2)
	}
}

func runRedisDB() {
	connectRedis()
	exampleClient()
}

