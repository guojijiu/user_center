package StoreUser

type StoreReq struct {
	Account  string `binding:"required,string,min:4,max:64" comment:"账号" json:"account"`
	Phone    string `binding:"required_mobile" comment:"手机号" json:"phone"`
	Email    string `binding:"required,string,email" comment:"邮箱" json:"email"`
	Passwd string `binding:"required,string,min:6,max:64" comment:"密码" json:"passwd"`
}
