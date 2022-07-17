package core

import (
	"testing"
)

type IntDomain struct {
	inputData    int
	inputPointer *int
	compare      int
	from         int
	to           int
	expectedData bool
	execute      func(instance IntValidator, domain IntDomain) bool
}

func TestIntValidator(t *testing.T) {
	useCases := map[string]IntDomain{
		"success: when value is null to IsNull method expectedData should to be true": {
			inputPointer: nil,
			expectedData: true,
			execute: func(instance IntValidator, domain IntDomain) bool {
				return instance.IsNull(domain.inputPointer)
			},
		},
		"error: when value is not null to IsNull method expectedData should to be false": {
			inputPointer: getIntPointer(0),
			expectedData: false,
			execute: func(instance IntValidator, domain IntDomain) bool {
				return instance.IsNull(domain.inputPointer)
			},
		},
		"success: when value is not null to IsNotNull method expectedData should to be true": {
			inputPointer: getIntPointer(0),
			expectedData: true,
			execute: func(instance IntValidator, domain IntDomain) bool {
				return instance.IsNotNull(domain.inputPointer)
			},
		},
		"error: when value is null to IsNotNull method expectedData should to be false": {
			inputPointer: nil,
			expectedData: false,
			execute: func(instance IntValidator, domain IntDomain) bool {
				return instance.IsNotNull(domain.inputPointer)
			},
		},
		"success: when value is greater than compare to IsGreaterThan method expectedData should to be true": {
			inputData:    10,
			compare:      9,
			expectedData: true,
			execute: func(instance IntValidator, domain IntDomain) bool {
				return instance.IsGreaterThan(domain.inputData, domain.compare)
			},
		},
		"error: when value are equal compare to IsGreaterThan method expectedData should to be false": {
			inputData:    10,
			compare:      10,
			expectedData: false,
			execute: func(instance IntValidator, domain IntDomain) bool {
				return instance.IsGreaterThan(domain.inputData, domain.compare)
			},
		},
		"error: when value is less than compare to IsGreaterThan method expectedData should to be false": {
			inputData:    9,
			compare:      10,
			expectedData: false,
			execute: func(instance IntValidator, domain IntDomain) bool {
				return instance.IsGreaterThan(domain.inputData, domain.compare)
			},
		},
		"success: when value is greater than compare to IsGreaterOrEqualThan method expectedData should to be true": {
			inputData:    10,
			compare:      9,
			expectedData: true,
			execute: func(instance IntValidator, domain IntDomain) bool {
				return instance.IsGreaterOrEqualThan(domain.inputData, domain.compare)
			},
		},
		"success: when value are equal compare to IsGreaterOrEqualThan method expectedData should to be true": {
			inputData:    10,
			compare:      10,
			expectedData: true,
			execute: func(instance IntValidator, domain IntDomain) bool {
				return instance.IsGreaterOrEqualThan(domain.inputData, domain.compare)
			},
		},
		"error: when value is less than compare to IsGreaterOrEqualThan method expectedData should to be false": {
			inputData:    9,
			compare:      10,
			expectedData: false,
			execute: func(instance IntValidator, domain IntDomain) bool {
				return instance.IsGreaterOrEqualThan(domain.inputData, domain.compare)
			},
		},
		"success: when value is less than compare to IsLessThan method expectedData should to be true": {
			inputData:    9,
			compare:      10,
			expectedData: true,
			execute: func(instance IntValidator, domain IntDomain) bool {
				return instance.IsLessThan(domain.inputData, domain.compare)
			},
		},
		"error: when value are equal compare to IsLessThan method expectedData should to be false": {
			inputData:    9,
			compare:      9,
			expectedData: false,
			execute: func(instance IntValidator, domain IntDomain) bool {
				return instance.IsLessThan(domain.inputData, domain.compare)
			},
		},
		"error: when value is less than compare to IsLessThan method expectedData should to be false": {
			inputData:    10,
			compare:      9,
			expectedData: false,
			execute: func(instance IntValidator, domain IntDomain) bool {
				return instance.IsLessThan(domain.inputData, domain.compare)
			},
		},
		"success: when value is greater than compare to IsLessOrEqualThan method expectedData should to be true": {
			inputData:    9,
			compare:      10,
			expectedData: true,
			execute: func(instance IntValidator, domain IntDomain) bool {
				return instance.IsLessOrEqualThan(domain.inputData, domain.compare)
			},
		},
		"success: when value are equal compare to IsLessOrEqualThan method expectedData should to be true": {
			inputData:    10,
			compare:      10,
			expectedData: true,
			execute: func(instance IntValidator, domain IntDomain) bool {
				return instance.IsLessOrEqualThan(domain.inputData, domain.compare)
			},
		},
		"error: when value is less than compare to IsLessOrEqualThan method expectedData should to be false": {
			inputData:    10,
			compare:      9,
			expectedData: false,
			execute: func(instance IntValidator, domain IntDomain) bool {
				return instance.IsLessOrEqualThan(domain.inputData, domain.compare)
			},
		},
		"success: when value are equal compare to AreEquals method expectedData should to be true": {
			inputData:    9,
			compare:      9,
			expectedData: true,
			execute: func(instance IntValidator, domain IntDomain) bool {
				return instance.AreEquals(domain.inputData, domain.compare)
			},
		},
		"error: when value are not equal compare to AreEquals method expectedData should to be false": {
			inputData:    9,
			compare:      8,
			expectedData: false,
			execute: func(instance IntValidator, domain IntDomain) bool {
				return instance.AreEquals(domain.inputData, domain.compare)
			},
		},
		"success: when value are not equal compare to AreNotEquals method expectedData should to be true": {
			inputData:    9,
			compare:      8,
			expectedData: true,
			execute: func(instance IntValidator, domain IntDomain) bool {
				return instance.AreNotEquals(domain.inputData, domain.compare)
			},
		},
		"error: when value are equal compare to AreNotEquals method expectedData should to be false": {
			inputData:    9,
			compare:      9,
			expectedData: false,
			execute: func(instance IntValidator, domain IntDomain) bool {
				return instance.AreNotEquals(domain.inputData, domain.compare)
			},
		},
		"success: when value is in range to IsBetween method expectedData should to be true": {
			inputData:    9,
			from:         1,
			to:           10,
			expectedData: true,
			execute: func(instance IntValidator, domain IntDomain) bool {
				return instance.IsBetween(domain.inputData, domain.from, domain.to)
			},
		},
		"error: when value is less than start value to IsBetween method expectedData should to be false": {
			inputData:    0,
			from:         1,
			to:           10,
			expectedData: false,
			execute: func(instance IntValidator, domain IntDomain) bool {
				return instance.IsBetween(domain.inputData, domain.from, domain.to)
			},
		},
		"error: when value is greater than end value to IsBetween method expectedData should to be false": {
			inputData:    11,
			from:         1,
			to:           10,
			expectedData: false,
			execute: func(instance IntValidator, domain IntDomain) bool {
				return instance.IsBetween(domain.inputData, domain.from, domain.to)
			},
		},
	}

	for name, useCase := range useCases {
		t.Run(name, func(t *testing.T) {
			instance := NewInt()

			response := useCase.execute(instance, useCase)

			if response != useCase.expectedData {
				t.Errorf("Expected %t, but got %t", useCase.expectedData, response)
			}
		})
	}
}

func getIntPointer(value int) *int {
	return &value
}
