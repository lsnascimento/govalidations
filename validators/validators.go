package validators

import "github.com/lsnascimento/govalidations/notifications"

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
