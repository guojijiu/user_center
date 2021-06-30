package validator

import (
	"github.com/gin-gonic/gin/binding"
	zh_cn "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
	"reflect"
)

var Trans ut.Translator
var Validate *validator.Validate

var bakedInValidatorsSelfDefined = map[string]struct {
	validateFunc    validator.Func
	transErrMsgFunc func(ut ut.Translator) error
	transFieldFunc  func(ut ut.Translator, fe validator.FieldError) string
}{
	"mobile": {
		validateFunc:    IsMobile,
		transErrMsgFunc: TranslateErrorShow,
		transFieldFunc:  TranslateField,
	},
	"required_mobile": {
		validateFunc:    IsRequiredMobile,
		transErrMsgFunc: TranslateErrorShow,
		transFieldFunc:  TranslateField,
	},
}

func init() {
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		Validate = validate
		zh := zh_cn.New()
		uni := ut.New(zh, zh)
		// this is usually know or extracted from http 'Accept-Language' header
		// also see uni.FindTranslator(...)
		Trans, _ = uni.GetTranslator("zh")
		_ = zh_translations.RegisterDefaultTranslations(Validate, Trans)

		// 将 comment 标签注册为参数的翻译
		Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			return fld.Tag.Get("comment")
		})

		// 注册自定义验证规则函数及翻译函数
		register()
	}
}

func register() {
	// TODO: waiting needed
	// must copy alias validators for separate validations to be used in each validator instance
	//for k, val := range bakedInAliases {
	//	v.RegisterAlias(k, val)
	//}

	// must copy validators for separate validations to be used in each instance
	for tag, v := range bakedInValidatorsSelfDefined {
		// no need to error check here, baked in will always be valid
		_ = Validate.RegisterValidation(tag, v.validateFunc)

		// 注册验证规则的翻译
		_ = Validate.RegisterTranslation(tag, Trans, v.transErrMsgFunc, v.transFieldFunc)
	}
}
