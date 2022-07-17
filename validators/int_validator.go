package validators

import (
	"github.com/lsnascimento/govalidations/notifications"
)

type Int struct {
	Notifier *notifications.Notifier
}

func NewInt(notifier *notifications.Notifier) Int {
	return Int{
		Notifier: notifier,
	}
}

func (validation *Int) IsRequired(value *int, property, message string) {
	if value == nil {
		validation.Notifier.AddMessage(property, message)
	}
}

func (validation *Int) IsGreaterThan(value, compare int, property, message string) {
	if value <= compare {
		validation.Notifier.AddMessage(property, message)
	}
}

func (validation *Int) IsGreaterOrEqualThan(value, compare int, property, message string) {
	if value < compare {
		validation.Notifier.AddMessage(property, message)
	}
}

func (validation *Int) IsLessThan(value, compare int, property, message string) {
	if value >= compare {
		validation.Notifier.AddMessage(property, message)
	}
}

func (validation *Int) IsLessOrEqualThan(value, compare int, property, message string) {
	if value > compare {
		validation.Notifier.AddMessage(property, message)
	}
}

func (validation *Int) AreEquals(value, compare int, property, message string) {
	if value != compare {
		validation.Notifier.AddMessage(property, message)
	}
}

func (validation *Int) AreNotEquals(value, compare int, property, message string) {
	if value == compare {
		validation.Notifier.AddMessage(property, message)
	}
}

func (validation *Int) IsBetween(value, from, to int, property, message string) {
	if value < from || value > to {
		validation.Notifier.AddMessage(property, message)
	}
}
