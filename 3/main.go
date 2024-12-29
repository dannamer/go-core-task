package main

import "fmt"

type StringIntMap struct {
	data map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{data: make(map[string]int)}
}

func (m *StringIntMap) Add(key string, value int) {
	m.data[key] = value
}

func (m *StringIntMap) Remove(key string) {
	delete(m.data, key)
}

// Метод для копирования карты
func (m *StringIntMap) Copy() map[string]int {
	copyMap := make(map[string]int)
	for k, v := range m.data {
		copyMap[k] = v
	}
	return copyMap
}

func (m *StringIntMap) Exists(key string) bool {
	_, exists := m.data[key]
	return exists
}

func (m *StringIntMap) Get(key string) (int, bool) {
	value, exists := m.data[key]
	return value, exists
}

func main() {
	m := NewStringIntMap()
	m.Add("apple", 5)
	m.Add("banana", 10)

	fmt.Println(m.Get("apple"))
	fmt.Println(m.Exists("grape"))
	m.Remove("apple")
	fmt.Println(m.Get("apple"))
}
