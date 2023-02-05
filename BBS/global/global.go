package global

import (
	"github.com/jmoiron/sqlx"
	"github.com/robot007num/go/bbs/config"
	"github.com/robot007num/go/bbs/pkg/log"
	"github.com/spf13/viper"
)

// "GVA global value"

var (
	GVA_LOG    *log.Logger
	GVA_CONFIG config.AllConfig
	GVA_VIPER  *viper.Viper
	GVA_DB     *sqlx.DB
)
