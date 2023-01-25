package user

import (
	"fmt"

	"github.com/robot007num/go/go-web/model/section"

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

//RootInsertSection 管理者使用
func RootInsertSection(name, introduction string) (bool, error) {
	//先判断有无此板块名字
	sqlStr := "SELECT count(id) FROM `community_section` WHERE BINARY section_name=?"
	var count int
	if err := repository.GetDb().Get(&count, sqlStr, name); err != nil {
		return false, err
	}

	//无则插入
	if count == 0 {
		sqlStr = "INSERT INTO `community_section`(section_name,introduction) VALUES (?,?)"
		_, err := repository.GetDb().Exec(sqlStr, name, introduction)
		if err != nil {
			return false, err
		}
	}

	return count == 0, nil
}

func GetALLSection(all *[]section.AllSectionList) error {
	sqlStr := "SELECT id,section_name,introduction FROM `community_section`"
	if err := repository.GetDb().Select(all, sqlStr); err != nil {
		return err
	}

	return nil
}

func GetSectionClass(all *[]section.SectionClassList, id string) error {
	sqlStr := "SELECT section_name,class_name FROM community_section AS cn INNER JOIN section_class AS se ON cn.id=se.class_id WHERE cn.id=?"
	if err := repository.GetDb().Select(all, sqlStr, id); err != nil {
		return err
	}
	return nil
}
