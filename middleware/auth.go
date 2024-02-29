package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoywu/budgetServer/dto"
	"github.com/hoywu/budgetServer/service"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		isValid, uid := service.IsTokenValid(token)
		if !isValid {
			c.JSON(http.StatusUnauthorized, dto.ErrorResp(401, "Unauthorized"))
			c.Abort()
			return
		}
		c.Set("uid", uid)
		c.Next()
		service.RefreshToken(token)
	}
}
