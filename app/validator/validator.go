package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
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

func Init() {

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		Validate = v
		uni := ut.New(zh.New())
		Trans, _ = uni.GetTranslator("zh")

		//注册翻译器
		_ = zh_translations.RegisterDefaultTranslations(v, Trans)

		//注册一个函数，获取struct tag里自定义的label作为字段名
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := fld.Tag.Get("comment")
			return name
		})
		//注册自定义函数
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
	for tag, valida := range bakedInValidatorsSelfDefined {
		// no need to error check here, baked in will always be valid
		_ = Validate.RegisterValidation(tag, valida.validateFunc)

		// 注册验证规则的翻译
		_ = Validate.RegisterTranslation(tag, Trans, valida.transErrMsgFunc, valida.transFieldFunc)
	}
}
