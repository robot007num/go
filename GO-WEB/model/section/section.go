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
