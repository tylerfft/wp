package myredis

import (
	"sync"
	"time"

	redis "github.com/garyburd/redigo/redis"
)

var singlleRedisDb redisClientStu
var singlleRedisDbOnce sync.Once

func GetRedisClientStu() *redisClientStu {
	singlleRedisDbOnce.Do(singlleRedisDb.construct)
	return &singlleRedisDb
}

type redisClientStu struct {
	*redis.Pool
}

func (r *redisClientStu) construct() {
	r.Pool = &redis.Pool{
		MaxIdle:     100,
		MaxActive:   0,
		IdleTimeout: 5 * time.Second,
		Dial: func() (conn redis.Conn, err error) {
			conn, err = redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {

				return nil, err
			}
			return
		},
	}
	return
}

func (r *redisClientStu) CmdExec(commandName string, args ...interface{}) (reply interface{}, err error) {
	conn := r.Get()
	defer conn.Close()
	if conn != nil {
		reply, err = conn.Do(commandName, args)
	}
	return
}
