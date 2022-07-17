package core

import "github.com/lsnascimento/govalidations/validators/utils"

type StringValidator struct{}

func NewString() StringValidator {
	return StringValidator{}
}

func (validator StringValidator) IsNotNullOrEmpty(value string) bool {
	return value != "" && len(value) > 0
}

func (validator StringValidator) IsNullOrEmpty(value string) bool {
	return value == "" && len(value) == 0
}

func (validator StringValidator) HasMinLength(value string, min int) bool {
	return len(value) >= min
}

func (validator StringValidator) HasMaxLength(value string, max int) bool {
	return len(value) <= max
}

func (validator StringValidator) HasLength(value string, compare int) bool {
	return len(value) == compare
}

func (validator StringValidator) Contains(value, compare string) bool {
	return utils.Contains(value, compare)
}

func (validator StringValidator) AreEquals(value, compare string) bool {
	return value == compare
}

func (validator StringValidator) AreNotEquals(value, compare string) bool {
	return value != compare
}

func (validator StringValidator) IsEmail(value string) bool {
	pattern := `^\w+([-+.']\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`

	return utils.Matches(value, pattern)
}

func (validator StringValidator) Match(value, pattern string) bool {
	return utils.Matches(value, pattern)
}
