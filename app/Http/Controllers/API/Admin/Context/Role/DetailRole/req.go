package DetailRole

type Req struct {
	ID uint `binding:"required" comment:"角色id" json:"id" form:"id"`
}
