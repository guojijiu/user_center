package GetBindPermission

type Result struct {
	ID   uint   `binding:"required" comment:"权限id" json:"id"`
	Name string `binding:"required" comment:"权限名称" json:"name"`
	Mark string `binding:"required" comment:"权限标识" json:"mark"`
}
