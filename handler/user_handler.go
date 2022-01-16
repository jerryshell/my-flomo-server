package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jerryshell/my-flomo-server/form"
	"github.com/jerryshell/my-flomo-server/model"
	"github.com/jerryshell/my-flomo-server/result"
	"github.com/jerryshell/my-flomo-server/service"
	"net/http"
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
