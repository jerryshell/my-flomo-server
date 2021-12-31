package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/jerryshell/my-flomo-server/config"
	"github.com/jerryshell/my-flomo-server/form"
	"github.com/jerryshell/my-flomo-server/result"
	"github.com/jerryshell/my-flomo-server/service"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func LoginOrRegister(c *gin.Context) {
	var formData = form.UserLoginOrRegisterForm{}
	if err := c.ShouldBindJSON(&formData); err != nil {
		c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
		return
	}

	user := service.UserGetByUsername(formData.Username)
	if user.ID == "" {
		userByRegister, err := service.Register(formData.Username, formData.Password)
		if err != nil {
			c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
			return
		}
		user = userByRegister
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(formData.Password))
	if err != nil {
		c.JSON(http.StatusOK, result.ErrorWithMessage("用户名或密码错误"))
		return
	}

	now := time.Now().Unix()
	expiresAt := time.Now().Add(time.Second * 3).Unix()
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
	formData := &form.UserLoginOrRegisterForm{}
	if err := c.ShouldBindJSON(formData); err != nil {
		c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
		return
	}

	_, err := service.Register(formData.Username, formData.Password)
	if err != nil {
		c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result.Success())
}

func VerifyToken(c *gin.Context) {
	tokenString := c.Param("token")
	mapClaims, err := service.VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.SuccessWithData(gin.H{
		"username": (*mapClaims)["sub"],
	}))
}
