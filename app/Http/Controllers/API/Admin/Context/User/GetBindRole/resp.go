package GetBindRole

type Result struct {
	ID   uint   `binding:"required" comment:"角色id" json:"id"`
	Name string `binding:"required" comment:"角色名称" json:"name"`
}
