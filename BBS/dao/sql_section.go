package dao

import (
	"github.com/robot007num/go/bbs/global"
	"github.com/robot007num/go/bbs/model/common/request"
	"github.com/robot007num/go/bbs/model/common/response"
)

//SQLSectionAdd 增加部门
func SQLSectionAdd(r request.AddSection) (bool, error) {
	//先判断有无此板块名字
	sqlStr := "SELECT count(id) FROM `section` WHERE BINARY class_name=?"
	var count int
	if err := global.GVA_DB.Get(&count, sqlStr, r.ClassName); err != nil {
		return false, err
	}

	//无则插入
	if count == 0 {
		sqlStr = "INSERT INTO `section`(section_id,class_name,introduction) VALUES (?,?,?)"
		_, err := global.GVA_DB.Exec(sqlStr, r.SectionId, r.ClassName, r.Introduction)
		if err != nil {
			return false, err
		}
	}

	return count == 0, nil
}

//SQLSectionGet 获取从属社区的部门
func SQLSectionGet(all *[]response.GetSection, id int) error {
	sqlStr := "SELECT * FROM `section` WHERE section_id=?"
	if err := global.GVA_DB.Select(all, sqlStr, id); err != nil {
		return err
	}

	return nil
}
