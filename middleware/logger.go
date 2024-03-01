package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hoywu/budgetServer/log"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		log.INFO(fmt.Sprintf("[%v ms] %s %s (%s)", latency.Milliseconds(), c.Request.Method, c.Request.RequestURI, c.ClientIP()))
	}
}
