package pkg

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/robot007num/go/bbs/global"
	"go.uber.org/zap"
	"os"
)

func Redis() {
	redisCfg := global.GVA_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.GVA_LOG.Error("redis connect ping failed, err:", zap.Error(err))
		os.Exit(-1)
	} else {
		global.GVA_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.GVA_REDIS = client
	}
}
