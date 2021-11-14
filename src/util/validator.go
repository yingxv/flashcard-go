package util

import (
	"reflect"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

// NewValidator 工厂方法
func NewValidator() *validator.Validate {
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")

	validate := validator.New()
	//注册一个函数，获取struct tag里自定义的label作为字段名
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})

	//注册翻译器
	err := zh_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		panic(err)
	}

	return validate

}
