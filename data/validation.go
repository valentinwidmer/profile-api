package data

import (
	"regexp"

	"github.com/go-playground/validator"
)

const NAME_REGEX = `[a-zA-Z][0-9a-zA-Z .,'-]*`

type Validation struct {
	validate *validator.Validate
}

func NewValidator() *Validation {
	v := validator.New()
	v.RegisterValidation("name", validateName)

	return &Validation{v}
}

func validateName(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(NAME_REGEX)
	matches := re.FindAllString(fl.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}

	return true
}

func (v *Validation) Validate(i interface{}) error {
	err := v.validate.Struct(i)
	if err != nil {
		return err
	}
	return nil
}
