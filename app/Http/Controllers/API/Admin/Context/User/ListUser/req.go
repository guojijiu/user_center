package ListUser

type Req struct {
	Page     int    `form:"page" binding:"required,max=999,min=1" comment:"当前页" json:"page"`
	Size     int    `form:"size" binding:"required,max=999,min=1" comment:"每页显示条数" json:"size"`
	Account  string `validate:"max=64,min=2" comment:"账号" json:"account"`
	Phone    string `validate:"mobile,len=11" comment:"手机号" json:"phone"`
	Email    string `validate:"email" validate:"required,email" comment:"邮箱" json:"email"`
	Nickname string `validate:"max=32,min=1" comment:"昵称" json:"nickname"`
}
