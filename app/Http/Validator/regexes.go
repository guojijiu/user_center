package validator

import "regexp"

// 在写我们uims的正则之前，先看一下validator包内置的正则
// gopkg.in/go-playground/validator.v9@v9.29.1/regexes.go
const (
	MobileRegexString = `^1[3-9][0-9]{9}$`
)

var (
	MobileRegex = regexp.MustCompile(MobileRegexString)
)
