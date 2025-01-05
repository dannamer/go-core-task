package main

import (
	"testing"
)

func TestDifference(t *testing.T) {
	tests := []struct {
		slice1, slice2 []string
		expected       []string
	}{
		{
			slice1:   []string{"apple", "banana", "cherry"},
			slice2:   []string{"banana", "date"},
			expected: []string{"apple", "cherry"},
		},
		{
			slice1:   []string{"apple", "banana", "cherry", "date", "lead"},
			slice2:   []string{"banana", "date", "fig"},
			expected: []string{"apple", "cherry", "lead"},
		},
		{
			slice1:   []string{"a", "b", "c"},
			slice2:   []string{"d", "e", "f"},
			expected: []string{"a", "b", "c"},
		},
		{
			slice1:   []string{"a", "b", "c", "c"},
			slice2:   []string{"a", "b"},
			expected: []string{"c", "c"},
		},
		{
			slice1:   []string{"apple", "banana", "cherry"},
			slice2:   []string{"apple", "banana", "cherry"},
			expected: []string{},
		},
	}

	for _, test := range tests {
		t.Run("Test", func(t *testing.T) {
			result := difference(test.slice1, test.slice2)
			if len(result) != len(test.expected) {
				t.Errorf("expected length %d, got %d", len(test.expected), len(result))
			}
			for i, item := range result {
				if item != test.expected[i] {
					t.Errorf("expected %s at index %d, got %s", test.expected[i], i, item)
				}
			}
		})
	}
}
