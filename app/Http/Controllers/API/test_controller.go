package API

import (
	"github.com/gin-gonic/gin"
	"user_center/app/Http/Controllers/API/Admin/Responses"
)

func Test(c *gin.Context) {
	Responses.Success(c, "生成数学图片验证码成功","243")
	return
}
