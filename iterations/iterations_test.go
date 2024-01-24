package main

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	t.Run("Repeat character 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected '%s' but got '%s'", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}
}

func ExampleIteration() {
	fmt.Println(Repeat("a", 3))
	// Output: aaa
}
