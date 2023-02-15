package dao

import (
	"github.com/robot007num/go/bbs/global"
	"github.com/robot007num/go/bbs/model/common/response"
	"strconv"
)

//SQLInsertUser 插入新用户
func SQLInsertUser(u response.SQLUser) (err error) {
	sqlStr := "INSERT INTO `user`(user_id,account,password,username) VALUES(?,?,?,?)"
	_, err = global.GVA_DB.Exec(sqlStr, u.UserId, u.Account, u.Password, u.UserName)
	return err
}

//SQLUserSelectToAccount 通过账户获取用户
func SQLUserSelectToAccount(account string) (response.SQLUser, error) {
	userInter := response.SQLUser{}
	//sqlStr := "SELECT user_id,account,password,username,email,type,enable,create_time,update_time FROM `user` WHERE user_id=?"
	sqlStr := "SELECT * FROM `user` WHERE account=?"
	err := global.GVA_DB.Get(&userInter, sqlStr, account)
	return userInter, err
}

//SQLUserSelectToUserID 通过UserID获取用户
func SQLUserSelectToUserID(account int64) (response.SQLUser, error) {
	userInter := response.SQLUser{}
	//sqlStr := "SELECT user_id,account,password,username,email,type,enable,create_time,update_time FROM `user` WHERE user_id=?"
	sqlStr := "SELECT * FROM `user` WHERE user_id=?"
	err := global.GVA_DB.Get(&userInter, sqlStr, account)
	return userInter, err
}

//SQLUserExits 根据账户判断是否存在
func SQLUserExits(account string) (bool, error) {
	sqlStr := "SELECT count(user_id) FROM `user` WHERE BINARY account=?"
	var count int
	if err := global.GVA_DB.Get(&count, sqlStr, account); err != nil {
		return false, err
	}
	return count > 0, nil
}

//SQLUserSelectPart 通过User_id得到原密码
func SQLUserSelectPart(UserID int64) (pword string, err error) {
	sqlStr := "SELECT password FROM `user` WHERE user_id=?"
	var pd string
	err = global.GVA_DB.Get(&pd, sqlStr, strconv.FormatInt(int64(UserID), 10))
	return pd, err
}

//SQLUserChange 更新密码
func SQLUserChange(password string, UserID int64) (err error) {
	sqlStr := "UPDATE `user` SET password=? WHERE user_id=?"
	_, err = global.GVA_DB.Exec(sqlStr, password, UserID)
	return err
}
