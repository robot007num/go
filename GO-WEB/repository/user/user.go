package user

import (
	"fmt"
	"github.com/robot007num/go/go-web/model/user"
	"github.com/robot007num/go/go-web/repository"
)

//VerifyUserExits 根据用户名找到是否有对应的id
func VerifyUserExits(username string) (bool, error) {
	sqlStr := "SELECT count(user_id) FROM `user` WHERE BINARY username=?"
	var count int
	if err := repository.GetDb().Get(&count, sqlStr, username); err != nil {
		return false, err
	}
	return count > 0, nil

}

//InsertUserTable 插入新的数据
func InsertUserTable(r user.RegisterTable) error {
	sqlStr := "INSERT INTO `user`(user_id,username,password) VALUES (?,?,?)"
	_, err := repository.GetDb().Exec(sqlStr, r.Userid, r.Username, r.Password)
	if err != nil {
		return err
	}
	//fmt.Println(ret.LastInsertId())
	return nil
}

func VerifyUserLogin(username string, w *user.Login) error {
	sqlStr := "SELECT username,password FROM `user` WHERE BINARY username=?"
	if err := repository.GetDb().Get(w, sqlStr, username); err != nil {
		fmt.Println("err:  " + err.Error())
		return err
	}

	return nil
}
