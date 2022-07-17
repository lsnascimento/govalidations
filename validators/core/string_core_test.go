package core

import (
	"testing"
)

type StringDomain struct {
	inputData    string
	compare      int
	contain      string
	pattern      string
	expectedData bool
	execute      func(instance StringValidator, domain StringDomain) bool
}

func TestStringValidator(t *testing.T) {
	useCases := map[string]StringDomain{
		"success: when value is not null to IsNotNullOrEmpty method expectedData should to be true": {
			inputData:    "a",
			expectedData: true,
			execute: func(instance StringValidator, domain StringDomain) bool {
				return instance.IsNotNullOrEmpty(domain.inputData)
			},
		},
		"error: when value is empty to IsNotNullOrEmpty method expectedData should to be false": {
			inputData:    "",
			expectedData: false,
			execute: func(instance StringValidator, domain StringDomain) bool {
				return instance.IsNotNullOrEmpty(domain.inputData)
			},
		},
		"success: when value is empty to IsNullOrEmpty method expectedData should to be true": {
			inputData:    "",
			expectedData: true,
			execute: func(instance StringValidator, domain StringDomain) bool {
				return instance.IsNullOrEmpty(domain.inputData)
			},
		},
		"error: when value is not empty to IsNullOrEmpty method expectedData should to be false": {
			inputData:    "a",
			expectedData: false,
			execute: func(instance StringValidator, domain StringDomain) bool {
				return instance.IsNullOrEmpty(domain.inputData)
			},
		},
		"success: when value is greater than compare to HasMinLength method expectedData should to be true": {
			inputData:    "abcd",
			compare:      3,
			expectedData: true,
			execute: func(instance StringValidator, domain StringDomain) bool {
				return instance.HasMinLength(domain.inputData, domain.compare)
			},
		},
		"success: when value are equal compare to HasMinLength method expectedData should to be true": {
			inputData:    "abc",
			compare:      3,
			expectedData: true,
			execute: func(instance StringValidator, domain StringDomain) bool {
				return instance.HasMinLength(domain.inputData, domain.compare)
			},
		},
		"error: when value is less than compare to HasMinLength method expectedData should to be false": {
			inputData:    "a",
			compare:      3,
			expectedData: false,
			execute: func(instance StringValidator, domain StringDomain) bool {
				return instance.HasMinLength(domain.inputData, domain.compare)
			},
		},
		"success: when value is less than compare to HasMaxLength method expectedData should to be true": {
			inputData:    "abcd",
			compare:      5,
			expectedData: true,
			execute: func(instance StringValidator, domain StringDomain) bool {
				return instance.HasMaxLength(domain.inputData, domain.compare)
			},
		},
		"success: when value are equal compare to HasMaxLength method expectedData should to be true": {
			inputData:    "abc",
			compare:      3,
			expectedData: true,
			execute: func(instance StringValidator, domain StringDomain) bool {
				return instance.HasMaxLength(domain.inputData, domain.compare)
			},
		},
		"error: when value is greater than compare to HasMaxLength method expectedData should to be false": {
			inputData:    "abcd",
			compare:      3,
			expectedData: false,
			execute: func(instance StringValidator, domain StringDomain) bool {
				return instance.HasMaxLength(domain.inputData, domain.compare)
			},
		},
		"success: when value are equal compare to HasLength method expectedData should to be true": {
			inputData:    "abcd",
			compare:      4,
			expectedData: true,
			execute: func(instance StringValidator, domain StringDomain) bool {
				return instance.HasLength(domain.inputData, domain.compare)
			},
		},
		"error: when value are not equal compare to HasLength method expectedData should to be false": {
			inputData:    "abcd",
			compare:      3,
			expectedData: false,
			execute: func(instance StringValidator, domain StringDomain) bool {
				return instance.HasLength(domain.inputData, domain.compare)
			},
		},
		"success: when value contain in text to Contains method expectedData should to be true": {
			inputData:    "string",
			contain:      "ing",
			expectedData: true,
			execute: func(instance StringValidator, domain StringDomain) bool {
				return instance.Contains(domain.inputData, domain.contain)
			},
		},
		"error: when value not contain in text to Contains method expectedData should to be false": {
			inputData:    "string",
			contain:      "ab",
			expectedData: false,
			execute: func(instance StringValidator, domain StringDomain) bool {
				return instance.Contains(domain.inputData, domain.contain)
			},
		},
		"success: when value are equal text to AreEquals method expectedData should to be true": {
			inputData:    "string",
			contain:      "string",
			expectedData: true,
			execute: func(instance StringValidator, domain StringDomain) bool {
				return instance.AreEquals(domain.inputData, domain.contain)
			},
		},
		"error: when value are not equal text to AreEquals method expectedData should to be false": {
			inputData:    "string",
			contain:      "ab",
			expectedData: false,
			execute: func(instance StringValidator, domain StringDomain) bool {
				return instance.AreEquals(domain.inputData, domain.contain)
			},
		},
		"success: when value are not equal text to AreNotEquals method expectedData should to be true": {
			inputData:    "string",
			contain:      "str",
			expectedData: true,
			execute: func(instance StringValidator, domain StringDomain) bool {
				return instance.AreNotEquals(domain.inputData, domain.contain)
			},
		},
		"error: when value are equal text to AreNotEquals method expectedData should to be false": {
			inputData:    "string",
			contain:      "string",
			expectedData: false,
			execute: func(instance StringValidator, domain StringDomain) bool {
				return instance.AreNotEquals(domain.inputData, domain.contain)
			},
		},
		"success: when value is email valid to IsEmail method expectedData should to be true": {
			inputData:    "leonardo@email.com",
			expectedData: true,
			execute: func(instance StringValidator, domain StringDomain) bool {
				return instance.IsEmail(domain.inputData)
			},
		},
		"error: when value is not email valid to IsEmail method expectedData should to be false": {
			inputData:    "leonardo@",
			expectedData: false,
			execute: func(instance StringValidator, domain StringDomain) bool {
				return instance.IsEmail(domain.inputData)
			},
		},
		"success: when value match pattern to Match method expectedData should to be true": {
			inputData:    "11:00",
			pattern:      "^[0-9]{2}:[0-9]{2}$",
			expectedData: true,
			execute: func(instance StringValidator, domain StringDomain) bool {
				return instance.Match(domain.inputData, domain.pattern)
			},
		},
		"error: when value not match pattern to Match method expectedData should to be false": {
			inputData:    "leonardo@",
			pattern:      "^[0-9]{2}:[0-9]{2}$",
			expectedData: false,
			execute: func(instance StringValidator, domain StringDomain) bool {
				return instance.Match(domain.inputData, domain.pattern)
			},
		},
	}

	for name, useCase := range useCases {
		t.Run(name, func(t *testing.T) {
			instance := NewString()

			response := useCase.execute(instance, useCase)

			if response != useCase.expectedData {
				t.Errorf("Expected %t, but got %t", useCase.expectedData, response)
			}
		})
	}
}
