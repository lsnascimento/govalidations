package validators

import (
	"testing"

	"github.com/lsnascimento/govalidations/notifications"
)

type IntDomain struct {
	inputData    int
	inputPointer *int
	comparer     int
	property     string
	message      string
	expectedData bool
	execute      func(instance Int, domain IntDomain)
}

func TestIntValidator(t *testing.T) {
	useCases := map[string]IntDomain{
		"success: IsRequired pass 1": {
			inputPointer: getIntPointer(1),
			property:     "IsRequired",
			message:      "",
			expectedData: false,
			execute: func(instance Int, domain IntDomain) {
				instance.IsRequired(domain.inputPointer, domain.property, domain.message)
			},
		},
		"success: IsRequired pass 0": {
			inputPointer: getIntPointer(0),
			property:     "IsRequired",
			message:      "",
			expectedData: false,
			execute: func(instance Int, domain IntDomain) {
				instance.IsRequired(domain.inputPointer, domain.property, domain.message)
			},
		},
		"error: IsRequired pass nil": {
			inputPointer: nil,
			property:     "IsRequired",
			message:      "",
			expectedData: true,
			execute: func(instance Int, domain IntDomain) {
				instance.IsRequired(domain.inputPointer, domain.property, domain.message)
			},
		},
		"success: IsGreaterThan": {
			inputData:    4,
			comparer:     3,
			property:     "IsGreaterThan",
			message:      "",
			expectedData: false,
			execute: func(instance Int, domain IntDomain) {
				instance.IsGreaterThan(domain.inputData, domain.comparer, domain.property, domain.message)
			},
		},
		"error: IsGreaterThan": {
			inputData:    3,
			comparer:     3,
			property:     "IsGreaterThan",
			message:      "",
			expectedData: true,
			execute: func(instance Int, domain IntDomain) {
				instance.IsGreaterThan(domain.inputData, domain.comparer, domain.property, domain.message)
			},
		},
		"success: IsGreaterOrEqualThan": {
			inputData:    3,
			comparer:     3,
			property:     "IsGreaterOrEqualThan",
			message:      "",
			expectedData: false,
			execute: func(instance Int, domain IntDomain) {
				instance.IsGreaterOrEqualThan(domain.inputData, domain.comparer, domain.property, domain.message)
			},
		},
		"error: IsGreaterOrEqualThan": {
			inputData:    2,
			comparer:     3,
			property:     "IsGreaterOrEqualThan",
			message:      "",
			expectedData: true,
			execute: func(instance Int, domain IntDomain) {
				instance.IsGreaterOrEqualThan(domain.inputData, domain.comparer, domain.property, domain.message)
			},
		},
		"success: IsLessThan": {
			inputData:    2,
			comparer:     3,
			property:     "IsLessThan",
			message:      "",
			expectedData: false,
			execute: func(instance Int, domain IntDomain) {
				instance.IsLessThan(domain.inputData, domain.comparer, domain.property, domain.message)
			},
		},
		"error: IsLessThan": {
			inputData:    3,
			comparer:     3,
			property:     "IsLessThan",
			message:      "",
			expectedData: true,
			execute: func(instance Int, domain IntDomain) {
				instance.IsLessThan(domain.inputData, domain.comparer, domain.property, domain.message)
			},
		},
		"success: IsLessOrEqualThan": {
			inputData:    3,
			comparer:     3,
			property:     "IsLessOrEqualThan",
			message:      "",
			expectedData: false,
			execute: func(instance Int, domain IntDomain) {
				instance.IsLessOrEqualThan(domain.inputData, domain.comparer, domain.property, domain.message)
			},
		},
		"error: IsLessOrEqualThan": {
			inputData:    4,
			comparer:     3,
			property:     "IsLessOrEqualThan",
			message:      "",
			expectedData: true,
			execute: func(instance Int, domain IntDomain) {
				instance.IsLessOrEqualThan(domain.inputData, domain.comparer, domain.property, domain.message)
			},
		},
		"success: AreEquals": {
			inputData:    3,
			comparer:     3,
			property:     "AreEquals",
			message:      "",
			expectedData: false,
			execute: func(instance Int, domain IntDomain) {
				instance.AreEquals(domain.inputData, domain.comparer, domain.property, domain.message)
			},
		},
		"error: AreEquals": {
			inputData:    4,
			comparer:     3,
			property:     "AreEquals",
			message:      "",
			expectedData: true,
			execute: func(instance Int, domain IntDomain) {
				instance.AreEquals(domain.inputData, domain.comparer, domain.property, domain.message)
			},
		},
		"success: AreNotEquals": {
			inputData:    4,
			comparer:     3,
			property:     "AreNotEquals",
			message:      "",
			expectedData: false,
			execute: func(instance Int, domain IntDomain) {
				instance.AreNotEquals(domain.inputData, domain.comparer, domain.property, domain.message)
			},
		},
		"error: AreNotEquals": {
			inputData:    3,
			comparer:     3,
			property:     "AreNotEquals",
			message:      "",
			expectedData: true,
			execute: func(instance Int, domain IntDomain) {
				instance.AreNotEquals(domain.inputData, domain.comparer, domain.property, domain.message)
			},
		},
		"success: IsBetween": {
			inputData:    4,
			comparer:     3,
			property:     "IsBetween",
			message:      "",
			expectedData: false,
			execute: func(instance Int, domain IntDomain) {
				instance.IsBetween(domain.inputData, 1, 6, domain.property, domain.message)
			},
		},
		"error: IsBetween": {
			inputData:    3,
			comparer:     3,
			property:     "IsBetween",
			message:      "",
			expectedData: true,
			execute: func(instance Int, domain IntDomain) {
				instance.IsBetween(domain.inputData, 4, 10, domain.property, domain.message)
			},
		},
	}

	for name, useCase := range useCases {
		t.Run(name, func(t *testing.T) {
			notifier := notifications.New()
			instance := NewInt(notifier)

			useCase.execute(instance, useCase)

			invalid := instance.Notifier.Invalid()

			if invalid != useCase.expectedData {
				t.Errorf("Expected %t, but got %t", useCase.expectedData, invalid)
			}
		})
	}
}

func getIntPointer(value int) *int {
	return &value
}
