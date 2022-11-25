package Register

type Req struct {
	Type        uint   `binding:"required,oneof=1 2" comment:"类型" json:"type"`
	Phone       string `binding:"required,mobile,len=11" comment:"手机号" json:"phone"`
	Email       string `binding:"required,email" comment:"邮箱" json:"email"`
	CaptchaCode string `binding:"required,len=6" comment:"验证码" json:"captcha_code"`
	VerifyCode  string `binding:"required,len=6" comment:"激活码" json:"verify_code"`
	Account     string `binding:"required,alphanum,max=16" comment:"账号" json:"account"`
	Passwd      string `binding:"required,alphanum,max=32" comment:"密码" json:"passwd"`
	Nickname    string `binding:"len=32" comment:"昵称" json:"nickname"`
}
