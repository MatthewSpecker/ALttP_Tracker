package menu

import (
	"testing"
)

func TestAlmostEqual(t *testing.T) {
	t.Run("confirm function returns true when floats are equal enough", func(t *testing.T) {
		got := almostEqual(1.000100, 1.000200)
		want := true

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
	t.Run("confirm function returns false when floats aren't equal enough", func(t *testing.T) {
		got := almostEqual(1.001000, 1.002000)
		want := false

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
}