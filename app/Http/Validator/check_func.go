package validator

import (
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
)

func IsMobile(fl validator.FieldLevel) bool {
	v := fl.Field().String()
	if len(v) != 0 {
		return MobileRegex.MatchString(v)
	}
	return true
}

// TranslationFunc is the function type used to register or override
//// custom translations
//type TranslationFunc func(ut ut.Translator, fe FieldError) string
//
//// RegisterTranslationsFunc allows for registering of translations
//// for a 'ut.Translator' for use within the 'TranslationFunc'
//type RegisterTranslationsFunc func(ut ut.Translator) error

func TranslateErrorShow(ut ut.Translator) error {
	return ut.Add("mobile", "{0}不是一个有效的手机号码(目前暂仅支持中国大陆手机号码)", true) // see universal-translator for details
}

func TranslateField(ut ut.Translator, fe validator.FieldError) string {
	t, _ := ut.T("mobile", fe.Field())
	return t
}

func IsRequiredMobile(fl validator.FieldLevel) bool {
	return MobileRegex.MatchString(fl.Field().String())
}
