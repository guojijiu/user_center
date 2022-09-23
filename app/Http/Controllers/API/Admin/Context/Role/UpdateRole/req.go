package UpdateRole

type Req struct {
	ID     uint   `binding:"required" comment:"角色id" json:"id"`
	Name   string `binding:"max=16,min=1" comment:"角色名称" json:"name"`
	Sort   uint   `binding:"max=999,min=1" validate:"required,email" comment:"排序序号" json:"sort"`
	Mark   string `binding:"max=32,min=1" comment:"角色唯一标识" json:"mark"`
	Remark string `validate:"max=255,min=1" comment:"备注" json:"remark"`
}
