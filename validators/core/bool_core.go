package core

type BoolValidator struct{}

func NewBool() BoolValidator {
	return BoolValidator{}
}

func (validator BoolValidator) IsNull(value *bool) bool {
	return value == nil
}

func (validator BoolValidator) IsNotNull(value *bool) bool {
	return value != nil
}

func (validator BoolValidator) IsTrue(value bool) bool {
	return value
}

func (validator BoolValidator) IsFalse(value bool) bool {
	return !value
}
