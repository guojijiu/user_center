package ListDepartment

type Req struct {
	Page int    `form:"page" binding:"required,max=999,min=1" comment:"当前页" json:"page"`
	Size int    `form:"size" binding:"required,max=999,min=1" comment:"每页显示条数" json:"size"`
	Name string `validate:"max=64,min=2" comment:"名称" json:"name"`
}
