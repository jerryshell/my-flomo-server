package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jerryshell/my-flomo-server/result"
	"github.com/jerryshell/my-flomo-server/service"
	"net/http"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("token")
		if tokenString == "" {
			c.JSON(200, result.ErrorWithMessage("token is empty"))
			return
		}

		mapClaims, err := service.VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
			return
		}

		username := (*mapClaims)["sub"].(string)
		c.Set("username", username)

		user := service.UserGetByUsername(username)
		if user == nil {
			c.JSON(http.StatusOK, result.ErrorWithMessage("用户不存在"))
			return
		}
		c.Set("user", user)

		c.Next()
	}
}
