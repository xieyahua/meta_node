
// 1.只出现一次的数字
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

//2.回文数
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

// 3.有效的括号
func isValid(s string) bool {
    n := len(s)
    if n%2 == 1 {
        return false
    }
    
    pairs := map[byte]byte{
        ')': '(',
        ']': '[',
        '}': '{',
    }
    
    stack := []byte{}
    for i := 0; i < n; i++ {
        if pairs[s[i]] > 0 {
            if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
                return false
            }
            stack = stack[:len(stack)-1]
        } else {
            stack = append(stack, s[i])
        }
    }
    return len(stack) == 0
}

// 4.最长公共前缀

func longestCommonPrefix(strs []string) string {
	// 首先检查空数组情况
	if len(strs) == 0 {
		return ""
	}

	// 然后以第一个字符串作为初始前缀，依次与其他字符串比较，逐步缩短前缀直到找到最长公共前缀
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(prefix); j++ {
			if j >= len(strs[i]) || strs[i][j] != prefix[j] {
				prefix = prefix[:j]
				break
			}
		}
		if prefix == "" {
			break
		}
	}
	return prefix
}

// 5.删除排序数组中的重复项 
func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    slow := 0
    for fast := 1; fast < len(nums); fast++ {
        if nums[fast] != nums[slow] {
            slow++
            nums[slow] = nums[fast]
        }
    }
    return slow + 1
}

// 6.给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			// 将该数字之后的所有数字置为 0
			for j := i + 1; j < n; j++ {
				digits[j] = 0
			}
			return digits
		}
	}
	// 如果所有数字都是 9，则创建一个新的数组
	result := make([]int, n+1)
	result[0] = 1
	return result
}

// 7.删除有序数组中的重复项
func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    slow := 0
    for fast := 1; fast < len(nums); fast++ {
        if nums[fast] != nums[slow] {
            slow++
            nums[slow] = nums[fast]
        }
    }
    return slow + 1
}


// 8.两数之和
func twoSum(nums []int, target int) []int {
    var re_nums []int
    for i:=0;i<len(nums)-1;i++{
        for j:=i+1;j<len(nums);j++{
            if nums[i]+nums[j] == target{
                re_nums = append(re_nums,i)
                re_nums = append(re_nums,j)
		break
            }
        }
    }
    return re_nums
}
