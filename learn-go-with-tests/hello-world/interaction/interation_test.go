package interaction

import (
	"fmt"
	"testing"
)

func TestInteraction(t *testing.T){
	verifyCorrectness := func (t *testing.T, repeat,  expected string) {
		t.Helper()

		if repeat != expected {
			t.Errorf("expected '%s' but got '%s'", expected, repeat)
		}
	}

	t.Run("should run using the default repeat number", func(t *testing.T) {
		repeat := Repeat("a", 5)
		expected := "aaaaa"

		verifyCorrectness(t, repeat, expected)
	})

	t.Run("should repeat the char according to the repetition", func(t *testing.T) {
		repeat := Repeat("m", 10)
		expected := "mmmmmmmmmm"

		verifyCorrectness(t, repeat, expected)
	})
}

func BenchmarkInteraction(b *testing.B) {
	for i:=0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeat := Repeat("b", 3)

	fmt.Println(repeat)
	// Output: bbb
}