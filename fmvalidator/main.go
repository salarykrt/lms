package fmvalidator

import "github.com/go-playground/validator/v10"

type Validator struct {
	Name    string
	Allowed []string
}

func (v Validator) IsValid(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	for _, a := range v.Allowed {
		if value == a {
			return true
		}
	}
	return false
}

func NewValidator(name string, allowed []string) Validator {
	return Validator{
		Name:    name,
		Allowed: allowed,
	}
}
