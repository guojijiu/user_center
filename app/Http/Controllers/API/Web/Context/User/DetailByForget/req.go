package DetailByForget

type Req struct {
	ForgetType string `binding:"required" comment:"类型" json:"forget_type"`
	ForgetData string `binding:"required" comment:"请求数据" json:"forget_data"`
}
