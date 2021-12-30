package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jerryshell/my-flomo-server/db"
	"github.com/jerryshell/my-flomo-server/form"
	"github.com/jerryshell/my-flomo-server/model"
	"github.com/jerryshell/my-flomo-server/result"
	"net/http"
	"time"
)

func Login(c *gin.Context) {
	var formData = form.UserLoginForm{}
	if err := c.ShouldBindJSON(&formData); err != nil {
		c.JSON(http.StatusBadRequest, result.ErrorWithMessage(err.Error()))
		return
	}

	var user = model.User{}
	db.DB.Where("username = ? AND password = ?", formData.Username, formData.Password).First(&user)
	if user == (model.User{}) {
		c.JSON(http.StatusBadRequest, result.ErrorWithMessage("用户名或密码错误"))
		return
	}

	expiresAt := time.Now().Add(time.Hour * 24).Unix()
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: expiresAt,
		Subject:   user.ID,
	}).SignedString([]byte("jwT_p@sSw0rd"))
	if err != nil {
		c.JSON(http.StatusBadRequest, result.ErrorWithMessage(err.Error()))
		return
	}

	c.JSON(200, result.SuccessWithData(gin.H{
		"username":  user.Username,
		"email":     user.Email,
		"token":     token,
		"expiresAt": expiresAt,
	}))
}
