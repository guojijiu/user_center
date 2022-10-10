package BindPermission

type Req struct {
	ID            uint   `binding:"required" comment:"角色id" json:"id"`
	PermissionIDs []uint `binding:"required" comment:"权限id集" json:"permission_ids"`
}
