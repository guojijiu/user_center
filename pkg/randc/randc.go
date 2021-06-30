package randc

import (
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// 产生指定长度的随机字符串
func RandStringN(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// 生成没有中划线的uuid
func UUID() string {
	return strings.ReplaceAll(uuid.NewV4().String(), "-", "")
}
