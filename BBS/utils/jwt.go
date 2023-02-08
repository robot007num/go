package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/robot007num/go/bbs/global"
	"github.com/robot007num/go/bbs/model/common/request"
	"time"
)

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.GVA_CONFIG.JWT.SigningKey),
	}
}

var (
	TokenExpired     = errors.New("token已到期")
	TokenNotValidYet = errors.New("token尚未生效")
	TokenMalformed   = errors.New("token格式不符合")
	TokenInvalid     = errors.New("token无效")
)

const (
	TokenEffectiveTime = time.Hour
)

//ParseToken 解析 token
func (j *JWT) ParseToken(tokenString string) (*request.TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.TokenClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	// 对token对象中的Claim进行类型断言
	if token != nil {
		if claims, ok := token.Claims.(*request.TokenClaims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, TokenInvalid

}

//CreateClaims  Access Token 和 Refresh Token
func (j *JWT) CreateClaims(baseClaims request.BaseClaims) (c1, c2 request.TokenClaims) {
	bf := global.GVA_CONFIG.JWT.AExpiresTime
	bf2 := global.GVA_CONFIG.JWT.RExpiresTime
	c1 = request.TokenClaims{
		BaseClaims: baseClaims,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Duration(1))),              // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(bf))), // 过期时间 1天  配置文件
			Issuer:    global.GVA_CONFIG.JWT.Issuer,                                      // 签名的发行者
		},
	}
	c2 = request.TokenClaims{
		BaseClaims: baseClaims,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Duration(1))),               // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(bf2))), // 过期时间 7天  配置文件
			Issuer:    global.GVA_CONFIG.JWT.Issuer,                                       // 签名的发行者
		},
	}

	return c1, c2
}

// 创建一个token
func (j *JWT) CreateToken(claims request.TokenClaims) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(j.SigningKey)
}
