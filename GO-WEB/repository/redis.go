package repository

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var redisCon redis.Conn

func StartRedis() {
	con, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		panic("redis inint fail,err :" + err.Error())
	}

	redisCon = con
}

func GetRedisCon() redis.Conn {
	return redisCon
}

func SismemberKey(key, member string) bool {
	res, err := redis.Bool(GetRedisCon().Do("SISMEMBER", key, member))
	if err != nil {
		fmt.Println("SismemberKey fail!....")
	}
	return res
}
