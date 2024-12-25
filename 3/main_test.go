package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	m := NewStringIntMap()
	m.Add("apple", 5)

	value, exists := m.Get("apple")
	if !exists {
		t.Errorf("Expected 'apple' to exist in map")
	}
	if value != 5 {
		t.Errorf("Expected value 5, got %d", value)
	}
}

func TestRemove(t *testing.T) {
	m := NewStringIntMap()
	m.Add("apple", 5)
	m.Remove("apple")

	_, exists := m.Get("apple")
	if exists {
		t.Errorf("Expected 'apple' to be removed from map")
	}
}

func TestCopy(t *testing.T) {
	m := NewStringIntMap()
	m.Add("apple", 5)
	m.Add("banana", 10)

	copiedMap := m.Copy()

	// Проверим, что элементы были скопированы
	if copiedMap["apple"] != 5 {
		t.Errorf("Expected copied map to contain 'apple' with value 5")
	}
	if copiedMap["banana"] != 10 {
		t.Errorf("Expected copied map to contain 'banana' with value 10")
	}

	// Проверим, что оригинальная карта не изменилась
	m.Remove("apple")
	if _, exists := copiedMap["apple"]; exists {
		t.Errorf("Expected 'apple' to remain in copied map")
	}
}

func TestExists(t *testing.T) {
	m := NewStringIntMap()
	m.Add("apple", 5)

	if !m.Exists("apple") {
		t.Errorf("Expected 'apple' to exist in map")
	}
	if m.Exists("banana") {
		t.Errorf("Expected 'banana' to not exist in map")
	}
}

func TestGet(t *testing.T) {
	m := NewStringIntMap()
	m.Add("apple", 5)

	value, exists := m.Get("apple")
	if !exists {
		t.Errorf("Expected 'apple' to exist in map")
	}
	if value != 5 {
		t.Errorf("Expected value 5 for 'apple', got %d", value)
	}

	_, exists = m.Get("banana")
	if exists {
		t.Errorf("Expected 'banana' to not exist in map")
	}
}
