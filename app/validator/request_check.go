package validator

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	zhTranslations "gopkg.in/go-playground/validator.v9/translations/zh"
	"reflect"
)

func RequestCheck(req interface{}) error {
	zhTrans := zh.New()
	trans, _ := ut.New(zhTrans, zhTrans).GetTranslator("zh")
	validate := validator.New()
	_ = zhTranslations.RegisterDefaultTranslations(validate, trans)
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return fld.Tag.Get("comment")
	})
	err := validate.Struct(req)
	if err != nil {

		// translate all error at once
		errs := err.(validator.ValidationErrors)

		// returns a map with key = namespace & value = translated error
		// NOTICE: 2 errors are returned and you'll see something surprising
		// translations are i18n aware!!!!
		// eg. '10 characters' vs '1 character'

		fmt.Println(errs.Translate(trans))
		return errs
	}
	return nil
}
