package Middleware

import (
	"github.com/gin-gonic/gin"
	resp "user_center/app/Http/Controllers/Responses"
	"user_center/pkg/jwt"
)

// api 路由组的中间件组示例
func BackJWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.Abort()
			resp.TokenFailed(c, "请求未携带token，无权限访问")
			return
		}
		//log.Print("get token: ", token)
		j := jwt.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		//fmt.Println("claims", claims)
		if err != nil {
			if err == jwt.TokenExpired {
				c.Abort()
				resp.TokenFailed(c, "授权已过期")
				return
			}
			c.Abort()
			resp.TokenFailed(c, err.Error())
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
		c.Next()
	}
}
