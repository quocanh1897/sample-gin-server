package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func IsUUID(fl validator.FieldLevel) bool {
	id, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	_, err := uuid.Parse(id)
	return err == nil
}
