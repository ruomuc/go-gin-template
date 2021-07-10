package validate

import (
	"reflect"
	"strings"
	"ticket-crawler/pkg/setting"

	en2 "github.com/go-playground/locales/en"
	zh2 "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en3 "github.com/go-playground/validator/v10/translations/en"
	zhTranslation "github.com/go-playground/validator/v10/translations/zh"
)

var (
	V     *validator.Validate
	trans ut.Translator
)

func InitTrans() {
	lang := setting.AppSetting.ValidatorLanguage
	// 多语言
	zh := zh2.New()
	en := en2.New()
	uni := ut.New(zh, en)

	switch lang {
	case "zh":
		trans, _ = uni.GetTranslator("zh")
	default:
		trans, _ = uni.GetTranslator("en")
	}

	//var ok bool
	//if V, ok = binding.Validator.Engine().(*validator.Validate); ok {
	// 注册翻译器
	V = validator.New()
	switch lang {
	case "zh":
		_ = zhTranslation.RegisterDefaultTranslations(V, trans)
	default:
		_ = en3.RegisterDefaultTranslations(V, trans)
	}
	//注册获取json tag的自定义方法
	V.RegisterTagNameFunc(func(sf reflect.StructField) string {
		name := strings.SplitN(sf.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

func translate(trans ut.Translator, fe validator.FieldError) string {
	msg, err := trans.T(fe.Tag(), fe.Field())
	if err != nil {
		panic(fe.(error).Error())
	}
	return msg
}

func Translate(errs validator.ValidationErrors) []string {
	var res []string
	for _, err := range errs {
		res = append(res, err.Translate(trans))
	}
	return res
}
