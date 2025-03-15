package api

import (
	"simple-bank/util"

	"github.com/go-playground/validator/v10"
)

// validCurrency a custom validator to check currency validity
var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// check currency is supported
		return util.IsSupportedCurrency(currency)
	}
	return false
}
