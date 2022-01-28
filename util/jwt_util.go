package util

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/jerryshell/my-flomo-server/config"
	"log"
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

func GetUsernameFromJWT(token string) (string, error) {
	mapClaims, err := VerifyToken(token)
	if err != nil {
		return "", err
	}
	username := (*mapClaims)["sub"].(string)
	return username, nil
}
