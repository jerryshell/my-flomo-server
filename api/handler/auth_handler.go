package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/jerryshell/my-flomo-server/config"
	"github.com/jerryshell/my-flomo-server/form"
	"github.com/jerryshell/my-flomo-server/result"
	"github.com/jerryshell/my-flomo-server/service"
	"github.com/jerryshell/my-flomo-server/util"
	"golang.org/x/crypto/bcrypt"
)

func LoginOrRegister(c *gin.Context) {
	var formData = form.UserLoginOrRegisterForm{}
	if err := c.ShouldBindJSON(&formData); err != nil {
		log.Println("c.ShouldBindJSON :: err", err)
		c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
		return
	}

	// 验证email格式
	if emailValidationMessage := util.GetEmailValidationMessage(formData.Email); emailValidationMessage != "" {
		c.JSON(http.StatusOK, result.ErrorWithMessage(emailValidationMessage))
		return
	}

	user, _ := service.UserGetByEmail(formData.Email)
	if user.ID == "" {
		userByRegister, err := service.Register(formData.Email, formData.Password)
		if err != nil {
			log.Println("service.Register :: err", err)
			c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
			return
		}
		user = userByRegister
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(formData.Password)); err != nil {
		log.Println("bcrypt.CompareHashAndPassword :: err", err)
		c.JSON(http.StatusOK, result.ErrorWithMessage("邮箱或密码错误"))
		return
	}

	now := time.Now().Unix()
	expiresAt := time.Now().Add(time.Hour * 24 * 7).Unix()
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		IssuedAt:  now,
		Issuer:    "my-flomo-server",
		ExpiresAt: expiresAt,
		Subject:   user.Email,
	}).SignedString([]byte(config.Data.JwtKey))
	if err != nil {
		log.Println("jwt.NewWithClaims :: err", err)
		c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result.SuccessWithData(gin.H{
		"email":     user.Email,
		"token":     token,
		"expiresAt": expiresAt,
	}))
}

func Register(c *gin.Context) {
	formData := &form.UserLoginOrRegisterForm{}
	if err := c.ShouldBindJSON(formData); err != nil {
		log.Println("c.ShouldBindJSON :: err", err)
		c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
		return
	}

	// 验证email格式
	if emailValidationMessage := util.GetEmailValidationMessage(formData.Email); emailValidationMessage != "" {
		c.JSON(http.StatusOK, result.ErrorWithMessage(emailValidationMessage))
		return
	}

	if _, err := service.Register(formData.Email, formData.Password); err != nil {
		log.Println("service.Register :: err", err)
		c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result.Success())
}

func VerifyToken(c *gin.Context) {
	tokenString := c.Param("token")
	mapClaims, err := util.VerifyToken(tokenString)
	if err != nil {
		log.Println("service.VerifyToken :: err", err)
		c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.SuccessWithData(gin.H{
		"email": (*mapClaims)["sub"],
	}))
}
