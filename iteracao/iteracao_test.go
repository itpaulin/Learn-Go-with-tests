package iteracao

import (
	"fmt"
	"testing"
)

func TestRepetir(t *testing.T) {
	verifyRepetir := func(t *testing.T, result, expected string) {
		if result != expected {
			t.Errorf("resultado '%s', esperado '%s'", result, expected)
		}
	}

	t.Run("repetir sem passar valor", func(t *testing.T) {
		result := Repetir("a", 0)
		expected := "aaaaa"
		verifyRepetir(t, result, expected)
	})

	t.Run("repetir com valor 3", func(t *testing.T) {
		result := Repetir("b", 3)
		expected := "bbb"

		verifyRepetir(t, result, expected)
	})

}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a", 5)
	}
}

func ExampleRepetir() {
	result := Repetir("a", 5)
	fmt.Println(result)
	// Output: aaaaa
}
