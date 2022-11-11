package GetBindOrganize

type Result struct {
	ID   uint   `binding:"required" comment:"组织id" json:"id"`
	Name string `binding:"required" comment:"组织名称" json:"name"`
}
