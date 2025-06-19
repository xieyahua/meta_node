
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
package main
import (
	"fmt"
)
func isPalindrome(x int) bool {
    // y := len(strconv.Itoa(x)) -1 //判断数字有几位数
    y := x
    m := 0 
    for y>0 {
        m = m*10 + y%10
        y = y/10
    }
    if m == x {
        return true
    } else {
        return false
    }
}
