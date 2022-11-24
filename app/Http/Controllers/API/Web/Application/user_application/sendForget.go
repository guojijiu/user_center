package user_application

import (
	"fmt"
	"time"
	"user_center/app/Domain/Cache"
	"user_center/app/Http/Controllers/API/Web/Context/User/SendForgetCode"
	conf "user_center/config"
	"user_center/pkg/encryption"
	"user_center/pkg/mail"
)

func SendForget(req *SendForgetCode.Req) error {

	if req.Type == 1 {
		sendForgetByEmail(req.Email)
	} else {
		sendForgetByPhone(req.Phone)
	}

	return nil
}

func sendForgetByEmail(email string) {
	codeCache := &Cache.CodeCache{}
	str, encrypErr := encryptStr(email)
	if encrypErr != nil {
		fmt.Println(encrypErr)
		return
	}
	forgetType, _ := encryptStr("1")
	forgetUrl := fmt.Sprintf("%s%s&forget_type=", "http://aaaa.com/forget?forget_data=", str, forgetType)
	codeCache.SetCacheKey("user:forget_code", email).Store(forgetUrl)
	now := time.Now().Format("2006-01-02 15:04:05")
	tilte := "忘记密码激活码"

	body := fmt.Sprintf(`
	<div>
		<div>
			尊敬的%s，您好！
		</div>
		<div style="padding: 8px 40px 8px 50px;">
			<p>您于 %s 提交的重置密码申请，点击或者复制链接<u><strong>%s</strong></u>访问，为了保证账号安全，此链接有效期为30分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。</p>
		</div>
		<div>
			<p>此邮箱为系统邮箱，请勿回复。</p>
		</div>
	</div>
	`, email, now, forgetUrl)
	err := mail.SendGoMail([]string{email}, tilte, body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("aaa")
}

func sendForgetByPhone(phone string) {
	fmt.Println("bbb")
}

func encryptStr(val string) (string, error) {
	cipher, err := encryption.NewEncryptor(conf.APPKey, encryption.CIPHER_AES_256_CBC)
	if err != nil {
		return "", err
	}
	cryptedStr := cipher.EncryptWithCBC(val)
	return cryptedStr, nil
}
