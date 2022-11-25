package validator

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/jabardigitalservice/utilities-go/array"
)

type Validator struct {
	validator *validator.Validate
}

func New() Validator {
	return Validator{validator.New()}
}

func (v Validator) Validate(args interface{}) interface{} {
	errors := map[string]string{}

	err := v.validator.Struct(args)

	if err == nil {
		return nil
	}

	if _, ok := err.(*validator.InvalidValidationError); ok {
		return err
	}

	errs := err.(validator.ValidationErrors)

	for _, err := range errs {
		message := getMessage(err.Error())
		key := getKey(err.StructNamespace())
		errors[key] = message
	}

	return errors
}

func getMessage(error string) string {
	re := regexp.MustCompile(`Error:(.+)`)
	match := re.FindStringSubmatch(error)
	message := match[1]
	return message
}

func getKey(keyError string) string {
	key := strings.ToLower(keyError)
	keys := strings.Split(key, ".")
	newKeys := array.RemoveIndex(keys, 0)
	key = strings.Join(newKeys, ".")

	return key
}
