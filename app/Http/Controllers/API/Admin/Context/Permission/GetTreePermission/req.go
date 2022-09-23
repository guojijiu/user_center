package GetTreePermission

type Req struct {
	ClientID uint `binding:"required" comment:"客户端id" query:"client_id"`
}
