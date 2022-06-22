package validators

import (
	"github.com/lsnascimento/govalidations/notifications"
	"github.com/lsnascimento/govalidations/validators/core"
)

type Slice struct {
	Core     core.SliceValidator
	Notifier *notifications.Notifier
}

func NewSlice(coreInstance core.SliceValidator, notifier *notifications.Notifier) Slice {
	return Slice{
		Core:     coreInstance,
		Notifier: notifier,
	}
}

func (validation *Slice) IsRequired(value interface{}, property, message string) {
	if validation.Core.IsEmpty(value) {
		validation.Notifier.AddNotification(property, message)
	}
}
