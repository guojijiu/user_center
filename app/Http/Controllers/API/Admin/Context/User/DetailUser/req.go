package DetailUser

type Req struct {
	ID uint `binding:"required" comment:"用户id" json:"id"`
}