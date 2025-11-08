package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jerryshell/my-flomo/api/form"
	"github.com/jerryshell/my-flomo/api/model"
	"github.com/jerryshell/my-flomo/api/result"
	"github.com/jerryshell/my-flomo/api/service"
)

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

func UpdateUserSettings(c *gin.Context) {
	user := c.MustGet("user").(*model.User)

	formData := &form.UserUpdateSettingsForm{}
	if err := c.BindJSON(formData); err != nil {
		c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
		return
	}

	if _, err := service.UserUpdateSettings(user.ID, formData.DailyReviewEnabled); err != nil {
		c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result.Success())
}
