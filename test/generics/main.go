package main

import (
	"errors"
	"fmt"
)

type CustomMap struct {
	data map[string]int
}

func (m *CustomMap) Insert(key string, v int) error {
	if _, ok := m.data[key]; ok {
		return errors.New("the key is exists")
	}
	m.data[key] = v
	return nil
}

func NewCustomMap() *CustomMap {
	return &CustomMap{
		data: make(map[string]int),
	}
}

type CusMap[K comparable, V any] struct {
	data map[K]V
}

func NewCusMap[K comparable, V any]() *CusMap[K, V] {
	return &CusMap[K, V]{
		data: make(map[K]V),
	}
}

func (m *CusMap[K, V]) Insert(key K, v V) error {
	if _, ok := m.data[key]; ok {
		return errors.New("the key exists")
	}
	m.data[key] = v
	return nil
}

func main() {
	cm := NewCustomMap()
	cm.Insert("hello", 100)
	fmt.Println(cm.data["hello"])

	// generics
	m1 := NewCusMap[string, int]()
	m1.Insert("Hello", 100)
	val := m1.data["Hello"]
	fmt.Println(val)
}
