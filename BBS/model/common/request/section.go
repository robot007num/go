package request

//属于社区里面的部门

type AddSection struct {
	SectionId    int    `json:"section_id" db:"section_id" binding:"required"`     //属于哪个社区的ID
	ClassName    string `json:"class_name" db:"class_name" binding:"required"`     //部门名字
	Introduction string `json:"introduction" db:"introduction" binding:"required"` //描述
}
