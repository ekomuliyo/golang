package redis

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

// redis-cli -u redis://default:upNTwQBzsHUDwadCDX6e9IJWLKiu0kq6@redis-18824.c244.us-east-1-2.ec2.redns.redis-cloud.com:18824
var client = redis.NewClient(&redis.Options{
	Addr:     "redis-18824.c244.us-east-1-2.ec2.redns.redis-cloud.com:18824",
	Password: "upNTwQBzsHUDwadCDX6e9IJWLKiu0kq6",
})

func TestConnection(t *testing.T) {
	assert.NotNil(t, client)

	err := client.Close()
	assert.Nil(t, err)
}

var ctx = context.Background()

func TestPing(t *testing.T) {
	pong, err := client.Ping(ctx).Result()
	assert.Nil(t, err)
	assert.Equal(t, "PONG", pong)
}

func TestString(t *testing.T) {
	client.SetEx(ctx, "name", "Eko Muliyo", 3*time.Second)

	result, err := client.Get(ctx, "name").Result()
	assert.Nil(t, err)
	assert.Equal(t, "Eko Muliyo", result)

	time.Sleep(5 * time.Second)
	_, err = client.Get(ctx, "name").Result()
	assert.NotNil(t, err)
}

func TestList(t *testing.T) {
	client.RPush(ctx, "names", "Eko Muliyo")
	client.RPush(ctx, "names", "Budi Santoso")
	client.RPush(ctx, "names", "Joko Widodo")

	assert.Equal(t, "Eko Muliyo", client.LPop(ctx, "names").Val())
	assert.Equal(t, "Budi Santoso", client.LPop(ctx, "names").Val())
	assert.Equal(t, "Joko Widodo", client.LPop(ctx, "names").Val())

	client.Del(ctx, "names")

}

func TestSet(t *testing.T) {
	client.SAdd(ctx, "names", "Eko")
	client.SAdd(ctx, "names", "Eko")
	client.SAdd(ctx, "names", "Eko")

	client.SAdd(ctx, "names", "Budi")
	client.SAdd(ctx, "names", "Budi")
	client.SAdd(ctx, "names", "Budi")

	client.SAdd(ctx, "names", "Joko")
	client.SAdd(ctx, "names", "Joko")
	client.SAdd(ctx, "names", "Joko")

	client.SAdd(ctx, "names", "Santoso")
	client.SAdd(ctx, "names", "Santoso")
	client.SAdd(ctx, "names", "Santoso")

	assert.Equal(t, int64(4), client.SCard(ctx, "names").Val())
	assert.Equal(t, []string{"Eko", "Budi", "Joko", "Santoso"}, client.SMembers(ctx, "names").Val())
}

func TestSortedSet(t *testing.T) {
	client.ZAdd(ctx, "scores", redis.Z{
		Score:  1,
		Member: "Eko",
	})
	client.ZAdd(ctx, "scores", redis.Z{
		Score:  2,
		Member: "Budi",
	})
	client.ZAdd(ctx, "scores", redis.Z{
		Score:  3,
		Member: "Joko",
	})
	client.ZAdd(ctx, "scores", redis.Z{
		Score:  4,
		Member: "Santoso",
	})

	assert.Equal(t, []string{"Eko", "Budi", "Joko", "Santoso"}, client.ZRange(ctx, "scores", 0, -1).Val())

	assert.Equal(t, "Santoso", client.ZPopMax(ctx, "scores").Val()[0].Member)
	assert.Equal(t, "Joko", client.ZPopMax(ctx, "scores").Val()[0].Member)
	assert.Equal(t, "Budi", client.ZPopMax(ctx, "scores").Val()[0].Member)
	assert.Equal(t, "Eko", client.ZPopMax(ctx, "scores").Val()[0].Member)
}

func TestHash(t *testing.T) {

	client.HSet(ctx, "user:1", "id", 1)
	client.HSet(ctx, "user:1", "name", "Eko Muliyo")
	client.HSet(ctx, "user:1", "age", 28)

	user := client.HGetAll(ctx, "user:1").Val()
	assert.Equal(t, "1", user["id"])
	assert.Equal(t, "Eko Muliyo", user["name"])
	assert.Equal(t, "28", user["age"])

}

