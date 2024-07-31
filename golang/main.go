package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	redis "github.com/redis/go-redis/v9"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // use default address
		Password: "",               // no password set
		DB:       0,                // use default DB
	})
	defer client.Close()

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		fmt.Println("Redis could not be reached")
		fmt.Println(err.Error())
	} else {
		fmt.Println("Connected to Redis....")
	}

	type User struct {
		Name     string
		Email    string
		Username string
	}

	user := User{
		Name:     "John Doe",
		Email:    "johndoe@gmail.com",
		Username: "johndoe",
	}

	err := SetData(client, "user.data", &user, 60)   // 60 seconds
	if err != nil {
		fmt.Println("Error while saving user data to Redis")
	}

	var userData User
	err = GetData(client, "user.data", &userData)
	if err != nil {
		fmt.Println("Error while getting user data from Redis")
	}
	fmt.Println(userData.Email)
}

func SetData(client *redis.Client, key string, data interface{}, ex time.Duration) error {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return client.Set(context.Background(), key, dataBytes, time.Second*ex).Err()
}

func GetData(client *redis.Client, key string, data interface{}) error {
	val, err := client.Get(context.Background(), key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), data)
}
