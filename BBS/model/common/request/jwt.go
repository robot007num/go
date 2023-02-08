package request

import "github.com/golang-jwt/jwt/v4"

//TokenClaims 自定义声明结构体并内嵌 jwt.StandardClaims (jwt包自带)
type TokenClaims struct {
	BaseClaims
	jwt.RegisteredClaims
}

type BaseClaims struct {
	UserID   string
	UserName string
}
