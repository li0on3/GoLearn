package main

import (
	"errors"
	"fmt"
)

//动态数组 用了一个泛型能够匹配所有的类型  并且做了相对应的基础方法  也创建了动态扩容的算法

type DynamicArray[T any] struct {
	data     []T
	size     int
	capacity int
}

func main() {
	arr := NewDynamicArray[int](2) // 关键修复点：类型参数 + 调用

	fmt.Println("Pushing 1, 2, 3 into DynamicArray:")
	arr.Push(1)
	arr.Push(2)
	arr.Push(3) // will trigger resize

	for i := 0; i < arr.size; i++ {
		val, _ := arr.Get(i)
		fmt.Printf("Index %d: %d\n", i, val)
	}

	fmt.Println("Setting index 1 to 99")
	arr.Set(1, 99)

	val, _ := arr.Get(1)
	fmt.Printf("Index 1 after Set: %d\n", val)

	fmt.Println("Popping values:")
	for i := 0; i < 3; i++ {
		v, err := arr.Pop()
		if err != nil {
			fmt.Println("Pop error:", err)
			break
		}
		fmt.Println("Popped:", v)
	}

	fmt.Println("Pop from empty array:")
	_, err := arr.Pop()
	if err != nil {
		fmt.Println("Expected error:", err)
	}
}

func NewDynamicArray[T any](capacity int) *DynamicArray[T] {
	if capacity < 1 {
		capacity = 1
	}
	return &DynamicArray[T]{
		data:     make([]T, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (t *DynamicArray[T]) Push(item T) {
	if t.size >= t.capacity {
		t.resize()
	}
	t.data[t.size] = item
	t.size++
}

func (t *DynamicArray[T]) Pop() (T, error) {
	if t.size == 0 {
		var zero T
		return zero, errors.New("empty")
	}
	item := t.data[t.size-1]
	t.size--
	return item, nil
}

func (t *DynamicArray[T]) Get(index int) (T, error) {
	if index < 0 || index >= t.size {
		var zero T
		return zero, errors.New("index out of bounds")
	}
	return t.data[index], nil
}

func (t *DynamicArray[T]) Set(index int, value T) error {
	if index < 0 || index >= t.size {
		return errors.New("index out of bounds")
	}
	t.data[index] = value
	return nil
}

func (t *DynamicArray[T]) resize() {
	var newCapacity int
	if t.capacity > 1024 {
		newCapacity = t.capacity + t.capacity/2 + 1
	} else {
		newCapacity = t.capacity * 2
	}
	newData := make([]T, newCapacity)
	copy(newData, t.data[:t.size])

	t.data = newData
	t.capacity = newCapacity
}

//这里面的泛型我不怎么会用
