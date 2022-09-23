package DeletePermission

type Resp struct {
	Account  string `comment:"账号" json:"account"`
	Phone    string `comment:"手机号" json:"phone"`
	Email    string `comment:"邮箱" json:"email"`
	Nickname string `comment:"昵称" json:"nickname"`
}
