package StoreUser

type StoreReq struct {
	Account string `binding:"required" validate:"max=2,min=2" comment:"账号" json:"account"`
	Phone   string `binding:"required" validate:"max=2,min=2" comment:"手机号" json:"phone"`
	Email   string `binding:"required,email" comment:"邮箱" json:"email"`
	Passwd  string `binding:"required" comment:"密码" json:"passwd"`
}
