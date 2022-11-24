package validator

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/jabardigitalservice/utilities-go/lang"
)

type Validator struct {
	Validate *validator.Validate
	Lang     lang.Lang
}

func New(lang lang.Lang) Validator {
	return Validator{validator.New(), lang}
}

func (v Validator) Validation(args interface{}) interface{} {
	errors := map[string]string{}

	err := v.Validate.Struct(args)

	if err == nil {
		return nil
	}

	if _, ok := err.(*validator.InvalidValidationError); ok {
		return err
	}

	errs := err.(validator.ValidationErrors)

	for _, err := range errs {
		msg, errMsg := v.Lang.GetMessage(err.Tag(), map[string]interface{}{
			"Field": err.Field(),
		})

		re := regexp.MustCompile(`Error:(.+)`)
		match := re.FindStringSubmatch(err.Error())
		message := match[1]

		if errMsg == nil {
			message = msg
		}

		key := strings.ToLower(err.StructNamespace())
		errors[key] = message
	}

	return errors
}
