package validators

import (
	"testing"

	"github.com/lnascimento1988/govalidations/notifications"
)

type StringDomain struct {
	inputData    string
	property     string
	message      string
	expectedData bool
	execute      func(instance String, domain StringDomain)
}

func TestStringValidator(t *testing.T) {
	useCases := map[string]StringDomain{
		"success: IsNotNullOrEmpty": {
			inputData:    "leonardo",
			property:     "IsNotNullOrEmpty",
			message:      "is not null or empty",
			expectedData: false,
			execute: func(instance String, domain StringDomain) {
				instance.IsNotNullOrEmpty(domain.inputData, domain.property, domain.message)
			},
		},
		"error: IsNotNullOrEmpty": {
			inputData:    "",
			property:     "IsNotNullOrEmpty",
			message:      "is not null or empty",
			expectedData: true,
			execute: func(instance String, domain StringDomain) {
				instance.IsNotNullOrEmpty(domain.inputData, domain.property, domain.message)
			},
		},
		"success: IsNullOrEmpty": {
			inputData:    "",
			property:     "IsNullOrEmpty",
			message:      "is null or empty",
			expectedData: false,
			execute: func(instance String, domain StringDomain) {
				instance.IsNullOrEmpty(domain.inputData, domain.property, domain.message)
			},
		},
		"error: IsNullOrEmpty": {
			inputData:    "leonardo",
			property:     "IsNullOrEmpty",
			message:      "is null or empty",
			expectedData: true,
			execute: func(instance String, domain StringDomain) {
				instance.IsNullOrEmpty(domain.inputData, domain.property, domain.message)
			},
		},
		"success: HasMinLength": {
			inputData:    "leonardo",
			property:     "HasMinLength",
			message:      "has min length",
			expectedData: false,
			execute: func(instance String, domain StringDomain) {
				instance.HasMinLength(domain.inputData, 3, domain.property, domain.message)
			},
		},
		"error: HasMinLength": {
			inputData:    "le",
			property:     "HasMinLength",
			message:      "has min length",
			expectedData: true,
			execute: func(instance String, domain StringDomain) {
				instance.HasMinLength(domain.inputData, 3, domain.property, domain.message)
			},
		},
		"success: HasMaxLength": {
			inputData:    "leonardo",
			property:     "HasMaxLength",
			message:      "has max length",
			expectedData: false,
			execute: func(instance String, domain StringDomain) {
				instance.HasMaxLength(domain.inputData, 10, domain.property, domain.message)
			},
		},
		"error: HasMaxLength": {
			inputData:    "12345678901",
			property:     "HasMaxLength",
			message:      "has max length",
			expectedData: true,
			execute: func(instance String, domain StringDomain) {
				instance.HasMaxLength(domain.inputData, 10, domain.property, domain.message)
			},
		},
		"success: HasLength": {
			inputData:    "leonardo",
			property:     "HasLength",
			message:      "has length",
			expectedData: false,
			execute: func(instance String, domain StringDomain) {
				instance.HasLength(domain.inputData, 8, domain.property, domain.message)
			},
		},
		"error: HasLength": {
			inputData:    "12345678901",
			property:     "HasLength",
			message:      "has length",
			expectedData: true,
			execute: func(instance String, domain StringDomain) {
				instance.HasLength(domain.inputData, 8, domain.property, domain.message)
			},
		},
		"success: Contains": {
			inputData:    "leonardo",
			property:     "Contains",
			message:      "contains",
			expectedData: false,
			execute: func(instance String, domain StringDomain) {
				instance.Contains(domain.inputData, "leo", domain.property, domain.message)
			},
		},
		"error: Contains": {
			inputData:    "leonardo",
			property:     "Contains",
			message:      "contains",
			expectedData: true,
			execute: func(instance String, domain StringDomain) {
				instance.Contains(domain.inputData, "silva", domain.property, domain.message)
			},
		},
		"success: AreEquals": {
			inputData:    "leonardo",
			property:     "AreEquals",
			message:      "are equals",
			expectedData: false,
			execute: func(instance String, domain StringDomain) {
				instance.AreEquals(domain.inputData, "leonardo", domain.property, domain.message)
			},
		},
		"error: AreEquals": {
			inputData:    "leonardo",
			property:     "AreEquals",
			message:      "are equals",
			expectedData: true,
			execute: func(instance String, domain StringDomain) {
				instance.AreEquals(domain.inputData, "silva", domain.property, domain.message)
			},
		},
		"success: AreNotEquals": {
			inputData:    "leonardo",
			property:     "AreNotEquals",
			message:      "are not equals",
			expectedData: false,
			execute: func(instance String, domain StringDomain) {
				instance.AreNotEquals(domain.inputData, "silva", domain.property, domain.message)
			},
		},
		"error: AreNotEquals": {
			inputData:    "leonardo",
			property:     "AreNotEquals",
			message:      "are not equals",
			expectedData: true,
			execute: func(instance String, domain StringDomain) {
				instance.AreNotEquals(domain.inputData, "leonardo", domain.property, domain.message)
			},
		},
		"success: IsEmail": {
			inputData:    "leonardo@email.com",
			property:     "IsEmail",
			message:      "is email",
			expectedData: false,
			execute: func(instance String, domain StringDomain) {
				instance.IsEmail(domain.inputData, domain.property, domain.message)
			},
		},
		"error: IsEmail": {
			inputData:    "leonardo",
			property:     "IsEmail",
			message:      "is email",
			expectedData: true,
			execute: func(instance String, domain StringDomain) {
				instance.IsEmail(domain.inputData, domain.property, domain.message)
			},
		},
		"success: Match": {
			inputData:    "10:00",
			property:     "Match",
			message:      "match",
			expectedData: false,
			execute: func(instance String, domain StringDomain) {
				instance.Match(domain.inputData, "^[0-9]{2}:[0-9]{2}$", domain.property, domain.message)
			},
		},
		"error: Match": {
			inputData:    "leonardo",
			property:     "Match",
			message:      "match",
			expectedData: true,
			execute: func(instance String, domain StringDomain) {
				instance.Match(domain.inputData, "^[0-9]{2}:[0-9]{2}$", domain.property, domain.message)
			},
		},
	}

	for name, useCase := range useCases {
		t.Run(name, func(t *testing.T) {
			notifier := notifications.New()
			instance := NewString(notifier)

			useCase.execute(instance, useCase)

			invalid := instance.Notifier.Invalid()

			if invalid != useCase.expectedData {
				t.Errorf("Expected %t, but got %t", useCase.expectedData, invalid)
			}
		})
	}
}
