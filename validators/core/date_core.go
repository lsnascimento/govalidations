package core

import (
	"time"

	"github.com/lsnascimento/govalidations/validators/utils"
)

type DateValidator struct{}

func NewDate() DateValidator {
	return DateValidator{}
}

func (validator DateValidator) IsValid(value string) bool {
	valid, _ := validator.validateDateString(value)

	return valid
}

func (validator DateValidator) IsGreaterThan(value, compare string) bool {
	valueIsValid, valueInDateTime := validator.validateDateString(value)
	compareIsValid, compareInDateTime := validator.validateDateString(compare)

	invalid := !valueIsValid || !compareIsValid

	if invalid || valueInDateTime.Before(compareInDateTime) || valueInDateTime.Equal(compareInDateTime) {
		return false
	}

	return true
}

func (validator DateValidator) IsGreaterOrEqualThan(value, compare string) bool {
	valueIsValid, valueInDateTime := validator.validateDateString(value)
	compareIsValid, compareInDateTime := validator.validateDateString(compare)

	invalid := !valueIsValid || !compareIsValid

	if invalid || valueInDateTime.Before(compareInDateTime) {
		return false
	}

	return true
}

func (validator DateValidator) IsLessThan(value, compare string) bool {
	valueIsValid, valueInDateTime := validator.validateDateString(value)
	compareIsValid, compareInDateTime := validator.validateDateString(compare)

	invalid := !valueIsValid || !compareIsValid

	if invalid || valueInDateTime.After(compareInDateTime) || valueInDateTime.Equal(compareInDateTime) {
		return false
	}

	return true
}

func (validator DateValidator) IsLessOrEqualThan(value, compare string) bool {
	valueIsValid, valueInDateTime := validator.validateDateString(value)
	compareIsValid, compareInDateTime := validator.validateDateString(compare)

	invalid := !valueIsValid || !compareIsValid

	if invalid || valueInDateTime.After(compareInDateTime) {
		return false
	}

	return true
}

func (validator DateValidator) IsBetween(value, from, to string) bool {
	valueIsValid, valueInDateTime := validator.validateDateString(value)
	fromIsValid, fromInDateTime := validator.validateDateString(from)
	toIsValid, toInDateTime := validator.validateDateString(to)

	invalid := !valueIsValid || !fromIsValid || !toIsValid

	if invalid || valueInDateTime.Before(fromInDateTime) || valueInDateTime.After(toInDateTime) {
		return false
	}

	return true
}

func (validator DateValidator) validateDateString(value string) (valid bool, dateTime time.Time) {
	valid, dateTime = utils.IsDate(value)
	return
}
