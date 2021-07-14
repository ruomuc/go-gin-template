package util

import (
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"ticket-crawler/pkg/setting"
	"time"
)

type customClaims struct {
	ID       string `json:"id"`
	Username string
	Phone    string
	jwt.StandardClaims
}

func GenerateToken(id int, username string, phone int) (string, error) {
	expireTime := time.Now().Add(time.Hour * time.Duration(setting.AppSetting.JwtExpireTime))

	claims := &customClaims{EncodeMd5(strconv.Itoa(id)), EncodeMd5(username), EncodeMd5(strconv.Itoa(phone)), jwt.StandardClaims{ExpiresAt: expireTime.Unix(), Issuer: "ticket-crawler"}}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenClaims.SignedString([]byte(setting.AppSetting.JwtSecret))
	return token, err
}

func ParseToken(token string) error {
	tokenClaims, err := jwt.ParseWithClaims(token, &customClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(setting.AppSetting.JwtSecret), nil
	})

	if tokenClaims != nil {
		if _, ok := tokenClaims.Claims.(*customClaims); ok && tokenClaims.Valid {
			return nil
		}
	}
	return err
}
