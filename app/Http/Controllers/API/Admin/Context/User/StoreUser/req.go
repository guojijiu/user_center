package StoreUser

type Req struct {
	Account         string `binding:"required,max=64,min=2" comment:"账号" json:"account"`
	Phone           string `binding:"required,mobile,len=11" comment:"手机号" json:"phone"`
	Email           string `binding:"required,email" validate:"required,email" comment:"邮箱" json:"email"`
	Passwd          string `binding:"required,max=64,min=6" comment:"密码" json:"passwd"`
	Nickname        string `validate:"max=32,min=1" comment:"昵称" json:"nickname"`
	HeaderImgPath   string `validate:"max=64,min=2" comment:"用户头像" json:"header_img_path"`
	Sex             string `validate:"len=1" comment:"性别" json:"sex"`
	Birthday        string `validate:"max=32" comment:"出生年月日" json:"birthday"`
	Address         string `validate:"max=64" comment:"出生年月日" json:"address"`
	Organization    string `validate:"max=32" comment:"公司，组织或院校" json:"organization"`
	PersonalProfile string `validate:"max=32" comment:"个人简介" json:"personal_profile"`
}
