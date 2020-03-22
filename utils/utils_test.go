package utils

import (
	"testing"
)

func twoArrayEqual(arr1, arr2 []float64) bool {
	if len(arr1) == len(arr2) {
		for i := 0; i < len(arr1); i++ {
			if arr1[i] != arr2[i] {
				return false
			}
		}
		return true
	}
	return false
}

func TestReverse(t *testing.T) {
	cases := []struct {
		value    []float64
		expected []float64
	}{
		{
			value:    []float64{1, 2, 3, 4, 5},
			expected: []float64{5, 4, 3, 2, 1},
		},
		{
			value:    []float64{56, 34, 78, 12, 222},
			expected: []float64{222, 12, 78, 34, 56},
		},
		{
			value:    []float64{-1, 2, -3, 5, 1000},
			expected: []float64{1000, 5, -3, 2, -1},
		},
	}

	for _, c := range cases {
		dt := c.value
		Reverse(&dt)
		if !twoArrayEqual(dt, c.expected) {
			t.Errorf("%s:\nexpected: %v\nactual: %v", "Reverse", c.expected, dt)
		}
	}
}
