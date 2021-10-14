package cache

import (
	"fmt"
	"loginExample/constants"
	"loginExample/util/logger"
	"strconv"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/spf13/viper"
)

var clients = map[string]*redis.Client{}

func Register() error {
	configs := viper.GetStringMap("component.redis")
	for key := range configs {
		config := viper.GetStringMapString("component.redis." + key)
		db, err := strconv.Atoi(config["db"])
		if err != nil {
			return err
		}
		client := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", config["host"], config["port"]),
			Password: config["password"],
			DB:       db,
		})
		_, err = client.Ping().Result()
		if err != nil {
			return err
		}
		clients[key] = client
	}
	return nil
}

func Get(key string) *redis.Client {
	client, ok := clients[key]
	if !ok {
		panic("redis配置不存在:" + key)
	}
	return client
}

func Log(logtype string, log string) error {
	redis := Get(constants.CacheName)
	today := time.Now().Format("2006-01-02")
	// login:yyyy-mm-dd  = xxx,2021-10-14 20:00:00.999
	err := redis.HSet(logtype+":"+today, time.Now().Format("2006-01-02 15:04:05"), log).Err()
	if err != nil {
		logger.Error("SetLog Error:", err, logtype)
	}
	return err
}
