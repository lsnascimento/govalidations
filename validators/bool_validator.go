package validators

import (
	"github.com/lsnascimento/govalidations/notifications"
	"github.com/lsnascimento/govalidations/validators/core"
)

type Bool struct {
	Core     core.BoolValidator
	Notifier *notifications.Notifier
}

func NewBool(coreInstance core.BoolValidator, notifier *notifications.Notifier) Bool {
	return Bool{
		Notifier: notifier,
		Core:     coreInstance,
	}
}

func (validation *Bool) IsRequired(value *bool, property, message string) {
	if validation.Core.IsNull(value) {
		validation.Notifier.AddNotification(property, message)
	}
}

func (validation *Bool) IsTrue(value bool, property, message string) {
	if validation.Core.IsFalse(value) {
		validation.Notifier.AddNotification(property, message)
	}
}

func (validation *Bool) IsFalse(value bool, property, message string) {
	if validation.Core.IsTrue(value) {
		validation.Notifier.AddNotification(property, message)
	}
}
