package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jerryshell/my-flomo/api/db"
	"github.com/jerryshell/my-flomo/api/model"
	"github.com/jerryshell/my-flomo/api/result"
)

func DeleteMyAccount(c *gin.Context) {
	user := c.MustGet("user").(*model.User)
	// remove memo
	db.DB.Unscoped().Delete(model.Memo{}, "user_id = ?", user.ID)
	// remove user
	db.DB.Unscoped().Delete(model.User{}, "id = ?", user.ID)
	c.JSON(http.StatusOK, result.SuccessWithDataAndMessage(nil, "此账户数据已被毫无痕迹的永久抹除！👋"))
}
