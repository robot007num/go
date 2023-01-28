package section

type AllSectionList struct {
	Id           int    `json:"Id" db:"id"`
	SectionName  string `json:"SectionName" db:"section_name"`
	Introduction string `json:"Introduction" db:"introduction"`
}

type SectionClassList struct {
	SectionName string `json:"SectionName" db:"section_name"`
	ClassName   string `json:"ClassName" db:"class_name"`
}

type NewPost struct {
	Title        string `json:"Title" binding:"required""`
	Content      string `json:"Content" binding:"required"`
	SectionClass int    `json:"SectionClass" `
}

//post_id,author_name,title,content,create_time
type SectionClassPost struct {
	PostId     string `json:"PostId" db:"post_id"`
	AuthorName string `json:"AuthorName" db:"author_name"`
	Title      string `json:"Title" binding:"required" db:"title"`
	Content    string `json:"Content" binding:"required" db:"content"`
	CreateTime string `json:"CreateTime" db:"create_time"`
}
