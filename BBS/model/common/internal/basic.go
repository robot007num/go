package internal

type Basic struct {
	ID         int    `json:"id" db:"id"`                   //下标
	CreateTime string `json:"create_time" db:"create_time"` //创建时间
	UpdateTime string `json:"update_time" db:"update_time"` //更新时间
}
