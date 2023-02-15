package dao

import (
	"github.com/robot007num/go/bbs/global"
	"github.com/robot007num/go/bbs/model/common/response"
)

func PostNew(p response.Post) error {
	sqlStr := "INSERT INTO `post`(section_id,post_id,user_id,title,content) VALUES(?,?,?,?,?)"
	_, err := global.GVA_DB.Exec(sqlStr, p.SectionId, p.PostId, p.Userid, p.Title, p.Content)
	return err
}

func PostGet(postId int64) (p response.Post, err error) {
	sqlStr := "SELECT * FROM `post` WHERE post_id=?"
	err = global.GVA_DB.Get(&p, sqlStr, postId)
	return
}

func PostGetAll(id int) (p []response.Post, err error) {
	sqlStr := "SELECT * FROM `post` WHERE section_id=?"
	err = global.GVA_DB.Select(&p, sqlStr, id)
	return
}
