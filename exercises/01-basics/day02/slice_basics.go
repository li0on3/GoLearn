package main

import (
	"fmt"
	"reflect"
)

//三种创造切片的方式  var  :=   make
//输出长度 容量   底层数组     len   cap  uintptr
//演示追加操作   主要是向slice中推送数据  并且跟踪len 和cap的变化
//切片的扩容机制  当新增数据发现容量不足时  查看容量大小  <256时  新容量为原来的2倍   >256则扩容1/2
//结果:
//nil slice
//Not nil slice
//Not nil slice
//slice len: 6, cap: 10
//slice len: 11, cap: 20
//slice len: 21, cap: 40
//slice len: 41, cap: 80

func main() {
	//1.声明切片
	var slice1 []int
	if slice1 == nil {
		fmt.Println("nil slice")
	} else {
		fmt.Println(" Not nil slice")
	}
	if slice1 == nil {
		slice1 = make([]int, 0, 5)
	}

	//2. :=
	slice2 := []int{1, 2, 3, 4, 5}
	if slice2 == nil {
		fmt.Println("nil slice")
	} else {
		fmt.Println("Not nil slice")
	}

	//3.make
	slice3 := make([]int, 0, 5)
	if slice3 == nil {
		fmt.Println("nil slice")
	} else {
		fmt.Println("Not nil slice")
	}
	printSliceLength(slice1)
	sliceAppendProcess(slice3)
}

// 输出slice为int的长度 容量  底层数组地址
func printSliceLength(s interface{}) (int, int, uintptr) {
	v := reflect.ValueOf(s)

	if v.Kind() != reflect.Slice {
		fmt.Println("not slice")
		return 0, 0, 0
	}

	ptr := v.UnsafePointer()
	length := v.Len()
	c := v.Cap()

	return length, c, uintptr(ptr)
}

func sliceAppendProcess(s []int) {

	limit := 50
	lastCaps := cap(s)
	for i := 0; i < limit; i++ {
		s = append(s, i)
		if cap(s) != lastCaps {
			fmt.Printf("slice len: %d, cap: %d\n", len(s), cap(s))
			lastCaps = cap(s)
		}
	}

}
