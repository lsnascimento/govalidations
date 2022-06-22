package validators

import (
	"github.com/lsnascimento/govalidations/notifications"
	"github.com/lsnascimento/govalidations/validators/core"
)

type String struct {
	Core     core.StringValidator
	Notifier *notifications.Notifier
}

func NewString(coreInstance core.StringValidator, notifier *notifications.Notifier) String {
	return String{
		Core:     coreInstance,
		Notifier: notifier,
	}
}

func (validation *String) IsNotNullOrEmpty(value, property, message string) {
	if validation.Core.IsNullOrEmpty(value) {
		validation.Notifier.AddNotification(property, message)
	}
}

func (validation *String) IsNullOrEmpty(value, property, message string) {
	if validation.Core.IsNotNullOrEmpty(value) {
		validation.Notifier.AddNotification(property, message)
	}
}

func (validation *String) HasMinLength(value string, min int, property, message string) {
	if !validation.Core.HasMinLength(value, min) {
		validation.Notifier.AddNotification(property, message)
	}
}

func (validation *String) HasMaxLength(value string, max int, property, message string) {
	if !validation.Core.HasMaxLength(value, max) {
		validation.Notifier.AddNotification(property, message)
	}
}

func (validation *String) HasLength(value string, compare int, property, message string) {
	if !validation.Core.HasLength(value, compare) {
		validation.Notifier.AddNotification(property, message)
	}
}

func (validation *String) Contains(value, compare, property, message string) {
	if !validation.Core.Contains(value, compare) {
		validation.Notifier.AddNotification(property, message)
	}
}

func (validation *String) AreEquals(value, compare, property, message string) {
	if validation.Core.AreNotEquals(value, compare) {
		validation.Notifier.AddNotification(property, message)
	}
}

func (validation *String) AreNotEquals(value, compare, property, message string) {
	if validation.Core.AreEquals(value, compare) {
		validation.Notifier.AddNotification(property, message)
	}
}

func (validation *String) IsEmail(value, property, message string) {
	if !validation.Core.IsEmail(value) {
		validation.Notifier.AddNotification(property, message)
	}
}

func (validation *String) Match(value, pattern, property, message string) {
	if !validation.Core.Match(value, pattern) {
		validation.Notifier.AddNotification(property, message)
	}
}
