package API

import (
	"github.com/gin-gonic/gin"
	responses2 "user_center/app/Http/Controllers/Responses"
)

func Test(c *gin.Context) {
	responses2.Success(c, "生成数学图片验证码成功","243")
	return
}
