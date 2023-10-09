package consumer_ucase

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"technical_test_go/technical_test_go/internal/presentations"
	"technical_test_go/technical_test_go/internal/validator"
)

func validateParams(param presentations.PayloadConsumer) error {
	rules := []*validation.FieldRules{
		validation.Field(&param.FullName, validation.Required, validation.Length(2, 60), validator.ValidateHumanName()),
		validation.Field(&param.NIK, validation.Required, validator.ValidateDigit()),
		validation.Field(&param.Salary, validation.Required),
		validation.Field(&param.Dob, validation.Required),
		validation.Field(&param.LegalName, validation.Required, validation.Length(2, 10)),
		validation.Field(&param.Pob, validation.Required),
	}

	err := validation.ValidateStruct(&param, rules...)
	var ve validation.Errors
	ok := errors.As(err, &ve)
	if !ok {
		ve = make(validation.Errors)
	}

	if len(ve) == 0 {
		return nil
	}

	return ve

}
