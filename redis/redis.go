package redis

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"hot/utils"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.source"),   // ip:port
		Password: viper.GetString("redis.password"), // no password set
		DB:       viper.GetInt("redis.db"),          // use default DB
	})
}

func Set(key string, value interface{}) {
	err := client.Set(key, value, 0).Err()
	if err != nil {
		utils.Error.Println(err)
	}
}

func Get(key string) string {
	value, err := client.Get(key).Result()
	if err != nil {
		utils.Error.Println(err)
	}
	return value
}

func Del(key string) {
	err := client.Del(key).Err()
	if err != nil {
		utils.Error.Println(err)
	}
}
