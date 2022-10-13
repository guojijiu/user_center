package ForbiddenClient

type Req struct {
	ID          uint `binding:"required" comment:"客户端id" json:"id"`
	IsForbidden uint `binding:"required" validate:"required" comment:"是否禁用，1：是，2：否；" json:"is_forbidden"`
}
