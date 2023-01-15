package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/robot007num/go/go-web/model/config"
	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
)

var db *sqlx.DB

func StartMySql() {
	dbCon := config.GetAllConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		dbCon.MySql.User, dbCon.MySql.PassWord, dbCon.MySql.Host, dbCon.MySql.Port, dbCon.DbName,
		"charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true")
	//fmt.Println(dsn)
	d, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic("sqlx Connect fail,err: " + err.Error())
	}
	db = d

}

func GetDb() *sqlx.DB {
	return db
}

//SelectTableIsExits 判断表是否存在
func SelectTableIsExits(tableName string) error {
	sqlStr := fmt.Sprintf("SELECT * FROM %s", tableName)
	//sqlStr := fmt.Sprintf("SHOW TABLES LIKE '%c%s%c'", '%', tableName, '%')
	//fmt.Println(sqlStr)
	_, err := db.Exec(sqlStr)
	if err != nil {
		return errors.New("table doesn't exist")
	}
	return nil

}

//CreateTableFromSQLFile 读取SQL文件并执行
func CreateTableFromSQLFile() error {
	fileName := "./repository/tables.sql"
	sqlBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic("Readfile fail,err: " + err.Error())
	}

	//fmt.Println(string(sqlBytes))
	sqlTable := string(sqlBytes)
	_, err = db.Exec(sqlTable)
	if err != nil {
		panic("CreateTable failed, err:" + err.Error())
	}
	return nil

}
