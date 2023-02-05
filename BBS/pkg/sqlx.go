package pkg

import (
	"github.com/jmoiron/sqlx"
	"github.com/robot007num/go/bbs/global"
)

// Sqlx 初始化数据库并产生数据库全局变量
func Sqlx() *sqlx.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return SqlxMysql()
	//case "pgsql":
	//	return GormPgSql()
	//case "oracle":
	//	return GormOracle()
	default:
		return SqlxMysql()
	}
}
