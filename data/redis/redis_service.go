package redis

import "github.com/go-redis/redis"

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "120.79.7.215:6379",
		Password: "redis", // 没有密码，默认值
		DB:       0,       // 默认DB 0
	})
}

func set(key string, value interface{}) error {
	return rdb.Set(key, value, 0).Err()
}

func setList(key string, value []interface{}) error {
	_, err := rdb.LPush(key, value...).Result()
	return err
}
func setPairs(pairs ...interface{}) error {
	_, err := rdb.MSet(pairs).Result()
	return err
}

func appendValue(key string, value interface{}) error {
	_, err := rdb.Do("APPEND", key, value).Result()
	return err
}

func get(key string) (string, error) {
	return rdb.Get(key).Result()
}

func mget(key ...string) ([]interface{}, error) {
	return rdb.MGet(key...).Result()
}

func RPush(key string, value ...interface{}) error {
	return rdb.RPush(key, value).Err()
}

func LRange(key string) ([]string, error) {
	resut, err := rdb.LRange(key, 0, -1).Result()
	return resut, err
}

func exists(key ...string) (int64, error) {
	return rdb.Exists(key...).Result()
}

func delete(key ...string) (int64, error) {
	deleted, err := rdb.Del(key...).Result()
	return deleted, err
}
