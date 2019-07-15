package httpvalidator

import "gopkg.in/go-playground/validator.v9"

type CustomValidator struct {
	validator *validator.Validate
}

func NewHTTPValidator() *CustomValidator {
	return &CustomValidator{
		validator: validator.New(),
	}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
