package redis

import (
	"fmt"

	"github.com/go-redis/redis"
)

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

// insert value to tail, maintian data inserted before
func setList[T interface{} | int | string](key string, value []T) error {
	if len(value) == 0 {
		return set(key, "")
	}
	arrayPairs := make([]interface{}, len(value))
	for index, value := range value {
		arrayPairs[index] = value
	}
	fmt.Println("setList,key=", key, "newValue=", arrayPairs)
	_, err := rdb.LPush(key, arrayPairs...).Result()
	return err
}

// todo
func setPairs[T interface{} | string](pairs map[string]T) error {
	arrayPairs := make([]interface{}, len(pairs)*2)
	var index = 0
	for key, value := range pairs {
		arrayPairs[index] = key
		index++
		arrayPairs[index] = value
		index++
	}
	_, err := rdb.MSet(arrayPairs...).Result()
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
	result, err := rdb.MGet(key...).Result()
	fmt.Println("redis,mget,key=", key, "result=", result)
	fmt.Println("err=", err)

	var filtedResult []interface{}
	for _, value := range result {
		if value != nil {
			filtedResult = append(filtedResult, value)
		}
	}
	return filtedResult, err
}

// insert value to tail, maintian data inserted before
func RPush[T interface{} | int | string](key string, value []T) error {

	arrayPairs := make([]interface{}, len(value))
	for index, value := range value {
		arrayPairs[index] = value
	}

	err := rdb.RPush(key, arrayPairs...).Err()
	return err
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
