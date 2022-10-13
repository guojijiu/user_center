package GetBindClient

type Result struct {
	ID   uint   `binding:"required" comment:"客户端id" json:"id"`
	Name string `binding:"required" comment:"客户端名称" json:"name"`
}
