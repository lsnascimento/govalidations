package validators

import "github.com/lsnascimento/govalidations/notifications"

type Bool struct {
	Notifier *notifications.Notifier
}

func NewBool(notifier *notifications.Notifier) Bool {
	return Bool{
		Notifier: notifier,
	}
}

func (validation *Bool) IsRequired(value *bool, property, message string) {
	if value == nil {
		validation.Notifier.AddMessage(property, message)
	}
}

func (validation *Bool) IsTrue(value bool, property, message string) {
	if !value {
		validation.Notifier.AddMessage(property, message)
	}
}

func (validation *Bool) IsFalse(value bool, property, message string) {
	if value {
		validation.Notifier.AddMessage(property, message)
	}
}
