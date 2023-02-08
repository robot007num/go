package request

// Login User login structure
type Login struct {
	Account  string `json:"account" db:"account" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
}

// Register User Register structure
type Register struct {
	Login
	Name  string
	Email string
}
