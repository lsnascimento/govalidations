package validators

import (
	"github.com/lsnascimento/govalidations/notifications"
)

type Date struct {
	Notifier *notifications.Notifier
}

func NewDate(notifier *notifications.Notifier) Date {
	return Date{
		Notifier: notifier,
	}
}

func (validation *Date) IsValid(value, property, message string) {
	valid, _ := IsDate(value)

	if !valid {
		validation.Notifier.AddNotification(property, message)
	}
}

func (validation *Date) IsGreaterThan(value, compare, property, message string) {
	valueValid, check := IsDate(value)
	compareValid, start := IsDate(compare)

	invalid := !valueValid || !compareValid

	if invalid || check.Before(start) || check.Equal(start) {
		validation.Notifier.AddNotification(property, message)
	}
}

func (validation *Date) IsGreaterOrEqualThan(value, compare, property, message string) {
	valueValid, check := IsDate(value)
	compareValid, start := IsDate(compare)

	invalid := !valueValid || !compareValid

	if invalid || check.Before(start) {
		validation.Notifier.AddNotification(property, message)
	}
}

func (validation *Date) IsLessThan(value, compare, property, message string) {
	valueValid, check := IsDate(value)
	compareValid, start := IsDate(compare)

	invalid := !valueValid || !compareValid

	if invalid || check.After(start) || check.Equal(start) {
		validation.Notifier.AddNotification(property, message)
	}
}

func (validation *Date) IsLessOrEqualThan(value, compare, property, message string) {
	valueValid, check := IsDate(value)
	compareValid, start := IsDate(compare)

	invalid := !valueValid || !compareValid

	if invalid || check.After(start) {
		validation.Notifier.AddNotification(property, message)
	}
}

func (validation *Date) IsBetween(value, from, to, property, message string) {
	valueValid, check := IsDate(value)
	fromValid, start := IsDate(from)
	toValid, end := IsDate(to)

	invalid := !valueValid || !fromValid || !toValid

	if invalid || check.Before(start) || check.After(end) {
		validation.Notifier.AddNotification(property, message)
	}
}
