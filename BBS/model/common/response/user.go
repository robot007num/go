package response

//SQLUser 存入数据库并在登录之后返回此结构
type SQLUser struct {
	Account    string `json:"account" db:"account" `        //登录账号
	Password   string `json:"password" db:"password" `      //密码
	Name       string `json:"name" db:"name"`               //用户昵称
	UserId     string `json:"user_id" db:"userid"`          //每个用户独有的ID
	Email      string `json:"email" db:"email"`             //邮件
	CreateTime string `json:"create_time" db:"create_time"` //创建时间
	UpdateTime string `json:"update_time" db:"update_time"` //更新时间
	Enable     int    `json:"enable" db:"enable"`           //用户是否被冻结 0正常 1冻结
}

type LoginResponse struct {
	User         SQLUser `json:"userinfo"`
	AccessToken  string  `json:"access_token"`
	RefreshToken string  `json:"refresh_token"`
}
