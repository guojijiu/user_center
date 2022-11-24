package ResetPasswd

type Req struct {
	ID        uint   `binding:"required" comment:"主键ID" json:"name"`
	Account   string `binding:"required" comment:"账号" json:"mark"`
	Email     string `binding:"omitempty,mobile,len=11" comment:"邮箱" json:"remark"`
	Phone     string `binding:"omitempty,email" comment:"手机" json:"phone"`
	Passwd    string `binding:"required" comment:"密码" json:"passwd"`
	PasswdTwo string `binding:"required" comment:"再次输入密码" json:"passwd_two"`
}
