package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/jerryshell/my-flomo-server/config"
	"github.com/jerryshell/my-flomo-server/db"
	"github.com/jerryshell/my-flomo-server/form"
	"github.com/jerryshell/my-flomo-server/model"
	"github.com/jerryshell/my-flomo-server/result"
	"github.com/jerryshell/my-flomo-server/service"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

func Login(c *gin.Context) {
	var formData = form.UserLoginForm{}
	if err := c.ShouldBindJSON(&formData); err != nil {
		c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
		return
	}

	var user = model.User{}
	db.DB.Where("username = ?", formData.Username).First(&user)
	if user == (model.User{}) {
		c.JSON(http.StatusOK, result.ErrorWithMessage("用户不存在"))
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(formData.Password))
	if err != nil {
		c.JSON(http.StatusOK, result.ErrorWithMessage("用户名或密码错误"))
		return
	}

	now := time.Now().Unix()
	expiresAt := time.Now().Add(time.Hour * 24).Unix()
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		IssuedAt:  now,
		Issuer:    "my-flomo-server",
		ExpiresAt: expiresAt,
		Subject:   user.Username,
	}).SignedString([]byte(config.Data.JwtKey))
	if err != nil {
		c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result.SuccessWithData(gin.H{
		"username":  user.Username,
		"email":     user.Email,
		"token":     token,
		"expiresAt": expiresAt,
	}))
}

func Register(c *gin.Context) {
	var formData = form.UserRegisterForm{}
	if err := c.ShouldBindJSON(&formData); err != nil {
		c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
		return
	}

	var user = model.User{}
	db.DB.Where("username = ?", formData.Username).First(&user)
	if user != (model.User{}) {
		c.JSON(http.StatusOK, result.ErrorWithMessage("用户已存在"))
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(formData.Password), 10)
	if err != nil {
		c.JSON(http.StatusOK, result.ErrorWithMessage("加密错误"))
		return
	}
	res, err := service.UserCreate(formData.Username, string(password))
	if err != nil {
		c.JSON(http.StatusOK, result.ErrorWithMessage("注册失败"))
		return
	}
	log.Println(res)
	c.JSON(http.StatusOK, result.SuccessWithData(gin.H{
		"username": formData.Username,
		"result":   "注册成功",
	}))
}

func VerifyToken(c *gin.Context) {
	tokenString := c.Param("token")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("不支持的签名方法：%v", token.Header["alg"])
		}
		return []byte(config.Data.JwtKey), nil
	})
	if err != nil {
		c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
		return
	}
	if !token.Valid {
		c.JSON(http.StatusOK, result.ErrorWithMessage("token 无效"))
		return
	}
	c.JSON(http.StatusOK, result.SuccessWithData(gin.H{
		"username": token.Claims.(jwt.MapClaims)["sub"],
	}))
}
