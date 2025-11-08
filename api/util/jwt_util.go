package util

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/jerryshell/my-flomo/api/config"
)

func VerifyToken(tokenString string) (*jwt.MapClaims, error) {
	logger := NewLogger("jwt_util")
	
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			logger.Warn("unsupported signing method", StringField("algorithm", fmt.Sprintf("%v", token.Header["alg"])))
			return nil, fmt.Errorf("不支持的签名方法：%v", token.Header["alg"])
		}
		return []byte(config.Data.JwtKey), nil
	})
	if err != nil {
		logger.Error("failed to parse jwt token", ErrorField(err))
		return nil, err
	}
	if !token.Valid {
		logger.Warn("invalid jwt token")
		return nil, errors.New("token无效")
	}
	jwtMapClaims := token.Claims.(jwt.MapClaims)
	
	logger.Debug("jwt token verified successfully")
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
