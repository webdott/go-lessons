package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Returns the sum of all numbers in an array", func(t *testing.T) {
		number := Sum([]int{3, 4, 5, 6, 7})
		expected := 25

		if number != expected {
			t.Errorf("Expected %d but got %d", expected, number)
		}
	})
}

func assertError(t testing.TB, got []int, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAll(t *testing.T) {
	t.Run("Returns each sum of all numbers in arrays of number list ", func(t *testing.T) {
		number := SumAll([]int{3, 4, 5, 6, 7}, []int{3, 4, 5, 6, 9})
		expected := ([]int{25, 27})

		assertError(t, number, expected)
	})
}

func TestTails(t *testing.T) {
	t.Run("Returns each sum of all tails in arrays of number list", func(t *testing.T) {
		number := SumTails([]int{3, 4, 5, 6, 7}, []int{3, 4, 5, 6, 9})
		expected := ([]int{22, 24})

		assertError(t, number, expected)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		assertError(t, got, want)
	})
}
