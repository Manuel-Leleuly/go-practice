package belajargolangredis

/*
WARNING:
Install and run Redis before starting these tests
*/

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

var client = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	DB:   0,
})

func TestConnection(t *testing.T) {
	assert.NotNil(t, client)
	// err := client.Close()
	// assert.Nil(t, err)
}

var ctx = context.Background()

func TestPing(t *testing.T) {
	result, err := client.Ping(ctx).Result()
	assert.Nil(t, err)
	assert.Equal(t, "PONG", result)
}

func TestString(t *testing.T) {
	client.SetEx(ctx, "name", "Manuel Leleuly", 3*time.Second)

	result, err := client.Get(ctx, "name").Result()
	assert.Nil(t, err)
	assert.Equal(t, "Manuel Leleuly", result)

	time.Sleep(5 * time.Second)
	result, err = client.Get(ctx, "name").Result()
	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func TestList(t *testing.T) {
	client.RPush(ctx, "names", "Manuel")
	client.RPush(ctx, "names", "Theodore")
	client.RPush(ctx, "names", "Leleuly")

	assert.Equal(t, "Manuel", client.LPop(ctx, "names").Val())
	assert.Equal(t, "Theodore", client.LPop(ctx, "names").Val())
	assert.Equal(t, "Leleuly", client.LPop(ctx, "names").Val())

	client.Del(ctx, "names")
}

func TestSet(t *testing.T) {
	client.SAdd(ctx, "students", "Manuel")
	client.SAdd(ctx, "students", "Manuel")
	client.SAdd(ctx, "students", "Theodore")
	client.SAdd(ctx, "students", "Theodore")
	client.SAdd(ctx, "students", "Leleuly")
	client.SAdd(ctx, "students", "Leleuly")

	assert.Equal(t, int64(3), client.SCard(ctx, "students").Val())
	assert.Equal(t, []string{"Manuel", "Theodore", "Leleuly"}, client.SMembers(ctx, "students").Val())
}

func TestSortedSet(t *testing.T) {
	client.ZAdd(ctx, "scores", redis.Z{Score: 100, Member: "Manuel"})
	client.ZAdd(ctx, "scores", redis.Z{Score: 85, Member: "Eko"})
	client.ZAdd(ctx, "scores", redis.Z{Score: 95, Member: "Budi"})

	assert.Equal(t, []string{"Eko", "Budi", "Manuel"}, client.ZRange(ctx, "scores", 0, 2).Val())
	assert.Equal(t, "Manuel", client.ZPopMax(ctx, "scores").Val()[0].Member)
	assert.Equal(t, "Eko", client.ZPopMax(ctx, "scores").Val()[0].Member)
	assert.Equal(t, "Budi", client.ZPopMax(ctx, "scores").Val()[0].Member)
}

func TestHash(t *testing.T) {
	client.HSet(ctx, "user:1", "id", "1")
	client.HSet(ctx, "user:1", "name", "Manuel")
	client.HSet(ctx, "user:1", "email", "manuel@example.com")

	user := client.HGetAll(ctx, "user:1").Val()
	assert.Equal(t, "1", user["id"])
	assert.Equal(t, "Manuel", user["name"])
	assert.Equal(t, "manuel@example.com", user["email"])

	client.Del(ctx, "user:1")
}

func TestGeoPoint(t *testing.T) {
	client.GeoAdd(ctx, "sellers", &redis.GeoLocation{
		Name:      "Toko A",
		Longitude: 106.822702,
		Latitude:  -6.177590,
	})
	client.GeoAdd(ctx, "sellers", &redis.GeoLocation{
		Name:      "Toko B",
		Longitude: 106.820889,
		Latitude:  -6.174964,
	})

	assert.Equal(t, 0.3543, client.GeoDist(ctx, "sellers", "Toko A", "Toko B", "km").Val())

	sellers := client.GeoSearch(ctx, "sellers", &redis.GeoSearchQuery{
		Longitude:  106.821825,
		Latitude:   -6.175105,
		Radius:     5,
		RadiusUnit: "km",
	}).Val()
	assert.Equal(t, []string{"Toko A", "Toko B"}, sellers)
}

func TestHyperLogLog(t *testing.T) {
	client.PFAdd(ctx, "visitors", "eko", "kurniawan", "khannedy")
	client.PFAdd(ctx, "visitors", "eko", "budi", "joko")
	client.PFAdd(ctx, "visitors", "budi", "joko", "rully")
	assert.Equal(t, int64(6), client.PFCount(ctx, "visitors").Val())
}

func TestPipeline(t *testing.T) {
	_, err := client.Pipelined(ctx, func(p redis.Pipeliner) error {
		p.SetEx(ctx, "name", "Manuel", 5*time.Second)
		p.SetEx(ctx, "address", "Indonesia", 5*time.Second)
		return nil
	})
	assert.Nil(t, err)

	assert.Equal(t, "Manuel", client.Get(ctx, "name").Val())
	assert.Equal(t, "Indonesia", client.Get(ctx, "address").Val())
}

func TestTransaction(t *testing.T) {
	_, err := client.TxPipelined(ctx, func(p redis.Pipeliner) error {
		p.SetEx(ctx, "name", "Joko", 5*time.Second)
		p.SetEx(ctx, "address", "Cirebon", 5*time.Second)
		return nil
	})
	assert.Nil(t, err)

	assert.Equal(t, "Joko", client.Get(ctx, "name").Val())
	assert.Equal(t, "Cirebon", client.Get(ctx, "address").Val())
}

func TestPublishStream(t *testing.T) {
	for i := 0; i < 10; i++ {
		err := client.XAdd(ctx, &redis.XAddArgs{
			Stream: "members",
			Values: map[string]any{
				"name":    "Manuel",
				"address": "Indonesia",
			},
		}).Err()
		assert.Nil(t, err)
	}
}

func TestCreateConsumerGroup(t *testing.T) {
	client.XGroupCreate(ctx, "members", "group-1", "0")
	client.XGroupCreateConsumer(ctx, "members", "group-1", "consumer-1")
	client.XGroupCreateConsumer(ctx, "members", "group-1", "consumer-2")
}

func TestGetStream(t *testing.T) {
	result := client.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    "group-1",
		Consumer: "consumer-1",
		Streams:  []string{"members", ">"},
		Count:    2,
		Block:    5 * time.Second,
	}).Val()

	for _, stream := range result {
		for _, message := range stream.Messages {
			fmt.Println(message)
		}
	}
}

func TestSubscribePubSub(t *testing.T) {
	pubSub := client.Subscribe(ctx, "channel-1")
	for i := 0; i < 10; i++ {
		message, _ := pubSub.ReceiveMessage(ctx)
		fmt.Println(message.Payload)
	}
	err := pubSub.Close()
	assert.Nil(t, err)
}

func TestPublishPubSub(t *testing.T) {
	for i := 0; i < 10; i++ {
		client.Publish(ctx, "channel-1", "Hello "+strconv.Itoa(i))
	}
}
