package req

import (
	"github.com/go-playground/validator/v10"
)

func IsValid[T any](payload T) error {
	validte := validator.New()
	err := validte.Struct(payload)
	return err
}
