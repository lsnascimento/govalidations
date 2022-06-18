package validators

import (
	"testing"

	"github.com/lsnascimento/govalidations/notifications"
)

type BoolDomain struct {
	inputData    bool
	inputPointer *bool
	property     string
	message      string
	expectedData bool
	execute      func(instance Bool, domain BoolDomain)
}

func TestBoolValidator(t *testing.T) {
	useCases := map[string]BoolDomain{
		"success: IsRequired pass true": {
			inputPointer: getBoolPointer(true),
			property:     "IsRequired",
			message:      "",
			expectedData: false,
			execute: func(instance Bool, domain BoolDomain) {
				instance.IsRequired(domain.inputPointer, domain.property, domain.message)
			},
		},
		"success: IsRequired pass false": {
			inputPointer: getBoolPointer(false),
			property:     "IsRequired",
			message:      "",
			expectedData: false,
			execute: func(instance Bool, domain BoolDomain) {
				instance.IsRequired(domain.inputPointer, domain.property, domain.message)
			},
		},
		"error: IsRequired pass nil": {
			inputPointer: nil,
			property:     "IsRequired",
			message:      "",
			expectedData: true,
			execute: func(instance Bool, domain BoolDomain) {
				instance.IsRequired(domain.inputPointer, domain.property, domain.message)
			},
		},
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

func getBoolPointer(value bool) *bool {
	return &value
}
