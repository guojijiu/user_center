package validator

import "regexp"

// 在写我们正则之前，先看一下validator包内置的正则
const (
	MobileRegexString = `^1[3-9][0-9]{9}$`
)

var (
	MobileRegex = regexp.MustCompile(MobileRegexString)
)
