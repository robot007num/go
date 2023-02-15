package request

//帖子

type PostNew struct {
	SectionId int    `json:"section_id" db:"section_id" binding:"required"` //部门ID
	Title     string `json:"title" db:"title" binding:"required"`           //标题
	Content   string `json:"content" db:"content" binding:"required"`       //内容
}
