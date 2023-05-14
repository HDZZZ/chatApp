package common

import "github.com/go-redis/redis"

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
}

func Set(key string, value interface{}) error {
	return rdb.Set(key, value, 0).Err()
}

func SetList(key string, value []interface{}) error {
	_, err := rdb.LPush(key, value...).Result()
	return err
}

func AppendValue(key string, value interface{}) error {
	_, err := rdb.Do("APPEND", key, value).Result()
	return err
}

func Get(key string) (string, error) {
	return rdb.Get(key).Result()
}

func Delete(key string) (int64, error) {
	deleted, err := rdb.Del(key).Result()
	return deleted, err
}
