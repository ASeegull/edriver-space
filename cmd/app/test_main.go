package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type User struct {
	Login string
	ID    int
}

func main() {
	//manager, _ := auth.NewJWTManager("artur")
	//
	//token, _ := manager.NewJWT(fmt.Sprintf("%v", 1), time.Duration(10)*time.Minute)
	//
	//userId, _ := manager.Parse(token)
	//
	//fmt.Println(userId)
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if err := getAllKeys(ctx, client); err != nil {
		panic(err)
	}
	//iter := client.Scan(ctx, 0, "refresh-token:*", 0).Iterator()

	//var id string
	//var sess models.RefreshSession
	//var count int

	//for iter.Next(ctx) {
	//fmt.Println("keys", iter.Val())
	//key := iter.Val()
	//
	//sessionBytes, err := client.Get(ctx, key).Bytes()
	//if err != nil {
	//	panic(err)
	//}
	//
	//if err := json.Unmarshal(sessionBytes, &sess); err != nil {
	//	panic(err)
	//}
	//
	//if id == "" {
	//	id = sess.UserId
	//}
	//
	//if id == sess.UserId {
	//	count++
	//}
}

func getAllKeys(ctx context.Context, c *redis.Client) error {
	iter := c.Scan(ctx, 0, "refresh-token:*", 0).Iterator()

	//var refreshToken string

	for iter.Next(ctx) {
		fmt.Println(iter.Val(), c.Get(ctx, iter.Val()).Val())
		//row, err := c.Get(ctx, iter.Val()).Bytes()
		//if err != nil {
		//	return err
		//}
		//
		//if err := json.Unmarshal(row, &refreshToken); err != nil {
		//	return err
		//}
		//
		//fmt.Printf("UserId: %s, RefreshToken: %s\n", iter.Val(),refreshToken)
	}

	return nil
}

func delAllKeys(ctx context.Context, c *redis.Client) {
	iter := c.Scan(ctx, 0, "refresh-token:*", 0).Iterator()

	for iter.Next(ctx) {
		c.Del(ctx, iter.Val())
	}
}
