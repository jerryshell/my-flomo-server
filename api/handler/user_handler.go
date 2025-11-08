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

func GetUserSettings(c *gin.Context) {
	user := c.MustGet("user").(*model.User)

	userWithSettings, err := service.UserGetSettings(user.ID)
	if err != nil {
		c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
		return
	}

	// 返回用户设置信息，排除敏感字段
	settings := map[string]interface{}{
		"dailyReviewEnabled": userWithSettings.DailyReviewEnabled,
		"telegramChatId":     userWithSettings.TelegramChatID,
		"telegramBotToken":   userWithSettings.TelegramBotToken,
	}

	c.JSON(http.StatusOK, result.SuccessWithData(settings))
}

func UpdateUserSettings(c *gin.Context) {
	user := c.MustGet("user").(*model.User)

	formData := &form.UserUpdateSettingsForm{}
	if err := c.BindJSON(formData); err != nil {
		c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
		return
	}

	if _, err := service.UserUpdateSettings(user.ID, formData.DailyReviewEnabled, formData.TelegramChatID, formData.TelegramBotToken); err != nil {
		c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result.Success())
}
