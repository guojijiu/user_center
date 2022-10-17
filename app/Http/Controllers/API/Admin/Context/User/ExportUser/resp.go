package ExportUser

import (
	"strconv"
	"time"
)

type Resp struct {
	ID         uint      ` comment:"用户id" json:"id"`
	Email      string    `comment:"用户邮箱" json:"email"`
	ClientName string    ` comment:"客户端名称" json:"client_name"`
	CreatedAt  time.Time ` comment:"创建时间" json:"created_at"`
	ForbadeAt  time.Time ` comment:"禁用时间" json:"forbade_at"`
}

func Result(resp []Resp) map[int][]string {
	data := make(map[int][]string)
	data[0] = []string{"用户id", "用户邮箱", "客户端名称", "创建时间", "禁用时间"}
	for k, v := range resp {
		var forbadeAt string
		if v.ForbadeAt.IsZero() == false {
			forbadeAt = v.ForbadeAt.Format("2006-01-02 15:04:05")
		}
		data[k+1] = []string{
			strconv.Itoa(int(v.ID)),
			v.Email,
			v.ClientName,
			v.CreatedAt.Format("2006-01-02 15:04:05"),
			forbadeAt,
		}
	}
	return data
}
