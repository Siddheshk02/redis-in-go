package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Author struct {
	Name string `redis:"name"`
	Age  int    `redis:"age"`
}

func main() {
	ctx := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pong)

	err = client.Set(ctx, "name", "Gopher", 0).Err() //key-value data
	if err != nil {
		fmt.Println(err)
	}

	client.Set(ctx, "expiring", 10, 10*time.Minute) //key-value data with expiry

	client.HSet(ctx, "struct", Author{"Gopher", 22}) //storing a struct

	fmt.Println("Fetching Data : ")

	val, err := client.Get(ctx, "name").Result()
	if err != nil {
		fmt.Println("Key name not found in Redis cache")
	} else {
		fmt.Printf("name has value %s\n", val)
	}

	res, err := client.Get(ctx, "expiring").Int()
	if err != nil {
		fmt.Println("Key expiring not found in Redis cache")
	} else {
		fmt.Printf("expiring has value %d\n", res)
	}

	var data Author
	err = client.HGetAll(ctx, "struct").Scan(&data)
	if err != nil {
		fmt.Println("Key struct not found in Redis cache")
	} else {
		fmt.Printf("struct has value %+v\n", data)
	}

	result, err := client.Get(ctx, "somekey").Result()
	if err != nil {
		fmt.Println("Key somekey not found in Redis cache")
	} else {
		fmt.Printf("somekey has value %s\n", result)
	}

	fmt.Println("Updating Data : ")

	client.Set(ctx, "expiring", 5, 10*time.Minute)
	res, err = client.Get(ctx, "expiring").Int()
	if err != nil {
		fmt.Println("Key expiring not found in Redis cache")
	} else {
		fmt.Printf("expiring has value %d\n", res)
	}

	fmt.Println("Deleting Data : ")

	client.Del(ctx, "name")
	result, err = client.Get(ctx, "name").Result()
	if err != nil {
		fmt.Println("Key name not found")
	} else {
		fmt.Printf("name has value %s\n", result)
	}
}
