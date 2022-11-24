package user_application

import (
	"fmt"
	"time"
	"user_center/app/Domain/Cache"
	"user_center/app/Http/Controllers/API/Web/Context/User/SendRegisterCode"
	"user_center/pkg/mail"
	"user_center/pkg/tool"
)

func SendRegister(req *SendRegisterCode.Req) error {

	if req.Type == 1 {
		sendByEmail(req.Email)
	} else {
		sendByPhone(req.Phone)
	}

	return nil
}

func sendByEmail(email string) {
	codeCache := &Cache.CodeCache{}
	code := tool.RandomNumber(6)
	codeCache.SetCacheKey("user:register_code", email).Store(code)
	now := time.Now().Format("2006-01-02 15:04:05")
	tilte := "邮箱激活码"
	body := fmt.Sprintf(`
	<div>
		<div>
			尊敬的%s，您好！
		</div>
		<div style="padding: 8px 40px 8px 50px;">
			<p>您于 %s 提交的邮箱验证，本次验证码为<u><strong>%s</strong></u>，为了保证账号安全，验证码有效期为30分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。</p>
		</div>
		<div>
			<p>此邮箱为系统邮箱，请勿回复。</p>
		</div>
	</div>
	`, email, now, code)
	err := mail.SendGoMail([]string{email}, tilte, body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("aaa")
}

func sendByPhone(phone string) {
	fmt.Println("bbb")
}
