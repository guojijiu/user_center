package GetBindRole

type Req struct {
	ID uint `binding:"required" comment:"用户id" query:"id"`
}
