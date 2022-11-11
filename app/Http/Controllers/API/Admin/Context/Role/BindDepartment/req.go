package BindDepartment

type Req struct {
	ID            uint   `binding:"required" comment:"角色id" json:"id"`
	DepartmentIDs []uint `binding:"required" validate:"required" comment:"部门id组" json:"department_ids"`
}
