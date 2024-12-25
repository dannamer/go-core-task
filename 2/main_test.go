package main

import (
	"testing"
)

func TestGenerateRandomSlice(t *testing.T) {
	slice := generateRandomSlice()
	if len(slice) != 10 {
		t.Errorf("Expected slice length to be 10, got %d", len(slice))
	}
}

func TestSliceExample(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{2, 4, 6, 8, 10}
	result := sliceExample(slice)
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, result[i])
		}
	}
}

func TestAddElements(t *testing.T) {
	slice := []int{1, 2, 3}
	expected := []int{1, 2, 3, 4}
	result := addElements(slice, 4)
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, result[i])
		}
	}
}

func TestCopySlice(t *testing.T) {
	slice := []int{1, 2, 3}
	copied := copySlice(slice)
	if &slice[0] == &copied[0] {
		t.Errorf("Expected slices to be different, but they are the same")
	}
}

func TestRemoveElement(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	expected := []int{1, 2, 4, 5}
	result, err := removeElement(slice, 2)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	for i, v := range expected {
		if result[i] != v {
			t.Errorf("Expected %d at index %d, but got %d", v, i, result[i])
		}
	}

	_, err = removeElement(slice, 10)
	if err == nil {
		t.Errorf("Expected error, but got nil")
	} else if err.Error() != "index out of range" {
		t.Errorf("Expected error 'index out of range', but got: %v", err)
	}
}
