package util

import (
	"github.com/dgrijalva/jwt-go"
	"ticket-crawler/pkg/setting"
	"time"
)

type customClaims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Phone    int    `json:"phone"`
	jwt.StandardClaims
}

func GenerateToken(id int, username string, phone int) (string, error) {
	expireTime := time.Now().Add(time.Hour * time.Duration(setting.AppSetting.JwtExpireTime))

	claims := &customClaims{id, username, phone, jwt.StandardClaims{ExpiresAt: expireTime.Unix(), Issuer: "ticket-crawler"}}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenClaims.SignedString([]byte(setting.AppSetting.JwtSecret))
	return token, err
}

func ParseToken(token string) (*customClaims, error) {
	var custom customClaims
	tokenClaims, err := jwt.ParseWithClaims(token, &custom, func(token *jwt.Token) (interface{}, error) {
		return []byte(setting.AppSetting.JwtSecret), nil
	})

	if tokenClaims != nil {
		if _, ok := tokenClaims.Claims.(*customClaims); ok && tokenClaims.Valid {
			return &custom, nil
		}
	}
	return nil, err
}
