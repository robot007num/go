package pkg

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/robot007num/go/bbs/global"
	"go.uber.org/zap"
	"os"
)

func SqlxMysql() *sqlx.DB {

	dbCon := global.GVA_CONFIG.Mysql

	if dbCon.Dbname == "" {
		return nil
	}

	//后面字符串 是为了可一次性执行SQL语句
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		dbCon.Username, dbCon.Password, dbCon.Host, dbCon.Port, dbCon.Dbname,
		"charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true")
	d, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		global.GVA_LOG.Error("Sqlx连接MySQL", zap.String("status", "失败"), zap.String("err", err.Error()))
		os.Exit(-1)
	}
	return d
}
