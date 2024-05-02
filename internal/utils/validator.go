package utils

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

var lock = &sync.Mutex{}

var instance *validator.Validate

func GetValidatorInstance() *validator.Validate {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = validator.New()
		}
	}

	return instance
}
