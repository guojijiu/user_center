package StoreRole

type Req struct {
	ClientID uint   `binding:"required" comment:"所属客户端" json:"client_id"`
	Name     string `binding:"required,max=16,min=1" comment:"角色名称" json:"name"`
	Sort     uint   `binding:"required,max=64,min=1" validate:"required,email" comment:"排序序号" json:"sort"`
	Mark     string `binding:"required,max=32,min=1" comment:"角色唯一标识" json:"mark"`
	Remark   string `validate:"max=255,min=1" comment:"备注" json:"remark"`
}
