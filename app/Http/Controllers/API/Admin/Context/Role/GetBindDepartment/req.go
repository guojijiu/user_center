package GetBindDepartment

type Req struct {
	ID uint `binding:"required" comment:"角色id" query:"id"`
}
