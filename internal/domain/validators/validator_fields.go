package validators

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	validateOnce     sync.Once
	validateInstance *validator.Validate
)

func getValidator() *validator.Validate {
	validateOnce.Do(func() {
		validateInstance = validator.New()
	})
	return validateInstance
}

func ValidatorFields(data interface{}) error {
	validate := getValidator()
	return validate.Struct(data)
}
