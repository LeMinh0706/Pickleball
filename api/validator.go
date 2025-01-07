package api

import (
	"github.com/LeMinh0706/simplebank/util"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fiedLevel validator.FieldLevel) bool {
	if currency, ok := fiedLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}
