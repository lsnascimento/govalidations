package validators

import "github.com/lnascimento1988/govalidations/notifications"

type Validator struct {
	Int
	Bool
	Date
	Slice
	String
}

func New(notifier *notifications.Notifier) Validator {
	return Validator{
		Int:    NewInt(notifier),
		Bool:   NewBool(notifier),
		Date:   NewDate(notifier),
		Slice:  NewSlice(notifier),
		String: NewString(notifier),
	}
}
