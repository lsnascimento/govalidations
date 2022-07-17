package core

import (
	"testing"
)

type SliceDomain struct {
	inputData    interface{}
	compare      int
	expectedData bool
	execute      func(instance SliceValidator, domain SliceDomain) bool
}

func TestSliceValidator(t *testing.T) {
	useCases := map[string]SliceDomain{
		"success: when slice is empty to IsEmpty method expectedData should to be true": {
			inputData:    []string{},
			expectedData: true,
			execute: func(instance SliceValidator, domain SliceDomain) bool {
				return instance.IsEmpty(domain.inputData)
			},
		},
		"error: when slice is not empty to IsEmpty method expectedData should to be false": {
			inputData:    []string{"a"},
			expectedData: false,
			execute: func(instance SliceValidator, domain SliceDomain) bool {
				return instance.IsEmpty(domain.inputData)
			},
		},
		"success: when slice is not empty to IsNotEmpty method expectedData should to be true": {
			inputData:    []string{"a"},
			expectedData: true,
			execute: func(instance SliceValidator, domain SliceDomain) bool {
				return instance.IsNotEmpty(domain.inputData)
			},
		},
		"error: when slice is not empty to IsNotEmpty method expectedData should to be false": {
			inputData:    []string{},
			expectedData: false,
			execute: func(instance SliceValidator, domain SliceDomain) bool {
				return instance.IsNotEmpty(domain.inputData)
			},
		},
		"success: when value is greater than compare in HasMinLength method expectedData should to be true": {
			inputData:    []string{"a", "b", "c"},
			compare:      2,
			expectedData: true,
			execute: func(instance SliceValidator, domain SliceDomain) bool {
				return instance.HasMinLength(domain.inputData, domain.compare)
			},
		},
		"success: when value are equal to compare in HasMinLength method expectedData should to be true": {
			inputData:    []string{"a", "b"},
			compare:      2,
			expectedData: true,
			execute: func(instance SliceValidator, domain SliceDomain) bool {
				return instance.HasMinLength(domain.inputData, domain.compare)
			},
		},
		"error: when value is less than compare in HasMinLength method expectedData should to be false": {
			inputData:    []string{"a"},
			compare:      2,
			expectedData: false,
			execute: func(instance SliceValidator, domain SliceDomain) bool {
				return instance.HasMinLength(domain.inputData, domain.compare)
			},
		},
		"success: when value is less than compare in HasMaxLength method expectedData should to be true": {
			inputData:    []string{"a", "b", "c"},
			compare:      4,
			expectedData: true,
			execute: func(instance SliceValidator, domain SliceDomain) bool {
				return instance.HasMaxLength(domain.inputData, domain.compare)
			},
		},
		"success: when value are equal to compare in HasMaxLength method expectedData should to be true": {
			inputData:    []string{"a", "b", "c", "d"},
			compare:      4,
			expectedData: true,
			execute: func(instance SliceValidator, domain SliceDomain) bool {
				return instance.HasMaxLength(domain.inputData, domain.compare)
			},
		},
		"error: when value is greater than compare in HasMaxLength method expectedData should to be false": {
			inputData:    []string{"a", "b", "c"},
			compare:      2,
			expectedData: false,
			execute: func(instance SliceValidator, domain SliceDomain) bool {
				return instance.HasMaxLength(domain.inputData, domain.compare)
			},
		},
		"success: when value are equal to compare in HasLength method expectedData should to be true": {
			inputData:    []string{"a", "b", "c", "d"},
			compare:      4,
			expectedData: true,
			execute: func(instance SliceValidator, domain SliceDomain) bool {
				return instance.HasLength(domain.inputData, domain.compare)
			},
		},
		"error: when value are not equal compare in HasLength method expectedData should to be false": {
			inputData:    []string{"a", "b", "c"},
			compare:      2,
			expectedData: false,
			execute: func(instance SliceValidator, domain SliceDomain) bool {
				return instance.HasLength(domain.inputData, domain.compare)
			},
		},
	}

	for name, useCase := range useCases {
		t.Run(name, func(t *testing.T) {
			instance := NewSlice()

			response := useCase.execute(instance, useCase)

			if response != useCase.expectedData {
				t.Errorf("Expected %t, but got %t", useCase.expectedData, response)
			}
		})
	}
}
