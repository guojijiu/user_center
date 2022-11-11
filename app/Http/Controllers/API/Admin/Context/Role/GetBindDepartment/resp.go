package GetBindDepartment

type Result struct {
	ID   uint   `binding:"required" comment:"部门id" json:"id"`
	Name string `binding:"required" comment:"部门名称" json:"name"`
}
