package validators

import (
	"github.com/lsnascimento/govalidations/notifications"
	"github.com/lsnascimento/govalidations/validators/core"
)

type Date struct {
	Core     core.DateValidator
	Notifier *notifications.Notifier
}

func NewDate(coreInstance core.DateValidator, notifier *notifications.Notifier) Date {
	return Date{
		Core:     coreInstance,
		Notifier: notifier,
	}
}

func (validation *Date) IsValid(value, property, message string) {
	if !validation.Core.IsValid(value) {
		validation.Notifier.AddMessage(property, message)
	}
}

func (validation *Date) IsGreaterThan(value, compare, property, message string) {
	if !validation.Core.IsGreaterThan(value, compare) {
		validation.Notifier.AddMessage(property, message)
	}
}

func (validation *Date) IsGreaterOrEqualThan(value, compare, property, message string) {
	if !validation.Core.IsGreaterOrEqualThan(value, compare) {
		validation.Notifier.AddMessage(property, message)
	}
}

func (validation *Date) IsLessThan(value, compare, property, message string) {
	if !validation.Core.IsLessThan(value, compare) {
		validation.Notifier.AddMessage(property, message)
	}
}

func (validation *Date) IsLessOrEqualThan(value, compare, property, message string) {
	if !validation.Core.IsLessOrEqualThan(value, compare) {
		validation.Notifier.AddMessage(property, message)
	}
}

func (validation *Date) IsBetween(value, from, to, property, message string) {
	if !validation.Core.IsBetween(value, from, to) {
		validation.Notifier.AddMessage(property, message)
	}
}
