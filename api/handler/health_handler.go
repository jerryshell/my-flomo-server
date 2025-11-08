package handler

import (
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/jerryshell/my-flomo/api/result"
)

var Commit = func() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				return setting.Value
			}
		}
	}

	return ""
}()

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, result.BaseResult{
		Code:    http.StatusOK,
		Success: true,
		Message: "Server Online",
		Data:    gin.H{"commit": Commit},
	})
}
