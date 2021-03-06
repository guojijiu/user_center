package tool

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

func Camel2Case(name string) string {
	buffer := bytes.Buffer{}
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.WriteRune('_')
			}
			buffer.WriteRune(unicode.ToLower(r))
		} else {
			buffer.WriteRune(r)
		}
	}
	return buffer.String()
}

// 下划线写法转为驼峰写法
func Case2Camel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

// 切片转字符串
func SliceReplaceStr(slice []string, spacer string) string {
	return strings.Replace(strings.Trim(fmt.Sprint(slice), "[]"), " ", spacer, -1)
}
