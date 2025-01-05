package main

import "testing"

func TestFindIntersection(t *testing.T) {
	tests := []struct {
		a, b          []int
		expectedBool  bool
		expectedSlice []int
	}{
		{
			a:              []int{65, 3, 58, 678, 64},
			b:              []int{64, 2, 3, 43},
			expectedBool:   true,
			expectedSlice:  []int{64, 3},
		},
		{
			a:              []int{1, 2, 3},
			b:              []int{4, 5, 6},
			expectedBool:   false,
			expectedSlice:  []int{},
		},
		{
			a:              []int{10, 20, 30},
			b:              []int{30, 40, 50},
			expectedBool:   true,
			expectedSlice:  []int{30},
		},
		{
			a:              []int{7, 8, 9},
			b:              []int{1, 2, 3},
			expectedBool:   false,
			expectedSlice:  []int{},
		},
	}

	for _, test := range tests {
		t.Run("Test", func(t *testing.T) {
			resultBool, resultSlice := findIntersection(test.a, test.b)

			// Проверка результата типа bool
			if resultBool != test.expectedBool {
				t.Errorf("expected %v, got %v", test.expectedBool, resultBool)
			}

			// Проверка результата среза
			if len(resultSlice) != len(test.expectedSlice) {
				t.Errorf("expected slice length %d, got %d", len(test.expectedSlice), len(resultSlice))
			}
			for i, item := range resultSlice {
				if item != test.expectedSlice[i] {
					t.Errorf("expected %d at index %d, got %d", test.expectedSlice[i], i, item)
				}
			}
		})
	}
}
