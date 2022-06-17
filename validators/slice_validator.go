package validators

import (
	"reflect"

	"github.com/lnascimento1988/govalidations/notifications"
)

type Slice struct {
	Notifier *notifications.Notifier
}

func NewSlice(notifier *notifications.Notifier) Slice {
	return Slice{
		Notifier: notifier,
	}
}

func (validation *Slice) IsRequired(value interface{}, property, message string) {
	val := toSlice(value)

	if len(val) == 0 {
		validation.Notifier.AddNotification(property, message)
	}
}

func toSlice(value interface{}) (slice []interface{}) {
	element := reflect.ValueOf(value)
	elementsOfSlice := reflect.ValueOf(element.Interface())
	quantityElements := elementsOfSlice.Len()

	for i := 0; i < quantityElements; i++ {
		slice = append(slice, elementsOfSlice.Index(i).Interface())
	}

	return
}
