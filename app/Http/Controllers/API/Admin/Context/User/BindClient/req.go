package BindClient

type Req struct {
	ID        uint   `binding:"required" comment:"用户id" json:"id"`
	ClientIDs []uint `binding:"required" validate:"required" comment:"客户端id组" json:"client_ids"`
}
