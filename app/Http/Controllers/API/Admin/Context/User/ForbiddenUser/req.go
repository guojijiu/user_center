package ForbiddenUser

type Req struct {
	ID          uint `binding:"required" comment:"用户id" json:"id"`
	IsForbidden bool `binding:"required" comment:"是否禁用，true：是，false：否" json:"is_forbidden"`
}
