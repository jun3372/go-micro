package tool

import (
	"errors"
	"reflect"

	"github.com/go-playground/validator"
)

func Validate(obj interface{}) error {
	v := validator.New()
	if errs := v.Struct(obj); errs != nil {
		elem := reflect.TypeOf(obj)
		for _, err := range errs.(validator.ValidationErrors) {
			if field, b := elem.FieldByName(err.Field()); b {
				msg := field.Tag.Get("message")
				if msg == "" {
					msg = field.Tag.Get("msg")
				}

				return errors.New(msg)
			}

			return errs
		}

		return errs
	}

	return nil
}
