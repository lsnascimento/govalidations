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

func (validation *Bool) IsTrue(value bool, property, message string) {
	if !value {
		validation.Notifier.AddNotification(property, message)
	}
}

func (validation *Bool) IsFalse(value bool, property, message string) {
	if value {
		validation.Notifier.AddNotification(property, message)
	}
}
