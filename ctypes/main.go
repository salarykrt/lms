package ctypes

import (
	"fmt"
	"salarykart/ctypes/leadsource"
	"salarykart/ctypes/leadstage"
	"salarykart/fmvalidator"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var validators = []fmvalidator.Validator{
	leadsource.Validator(),
	leadstage.Validator(),
}

func RegisterValidators() error {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		for _, va := range validators {
			err := v.RegisterValidation(va.Name, va.IsValid)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func GetErrorMessage(fe validator.FieldError) (string, bool) {
	for _, v := range validators {
		if v.Name == fe.Tag() {
			return fmt.Sprintf("Has to be one of %v", v.Allowed), true
		}
	}
	return "", false
}
