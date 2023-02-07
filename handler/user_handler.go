package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jerryshell/my-flomo-server/form"
	"github.com/jerryshell/my-flomo-server/model"
	"github.com/jerryshell/my-flomo-server/result"
	"github.com/jerryshell/my-flomo-server/service"
)

func UpdateUserEmail(c *gin.Context) {
	user := c.MustGet("user").(*model.User)

	formData := &form.UserUpdateEmailForm{}
	if err := c.BindJSON(formData); err != nil {
		c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
		return
	}

	if _, err := service.UserUpdateEmail(user.ID, formData.Email); err != nil {
		c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result.Success())
}

func UpdateUserPassword(c *gin.Context) {
	user := c.MustGet("user").(*model.User)

	formData := &form.UserUpdatePasswordForm{}
	if err := c.BindJSON(formData); err != nil {
		c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
		return
	}

	password := strings.TrimSpace(formData.Password)
	if len(password) == 0 {
		c.JSON(http.StatusOK, result.BaseResult{
			Code:    -1,
			Success: false,
			Message: "密码不能为空",
		})
		return
	}

	if _, err := service.UserUpdatePassword(user.ID, password); err != nil {
		c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result.Success())
}
