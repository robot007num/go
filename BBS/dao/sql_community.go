package dao

import (
	"github.com/robot007num/go/bbs/global"
	"github.com/robot007num/go/bbs/model/common/request"
	"github.com/robot007num/go/bbs/model/common/response"
)

//SQLCommunityAdd 增加社区
func SQLCommunityAdd(r request.AddCommunity) (bool, error) {
	//先判断有无此板块名字
	sqlStr := "SELECT count(id) FROM `community` WHERE BINARY section_name=?"
	var count int
	if err := global.GVA_DB.Get(&count, sqlStr, r.SectionName); err != nil {
		return false, err
	}

	//无则插入
	if count == 0 {
		sqlStr = "INSERT INTO `community`(section_name,introduction) VALUES (?,?)"
		_, err := global.GVA_DB.Exec(sqlStr, r.SectionName, r.Introduction)
		if err != nil {
			return false, err
		}
	}

	return count == 0, nil
}

//SQLCommunityGet 获取所有社区
func SQLCommunityGet(all *[]response.Community) error {
	sqlStr := "SELECT * FROM `community`"
	if err := global.GVA_DB.Select(all, sqlStr); err != nil {
		return err
	}

	return nil
}
