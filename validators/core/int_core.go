package core

type IntValidator struct{}

func NewInt() IntValidator {
	return IntValidator{}
}

func (validator IntValidator) IsNull(value *int) bool {
	return value == nil
}

func (validator IntValidator) IsNotNull(value *int) bool {
	return value != nil
}

func (validator IntValidator) IsGreaterThan(value, compare int) bool {
	return value > compare
}

func (validator IntValidator) IsGreaterOrEqualThan(value, compare int) bool {
	return value >= compare
}

func (validator IntValidator) IsLessThan(value, compare int) bool {
	return value < compare
}

func (validator IntValidator) IsLessOrEqualThan(value, compare int) bool {
	return value <= compare
}

func (validator IntValidator) AreEquals(value, compare int) bool {
	return value == compare
}

func (validator IntValidator) AreNotEquals(value, compare int) bool {
	return value != compare
}

func (validator IntValidator) IsBetween(value, from, to int) bool {
	return value >= from && value <= to
}
