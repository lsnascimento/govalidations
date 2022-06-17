package validations

import (
	"reflect"

	"github.com/lnascimento1988/govalidations/notifications"
	"github.com/lnascimento1988/govalidations/validators"
)

type Parameters func(...interface{})

type Validation struct {
	Notifier   *notifications.Notifier
	Validators validators.Validator
}

func New() Validation {
	notifier := notifications.New()

	return Validation{
		Notifier:   notifier,
		Validators: validators.New(notifier),
	}
}

func (validation Validation) All(args ...interface{}) {
	for _, fn := range args {
		switch reflect.TypeOf(fn).Kind() {
		case reflect.Func:
			reflect.ValueOf(fn).Call([]reflect.Value{})
		default:
			continue
		}
	}
}

func (validation Validation) All2(args ...Parameters) {
	for _, fn := range args {
		fn()
	}
}
