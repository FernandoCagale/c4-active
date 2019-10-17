package entity

import (
	"github.com/go-ozzo/ozzo-validation"
)

type Active struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

func (e Active) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.Name, validation.Required, validation.Length(5, 50)),
		validation.Field(&e.Code, validation.Required, validation.Length(2, 50)),
	)
}
