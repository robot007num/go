package section

import (
	"github.com/robot007num/go/go-web/model/response"
	"github.com/robot007num/go/go-web/model/section"
	"github.com/robot007num/go/go-web/model/user"
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
