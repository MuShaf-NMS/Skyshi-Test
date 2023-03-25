package helper

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

var Validator = validator.New()

// Helper to validate struct
func Validate(s interface{}) error {
	return Validator.Struct(s)
}

func msgForTag(e validator.FieldError) string {
	switch e.ActualTag() {
	case "required":
		// return fmt.Sprintf("%s cannot be null", strings.ToLower(e.Field()))
		return fmt.Sprintf("%s cannot be null", ToSnakeCase(e.Field()))
	case "email":
		return "invalid email"
	}
	return ""
}

// Helper to extrac errors from struct error
func ValidationError(err error) []string {
	errMsg := []string{}
	for _, e := range err.(validator.ValidationErrors) {
		errMsg = append(errMsg, msgForTag(e))
	}
	return errMsg
}
