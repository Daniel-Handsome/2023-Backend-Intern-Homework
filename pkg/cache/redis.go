package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/utils"
	"github.com/go-redis/redis/v8"
)

type redisStore struct {
	redis *redis.Client
}

func NewRedisStore(client *redis.Client) Contract {
	return &redisStore{
		redis: client,
	}
}

func New() *redis.Client {
	fmt.Println()
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			utils.GetConfigToString("Redis_ip"),
			utils.GetConfigToInt("Redis_port"),
		),
		Username: utils.GetConfigToString("Redis_username"),
		Password: utils.GetConfigToString("redis_password"),
		DB:       int(utils.GetConfigToInt("Redis_db_id")),
	})
	ok := make(chan bool, 1)
	go func() { // retry for conn to redis
		for {
			_, err := client.Ping(context.Background()).Result()
			if err != nil {
				fmt.Println(err)
				continue
			}
			ok <- true
			return
		}
	}()
	select {
	case <-time.After(5 * time.Second):
		panic("connect to redis timeout 5 second")
	case <-ok:
		fmt.Println("redis connection success")
		return client
	}
}

func (m *redisStore) GetMarshal(key string, unMarshal interface{}) error {
	cached, err := m.GetOrErr(key)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(cached), unMarshal)
}

func (m *redisStore) SetMarshal(key string, canMarshalVal interface{}, seconds int) error {
	bytes, err := json.Marshal(canMarshalVal)
	if err != nil {
		return err
	}
	m.Set(key, string(bytes), seconds)
	return nil
}

func (m *redisStore) Get(key string, fallback string) string {
	val, err := m.GetOrErr(key)
	if err != nil {
		return fallback
	}
	return val
}

func (m *redisStore) Exist(key string) bool {
	result, err := m.redis.Exists(context.Background(), key).Result()
	return err == nil && result == 1
}

func (m *redisStore) Set(key string, value string, seconds int) {
	m.redis.Set(context.Background(), key, value, time.Duration(seconds)*time.Second)
}

func (m *redisStore) GetOrErr(key string) (string, error) {
	return m.redis.Get(context.Background(), key).Result()
}
