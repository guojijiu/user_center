package StoreUser

type StoreReq struct {
	Account string `binding:"required,max=64,min=2" comment:"账号" json:"account"`
	Phone   string `binding:"required,mobile,len=11" comment:"手机号" json:"phone"`
	Email   string `binding:"required,email" validate:"required,email" comment:"邮箱" json:"email"`
	Passwd  string `binding:"required,max=64,min=6" comment:"密码" json:"passwd"`
}
