package interaction

import "testing"

func TestInteraction(t *testing.T){
	repeat := Repeat("a")
	expected := "aaaaa"

	if repeat != expected {
		t.Errorf("expected '%s' but got '%s'", expected, repeat)
	}
}

func BenchmarkInteraction(b *testing.B) {
	for i:=0; i < b.N; i++ {
		Repeat("a")
	}
}