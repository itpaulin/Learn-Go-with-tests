package inteiros

import (
	"fmt"
	"testing"
)

func TestAdicionadores(t *testing.T) {
	sum := Adiciona(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("resultado '%d', esperado '%d'", sum, expected)
	}
}

func ExampleAdiciona() {
	soma := Adiciona(1, 5)
	fmt.Println(soma)
	// Output: 6
}
