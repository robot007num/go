package jwt

import (
	"errors"
	"fmt"
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

var MySecret = []byte("Robot007num")

var (
	accessToken  string
	refreshToken string
)

type AllToken struct {
	AccessToken  string `json:"access-token"`
	RefreshToken string `json:"refresh-token"`
}

const (
	TokenAccessEffectiveTime  = time.Hour
	TokenRefreshEffectiveTime = time.Hour * 15
	TokenIssuer               = "robot007num"
)

func GetAccessToken() string {
	return accessToken
}

func GetRefreshToken() string {
	return refreshToken
}

func CreateAccessToken(u user.Login) (token string, err error) {
	claims := TokenClaims{
		Login: user.Login{
			Username: u.Username,
			Password: u.Password,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenAccessEffectiveTime * time.Duration(1))),
			Issuer:    TokenIssuer,
		},
	}

	if token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(MySecret); err != nil {
		return "", err
	}
	return
}

func CreateRefreshToken() (token string, err error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenRefreshEffectiveTime * time.Duration(1))),
		Issuer:    TokenIssuer,
	}

	if token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(MySecret); err != nil {
		return "", err
	}
	return
}

//NewToken 生成access_token和refresh_token
func NewToken(u user.Login, Token *AllToken) (log string) {
	var err error
	Token.AccessToken, err = CreateAccessToken(u)
	if err != nil {
		log = response.InfoTokenCreate
		return
	}
	Token.RefreshToken, err = CreateRefreshToken()
	if err != nil {
		log = response.InfoTokenCreate
		return
	}

	return
}

func ParsingToken(tokenString string) (*TokenClaims, error) {
	var token *jwt.Token
	var err error
	//解析Token
	//if !b {
	//	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
	//		return MySecret, nil
	//	})
	//} else {
	token, err = jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	//}

	if err != nil {
		return nil, err
	}

	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid { // 校验token
		return claims, nil
	}

	return nil, errors.New("Invalid token")

}

func ParseDefaultToken(tokenString string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})

	if err != nil {
		fmt.Println("err :" + err.Error())
		return nil, err
	}

	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid { // 校验token
		return claims, nil
	}

	return nil, errors.New("Invalid token")

}

//场景：用于登录之后在一定时间内不用重复登录
// Refresh Token,Access Token
// 1. 先验证R是否存在有效期. 无效则直接返回需要登录;
// 2. 再验证下R 离过期时间有多久,少于10分钟则重新生成Refresh Token.
// 3. 再验证下A是否存在有效期,无效则重新生成。
// 4. 然后返回R和A.
