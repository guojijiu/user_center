package BindOrganize

type Req struct {
	ID          uint   `binding:"required" comment:"客户端id" json:"id"`
	OrganizeIDs []uint `binding:"required" validate:"required" comment:"组织id组" json:"organize_ids"`
}
