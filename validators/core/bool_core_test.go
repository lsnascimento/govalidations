package core

import (
	"testing"
)

type BoolDomain struct {
	inputData    bool
	inputPointer *bool
	expectedData bool
	execute      func(instance BoolValidator, domain BoolDomain) bool
}

func TestBoolValidator(t *testing.T) {
	useCases := map[string]BoolDomain{
		"success: when passing nil to IsNull method expectedData should to be true": {
			inputPointer: nil,
			expectedData: true,
			execute: func(instance BoolValidator, domain BoolDomain) bool {
				return instance.IsNull(domain.inputPointer)
			},
		},
		"error: when passing false to IsNull method expectedData should to be false": {
			inputPointer: getBoolPointer(false),
			expectedData: false,
			execute: func(instance BoolValidator, domain BoolDomain) bool {
				return instance.IsNull(domain.inputPointer)
			},
		},
		"error: when passing true to IsNull method expectedData should to be false": {
			inputPointer: getBoolPointer(true),
			expectedData: false,
			execute: func(instance BoolValidator, domain BoolDomain) bool {
				return instance.IsNull(domain.inputPointer)
			},
		},
		"success: when passing false to IsNotNull method expectedData should to be true": {
			inputPointer: getBoolPointer(false),
			expectedData: true,
			execute: func(instance BoolValidator, domain BoolDomain) bool {
				return instance.IsNotNull(domain.inputPointer)
			},
		},
		"success: when passing true to IsNotNull method expectedData should to be true": {
			inputPointer: getBoolPointer(true),
			expectedData: true,
			execute: func(instance BoolValidator, domain BoolDomain) bool {
				return instance.IsNotNull(domain.inputPointer)
			},
		},
		"error: when passing nil to IsNotNull method expectedData should to be false": {
			inputPointer: nil,
			expectedData: false,
			execute: func(instance BoolValidator, domain BoolDomain) bool {
				return instance.IsNotNull(domain.inputPointer)
			},
		},
		"success: when passing true to IsTrue method expectedData should to be true": {
			inputData:    true,
			expectedData: true,
			execute: func(instance BoolValidator, domain BoolDomain) bool {
				return instance.IsTrue(domain.inputData)
			},
		},
		"error: when passing false to IsTrue method expectedData should to be false": {
			inputData:    false,
			expectedData: false,
			execute: func(instance BoolValidator, domain BoolDomain) bool {
				return instance.IsTrue(domain.inputData)
			},
		},
		"success: when passing false to IsFalse method expectedData should to be true": {
			inputData:    false,
			expectedData: true,
			execute: func(instance BoolValidator, domain BoolDomain) bool {
				return instance.IsFalse(domain.inputData)
			},
		},
		"error: when passing true to IsFalse method expectedData should to be false": {
			inputData:    true,
			expectedData: false,
			execute: func(instance BoolValidator, domain BoolDomain) bool {
				return instance.IsFalse(domain.inputData)
			},
		},
	}

	for name, useCase := range useCases {
		t.Run(name, func(t *testing.T) {
			instance := NewBool()

			response := useCase.execute(instance, useCase)

			if response != useCase.expectedData {
				t.Errorf("Expected %t, but got %t", useCase.expectedData, response)
			}
		})
	}
}

func getBoolPointer(value bool) *bool {
	return &value
}
