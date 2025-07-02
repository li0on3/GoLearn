package main

import (
	"fmt"
	"sort"
)

// InPlaceDeduplication 原地去重要判断切片是否是有序的   如果是有序的那么复杂度就是n   如果不是有序的  先排序再去重nlogn+n  或者 n²
// InPlaceDeduplication 原地去重
func InPlaceDeduplication(num []int) []int {
	if len(num) <= 1 {
		return num
	}
	sort.Ints(num)

	clear := 0
	for i := 1; i < len(num); i++ {
		if num[clear] != num[i] {
			clear++
			num[clear] = num[i]
		}
	}
	return num[:clear+1]
}

func main() {
	num := []int{1, 2, 1, 3, 5, 5, 7, 7, 8, 10}
	inplace := InPlaceDeduplication(num)
	fmt.Println(inplace)
}