func TestGeoPoint(t *testing.T) {
	client.GeoAdd(ctx, "location", &redis.GeoLocation{
		Name:      "Eko Muliyo",
		Longitude: 101.84513,
		Latitude:  -1.208763,
	})
	client.GeoAdd(ctx, "location", &redis.GeoLocation{
		Name:      "Budi Santoso",
		Longitude: 102.84513,
		Latitude:  -2.208763,
	})
	client.GeoAdd(ctx, "location", &redis.GeoLocation{
		Name:      "Joko Widodo",
		Longitude: 103.84513,
		Latitude:  -3.208763,
	})

	distance := client.GeoDist(ctx, "location", "Eko Muliyo", "Budi Santoso", "km").Val()
	assert.Equal(t, float64(157.2614), distance)

	people := client.GeoSearch(ctx, "location", &redis.GeoSearchQuery{
		Longitude:  101.84513,
		Latitude:   -1.208763,
		Radius:     200,
		RadiusUnit: "km",
	})

	assert.Equal(t, []string{"Budi Santoso", "Eko Muliyo"}, people.Val())
}

func TestHyperLogLog(t *testing.T) {
	client.PFAdd(ctx, "users", "Eko", "Muliyo")
	client.PFAdd(ctx, "users", "Eko", "Pratama")
	client.PFAdd(ctx, "users", "Pratama", "Santoso")

	assert.Equal(t, int64(4), client.PFCount(ctx, "users").Val())
}

func TestPipeline(t *testing.T) {
	_, err := client.Pipelined(ctx, func(pipeliner redis.Pipeliner) error {
		pipeliner.Set(ctx, "name", "Eko Muliyo", 0)
		pipeliner.Set(ctx, "age", "28", 0)
		pipeliner.Set(ctx, "address", "Jakarta", 0)
		return nil
	})
	assert.Nil(t, err)

	assert.Equal(t, "Eko Muliyo", client.Get(ctx, "name").Val())
	assert.Equal(t, "28", client.Get(ctx, "age").Val())
	assert.Equal(t, "Jakarta", client.Get(ctx, "address").Val())
}

func TestTransaction(t *testing.T) {
	_, err := client.TxPipelined(ctx, func(pipeliner redis.Pipeliner) error {
		pipeliner.Set(ctx, "name", "Eko Muliyo", 0)
		pipeliner.Set(ctx, "age", "28", 0)
		pipeliner.Set(ctx, "address", "Jakarta", 0)
		return nil
	})
	assert.Nil(t, err)

	assert.Equal(t, "Eko Muliyo", client.Get(ctx, "name").Val())
	assert.Equal(t, "28", client.Get(ctx, "age").Val())
	assert.Equal(t, "Jakarta", client.Get(ctx, "address").Val())
}

func TestPublishStream(t *testing.T) {
	for i := 0; i < 10; i++ {
		err := client.XAdd(ctx, &redis.XAddArgs{
			Stream: "members",
			Values: map[string]interface{}{
				"name":  "userstream" + strconv.Itoa(i),
				"email": "userstream" + strconv.Itoa(i) + "@example.com",
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

// go test -v -timeout 30s -count=1 -run ^TestGetStream$
func TestGetStream(t *testing.T) {
	// Read stream
	result := client.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    "group-1",
		Consumer: "consumer-1",
		Streams:  []string{"members", ">"},
		Count:    2,
		Block:    5 * time.Second,
	}).Val()

	for _, stream := range result {
		for _, message := range stream.Messages {
			fmt.Println(message.ID)
			fmt.Println(message.Values)
		}
	}
}

// go test -v -timeout 30s -count=1 -run ^TestSubcribePubSub$
func TestSubcribePubSub(t *testing.T) {
	subscriber := client.Subscribe(ctx, "channel-1")
	defer subscriber.Close()
	for i := 0; i < 10; i++ {
		message, err := subscriber.ReceiveMessage(ctx)
		assert.Nil(t, err)
		fmt.Println("Message from channel-1: ", message.Payload)
	}
}

// go test -v -timeout 30s -count=1 -run ^TestPublishPubSub$
func TestPublishPubSub(t *testing.T) {
	for i := 0; i < 10; i++ {
		err := client.Publish(ctx, "channel-1", "Hello World "+strconv.Itoa(i)).Err()
		assert.Nil(t, err)
	}
}
