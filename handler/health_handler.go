package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jerryshell/my-flomo-server/result"
	"net/http"
)

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, result.BaseResult{
		Code:    http.StatusOK,
		Success: true,
		Message: "Server Online",
	})
}
