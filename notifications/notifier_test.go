package notifications

import "testing"

func TestNotifier(t *testing.T) {
	t.Run("Should return true when notifications is invalid", func(t *testing.T) {
		instance := New()
		expected := true

		instance.AddMessage("property", "message")

		actual := instance.Invalid()

		if expected != actual {
			t.Errorf("Expected %t, but got %t", expected, actual)
		}
	})

	t.Run("Should return true when notifications is valid", func(t *testing.T) {
		instance := New()
		expected := true

		actual := instance.Valid()

		if expected != actual {
			t.Errorf("Expected %t, but got %t", expected, actual)
		}
	})

	t.Run("Should return true when notifications has elements", func(t *testing.T) {
		instance := New()
		expected := 1

		instance.AddMessage("property", "message")

		actual := len(instance.GetMessages())

		if expected != actual {
			t.Errorf("Expected %d, but got %d", expected, actual)
		}
	})
}
