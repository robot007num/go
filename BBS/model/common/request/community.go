package request

//社区

type AddCommunity struct {
	SectionName  string `json:"section_name" db:"section_name" binding:"required"` //社区名字
	Introduction string `json:"introduction" db:"introduction" binding:"required"` //简单描述
}
