package api

import (
	"RD-Clone-NAPI/internal/api/validator"
)

type API struct {
	validator *validator.CustomValidator
}

func New() *API {
	v := validator.GetValidator()
	err := validator.AddValidators(v.Validate)
	if err != nil {
		panic(err)
	}

	return &API{validator: v}
}
