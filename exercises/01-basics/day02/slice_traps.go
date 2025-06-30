package main

import "fmt"

// 切片共享数组的问题  当切片是由数组转换而来   两个切片共用一个底层数组  这样一个切片修改的时候 会影响另外的切片  因为底层是一个共享数组
// 切片作为函数参数传递是值传递  切片的值传递有何其他的有一些差别  他会复制切片结构体的底层数组地址  长度  容量等信息作为一个新的切片
// 新的切片修改元素还是会影响原有的切片  但是如果增加了值以后  原有的切片看不到  因为切片能够看到的范围和len有关 原切片的len没有变 因此新增的元素看不到  如果修改这个值就可以看到了
// nil和空切片   nil切片不会分配数据指针  但是空指针会  而且空指针JSON序列化后是[]   还有什么不同呢
func main() {
	nilVsEmptySlice()
	funcParamTrap()
	sharedSliceTrap()
}

// Slice中共享数组的影响
func sharedSliceTrap() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	slice1 := arr[2:5]
	slice2 := arr[2:6]

	fmt.Printf("orignal array: %v\n", arr)
	fmt.Printf("first slice: %v\n", slice1)
	fmt.Printf("second slice: %v\n", slice2)

	//这样会报错  因为只有四个值   最大只能到3
	//slice2[4] = 10
	//fmt.Println(slice2[4])
	slice2[2] = 10
	fmt.Println(slice2[2])
	fmt.Println(slice1[2])

	if slice1[2] == slice2[2] {
		fmt.Println("The modification of the second slice affects the first slice. ")
	} else {
		fmt.Println("The modification of the first slice will not affect the second slice.")
	}

}

// Slice作为函数传入参数  不是值传递   是引用类型 但是修改原有的元素会影响

func modifySlice(s []int) {
	if len(s) > 0 {
		s[0] = 100
	}
	s = append(s, 1001)
	fmt.Printf("函数内: %v (len=%d, cap=%d)\n", s, len(s), cap(s))
}

func funcParamTrap() {
	s := []int{1, 2, 3}
	fmt.Printf("调用前: %v (len=%d, cap=%d)\n", s, len(s), cap(s))
	modifySlice(s)
	fmt.Printf("调用后: %v (len=%d, cap=%d)\n", s, len(s), cap(s))
}

func nilVsEmptySlice() {
	var s1 []int
	s2 := []int{1, 2, 3}
	s3 := make([]int, 0)

	fmt.Printf("s1: %v, len=%d, cap=%d, is nil? %v\n", s1, len(s1), cap(s1), s1 == nil)
	fmt.Printf("s2: %v, len=%d, cap=%d, is nil? %v\n", s2, len(s2), cap(s2), s2 == nil)
	fmt.Printf("s3: %v, len=%d, cap=%d, is nil? %v\n", s3, len(s3), cap(s3), s3 == nil)
}
