package core

import (
	"testing"
)

type DateDomain struct {
	inputData    string
	compare      string
	from         string
	to           string
	expectedData bool
	execute      func(instance DateValidator, domain DateDomain) bool
}

func TestDateValidator(t *testing.T) {
	useCases := map[string]DateDomain{
		"success: when passing string date to IsValid method expectedData should to be true": {
			inputData:    "2022-06-16",
			expectedData: true,
			execute: func(instance DateValidator, domain DateDomain) bool {
				return instance.IsValid(domain.inputData)
			},
		},
		"error: when passing invalid date to IsValid method expectedData should to be false": {
			inputData:    "2022-06",
			expectedData: false,
			execute: func(instance DateValidator, domain DateDomain) bool {
				return instance.IsValid(domain.inputData)
			},
		},
		"success: when first date is greater than second date in IsGreaterThan method expectedData should to be true": {
			inputData:    "2022-06-16",
			compare:      "2022-06-10",
			expectedData: true,
			execute: func(instance DateValidator, domain DateDomain) bool {
				return instance.IsGreaterThan(domain.inputData, domain.compare)
			},
		},
		"error: when first date are equal to second date in IsGreaterThan method expectedData should to be false": {
			inputData:    "2022-06-16",
			compare:      "2022-06-16",
			expectedData: false,
			execute: func(instance DateValidator, domain DateDomain) bool {
				return instance.IsGreaterThan(domain.inputData, domain.compare)
			},
		},
		"error: when first date is less than second date in IsGreaterThan method expectedData should to be false": {
			inputData:    "2022-06-10",
			compare:      "2022-06-16",
			expectedData: false,
			execute: func(instance DateValidator, domain DateDomain) bool {
				return instance.IsGreaterThan(domain.inputData, domain.compare)
			},
		},
		"success: when first date is greater or equal than second date in IsGreaterOrEqualThan method expectedData should to be true": {
			inputData:    "2022-06-16",
			compare:      "2022-06-10",
			expectedData: true,
			execute: func(instance DateValidator, domain DateDomain) bool {
				return instance.IsGreaterOrEqualThan(domain.inputData, domain.compare)
			},
		},
		"success: when first date are equal to second date in IsGreaterOrEqualThan method expectedData should to be true": {
			inputData:    "2022-06-16",
			compare:      "2022-06-16",
			expectedData: true,
			execute: func(instance DateValidator, domain DateDomain) bool {
				return instance.IsGreaterOrEqualThan(domain.inputData, domain.compare)
			},
		},
		"error: when first date is less than second date in IsGreaterOrEqualThan method expectedData should to be false": {
			inputData:    "2022-06-10",
			compare:      "2022-06-16",
			expectedData: false,
			execute: func(instance DateValidator, domain DateDomain) bool {
				return instance.IsGreaterOrEqualThan(domain.inputData, domain.compare)
			},
		},
		"success: when first date is greater than second date in IsLessThan method expectedData should to be true": {
			inputData:    "2022-06-10",
			compare:      "2022-06-16",
			expectedData: true,
			execute: func(instance DateValidator, domain DateDomain) bool {
				return instance.IsLessThan(domain.inputData, domain.compare)
			},
		},
		"error: when first date are equal to second date in IsLessThan method expectedData should to be false": {
			inputData:    "2022-06-16",
			compare:      "2022-06-16",
			expectedData: false,
			execute: func(instance DateValidator, domain DateDomain) bool {
				return instance.IsLessThan(domain.inputData, domain.compare)
			},
		},
		"error: when first date is less than second date in IsLessThan method expectedData should to be false": {
			inputData:    "2022-06-16",
			compare:      "2022-06-10",
			expectedData: false,
			execute: func(instance DateValidator, domain DateDomain) bool {
				return instance.IsLessThan(domain.inputData, domain.compare)
			},
		},
		"success: when first date is greater or equal than second date in IsLessOrEqualThan method expectedData should to be true": {
			inputData:    "2022-06-10",
			compare:      "2022-06-16",
			expectedData: true,
			execute: func(instance DateValidator, domain DateDomain) bool {
				return instance.IsLessOrEqualThan(domain.inputData, domain.compare)
			},
		},
		"success: when first date are equal to second date in IsLessOrEqualThan method expectedData should to be true": {
			inputData:    "2022-06-16",
			compare:      "2022-06-16",
			expectedData: true,
			execute: func(instance DateValidator, domain DateDomain) bool {
				return instance.IsLessOrEqualThan(domain.inputData, domain.compare)
			},
		},
		"error: when first date is less than second date in IsLessOrEqualThan method expectedData should to be false": {
			inputData:    "2022-06-16",
			compare:      "2022-06-10",
			expectedData: false,
			execute: func(instance DateValidator, domain DateDomain) bool {
				return instance.IsLessOrEqualThan(domain.inputData, domain.compare)
			},
		},
		"success: when date is between in range dates in IsBetween method expectedData should to be true": {
			inputData:    "2022-06-10",
			from:         "2022-06-01",
			to:           "2022-06-30",
			expectedData: true,
			execute: func(instance DateValidator, domain DateDomain) bool {
				return instance.IsBetween(domain.inputData, domain.from, domain.to)
			},
		},
		"success: when date are equal start date in IsBetween method expectedData should to be true": {
			inputData:    "2022-06-01",
			from:         "2022-06-01",
			to:           "2022-06-30",
			expectedData: true,
			execute: func(instance DateValidator, domain DateDomain) bool {
				return instance.IsBetween(domain.inputData, domain.from, domain.to)
			},
		},
		"success: when date is in range in IsBetween method expectedData should to be true": {
			inputData:    "2022-06-30",
			from:         "2022-06-01",
			to:           "2022-06-30",
			expectedData: true,
			execute: func(instance DateValidator, domain DateDomain) bool {
				return instance.IsBetween(domain.inputData, domain.from, domain.to)
			},
		},
		"error: when date is less than start date in IsBetween method expectedData should to be false": {
			inputData:    "2022-05-31",
			from:         "2022-06-01",
			to:           "2022-06-30",
			expectedData: false,
			execute: func(instance DateValidator, domain DateDomain) bool {
				return instance.IsBetween(domain.inputData, domain.from, domain.to)
			},
		},
		"error: when date is greater than end date in IsBetween method expectedData should to be false": {
			inputData:    "2022-07-01",
			from:         "2022-06-01",
			to:           "2022-06-30",
			expectedData: false,
			execute: func(instance DateValidator, domain DateDomain) bool {
				return instance.IsBetween(domain.inputData, domain.from, domain.to)
			},
		},
	}

	for name, useCase := range useCases {
		t.Run(name, func(t *testing.T) {
			instance := NewDate()

			response := useCase.execute(instance, useCase)

			if response != useCase.expectedData {
				t.Errorf("Expected %t, but got %t", useCase.expectedData, response)
			}
		})
	}
}
