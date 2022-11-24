package SendForgetCode

type Req struct {
	Type  uint   `binding:"required,oneof=1 2" comment:"类型" json:"type"`
	Phone string `binding:"omitempty,mobile,len=11" comment:"手机号" json:"phone"`
	Email string `binding:"omitempty,email" comment:"邮箱" json:"email"`
}
