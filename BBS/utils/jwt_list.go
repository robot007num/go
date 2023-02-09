package utils

import (
	"context"
	"github.com/robot007num/go/bbs/global"
)

//jwt 功能函数

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetRedisJWT
//@description: 从redis取jwt
//@param: userName string
//@return: redisJWT string, err error

func GetRedisJWT(userID string) (redisJWT string, err error) {
	redisJWT, err = global.GVA_REDIS.Get(context.Background(), userID).Result()
	return redisJWT, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetRedisJWT
//@description: jwt存入redis并设置过期时间
//@param: jwt string, userID string
//@return: err error

func SetRedisJWT(jwt string, userID string) (err error) {
	// 此处过期时间等于jwt过期时间(先定为永久)
	err = global.GVA_REDIS.Set(context.Background(), userID, jwt, 0).Err()
	return err
}
