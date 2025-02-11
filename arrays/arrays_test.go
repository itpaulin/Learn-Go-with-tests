package arrays

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection with size 5", func(t *testing.T) {

		numbers := [5]int{1, 2, 3, 4, 5}
		result := SumArray(numbers)
		expected := 15

		if result != expected {
			t.Errorf("resultado '%d', esperado '%d', enviado %v", result, expected, numbers)
		}
	})

	t.Run("collection with any syze", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7}
		expected := 28
		result := SumSlice(numbers)

		if result != expected {
			t.Errorf("resultado '%d', esperado '%d', enviado %v", result, expected, numbers)

		}
	})
}
