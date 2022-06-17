package validators

import (
	"testing"

	"github.com/lsnascimento/govalidations/notifications"
)

type DateDomain struct {
	inputData    string
	property     string
	message      string
	expectedData bool
	execute      func(instance Date, domain DateDomain)
}

func TestDateValidator(t *testing.T) {
	useCases := map[string]DateDomain{
		"success: IsValid": {
			inputData:    "2022-06-16",
			property:     "IsValid",
			message:      "",
			expectedData: false,
			execute: func(instance Date, domain DateDomain) {
				instance.IsValid(domain.inputData, domain.property, domain.message)
			},
		},
		"error: IsValid": {
			inputData:    "2022-06-36",
			property:     "IsValid",
			message:      "",
			expectedData: true,
			execute: func(instance Date, domain DateDomain) {
				instance.IsValid(domain.inputData, domain.property, domain.message)
			},
		},
		"success: IsGreaterThan": {
			inputData:    "2022-06-16T23:01:00",
			property:     "IsGreaterThan",
			message:      "",
			expectedData: false,
			execute: func(instance Date, domain DateDomain) {
				instance.IsGreaterThan(domain.inputData, "2022-06-16T23:00:00", domain.property, domain.message)
			},
		},
		"error: IsGreaterThan": {
			inputData:    "2022-06-16T23:00:00",
			property:     "IsGreaterThan",
			message:      "",
			expectedData: true,
			execute: func(instance Date, domain DateDomain) {
				instance.IsGreaterThan(domain.inputData, "2022-06-16T23:00:00", domain.property, domain.message)
			},
		},
		"success: IsGreaterOrEqualThan": {
			inputData:    "2022-06-16T23:01:00",
			property:     "IsGreaterOrEqualThan",
			message:      "",
			expectedData: false,
			execute: func(instance Date, domain DateDomain) {
				instance.IsGreaterOrEqualThan(domain.inputData, "2022-06-16T23:00:00", domain.property, domain.message)
			},
		},
		"error: IsGreaterOrEqualThan": {
			inputData:    "2022-06-16T22:00:00",
			property:     "IsGreaterOrEqualThan",
			message:      "",
			expectedData: true,
			execute: func(instance Date, domain DateDomain) {
				instance.IsGreaterOrEqualThan(domain.inputData, "2022-06-16T23:00:00", domain.property, domain.message)
			},
		},
		"success: IsLessThan": {
			inputData:    "2022-06-16T22:59:59",
			property:     "IsLessThan",
			message:      "",
			expectedData: false,
			execute: func(instance Date, domain DateDomain) {
				instance.IsLessThan(domain.inputData, "2022-06-16T23:00:00", domain.property, domain.message)
			},
		},
		"error: IsLessThan": {
			inputData:    "2022-06-16T23:00:00",
			property:     "IsLessThan",
			message:      "",
			expectedData: true,
			execute: func(instance Date, domain DateDomain) {
				instance.IsLessThan(domain.inputData, "2022-06-16T23:00:00", domain.property, domain.message)
			},
		},
		"success: IsLessOrEqualThan": {
			inputData:    "2022-06-16T23:00:00",
			property:     "IsLessOrEqualThan",
			message:      "",
			expectedData: false,
			execute: func(instance Date, domain DateDomain) {
				instance.IsLessOrEqualThan(domain.inputData, "2022-06-16T23:00:00", domain.property, domain.message)
			},
		},
		"error: IsLessOrEqualThan": {
			inputData:    "2022-06-16T23:01:00",
			property:     "IsLessOrEqualThan",
			message:      "",
			expectedData: true,
			execute: func(instance Date, domain DateDomain) {
				instance.IsLessOrEqualThan(domain.inputData, "2022-06-16T23:00:00", domain.property, domain.message)
			},
		},
		"success: IsBetween": {
			inputData:    "2022-06-16T23:00:00",
			property:     "IsBetween",
			message:      "",
			expectedData: false,
			execute: func(instance Date, domain DateDomain) {
				instance.IsBetween(domain.inputData, "2022-06-16T23:00:00", "2022-06-16T23:59:59", domain.property, domain.message)
			},
		},
		"error: IsBetween (Before)": {
			inputData:    "2022-06-16T22:00:00",
			property:     "IsBetween",
			message:      "",
			expectedData: true,
			execute: func(instance Date, domain DateDomain) {
				instance.IsBetween(domain.inputData, "2022-06-16T23:00:00", "2022-06-16T23:59:59", domain.property, domain.message)
			},
		},
		"error: IsBetween (After)": {
			inputData:    "2022-06-17T00:00:00",
			property:     "IsBetween",
			message:      "",
			expectedData: true,
			execute: func(instance Date, domain DateDomain) {
				instance.IsBetween(domain.inputData, "2022-06-16T23:00:00", "2022-06-16T23:59:59", domain.property, domain.message)
			},
		},
	}

	for name, useCase := range useCases {
		t.Run(name, func(t *testing.T) {
			notifier := notifications.New()
			instance := NewDate(notifier)

			useCase.execute(instance, useCase)

			invalid := instance.Notifier.Invalid()

			if invalid != useCase.expectedData {
				t.Errorf("Expected %t, but got %t", useCase.expectedData, invalid)
			}
		})
	}
}
