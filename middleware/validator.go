package middleware

import "gopkg.in/go-playground/validator.v9"

// CustomValidator
type CustomValidator struct {
	Validator *validator.Validate
}

// Validate : Validate Data
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
