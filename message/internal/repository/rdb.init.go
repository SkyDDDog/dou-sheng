package repository

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"strconv"
)

var RDB *redis.Client

func InitRedisClient() error {
	addr := viper.GetString("redis.host") + ":" + viper.GetString("redis.port")
	RDB = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: viper.GetString("redis.password"),
		DB:       0,
	})

	_, err := RDB.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func BuildMessageKey(userId int64) string {
	return strconv.FormatInt(userId, 10) + "_message"
}
