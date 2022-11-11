package DetailDepartment

type Req struct {
	ID uint `binding:"required" comment:"用户id" json:"id" form:"id"`
}
