package core

import "reflect"

type SliceValidator struct{}

func NewSlice() SliceValidator {
	return SliceValidator{}
}

func (validator SliceValidator) IsEmpty(value interface{}) bool {
	valueInSlice := validator.toSlice(value)

	return len(valueInSlice) == 0
}

func (validator SliceValidator) IsNotEmpty(value interface{}) bool {
	valueInSlice := validator.toSlice(value)

	return len(valueInSlice) > 0
}

func (validator SliceValidator) HasMinLength(value interface{}, compare int) bool {
	valueInSlice := validator.toSlice(value)

	return len(valueInSlice) >= compare
}

func (validator SliceValidator) HasMaxLength(value interface{}, compare int) bool {
	valueInSlice := validator.toSlice(value)

	return len(valueInSlice) <= compare
}

func (validator SliceValidator) HasLength(value interface{}, compare int) bool {
	valueInSlice := validator.toSlice(value)

	return len(valueInSlice) == compare
}

func (validator SliceValidator) toSlice(value interface{}) (slice []interface{}) {
	element := reflect.ValueOf(value)
	elementsOfSlice := reflect.ValueOf(element.Interface())
	quantityElements := elementsOfSlice.Len()

	for i := 0; i < quantityElements; i++ {
		slice = append(slice, elementsOfSlice.Index(i).Interface())
	}

	return
}
