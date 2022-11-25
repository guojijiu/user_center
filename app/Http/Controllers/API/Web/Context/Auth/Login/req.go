package Login

type Req struct {
	Account        string `binding:"required,alphanum,max=16" comment:"账号" json:"account"`
	Passwd         string `binding:"required,alphanum,max=32" comment:"密码" json:"passwd"`
	CaptchaCode    string `binding:"required,len=6" comment:"验证码" json:"captcha_code"`
	CaptchaCodeKey string `binding:"required" comment:"验证码key" json:"captcha_code_key"`
}
