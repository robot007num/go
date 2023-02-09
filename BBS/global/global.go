package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/robot007num/go/bbs/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// "GVA global value"

var (
	GVA_LOG    *zap.Logger
	GVA_CONFIG config.AllConfig
	GVA_VIPER  *viper.Viper
	GVA_DB     *sqlx.DB
	GVA_REDIS  *redis.Client
)
