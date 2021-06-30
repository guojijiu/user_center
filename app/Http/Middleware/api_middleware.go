package Middleware

import (
	"github.com/gin-gonic/gin"
)

// api 路由组的中间件组示例
func ApiTestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
