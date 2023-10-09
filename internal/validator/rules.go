// Package validator
package validator

import (
	"regexp"
	"time"

	"github.com/go-ozzo/ozzo-validation/v4"
)

const (
	AlphaNumericDash     = `^[0-9a-zA-Z\-]+$`
	AlphaNumeric         = `^[0-9a-zA-Z]+$`
	Numeric              = `^[0-9]+$`
	AlphaNumericSpace    = `^[0-9a-zA-Z\s]+$`
	Alpha                = `^[a-zA-Z]+$`
	AlphaSpace           = `^[a-zA-Z\s]+$`
	AlphaDashSpace       = `^[a-zA-Z\-\s]+$`
	IndonesianPeopleName = `^[a-zA-Z\'’.,\s]+$`
	RtRw                 = `^\d{1,3}\/\d{1,3}$`
	SubDistrict          = `^[0-9a-zA-Z\-\s\(\)]+$`
	Address              = `^[A-Za-z0-9'\.\-\s\,/#_()\[\]]+$`
	Pob                  = `^[A-Za-z'\.\-\s\,/#_()\[\]]+$`
	Email                = "^(((([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|((\\x22)((((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(([\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(\\([\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(\\x22)))@((([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
)

var (
	humanNameRegex = regexp.MustCompile(`^[a-zA-Z\'’.,\s]+$`)
	digit          = regexp.MustCompile(`^[0-9]+$`)
)

// ValidateAlphaNumericDash for reference transaction id
func validateAlphaNumericDash(v string) bool {
	pattern := `^[0-9a-zA-Z\-]+$`

	rgx, err := regexp.Compile(pattern)

	if err != nil {
		return false
	}

	return rgx.MatchString(v)
}

// validateBearerJWT for validate jwt string
func validateBearerJWT(v string) bool {
	pattern := `^Bearer ([a-zA-Z0-9_=]+)\.([a-zA-Z0-9_=]+)\.([a-zA-Z0-9_\-\+\/=]*)`

	rgx, err := regexp.Compile(pattern)

	if err != nil {
		return false
	}

	return rgx.MatchString(v)
}

func Regex(pattern string) func(v string) bool {
	return func(v string) bool {
		if len(v) == 0 {
			return true
		}
		rgx, err := regexp.Compile(pattern)

		if err != nil {
			return false
		}

		return rgx.MatchString(v)
	}
}

func validDOB(v string) bool {
	var (
		f  = []string{`2006-01-02`, `02-01-2006`}
		tm time.Time
	)

	if len(v) == 0 {
		return true
	}

	for i := 0; i < len(f); i++ {
		t, err := time.Parse(f[i], v)

		if err != nil && i == 0 {
			continue
		}

		if err != nil {
			return false
		}

		tm = t
		break
	}

	if tm.After(time.Now().AddDate(-15, 0, 0)) {
		return false
	}

	return true
}

func BearerJWT() validation.StringRule {
	return validation.NewStringRuleWithError(
		validateBearerJWT,
		validation.NewError("validation_is_bearer_jwt", "must valid Bearer JWT"))
}

func validateHumanName(v string) bool {
	if v == "" {
		return true
	}
	return humanNameRegex.MatchString(v)
}

func validateDigit(v string) bool {
	if v == "" {
		return true
	}
	return digit.MatchString(v)
}

func ValidateHumanName() validation.StringRule {
	return validation.NewStringRuleWithError(
		validateHumanName,
		validation.NewError("validation_name", "Invalid format. This field only allow these following characters: alphabet, single quote ('), space, comma(,), and period(.)."))
}

func ValidateDigit() validation.StringRule {
	return validation.NewStringRuleWithError(
		validateDigit,
		validation.NewError("validation_identity_id", "Invalid format. This field only allow these following characters: 0-9"))
}
