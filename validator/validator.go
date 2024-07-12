package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func PasswordRegex(fl validator.FieldLevel) bool {
	return regexp.MustCompile(PasswordRegexString).MatchString(fl.Field().String())
}

func DateTimeRegex(fl validator.FieldLevel) bool {
	return regexp.MustCompile(DateTimeRegexString).MatchString(fl.Field().String())
}

func NotContainFourByteCharacterRegex(fl validator.FieldLevel) bool {
	return !regexp.MustCompile(ContainFourByteCharacterRegexString).MatchString(fl.Field().String())
}
