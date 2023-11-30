package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	RefreshTokenExpireDuration = 7 * 24 * time.Hour
	AccessTokenExpireDuration  = time.Hour
)

var (
	mySecret          = []byte("bluebell")
	ErrorInvalidToken = errors.New("invalid token")
)

type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func GenToken(userID int64, username string) (aToken, rToken string, err error) {
	c := MyClaims{
		//自定义字段
		userID,
		username,

		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(AccessTokenExpireDuration).Unix(),
			Issuer:    "bluebell",
		},
	}
	//创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串aToken
	aToken, err = token.SignedString(mySecret)

	if err != nil {
		return "", "", err
	}
	// refresh token 不需要存任何自定义数据
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(RefreshTokenExpireDuration).Unix(), // 过期时间
		Issuer:    "bluebell",                                        // 签发人
	}).SignedString(mySecret)
	return
}

// ParseToken 解析Token
func ParseToken(tokenString string) (*MyClaims, error) {
	mc := new(MyClaims)
	//解析Token
	claims, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})

	if err != nil {
		return nil, err
	}
	//校验token
	if !claims.Valid {
		return nil, errors.New("token is invalid")
	}
	return mc, nil
}

// RefreshToken 刷新token
func RefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
	_, err = jwt.Parse(rToken, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	if err != nil {
		return
	}

	//从旧token解析claims数据
	var claims MyClaims
	_, err = jwt.ParseWithClaims(aToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	v, _ := err.(*jwt.ValidationError)

	if v.Errors == jwt.ValidationErrorExpired {
		return GenToken(claims.UserID, claims.Username)
	}
	return
}
