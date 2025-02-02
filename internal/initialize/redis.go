package initialize

import (
	"Golang-Masterclass/simplebank/global"
	"log"

	"github.com/hibiken/asynq"
)

type RedisClient struct {
	RedisOpt *asynq.RedisClientOpt
}

func (r *RedisClient) GetRedisOpt() *asynq.RedisClientOpt {
	return r.RedisOpt
}

func InitRedis() {
	client := NewRedis(global.Config.RedisAddress)
	if client == nil {
		log.Fatal("cannot connect to redis")
	}
	global.Redis = client
	log.Println("Successfully connected to redis")

}
func NewRedis(addr string) *RedisClient {

	redisOpt := asynq.RedisClientOpt{
		Addr: addr,
	}
	return &RedisClient{RedisOpt: &redisOpt}

}
