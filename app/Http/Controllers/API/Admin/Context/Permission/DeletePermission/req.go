package DeletePermission

type Req struct {
	ID uint `binding:"required" comment:"权限id" json:"id" form:"id"`
}
