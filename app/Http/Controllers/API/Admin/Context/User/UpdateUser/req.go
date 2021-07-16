package UpdateUser

type Req struct {
	ID      uint   `binding:"required" comment:"用户id" json:"id"`
	Account string `binding:"max=64,min=2" comment:"账号" json:"account"`
	Phone   string `binding:"mobile,len=11" comment:"手机号" json:"phone"`
	Email   string `binding:"email" validate:"required,email" comment:"邮箱" json:"email"`
}
