package validators

import (
	"github.com/lsnascimento/govalidations/notifications"
	"github.com/lsnascimento/govalidations/validators/core"
)

type Int struct {
	Core     core.IntValidator
	Notifier *notifications.Notifier
}

func NewInt(coreInstance core.IntValidator, notifier *notifications.Notifier) Int {
	return Int{
		Core:     coreInstance,
		Notifier: notifier,
	}
}

func (validation *Int) IsRequired(value *int, property, message string) {
	if validation.Core.IsNull(value) {
		validation.Notifier.AddMessage(property, message)
	}
}

func (validation *Int) IsGreaterThan(value, compare int, property, message string) {
	if validation.Core.IsLessOrEqualThan(value, compare) {
		validation.Notifier.AddMessage(property, message)
	}
}

func (validation *Int) IsGreaterOrEqualThan(value, compare int, property, message string) {
	if validation.Core.IsLessThan(value, compare) {
		validation.Notifier.AddMessage(property, message)
	}
}

func (validation *Int) IsLessThan(value, compare int, property, message string) {
	if validation.Core.IsGreaterOrEqualThan(value, compare) {
		validation.Notifier.AddMessage(property, message)
	}
}

func (validation *Int) IsLessOrEqualThan(value, compare int, property, message string) {
	if validation.Core.IsGreaterThan(value, compare) {
		validation.Notifier.AddMessage(property, message)
	}
}

func (validation *Int) AreEquals(value, compare int, property, message string) {
	if validation.Core.AreNotEquals(value, compare) {
		validation.Notifier.AddMessage(property, message)
	}
}

func (validation *Int) AreNotEquals(value, compare int, property, message string) {
	if validation.Core.AreEquals(value, compare) {
		validation.Notifier.AddMessage(property, message)
	}
}

func (validation *Int) IsBetween(value, from, to int, property, message string) {
	if !validation.Core.IsBetween(value, from, to) {
		validation.Notifier.AddMessage(property, message)
	}
}
