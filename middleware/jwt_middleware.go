package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jerryshell/my-flomo-server/result"
	"github.com/jerryshell/my-flomo-server/service"
	"log"
	"net/http"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("token")
		if tokenString == "" {
			c.Abort()
			c.JSON(http.StatusOK, result.TokenError())
			return
		}

		mapClaims, err := service.VerifyToken(tokenString)
		if err != nil {
			log.Println("service.VerifyToken :: err", err)
			c.Abort()
			c.JSON(http.StatusOK, result.TokenErrorWithMessage(err.Error()))
			return
		}

		username := (*mapClaims)["sub"].(string)
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
