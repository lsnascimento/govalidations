package validators

import (
	"testing"

	"github.com/lnascimento1988/govalidations/notifications"
)

type BoolDomain struct {
	inputData    bool
	property     string
	message      string
	expectedData bool
	execute      func(instance Bool, domain BoolDomain)
}

func TestBoolValidator(t *testing.T) {
	useCases := map[string]BoolDomain{
		"success: IsTrue": {
			inputData:    true,
			property:     "IsTrue",
			message:      "",
			expectedData: false,
			execute: func(instance Bool, domain BoolDomain) {
				instance.IsTrue(domain.inputData, domain.property, domain.message)
			},
		},
		"error: IsTrue": {
			inputData:    false,
			property:     "IsTrue",
			message:      "",
			expectedData: true,
			execute: func(instance Bool, domain BoolDomain) {
				instance.IsTrue(domain.inputData, domain.property, domain.message)
			},
		},
		"success: IsFalse": {
			inputData:    false,
			property:     "IsFalse",
			message:      "",
			expectedData: false,
			execute: func(instance Bool, domain BoolDomain) {
				instance.IsFalse(domain.inputData, domain.property, domain.message)
			},
		},
		"error: IsFalse": {
			inputData:    true,
			property:     "IsFalse",
			message:      "",
			expectedData: true,
			execute: func(instance Bool, domain BoolDomain) {
				instance.IsFalse(domain.inputData, domain.property, domain.message)
			},
		},
	}

	for name, useCase := range useCases {
		t.Run(name, func(t *testing.T) {
			notifier := notifications.New()
			instance := NewBool(notifier)

			useCase.execute(instance, useCase)

			invalid := instance.Notifier.Invalid()

			if invalid != useCase.expectedData {
				t.Errorf("Expected %t, but got %t", useCase.expectedData, invalid)
			}
		})
	}
}
