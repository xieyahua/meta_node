package main

import "fmt"



// 指针2
// DoubleSlice 接收切片指针并将每个元素乘以2
func DoubleSlice(slicePtr *[]int) {
    if slicePtr == nil {
        return
    }
    
    slice := *slicePtr
    for i := 0; i < len(slice); i++ {
        slice[i] *= 2
    }
}

func main() {
    nums := []int{1, 2, 3, 4, 5}
    fmt.Println("原切片:", nums)
    
    DoubleSlice(&nums)
    fmt.Println("处理后:", nums)
}
