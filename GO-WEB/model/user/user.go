package user

type Login struct {
	Username string `json:"username" db:"username" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
}

type Register struct {
	Login
	ConfirmPassword string `json:"confirmpassword" binding:"required,eqfield=Password"`
}

type RegisterTable struct {
	Register
	Userid int64 `db:"user_id"`
}
