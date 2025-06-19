
// 只出现一次的数字
package main

import (
	"fmt"
)
func Task1() {
	method := func(nums []int) int {
		numCount := make(map[int]int)
		for _, num := range nums {
			numCount[num]++
		}

		for num, count := range numCount {
			if count == 1 {
				return num
			}
		}

		panic("no unique element found")
	}
	nums := []int{2, 2, 1}
	fmt.Println(method(nums))
}

//回文数

