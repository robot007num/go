package jwt

import (
	"errors"
	"time"

	"github.com/robot007num/go/go-web/model/response"

	"github.com/robot007num/go/go-web/model/user"

	"github.com/golang-jwt/jwt/v4"
)

//TokenClaims 自定义声明结构体并内嵌 jwt.StandardClaims
//jwt 包自带的 jwt.StandardClaims 只包含了官方字段
type TokenClaims struct {
	user.Login
	jwt.RegisteredClaims
}

var MySecret = []byte("测试Token")

func NewToken(u user.Login) (string, string) {
	claims := TokenClaims{
		Login: user.Login{
			Username: u.Username,
			Password: u.Password,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(1))),
			Issuer:    "GO-WEB",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(MySecret)
	if err != nil {
		return "", response.InfoTokenCreate
	}
	return tokenStr, ""
}

func ParsingToken(tokenString string) (*TokenClaims, error) {
	//解析Token
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}

	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid { // 校验token
		return claims, nil
	}

	return nil, errors.New("Invalid token")

}
