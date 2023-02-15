package request

// Login /Register User login structure
type Login struct {
	Account  string `json:"account" db:"account" binding:"required,alphanum"` // 只能数字或字母
	Password string `json:"password" db:"password" binding:"required,min=3,max=10"`
}

//ChangePasswordReq  Modify password structure
type ChangePasswordReq struct {
	ID          uint   `json:"-"`                                           // 从 JWT 中提取 user id，避免越权
	Password    string `json:"password" binding:"required,min=3,max=10"`    // 密码
	NewPassword string `json:"newpassword" binding:"required,min=3,max=10"` // 新密码
}
