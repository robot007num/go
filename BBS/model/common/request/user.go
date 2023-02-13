package request

// Login /Register User login structure
type Login struct {
	Account  string `json:"account" db:"account" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
}
