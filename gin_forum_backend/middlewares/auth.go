package middlewares

import (
	"gin_forum/controller"
	"gin_forum/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)


func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		
		// Token在Header的Authorization中

		authHeader := c.Request.Header.Get("Authorization")

		if authHeader == "" {
			controller.ResponseError(c, controller.CodeNeedLogin)
			c.Abort()
			return
		}

	
		parts := strings.SplitN(authHeader, " ", 2)

		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}

		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}

		c.Set(controller.CtxUserIDKey, mc.UserID)

		c.Next() 
	}
}
