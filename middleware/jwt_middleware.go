package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jerryshell/my-flomo-server/result"
	"github.com/jerryshell/my-flomo-server/service"
	"github.com/jerryshell/my-flomo-server/util"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			c.Abort()
			c.JSON(http.StatusOK, result.TokenError())
			return
		}

		username, err := util.GetUsernameFromJWT(token)
		if err != nil {
			log.Println("util.GetUsernameFromJWT :: err", err)
			c.Abort()
			c.JSON(http.StatusOK, result.TokenErrorWithMessage(err.Error()))
			return
		}
		c.Set("username", username)

		user, err := service.UserGetByUsername(username)
		if err != nil {
			log.Println("service.UserGetByUsername :: err", err)
			c.Abort()
			c.JSON(http.StatusOK, result.TokenErrorWithMessage(err.Error()))
			return
		}
		if user.ID == "" {
			c.Abort()
			c.JSON(http.StatusOK, result.TokenErrorWithMessage("用户不存在"))
			return
		}
		c.Set("user", user)

		c.Next()
	}
}
