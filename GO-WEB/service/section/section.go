package section

import (
	"fmt"

	"github.com/robot007num/go/go-web/model/response"
	"github.com/robot007num/go/go-web/model/section"
	"github.com/robot007num/go/go-web/model/user"
	"github.com/robot007num/go/go-web/pkg/snowflake"
	sqluser "github.com/robot007num/go/go-web/repository/user"
)

const (
	logSuccess = "success"
	logError   = "error"
)

func AddSection(UserRe *user.RootAddSection) (response.ResCode, string, string) {
	ok, err := sqluser.RootInsertSection(UserRe.SectionName, UserRe.Introduction)
	if err != nil {
		return response.CodeInsertSectionError, logError, response.InfoUserVerify
	}
	if !ok {
		return response.CodeInsertSectionError, logError, response.InfoSectionSame
	}

	return response.CodeInsertSectionSuccess, logSuccess, ""
}

func GetSection(all *[]section.AllSectionList) (response.ResCode, string, string) {

	if err := sqluser.GetALLSection(all); err != nil {
		return response.CodeGetSectionError, logError, response.InfoSectionFail
	}

	return response.CodeGetSectionSuccess, logSuccess, ""
}

func GetSectionClass(all *[]section.SectionClassList, id string) (response.ResCode, string, string) {

	if err := sqluser.GetSectionClass(all, id); err != nil {
		return response.CodeGetSectionError, logError, response.InfoSectionFail
	}

	return response.CodeGetSectionSuccess, logSuccess, ""
}

func PostNewPost(p section.NewPost, user string) (response.ResCode, string, string) {
	//1. 生成帖子ID(和生成用户一样)
	id, err := snowflake.CreateSnowID()
	if err != nil {
		fmt.Println("生成帖子ID失败")
	}

	if err := sqluser.InsertNewPost(p, id, user); err != nil {
		return response.CodeNewPostError, logError, response.InfoUserInsert
	}

	return response.CodeNewPostSuccess, logSuccess, ""
}

func GetAllPost(se *[]section.SectionClassPost, id int64) (response.ResCode, string, string) {
	if err := sqluser.GetSectionPost(se, id); err != nil {
		return response.CodeGetPostError, logError, response.InfoSectionFail
	}

	return response.CodeGetPostSuccess, logSuccess, ""
}

func GetSpecifyPost(se *section.SectionClassPost, sid int64, pid int64) (response.ResCode, string, string) {
	if err := sqluser.GetSpecifyPost(se, sid, pid); err != nil {
		return response.CodeGetPostError, logError, response.InfoSectionFail
	}

	return response.CodeGetPostSuccess, logSuccess, ""
}
