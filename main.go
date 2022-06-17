package validations

import (
	"github.com/lsnascimento/govalidations/notifications"
	"github.com/lsnascimento/govalidations/validators"
)

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
