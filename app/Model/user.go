package Model

type User struct {
	ID             int     `gorm:"column:id;type:int(11) unsigned auto_increment;primary_key;comment:'会员或者用户ID'" json:"id"`
	OpenID         string  `gorm:"column:open_id;type:char(32);not null;default:'';unique;comment:'用户唯一ID标示'" json:"open_id"`
	UserType       string  `gorm:"column:user_type;type:char(32);not null;default:'';comment:'用户类型，取client_type字段的值，UIMS：UIMS下面的用户，VDK：微桌全平台下面的用户，CASS：结算系统下面的用户，MP：任务系统下面的用户, ESIGN: 电签系统的用户'" json:"user_type"`
	Account        string  `gorm:"column:account;type:varchar(30);not null;default:'';comment:'普通登录账号，允许最大长度16'" json:"account"`
	UserCode       string  `gorm:"column:user_code;type:char(19);not null;default:'';comment:'用户编码，全数字19位唯一id，不同的组织下可以重复'" json:"user_code"`
	NaCode         string  `gorm:"column:na_code;type:char(5);default:'';comment:'国家代码，中国：+86'" json:"na_code"`
	Phone          *string `gorm:"column:phone;type:char(11);unique;comment:'手机号'" json:"phone"`
	Email          string  `gorm:"column:email;type:varchar(32);not null;default:'';comment:'邮箱'" json:"email"`
	Salt           string  `gorm:"column:salt;type:varchar(255);not null;default:'';comment:'盐值'" json:"-"`
	EncryptType    int     `gorm:"column:encrypt_type;type:tinyint(1);not null;default:0;comment:'结算传0, 微桌系统传1'" json:"encrypt_type"`
	Passwd         string  `gorm:"column:passwd;type:varchar(512);not null;default:'';comment:'密码密文'" json:"-"`
	Passwd2        string  `gorm:"column:passwd2;type:varchar(512);not null;default:'';comment:'微桌密码密文'" json:"-"`
	Status         string  `gorm:"column:status;type:char(1);not null;default:'Y';comment:'账号的状态，默认Y：正常；N：已冻结，禁止登录'" json:"status"`
	Isdel          string  `gorm:"column:isdel;type:char(1);not null;default:'N';comment:'是否软删除，默认N：未软删除；Y：已软删除'" json:"-"`
	IdentityCardNo *string `gorm:"column:identity_card_no;type:varchar(20);unique;default:null;comment:'用户的身份证号'" json:"identity_card_no"`
}

func (User) TableName() string {
	return "uims_user_auth"
}
