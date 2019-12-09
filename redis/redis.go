package redis

import (
	"github.com/go-redis/redis"
	"hot/common/util"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "118.25.84.140:6379", // ip:port
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func Set(key string, value interface{})  {
	err := client.Set(key, value, 0).Err()
	if err != nil {
		util.Error.Fatalln(err)
	}
}

func Get(key string) string {
	value, err := client.Get(key).Result()
	if err != nil {
		util.Error.Fatalln(err)
	}
	return value
}

func Exist()  {
}

func Del(key string)  {
	client.Del(key)
}