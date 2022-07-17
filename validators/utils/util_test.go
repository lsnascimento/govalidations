package utils

import (
	"testing"
)

func TestUtil(t *testing.T) {
	t.Run("Should be able to reverse string", func(t *testing.T) {
		input := "123456"
		expected := "654321"
		actual := Reverse(input)

		if expected != actual {
			t.Errorf("Expected %s, but got %s", expected, actual)
		}
	})

	t.Run("Should be able to replace string in pattern", func(t *testing.T) {
		str := "123"
		pattern := "[0-9]"
		expected := "aaa"
		actual := ReplacePattern(str, pattern, "a")

		if expected != actual {
			t.Errorf("Expected %s, but got %s", expected, actual)
		}
	})

	t.Run("Should be able to find string in string", func(t *testing.T) {
		str := "123456"
		substring := "456"
		expected := true
		ok := Contains(str, substring)

		if !ok {
			t.Errorf("Expected %t, but got %t", expected, ok)
		}
	})

	t.Run("Should be able to find string in pattern", func(t *testing.T) {
		str := "123456"
		substring := "[0-9]"
		expected := true
		ok := Matches(str, substring)

		if !ok {
			t.Errorf("Expected %t, but got %t", expected, ok)
		}
	})

	t.Run("Should be able to validate date", func(t *testing.T) {
		str := "2022-06-17"
		expected := true
		ok, _ := IsDate(str)

		if !ok {
			t.Errorf("Expected %t, but got %t", expected, ok)
		}
	})
}
