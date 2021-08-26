package UpdateUser

type Req struct {
	ID              uint   `binding:"required" comment:"用户id" json:"id"`
	Account         string `validate:"max=64,min=2" comment:"账号" json:"account"`
	Phone           string `validate:"mobile,len=11" comment:"手机号" json:"phone"`
	Email           string `validate:"email" validate:"required,email" comment:"邮箱" json:"email"`
	Nickname        string `validate:"max=32,min=1" comment:"昵称" json:"nickname"`
	HeaderImgPath   string `validate:"max=64,min=2" comment:"用户头像" json:"header_img_path"`
	Sex             string `validate:"len=1" comment:"性别" json:"sex"`
	Birthday        string `validate:"max=32" comment:"出生年月日" json:"birthday"`
	Address         string `validate:"max=64" comment:"出生年月日" json:"address"`
	Organization    string `validate:"max=32" comment:"公司，组织或院校" json:"organization"`
	PersonalProfile string `validate:"max=32" comment:"个人简介" json:"personal_profile"`
}
