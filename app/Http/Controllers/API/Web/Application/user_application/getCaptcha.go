package user_application

import (
	"fmt"
	"github.com/wenlng/go-captcha/captcha"
	"user_center/app/Http/Controllers/API/Web/Context/User/GetCaptcha"
)

func Captcha(req *GetCaptcha.Req) (string, error) {
	return "", nil
}

// 滑块验证码
func sliderCaptcha() (string, error) {
	// Captcha Single Instances
	capt := captcha.GetCaptcha()

	// 生成验证码
	dots, b64, tb64, key, err := capt.Generate()
	if err != nil {
		return "", err
	}

	// 主图base64
	fmt.Println(len(b64))

	// 缩略图base64
	fmt.Println(len(tb64))

	// 唯一key
	fmt.Println(key)

	// 文本位置验证数据
	fmt.Println(dots)
	return b64, nil
}
