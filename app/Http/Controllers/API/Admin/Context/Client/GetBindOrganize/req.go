package GetBindOrganize

type Req struct {
	ID uint `binding:"required" comment:"客户端id" query:"id"`
}
