package validators

import (
	"testing"

	"github.com/lsnascimento/govalidations/notifications"
	"github.com/lsnascimento/govalidations/validators/core"
)

type SliceDomain struct {
	inputData    interface{}
	property     string
	message      string
	expectedData bool
	execute      func(instance Slice, domain SliceDomain)
}

func TestSliceValidator(t *testing.T) {
	useCases := map[string]SliceDomain{
		"success: IsRequired": {
			inputData:    []string{"leo"},
			property:     "IsRequired",
			message:      "",
			expectedData: false,
			execute: func(instance Slice, domain SliceDomain) {
				instance.IsRequired(domain.inputData, domain.property, domain.message)
			},
		},
		"error: IsRequired": {
			inputData:    []string{},
			property:     "IsRequired",
			message:      "",
			expectedData: true,
			execute: func(instance Slice, domain SliceDomain) {
				instance.IsRequired(domain.inputData, domain.property, domain.message)
			},
		},
	}

	for name, useCase := range useCases {
		t.Run(name, func(t *testing.T) {
			notifier := notifications.New()
			instance := NewSlice(core.NewSlice(), notifier)

			useCase.execute(instance, useCase)

			invalid := instance.Notifier.Invalid()

			if invalid != useCase.expectedData {
				t.Errorf("Expected %t, but got %t", useCase.expectedData, invalid)
			}
		})
	}
}
