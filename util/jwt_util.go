package util

import (
	"errors"
	"fmt"
	"log"

	"github.com/golang-jwt/jwt"
	"github.com/jerryshell/my-flomo-server/config"
)

func VerifyToken(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("不支持的签名方法：%v", token.Header["alg"])
		}
		return []byte(config.Data.JwtKey), nil
	})
	if err != nil {
		log.Println("jwt.Parse :: err", err)
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("token无效")
	}
	jwtMapClaims := token.Claims.(jwt.MapClaims)
	return &jwtMapClaims, nil
}

func GetEmailFromJWT(token string) (string, error) {
	mapClaims, err := VerifyToken(token)
	if err != nil {
		return "", err
	}
	email := (*mapClaims)["sub"].(string)
	return email, nil
}
