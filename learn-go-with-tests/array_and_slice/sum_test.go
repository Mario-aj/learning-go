package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("should sum the collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		result := Sum(numbers)
		expected := 15

		if expected != result {
			t.Errorf("Result '%d', expected '%d', data %v", result, expected, numbers)
		}
	})

	t.Run("should sum collect of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		result := Sum(numbers)
		expected := 6

		if expected != result {
			t.Errorf("Result '%d', expected '%d', data %v", result, expected, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	result := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result '%v' , expected '%v'", result, expected)
	}
}
