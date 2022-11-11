package StoreDepartment

type Req struct {
	Name   string `binding:"required,max=64,min=2" comment:"名称" json:"name"`
	Mark   string `binding:"required,max=32,min=2" comment:"客户端唯一标识" json:"mark"`
	Remark string `binding:"max=255,min=2" comment:"备注" json:"remark"`
}
