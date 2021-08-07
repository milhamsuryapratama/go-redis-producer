package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"milhamsuryapratama/go-redis-producer/config"
	"milhamsuryapratama/go-redis-producer/models"
)

const (
	stream = "stream-tes"
)

var client *redis.Client

func init() {
	var err error
	client, err = config.NewRedisClient()
	if err != nil {
		panic(err)
	}
}

func main() {

	fmt.Println("haloo", models.CategoryEvent)

	var category = &models.Category{}
	category = category.SetCategoryName("Coffee")

	categoryMarshal, _ := json.Marshal(category)

	result, err := client.XAdd(context.Background(), &redis.XAddArgs{
		Stream: stream,
		Values: map[string]interface{}{
			"type": string(models.CategoryEvent),
			"data": categoryMarshal,
		},
	}).Result()

	if err != nil {
		panic(err)
	}

	fmt.Println("result", result)
}
