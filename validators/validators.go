package validators

import (
	"github.com/lsnascimento/govalidations/notifications"
	"github.com/lsnascimento/govalidations/validators/core"
)

type Validator struct {
	Int
	Bool
	Date
	Slice
	String
}

func New(notifier *notifications.Notifier) Validator {
	return Validator{
		Int:    NewInt(core.NewInt(), notifier),
		Bool:   NewBool(core.NewBool(), notifier),
		Date:   NewDate(core.NewDate(), notifier),
		Slice:  NewSlice(core.NewSlice(), notifier),
		String: NewString(core.NewString(), notifier),
	}
}
