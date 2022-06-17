package validations

import (
	"github.com/lnascimento1988/govalidations/notifications"
	"github.com/lnascimento1988/govalidations/validators"
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
