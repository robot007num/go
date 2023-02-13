package dao

import (
	"github.com/robot007num/go/bbs/global"
	"github.com/robot007num/go/bbs/model/common/response"
)

//SQLInsertUser 插入新用户
func SQLInsertUser(u response.SQLUser) (err error) {
	sqlStr := "INSERT INTO `user`(user_id,account,password,username) VALUES(?,?,?,?)"
	_, err = global.GVA_DB.Exec(sqlStr, u.UserId, u.Account, u.Password, u.UserName)
	return err
}

//SQLUserSelect 获取用户
func SQLUserSelect(account string) (response.SQLUser, error) {
	userInter := response.SQLUser{}
	//sqlStr := "SELECT user_id,account,password,username,email,type,enable,create_time,update_time FROM `user` WHERE user_id=?"
	sqlStr := "SELECT * FROM `user` WHERE account=?"
	err := global.GVA_DB.Get(&userInter, sqlStr, account)
	return userInter, err
}

//SQLUserExits 根据账户判断是否存在
//func SQLUserExits(account string) (bool, error) {
//	sqlStr := "SELECT count(user_id) FROM `user` WHERE BINARY account=?"
//	var count int
//	if err := global.GVA_DB.Get(&count, sqlStr, account); err != nil {
//		return false, err
//	}
//	return count > 0, nil
//}
