package validate

import (
	en2 "github.com/go-playground/locales/en"
	zh2 "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en3 "github.com/go-playground/validator/v10/translations/en"
	zhTranslation "github.com/go-playground/validator/v10/translations/zh"
	"ticket-crawler/pkg/setting"
)

var (
	Validator *validator.Validate
	trans     ut.Translator
)

func InitTrans() {
	lang := setting.AppSetting.ValidatorLanguage
	switch lang {
	case "zh":
		zh := zh2.New()
		uni := ut.New(zh, zh)
		trans, _ = uni.GetTranslator("zh")
	default:
		en := en2.New()
		uni := ut.New(en, en)
		trans, _ = uni.GetTranslator("en")
	}

	Validator = validator.New()
	// 注册翻译器
	switch lang {
	case "zh":
		_ = zhTranslation.RegisterDefaultTranslations(Validator, trans)
	default:
		_ = en3.RegisterDefaultTranslations(Validator, trans)
	}
	// 自定义验证方法
}

func Translate(errs validator.ValidationErrors) []string {
	var res []string
	for _, err := range errs {
		res = append(res, err.Translate(trans))
	}
	return res
}
