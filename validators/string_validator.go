package validators

import (
	"github.com/lnascimento1988/govalidations/notifications"
)

type String struct {
	Notifier *notifications.Notifier
}

func NewString(notifier *notifications.Notifier) String {
	return String{
		Notifier: notifier,
	}
}

func (validation *String) IsNotNullOrEmpty(value, property, message string) {
	if value == "" || len(value) == 0 {
		validation.Notifier.AddNotification(property, message)
	}
}

func (validation *String) IsNullOrEmpty(value, property, message string) {
	if value != "" || len(value) > 0 {
		validation.Notifier.AddNotification(property, message)
	}
}

func (validation *String) HasMinLength(value string, min int, property, message string) {
	if len(value) < min {
		validation.Notifier.AddNotification(property, message)
	}
}

func (validation *String) HasMaxLength(value string, max int, property, message string) {
	if len(value) > max {
		validation.Notifier.AddNotification(property, message)
	}
}

func (validation *String) HasLength(value string, num int, property, message string) {
	if len(value) != num {
		validation.Notifier.AddNotification(property, message)
	}
}

func (validation *String) Contains(value, compare, property, message string) {
	if !Contains(value, compare) {
		validation.Notifier.AddNotification(property, message)
	}
}

func (validation *String) AreEquals(value, compare, property, message string) {
	if value != compare {
		validation.Notifier.AddNotification(property, message)
	}
}

func (validation *String) AreNotEquals(value, compare, property, message string) {
	if value == compare {
		validation.Notifier.AddNotification(property, message)
	}
}

func (validation *String) IsEmail(value, property, message string) {
	pattern := `^\w+([-+.']\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`

	if !Matches(value, pattern) {
		validation.Notifier.AddNotification(property, message)
	}
}

func (validation *String) Match(value, pattern, property, message string) {
	if !Matches(value, pattern) {
		validation.Notifier.AddNotification(property, message)
	}
}
