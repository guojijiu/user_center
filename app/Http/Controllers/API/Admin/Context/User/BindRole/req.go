package BindRole

type Req struct {
	ID      uint   `binding:"required" comment:"用户id" json:"id"`
	RoleIDs []uint `binding:"required" validate:"required" comment:"角色id组" json:"role_ids"`
}
