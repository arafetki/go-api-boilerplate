package validator

import (
	"github.com/go-playground/validator/v10"
)

type Wrapper struct {
	validator *validator.Validate
}

func (v *Wrapper) Validate(i any) error {
	if err := v.validator.Struct(i); err != nil {
		return err
	}
	return nil
}

func New() *Wrapper {
	return &Wrapper{validator: validator.New()}
}
