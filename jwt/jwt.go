package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	UserID int64  `json:"user_id"`
	Phone  string `json:"phone"`
	jwt.RegisteredClaims
}

// 定义jwtSecret 用于加密的字符串
var jwtSecret = []byte("secret")

// keyFunc 用于解析token时获取密钥
func keyFunc(_ *jwt.Token) (i interface{}, err error) {
	return jwtSecret, nil
}

// TokenExpireDuration 定义token的过期时间
const TokenExpireDuration = time.Hour * 24

// AccessTokenExpiredDuration 定义token的过期时间
const AccessTokenExpiredDuration = time.Hour * 24     //access_token过期时间
const RefreshTokenExpireDuration = time.Hour * 24 * 7 //refresh_token过期时间

// // GenerateToken 生成jwt （的access_token 和 refresh_token）
// func GenerateToken(userID int, phone string) (aToken, tToken string, err error) { }

func GenerateToken(userID int64, phone string) (aToken, rToken string, err error) {
	//claims := &Claims{
	//	UserID: UserID,
	//	Phone:  phone,
	//	RegisteredClaims: jwt.RegisteredClaims{
	//		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	//	},
	//}
	//创建一个声明
	c := Claims{
		userID, //自定义字段

		phone, //自定义字段
		jwt.RegisteredClaims{
			Issuer: "YiYuEcho",
			//Subject:   "",
			//Audience:  nil,
			//todo : 过期时间
			//ExpiresAt: time.Now().Add(AccessTokenExpiredDuration).Unix(),
			//NotBefore: nil,
			//IssuedAt:  nil,
			//ID:        "",
		},
	}
	//加密并获得完整的编码后的字符串token
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(jwtSecret)

	//refresh token 不需要存任何自定义数据
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer: "YiYuEcho"}).SignedString(jwtSecret) //使用指定的secret签名并获得完整的编码后的字符串token

	return
}

// ParseToken 解析jwt
func ParseToken(tokenString string) (claims *Claims, err error) {
	//解析token
	var token *jwt.Token
	claims = new(Claims)
	token, err = jwt.ParseWithClaims(tokenString, claims, keyFunc)
	if err != nil {
		return
	}
	if !token.Valid { //如果token是无效的，则返回错误
		err = errors.New("token无效")
	}
	return
}

// RefreshToken 刷新AccessToken
func RefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
	//refresh token无效直接返回
	if _, err = ParseToken(rToken); err != nil {
		return
	}

	//从旧access token中解析出claims数据 payload负载信息
	var claims Claims
	_, err = jwt.ParseWithClaims(aToken, &claims, keyFunc)
	//todo : 错误处理
	//v, _ := err.(*jwt.ValidationError)
	//当access token是过期错误时，刷新access token和refresh token
	//	if v.Errors == jwt.ValidationErrorExpired {
	//		return GenToken(claims.UserID, claims.Username)
	//	}
	return
}
