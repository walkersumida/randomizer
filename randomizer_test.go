package randomizer

import "testing"

func TestRun(t *testing.T) {
	t.Run("with a,b,c", func(t *testing.T){
		expected := 1
		got := len(Run("a,b,c"))

		if got != expected {
			t.Errorf("\nexpected: %d\n     got: %d", expected, got)
		}
	})

	t.Run("with 1,2,3", func(t *testing.T){
		expected := 1
		got := len(Run("1,2,3,4"))

		if got != expected {
			t.Errorf("\nexpected: %d\n     got: %d", expected, got)
		}
	})

	t.Run("with 10,20,30", func(t *testing.T){
		expected := 2
		got := len(Run("10,20,30"))

		if got != expected {
			t.Errorf("\nexpected: %d\n     got: %d", expected, got)
		}
	})
}